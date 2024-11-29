package dgra

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
  "µU/vtx"
  "µU/fmon"
)

func (x *distributedGraph) a1 (a any, i uint) any {
  x.awaitAllMonitors()
  x.tree = x.decodedGraph(a.(Stream))
  x.tree.Write()
  pause()
  s := x.tree.Get().(vtx.Vertex).Val()
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    x.tree.Ins (x.actVertex) // x.nb[j] colocal, x.actVertex local
    x.tree.Edge (x.edge(x.nb[j], x.actVertex))
    x.tree.Write()
    pause()
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
        x.tree = x.decodedGraph(x.mon[k].F(x.tree, discover).(Stream))
        x.tree.Write()
        pause()
      }
    }
    return x.tree.Encode()
  case distribute:
    x.tree.Ex (x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, distribute)
      }
    }
    x.tree.Write()
    done <- 0
  }
  return nil
}

func (x *distributedGraph) Awerbuch1() {
  scr.Cls()
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
    x.tree.Mark (x.actVertex, true)
    x.tree.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F(x.tree, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.tree.Ex (x.actVertex)
        s := x.mon[k].F(x.tree, discover).(Stream)
        x.tree = x.decodedGraph(s)
        x.tree.Write()
        pause()
      }
    }
    x.tree.Ex (x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, distribute)
      }
    }
    x.tree.Write()
    pause()
  } else {
    <-done
  }
}
