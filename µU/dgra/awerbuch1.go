package dgra

// (c) Christian Maurer   v. 171203 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) awerbuch1 (o Op) {
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
  x.Op = o
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
// x.log ("call discover", x.nr[k])
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
        x.tree.Write()
      }
    }
    x.Op(x.actVertex)
  } else {
    <-done
  }
}

func (x *distributedGraph) a1 (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.tree = x.decodedGraph(bs)
  s := nrLocal(x.tree)
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
// x.log ("call discover", x.nr[k])
        bs = x.mon[k].F(x.tree, discover).(Stream)
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
      }
    }
  case distribute:
    x.tree.Ex(x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
// x.log ("call distribute", x.nr[k])
        x.mon[k].F(x.tree, distribute)
      }
      x.tree.Write()
    }
    x.Op(x.actVertex)
    done <- 0
  }
  return x.tree.Encode()
}
