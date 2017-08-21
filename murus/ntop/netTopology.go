package ntop

// (c) murus.org  v. 170103 - license see murus.go

import (
//  "murus/ker"
  . "murus/obj"
  "murus/scr"; "murus/errh"
  "murus/nat"
  "murus/bnat"
  "murus/host"
  "murus/node"
  "murus/nchan"
  "murus/fmon"
  "murus/adj"; "murus/gra"
)

type
  netTopology struct {
                     uint "number of hosts involved"
                   n uint // identities of the neighbours of the current process
               graph gra.Graph // graph of all hosts involved and their neighbourhood relations
//              tree gra.Graph
                  me, // identity of the current process
                root uint
              meNode, // the node in the graph representing the current process
             tmpNode node.Node
//               mon []Monitor
                  nb []node.Node // the neighbour nodes
                  nn []bnat.Natural // their contents
//                   host.Host
                   h []host.Host // the neighbour hosts
                  nr []uint // and their identities (nh[i] must have identity i)
                port,
               port1 []uint16
                 top adj.AdjacencyMatrix //
            diameter uint // of the graph
                     bool "demo mode on"
                     Algorithm // used to construct the topology
                     }

func new_(g gra.Graph, h []host.Host, n, d uint) NetTopology {
  x := new(netTopology)
  x.uint = n
  x.graph = g.Clone().(gra.Graph)
  x.diameter = d
  x.n = x.graph.NumLoc()
  x.meNode = x.graph.Get().(node.Node)
  x.me = x.meNode.Content().(bnat.Natural).Val()
//  x.root = root
  x.nb = make([]node.Node, x.n)
  x.nn = make([]bnat.Natural, x.n)
//  x.Host = host.Localhost()
  x.h = make([]host.Host, x.uint)
  x.nr = make([]uint, x.n)
  x.port = make ([]uint16, x.n)


  for i := uint(0); i < x.n; i++ {
    x.nb[i] = x.graph.Neighbour(i).(node.Node)
    x.nn[i] = x.nb[i].Content().(bnat.Natural)
    x.nr[i] = x.nn[i].Val()
    x.h[i]  = h[x.nr[i]]
    x.port[i] = nchan.Port (x.uint, x.me, x.nr[i])
  }
  x.port1 = make ([]uint16, x.uint)
  const p0 = 50000 // nchan.Port0
  for i := uint(0); i < x.uint; i++ {
    x.port1[i] = p0 + uint16(i)
  }
  x.top = adj.New (x.uint, uint(1))
  for i := uint(0); i < x.n; i++ {
    x.top.Set (x.me, x.nr[i], uint(1))
    x.top.Set (x.nr[i], x.me, uint(1))
  }
  return x
}

func (x *netTopology) Demo () {
  x.bool = true
}

func (x *netTopology) enter (r uint) {
  if x.bool {
    errh.Error ("start round", r)
  }
}

func (x *netTopology) add (a Any, i uint) Any {
  x.top.Add (a.(adj.AdjacencyMatrix))
  return x.top
}

func (x *netTopology) Do(alg Algorithm) {
  x.Algorithm = alg
  switch x.Algorithm {
  case Matrix:
    if x.bool { x.Write() }
    ch := make ([]nchan.NetChannel, x.n)
    for i := uint(0); i < x.n; i++ {
      ch[i] = nchan.New (x.top, x.me, x.nn[i].Val(), x.h[i], x.port[i])
    }
    for r:= uint(0); r < x.diameter; r++ {
      x.enter (r + 1)
      for i := uint(0); i < x.n; i++ {
        ch[i].Send (x.top)
      }
      for i := uint(0); i < x.n; i++ {
        x.top.Add (ch[i].Recv().(adj.AdjacencyMatrix))
      }
      if x.bool { x.top.Write (0, scr.NColumns() / 2) }
    }
    for i := uint(0); i < x.n; i++ { ch[i].Fin() }
  case Graph:
    if x.bool { x.Write() }
    ch := make ([]nchan.NetChannel, x.n)
    for i := uint(0); i < x.n; i++ {
      ch[i] = nchan.New (nil, x.me, x.nn[i].Val(), x.h[i], x.port[i])
    }
    for r:= uint(0); r < x.diameter; r++ {
      x.enter (r + 1)
      for i := uint(0); i < x.n; i++ {
        ch[i].Send (x.graph)
      }
      for i := uint(0); i < x.n; i++ {
        nh := bnat.New (nat.Wd(x.uint))
        nh.SetVal (i)
        g := Decode (gra.New (false, node.New(nh, 1, 1), nil), ch[i].Recv().([]byte)).(gra.Graph)
        x.graph.Add (g)
      }
      if x.bool { x.Write() }
    }
    for i := uint(0); i < x.n; i++ { ch[i].Fin() }
  case MatrixFarMon:
    if x.bool { x.Write() }
    go func() { // start server
      fmon.New (x.top, 1, x.add, AllTrueSp, x.h[x.me], x.port1[x.me], true)
    }()
    mon:= make ([]fmon.FarMonitor, x.n)
    for i := uint(0); i < x.n; i++ { // start clients
      mon[i] = fmon.New (x.top, 1, NilSp, AllTrueSp, x.h[i], x.port1[x.nr[i]], false) // XXX
    }
    for r:= uint(0); r < 1 * x.diameter; r ++ {
      x.enter (r + 1)
      for i := uint(0); i < x.n; i++ {
        top := mon[i].F (x.top, 0).(adj.AdjacencyMatrix)
        x.top.Add (top)
      }
      if x.bool { x.top.Write (0, scr.NColumns() / 2) }
    }
    for i := uint(0); i < x.n; i++ { mon[i].Fin() }
  }
}

func (x *netTopology) Write() {
  if x.Algorithm == Graph {
    x.graph.Trav3Cond (node.O, node.O3)
  } else {
    x.top.Write (0, 0)
  }
}

func (x *netTopology) Topology() gra.Graph {
  return x.graph
}
