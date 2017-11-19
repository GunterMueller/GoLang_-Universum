package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/nat"
  "µU/errh"
  "µU/vtx"
  "µU/edg"
  "µU/host"
  "µU/nchan"
  "µU/fmon"
  "µU/gra"
  "µU/adj"
  "µU/ego"
)
type
  distributedGraph struct {
                   gra.Graph // neighbourhood of the current vertex
                             // must not be changed by any application
          tmpGraph gra.Graph // a clone of gra.Graph; these three graphs are only
                   bool "Graph.Directed"
         actVertex vtx.Vertex // the vertex representing the actual process
                me uint // and its identity
           actHost host.Host // the host running the actual process
                 n uint // the number of neighbours of the actual vertex
                nb []vtx.Vertex // the neighbour vertices
                nr []uint // and their identities
              host []host.Host // the hosts running the neighbour processes
                ch []nchan.NetChannel // the channels to the neighbours
              port []uint16 // ports to connect the vertices pairwise
        tree, ring gra.Graph // temporary graphs for the algorithms to work with
              root uint
         tmpVertex vtx.Vertex
           tmpEdge edg.Edge
          chanuint []chan uint // internal channels
             chan1 chan uint // internal channel
           visited,
            sendTo []bool
           labeled bool
               mon []fmon.FarMonitor
              monM []fmon.FarMonitorM
            parent uint
             child []bool
       time, time1 uint
      sport, cport []uint16 // server- and client ports for farMonitorM
              demo bool
               top adj.AdjacencyMatrix
          diameter,
          distance,
            leader uint
                   TopAlg; ElectAlg; TravAlg
                   Op
                   }
const
  inf = uint(1 << 16)
var (
  done = make(chan int, 1)
//  p0 = uint16(50000) // nchan.Port0()
)

func value (a Any) uint {
  return a.(vtx.Vertex).Content().(Valuator).Val()
}

func new_(g gra.Graph) DistributedGraph {
  x := new(distributedGraph)
  x.Graph = g
  x.Graph.SetWrite (vtx.W, edg.W)
  x.tmpGraph = Clone(x.Graph).(gra.Graph)
  x.bool = x.Graph.Directed()
  x.actVertex = x.Graph.Get().(vtx.Vertex)
  x.me = value(x.actVertex)
  if x.me != ego.Me() { errh.Error2 ("x.me", x.me, "Me", ego.Me()) }
  x.actHost = host.Localhost()
  x.n = x.Graph.Num() - 1
  x.nb = make([]vtx.Vertex, x.n)
  x.nr = make([]uint, x.n)
  x.host = make([]host.Host, x.n)
  x.port = make([]uint16, x.n)
  x.ch = make([]nchan.NetChannel, x.n)
  for i := uint(0); i < x.n; i++ {
    g.Ex (x.actVertex)
    x.nb[i] = g.Neighbour(i).(vtx.Vertex)
    x.nr[i] = x.nb[i].(Valuator).Val()
    x.Graph.Ex2 (x.actVertex, x.nb[i])
    v := x.Graph.Get1().(edg.Edge).(Valuator).Val()
    ps := v / (1<<20)
    pc := (v - 1<<20 * ps) / (1<<10)
//    x.port[i] = nchan.Port0 + uint16(v - 1<<10 * pc - 1<<20 * ps)
    x.port[i] = uint16(v - 1<<10 * pc - 1<<20 * ps)
  }
  x.tmpVertex = vtx.New (x.actVertex.Content(), x.actVertex.Wd(), x.actVertex.Ht())
  v0 := g.Neighbour(0).(vtx.Vertex)
  g.Ex2 (x.actVertex, v0)
  if g.Edged() {
    x.tmpEdge = g.Get1().(edg.Edge)
  } else {
    g.Ex2 (v0, x.actVertex)
    x.tmpEdge = g.Get1().(edg.Edge)
  }
  x.tree = gra.New (true, x.tmpVertex, x.tmpEdge); x.tree.SetWrite (x.Graph.Writes())
  x.ring = gra.New (true, x.tmpVertex, x.tmpEdge); x.ring.SetWrite (x.Graph.Writes())
  x.visited = make([]bool, x.n)
  x.sendTo = make([]bool, x.n)
  x.chanuint = make([]chan uint, x.n)
  x.chan1 = make(chan uint)
  x.parent, x.child = inf, make([]bool, x.n)
  x.sport, x.cport = make([]uint16, x.n), make([]uint16, x.n)
  x.mon = make([]fmon.FarMonitor, x.n)
  x.monM = make([]fmon.FarMonitorM, x.n)
  for i := uint(0); i < x.n; i++ {
    v := x.Graph.Get1().(edg.Edge).(Valuator).Val()
    ps := v / (1<<20)
    pc := (v - 1<<20 * ps) / (1<<10)
//    x.sport[i], x.cport[i] = nchan.Port0 + uint16(ps), nchan.Port0 + uint16(pc)
    x.sport[i], x.cport[i] = uint16(ps), uint16(pc)
    if x.nr[i] < x.me {
      x.sport[i], x.cport[i] = x.cport[i], x.sport[i]
    }
    x.chanuint[i] = make(chan uint)
  }
  g.Ex (x.actVertex)
  x.TopAlg, x.ElectAlg, x.TravAlg = PassGraph, ChangRoberts, DFS
  x.leader = x.me
  x.Op = Ignore
  return x
}

func (x *distributedGraph) SetHosts (h []host.Host) {
  if uint(len(h)) != x.n { ker.Shit() }
  for i := uint(0); i < x.n; i++ {
    if h[i] == nil { ker.Oops() }
    x.host[i] = h[i]
  }
}

func (x *distributedGraph) SetRoot (r uint) {
  x.root = r
}

func (x *distributedGraph) SetDiameter (d uint) {
  x.diameter = d
}

func (x *distributedGraph) SetN (n uint) {
  x.top = adj.New (n, uint(0), uint(0))
  x.top.Set (x.me, x.me, uint(x.me), uint(0))
  for i := uint(0); i < x.n; i++ {
    x.top.Set (x.me, x.nr[i], uint(0), uint(1))
    x.top.Set (x.nr[i], x.me, uint(0), uint(1))
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

func (x *distributedGraph) finMonM() {
  for i := uint(0); i < x.n; i++ {
    x.monM[i].Fin()
  }
}

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

func (x *distributedGraph) decodedGraph (bs []byte) gra.Graph {
  g := gra.New (x.tmpGraph.Directed(), x.tmpVertex, x.tmpEdge)
  g.Decode(bs)
  g.SetWrite (vtx.W, edg.W)
  return g
}

func (x *distributedGraph) directedEdge (v, v1 vtx.Vertex) edg.Edge {
  x.tmpGraph.Ex2 (v, v1)
  e := x.tmpGraph.Get1().(edg.Edge)
  e.Direct (true)
  e.SetPos0 (v.Pos()); e.SetPos1 (v1.Pos())
  e.Label (false)
  return e
}

func nrLocal (g gra.Graph) uint {
  return value (g.Get())
}

func valueMax (g gra.Graph) uint {
  m := uint(0)
  g.Trav (func (a Any) {
    v := value (a.(vtx.Vertex))
    if v > m { m = v }
  })
  return m
}

func exValue (g gra.Graph, v uint) bool {
  return g.ExPred (func (a Any) bool { return a.(vtx.Vertex).Val() == v })
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

func (x *distributedGraph) Diameter() uint {
  return x.diameter
}

func (x *distributedGraph) Demo() {
  x.demo = true
}

func (x *distributedGraph) ParentChildren() string {
  s := "parent " + nat.String(x.parent)
  cs := make([]uint, 0)
  for i := uint(0); i < x.n; i++ {
    if x.child[i] {
      cs = append(cs, x.nr[i])
    }
  }
  n := len(cs)
  if n > 0 {
    s += ", child"
    if n > 1 {
      s += "ren"
    }
    for _, c := range cs {
      s += " " + nat.String(c)
    }
  }
  return s
}
