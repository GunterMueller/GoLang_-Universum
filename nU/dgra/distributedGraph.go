package dgra

// (c) Christian Maurer   v. 171129 - license see nU.go

import ("strconv"; . "nU/obj"; "nU/env"; "nU/nchan"; "nU/fmon"; "nU/adj"; "nU/gra")

type distributedGraph struct {
  gra.Graph
  tmpGraph gra.Graph
  bool "gra.Graph.Directed"
  me uint
  actHost string
  n uint
  nr []uint
  host []string
  port []uint16
  ch []nchan.NetChannel
  tree, ring gra.Graph
  root uint
  tmpVertex uint
  tmpEdge uint16
  chan1 chan uint
  visited, sendTo []bool
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
}
const (
  p0 = nchan.Port0
  inf = uint(1 << 16)
)
var done = make(chan int, 1)

func value (a Any) uint {
  return a.(uint)
}

func new_(g gra.Graph) DistributedGraph {
  x := new(distributedGraph)
  x.Graph = g
  x.tmpGraph = Clone(x.Graph).(gra.Graph)
  x.bool = x.Graph.Directed()
  x.me = x.Graph.Get().(uint)
  x.me = x.me
  x.actHost = env.Localhost()
  x.n = x.Graph.Num() - 1
  x.nr = make([]uint, x.n)
  x.host = make([]string, x.n)
  x.ch = make([]nchan.NetChannel, x.n)
  x.port = make([]uint16, x.n)
  for i := uint(0); i < x.n; i++ {
    g.Ex (x.me)
    x.nr[i] = g.Neighbour(i).(uint)
    x.Graph.Ex2 (x.me, x.nr[i])
    x.port[i] = p0 + x.Graph.Get1().(uint16)
  }
  x.tmpVertex = x.me
  v0 := g.Neighbour(0).(uint)
  g.Ex2 (x.me, v0)
  if g.Edged() {
    x.tmpEdge = g.Get1().(uint16)
  } else {
    g.Ex2 (v0, x.me)
    x.tmpEdge = g.Get1().(uint16)
  }
  x.tree = gra.New (true, x.tmpVertex, x.tmpEdge)
  x.ring = gra.New (true, x.tmpVertex, x.tmpEdge)
  x.visited = make([]bool, x.n)
  x.sendTo = make([]bool, x.n)
  x.chan1 = make(chan uint)
  x.parent, x.child = inf, make([]bool, x.n)
  x.mon = make([]fmon.FarMonitor, x.n)
  g.Ex (x.me)
  x.PulseAlg, x.ElectAlg, x.TravAlg = PulseGraph, ChangRoberts, DFS
  x.leader = x.me
  x.Op = Ignore
  return x
}

func newg (dir bool, e [][]uint, m, id uint) DistributedGraph {
  g := gra.New (dir, uint(0), uint16(1)) // g ist ein neuer
       // Graph mit Eckentyp uint und Kanten vom Typ von ports
  n := uint(len(e))
  for i := uint(0); i < n; i++ { // die Identitäten aller
    g.Ins(i)           // Prozesse als Ecken in g einfügen
  }
  for i := uint(0); i < n; i++ {
    for _, j := range e[i] {
      g.Ex2 (i, j)     // i ist jetzt die kolokale
                       // und j die lokale Ecke in g
      if ! g.Edged() { // wenn es noch keine Kante
                       // von i nach j gibt,
        v := nchan.Port (n, i, j, 0)
        g.Edge (v)     // i mit j durch eine Kante mit dem Wert
                       // des o.a. Ports verbinden
      }
    }
  }
  g.Ex (id)      // id ist jetzt die lokale Ecke in g
  g = g.Star()   // und g ist jetzt nur noch der Stern von id
  d := new_(g).(*distributedGraph) // d ist der verteilte
                 // Graph mit g als zugrundeliegendem Graphen
  d.setSize (n)  // Zeilen/Spaltenzahl der Adjazenzmatrix 
                 // = Anzahl der Ecken von g
  h := make([]string, n)
  for i := uint(0); i < n; i++ {
    h[i] = env.Localhost() // ggf. Namen anderer Rechner
  }
  d.setHosts (h)
  d.diameter = m // Durchmesser von g
  return d
}

func (x *distributedGraph) setHosts (h []string) {
  for i := uint(0); i < x.n; i++ {
    x.host[i] = h[x.nr[i]]
  }
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
  return g
}

func (x *distributedGraph) decodedGraph (bs Stream) gra.Graph {
  g := x.emptyGraph()
  g.Decode(bs)
  return g
}

func (x *distributedGraph) edge (v, v1 uint) uint {
  x.tmpGraph.Ex2 (v, v1)
  return x.tmpGraph.Get1().(uint)
}

func nrLocal (g gra.Graph) uint {
  return value (g.Get())
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

func (x *distributedGraph) Time() uint {
  return x.time
}

func (x *distributedGraph) Time1() uint {
  return x.time1
}

func (x *distributedGraph) ParentChildren() string {
  s := "Vater " + strconv.Itoa(int(x.parent))
  cs := make([]uint, 0)
  for i := uint(0); i < x.n; i++ {
    if x.child[i] {
      cs = append(cs, x.nr[i])
    }
  }
  n := len(cs)
  if n > 0 {
    s += ", Kind"
    if n > 1 {
      s += "er"
    }
    for _, c := range cs {
      s += " " + strconv.Itoa(int(c))
    }
  }
  return s
}
