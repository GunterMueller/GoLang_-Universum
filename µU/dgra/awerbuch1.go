package dgra

// (c) Christian Maurer   v. 231215 - license see µU.go

import (
  . "µU/obj"
//  "µU/vtx"
  "µU/fmon"
//  "µU/errh"
)

func (x *distributedGraph) a1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.tree = x.decodedGraph(bs)
  x.tree.Write()

  s := nrLocal(x.tree) // s := x.tree.Get().(vtx.Vertex).Val()
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    x.tree.Ins (x.actVertex) // x.nb[j] colocal, x.actVertex local
    x.tree.Edge (x.edge(x.nb[j], x.actVertex))
    x.tree.Write()

    x.parent = x.nr[j]
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.tree.Ex (x.actVertex)
// x.log ("call visit", x.nr[k])
        x.mon[k].F(x.tree, visit)
      }
      x.tree.Write()
    }
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.tree.Ex (x.actVertex)
// x.log ("call discover1", x.nr[k])
        bs = x.mon[k].F(x.tree, discover).(Stream)
        x.tree = x.decodedGraph(bs)
        x.tree.Write()

      }
    }
//  return x.tree.Encode()
  case distribute:
    x.tree.Ex (x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
// x.log ("call distribute", x.nr[k])
        x.mon[k].F(x.tree, distribute)
      }
//    x.tree.Write()
    }
    x.tree.Write()
    done <- 0
  }
  return nil // XXX
  return x.tree.Encode() // nil // XXX
}

func (x *distributedGraph) Awerbuch1() {
  go func() {
    fmon.New (nil, 3, x.a1, AllTrueSp,
              x.actHost, p0 + uint16(3 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 3, x.a1, AllTrueSp,
                         x.host[i], p0 + uint16(3 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tree.Clr()
  if x.me == x.root {
    x.parent = x.root
    x.tree.Ins (x.actVertex)
    x.tree.Mark (x.actVertex)
    x.tree.Write()

    for k := uint(0); k < x.n; k++ {
// x.log ("call visit", x.nr[k])
      x.mon[k].F(x.tree, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.tree.Ex (x.actVertex)
// x.log ("call discover1", x.nr[k])
        bs := x.mon[k].F(x.tree, discover).(Stream)
        x.tree = x.decodedGraph(bs)
        x.tree.Write()

      }
    }
    x.tree.Ex (x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
// x.log ("call distribute", x.nr[k])
        x.mon[k].F(x.tree, distribute)
//      x.tree.Write()
      }
    }
    x.tree.Write()

    x.Op (x.me)
  } else {
    <-done
  }
}
