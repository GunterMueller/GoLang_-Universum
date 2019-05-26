package dgra

// (c) Christian Maurer   v. 190402 - license see nU.go

import ("strconv"; . "nU/obj"; "nU/env"; "nU/nchan"
        "nU/fmon"; "nU/adj"; "nU/vtx"; "nU/gra"
        "nU/dgra/status" )

type distributedGraph struct {
  gra.Graph
  tmpGraph gra.Graph
  bool "gra.Graph.Directed"
  actVertex vtx.Vertex
  me uint
  actHost string
  n uint
  nb []vtx.Vertex
  nr []uint
  host []string
  port []uint16
  ch []nchan.NetChannel
  tree, cycle gra.Graph
  root uint
  tmpVertex vtx.Vertex
  tmpEdge uint16
  chan1 chan uint
  visited, sendTo, echoed []bool
  labeled bool
  mon []fmon.FarMonitor
  parent uint
  child []bool
  time, time1 uint
  matrix adj.AdjacencyMatrix
  size uint
  diameter, distance, leader uint
  PulseAlg; ElectAlg; TravAlg
  Op
// for KochMoranZaks:
  state byte // King or Citizen
  pik status.Status // phase and identity of the king of the calling process
  unused []bool
  parentChannel uint // channel-number to parent
  msgchan []chan Stream
}
const (
  p0 = nchan.Port0
  inf = uint(1 << 16)
)
var (
  c0 = C0()
  done = make(chan int, 1)
)

func new_(g gra.Graph) DistributedGraph {
  x := new(distributedGraph)
  x.Graph = g
  x.Graph.SetWrite (vtx.W, vtx.W2)
  x.tmpGraph = Clone(x.Graph).(gra.Graph)
  x.bool = x.Graph.Directed()
  x.actVertex = x.Graph.Get().(vtx.Vertex)
  x.me = x.actVertex.Val()
  x.actHost = env.Localhost()
  x.n = x.Graph.Num() - 1
  x.nb = make([]vtx.Vertex, x.n)
  x.nr = make([]uint, x.n)
  x.host = make([]string, x.n)
  x.ch = make([]nchan.NetChannel, x.n)
  x.port = make([]uint16, x.n)
  for i := uint(0); i < x.n; i++ {
    g.Ex (x.actVertex)
    x.nb[i] = g.Neighbour(i).(vtx.Vertex)
    x.nr[i] = x.nb[i].(Valuator).Val()
    x.Graph.Ex2 (x.actVertex, x.nb[i])
    x.port[i] = p0 + x.Graph.Get1().(uint16)
  }
  x.tmpVertex = vtx.New (x.me)
  v0 := g.Neighbour(0).(vtx.Vertex)
  g.Ex2 (x.actVertex, v0)
  if g.Edged() {
    x.tmpEdge = g.Get1().(uint16)
  } else {
    g.Ex2 (v0, x.actVertex)
    x.tmpEdge = g.Get1().(uint16)
  }
  x.tree = gra.New (false, x.tmpVertex, x.tmpEdge)
  x.tree.SetWrite (vtx.W, vtx.W2)
  x.cycle = gra.New (true, x.tmpVertex, x.tmpEdge)
  x.cycle.SetWrite (vtx.W, vtx.W2)
  x.visited = make([]bool, x.n)
  x.sendTo = make([]bool, x.n)
  x.echoed = make([]bool, x.n)
  x.chan1 = make(chan uint, 1)
  x.parent, x.child = inf, make([]bool, x.n)
  x.mon = make([]fmon.FarMonitor, x.n)
  g.Ex (x.actVertex)
  x.PulseAlg, x.ElectAlg, x.TravAlg = PulseGraph, ChangRoberts, DFS
  x.leader = x.me
  x.Op = Ignore
  return x
}

func newg (dir bool, l, c []uint, e [][]uint, h []string, m, id uint) DistributedGraph {
  g := gra.New (dir, vtx.New(uint(0)), uint16(1)) // g ist ein neuer
       // Graph mit Eckentyp uint und Kanten vom Typ von ports
  g.SetWrite (vtx.W, vtx.W2)
  n := uint(len(e))
  v := make([]vtx.Vertex, n)     // Ecken aller Prozesse erzeugen
  for i := uint(0); i < n; i++ { // und für alle i < n
    v[i] = vtx.New(i)            // mit i als Wert füllen, mit der
    v[i].Set (l[i], c[i])        // Position (l[i],c[i]) versehen
    g.Ins(v[i])                  // und in g einfügen 
  }
  for i := uint(0); i < n; i++ {
    for _, j := range e[i] {
      g.Ex2 (v[i], v[j])     // i ist jetzt die kolokale
                             // und j die lokale Ecke in g
      if ! g.Edged() { // wenn es noch keine Kante
                       // von i nach j gibt,
        p := nchan.Port (n, i, j, 0)
        g.Edge (p)     // i mit j durch eine Kante mit dem Wert
                       // des o.a. Ports verbinden
      }
    }
  }
  g.Ex (v[id])      // v[id] ist jetzt die lokale Ecke in g
  g.SetWrite (vtx.W, vtx.W2)
  g = g.Star()   // g ist jetzt nur noch der Stern von v[id]
  d := new_(g).(*distributedGraph) // d ist der verteilte
                 // Graph mit g als zugrundeliegendem Graphen
  d.setSize (n)  // Zeilen/Spaltenzahl der Adjazenzmatrix 
                 // = Anzahl der Ecken von g
  if h == nil {  // damit examples.go einfacher wird
    h = make([]string, n)
    for i := uint(0); i < n; i++ {
      h[i] = env.Localhost()
    }
  }
  for i := uint(0); i < d.n; i++ { // die Namen der Rechner,
    d.host[i] = h[d.nr[i]] // auf denen die Nachbarprozesse
  } // laufen, aus examples.go übernehmen (oder s.o.)
  d.diameter = m // Durchmesser von g
  return d
}

func (x *distributedGraph) SetRoot (r uint) {
  x.root = r
}

func (x *distributedGraph) setSize (n uint) {
  x.size = n
  x.matrix = adj.New (x.size, uint(0), uint(0))
  for i := uint(0); i < x.n; i++ {
    x.matrix.Set (x.me, x.nr[i], uint(0), uint(1))
    x.matrix.Set (x.nr[i], x.me, uint(0), uint(1))
  }
}

func (x *distributedGraph) connect (a Any) {
  for i := uint(0); i < x.n; i++ {
    x.ch[i] = nchan.New (a, x.me, x.nr[i], x.host[i], x.port[i])
  }
}

func (x *distributedGraph) fin() {
  for i := uint(0); i < x.n; i++ {
    x.ch[i].Fin()
  }
}

func (x *distributedGraph) finMon() {
  for i := uint(0); i < x.n; i++ {
    x.mon[i].Fin()
  }
}

// Liefert die Identität des Nachbarn, mit dem der aufrufende
// Prozess über den j-ten Netzkanal (j < x.n) verbunden ist;
// für alle j < x.n gilt also genau dann x.channel(i) == j,
// wenn x.nr[j] == i.
func (x *distributedGraph) channel (id uint) uint {
  j := x.n
  for i := uint(0); i < x.n; i++ {
    if x.nr[i] == id {
      j = i
      break
    }
  }
  return j
}

// Liefert einen neuen leeren Graphen
// vom Typ des zugrundliegenden Graphen von x.
func (x *distributedGraph) emptyGraph() gra.Graph {
  g := gra.New (x.tmpGraph.Directed(), x.tmpVertex, x.tmpEdge)
  g.SetWrite (vtx.W, vtx.W2)
  return g
}

func (x *distributedGraph) decodedGraph (bs Stream) gra.Graph {
  g := x.emptyGraph()
  g.Decode(bs)
  return g
}

func (x *distributedGraph) edge (v, v1 vtx.Vertex) uint16 {
  x.tmpGraph.Ex2 (v, v1)
  return x.tmpGraph.Get1().(uint16)
}

func valueMax (g gra.Graph) uint {
  m := uint(0)
  g.Trav (func (a Any) {
    v := a.(uint)
    if v > m { m = v }
  })
  return m
}

func exValue (g gra.Graph, v uint) bool {
  return g.ExPred (func (a Any) bool { return a.(uint) == v })
}

func (x *distributedGraph) next (i uint) uint {
  for u := uint(0); u < x.n; u++ {
    if u != i && ! x.visited[u] {
      return u
    }
  }
  return x.n
}

func (x *distributedGraph) Me() uint {
  return x.me
}

func (x *distributedGraph) Root() uint {
  return x.root
}

func (x *distributedGraph) Parent() uint {
  return x.parent
}

func (x *distributedGraph) Children() string {
  s := ""
  n := uint(0)
  for i := uint(0); i < x.n; i++ {
    if x.child[i] {
      n++
      if n > 1 {
        s += ", "
      }
      s += strconv.Itoa(int(x.nr[i]))
    }
  }
  if n == 0 { s = "-" }
  return s
}

func (x *distributedGraph) Time() uint {
  return x.time
}

func (x *distributedGraph) Time1() uint {
  return x.time1
}
