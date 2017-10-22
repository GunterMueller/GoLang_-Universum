package dgra

// (c) Christian Maurer   v. 170510 - license see µU.go
//
// >>> simplified version of the algorithm of B. Awerbuch:
//     A New Distributed Depth-First-Search Algorithm, Inf, Proc. Letters 28 (1985) 147-160 

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) fmdfsa1 (o Op) {
  go func() { fmon.New (nil, 3, x.da1, AllTrueSp, x.actHost, p0 + uint16(3 * x.me), true) }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 3, x.da1, AllTrueSp, x.host[i], p0 + uint16(3 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.tree.Clr()
  x.Op = o
  if x.me == x.root {
    x.parent = x.me
    x.tree.Ins(x.actVertex)
    x.tree.SubLocal()
    x.tree.Write()
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F(x.tree, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.tree.Ex(x.actVertex)
        bs := x.mon[k].F(x.tree, discover).([]byte)
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
      }
    }
    x.tree.Ex(x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, 2)
      }
    }
    x.Op(x.actVertex)
  } else {
    <-done
  }
}

func (x *distributedGraph) da1 (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.([]byte)
  x.tree = x.decodedGraph(bs)
  s := nrLocal(x.tree)
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    x.tree.Ins (x.actVertex) // x.nb[j] colocal, x.actVertex local
    x.tree.Edge (x.directedEdge (x.nb[j], x.actVertex))
    x.tree.Write()
    x.parent = x.nr[j]
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.tree.Ex (x.actVertex)
        x.mon[k].F(x.tree, visit)
      }
    }
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.tree.Ex (x.actVertex)
        bs = x.mon[k].F(x.tree, discover).([]byte)
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
      }
    }
  case 2:
    x.tree.Ex(x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, 2)
      }
    }
    x.tree.Write()
    x.Op(x.actVertex)
    done <- 0
  }
  return x.tree.Encode()
}
