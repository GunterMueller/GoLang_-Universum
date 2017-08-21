package ntrav

// (c) murus.org  v. 170106 - license see murus.go

import (
//  "murus/ker"
  . "murus/obj"
  "murus/errh"
  "murus/nat"
  "murus/bnat"
  "murus/host"
  "murus/node"
  "murus/nchan"
  "murus/gra"
)

type
  netTraversal struct {
                      uint "number of hosts involved"
                    n uint // identities of the neighbours of the current process
                graph, // graph of all hosts involved and their neighbourhood relations
                 tree gra.Graph
                   me, // identity of the current process
                 root uint
               meNode, // the node in the graph representing the current process
              tmpNode node.Node
                      host.Host // localhost
                  mon []Monitor
                   nb []node.Node // the neighbour nodes
                   nn []bnat.Natural // their contents
                    h []host.Host // the neighbour hosts
                   nr []uint // and their identities (nh[i] must have the identity i)
                 port,
                port1 []uint16


                      bool "demo mode on"

                      }

func new_(g gra.Graph, h[]host.Host, n, root uint) NetTraversal {
  x := new(netTraversal)
  x.uint = n
  x.graph = g.Clone().(gra.Graph)
  x.n = x.graph.NumLoc()
  x.meNode = x.graph.Get().(node.Node) // the center of the star is the local vertex in x.graph !
  x.me = x.meNode.Content().(bnat.Natural).Val()
  x.root = root
  x.nb = make([]node.Node, x.n)
  x.nn = make([]bnat.Natural, x.n)
  x.Host = host.Localhost()
  x.h = make([]host.Host, x.n)
  x.nr = make([]uint, x.n)
  x.port = make([]uint16, x.n)
  x.mon = make([]Monitor, x.n)
  for i := uint(0); i < x.n; i++ {
    x.nb[i] = x.graph.Neighbour(i).(node.Node)
    x.nn[i] = x.nb[i].Content().(bnat.Natural)
    x.nr[i] = x.nn[i].Val()
    x.h[i] = h[x.nr[i]]
    x.port[i] = nchan.Port (x.uint, x.me, x.nr[i])
  }
  x.port1 = make([]uint16, x.uint)
  const p0 = 50000; // nchan.Port0
  for i := uint(0); i < x.uint; i++ {
    x.port1[i] = p0 + uint16(i)
  }
  x.tmpNode = node.New (bnat.New(nat.Wd(x.uint)), nat.Wd(x.uint), 1)
  x.tree = gra.New (true, x.tmpNode, nil)
  return x
}

func (x *netTraversal) Demo() {
  x.bool = true
}

func (x *netTraversal) Trav (alg Algorithm, op Op) {
  var done chan int
  x.graph.Locate (true)
  switch alg {
  case StartEndTimes:
    x.Write()
    t0, t1 := x.dfs (op)
    errh.Error2 ("Anfangs-/Endzeit", t0, "/", t1)
  case SpanningTree:
    x.spt (op)
  case Moni:
    go func() { NewMonitor(x, x.me, true) }()
//    ker.Sleep(10)
    for i := uint(0); i < x.n; i++ {
      x.mon[i] = NewMonitor (x, x.nr[i], false)
    }
    if x.me == x.root {
      r := x.mon[0].Probe(x.root)
      errh.Error("probe result", r)
    } else {
      <-done
    }
  }
}

func (x *netTraversal) Graph() gra.Graph {
  return x.graph
}

func (x *netTraversal) Write() {
  x.graph.Trav3Cond (node.O, node.O3)
}
