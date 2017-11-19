package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import (
//  "µU/ker"
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) fmdfs1 (o Op) {
//  go func() { fmon.New (nil, 2, x.d1, AllTrueSp, x.actHost, p0 + uint16(2 * x.me), true) }()
  go func() { fmon.New (nil, 2, x.d1, AllTrueSp, x.actHost, uint16(2 * x.me), true) }()
//ker.Sleep (3)
  for i := uint(0); i < x.n; i++ {
//println (i, " of ", x.n, "(")
//    x.mon[i] = fmon.New (nil, 2, x.d1, AllTrueSp, x.host[i], p0 + uint16(2 * x.nr[i]), false)
    x.mon[i] = fmon.New (nil, 2, x.d1, AllTrueSp, x.host[i], uint16(2 * x.nr[i]), false)
//println (")")
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.tree.Clr()
  x.Op = o
  x.tree.Ins (x.actVertex)
  x.tree.Sub (x.actVertex)
  x.tree.Write()
  if x.me == x.root {
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.child[k] = true
      x.tree.Ex(x.actVertex) // actVertex local in x.tree
      bs := x.mon[k].F(x.tree, 0).([]byte)
      if len(bs) == 0 {
        x.visited[k] = true
      } else {
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
      }
    }
    x.tree.Ex(x.actVertex) // actVertex local in x.tree
    x.tree.Write()
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.mon[k].F(x.tree, 1)
      }
    }
    x.Op(x.actVertex)
  } else {
    <-done // wait for root's result
  }
  x.tree.Write()
}

func (x *distributedGraph) d1 (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.([]byte)
  x.tree = x.decodedGraph(bs)
  switch i {
  case 0:
    j := x.channel(nrLocal(x.tree))
    if x.tree.Ex (x.actVertex) {
      return nil
    }
    x.parent = x.nr[j]
    x.tree.Ins (x.actVertex)
    x.tree.Edge (x.directedEdge(x.nb[j], x.actVertex))
    x.tree.Write()
    for k := uint(0); k < x.n; k++ {
      if k != j {
        if ! x.tree.Ex (x.nb[k]) {
          x.tree.Ex (x.actVertex)
          bs = x.mon[k].F(x.tree, 0).([]byte)
          if len(bs) == 0 {
            return nil
          } else {
            x.child[k] = true
            x.tree = x.decodedGraph (bs)
            x.tree.Write()
          }
        }
      }
    }
  case 1:
    x.tree.Ex(x.actVertex)
    x.tree.Write()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, 1)
      }
    }
    x.Op (x.actVertex)
    done <- 0
  }
  return x.tree.Encode()
}
