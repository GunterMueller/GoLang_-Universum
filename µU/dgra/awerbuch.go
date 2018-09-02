package dgra

// (c) Christian Maurer   v. 171130 - license see µU.go
//
// >>> simplified version of the algorithm of B. Awerbuch:
//     A New Distributed Depth-First-Search Algorithm, Inf, Proc. Letters 28 (1985) 147-160 

import (
  . "µU/obj"
  "µU/fmon"
)
const (
  visit = uint(iota)
  discover
  distribute
)

func (x *distributedGraph) awerbuch (o Op) {
  go func() {
    fmon.New (uint(0), 2, x.a, AllTrueSp,
              x.actHost, p0 + uint16(2 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 2, x.a, AllTrueSp,
                         x.host[i], p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  if x.me == x.root {
    x.parent = x.root
    for k := uint(0); k < x.n; k++ {
x.log ("call visit", x.nr[k])
      x.mon[k].F(x.me, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
// x.log ("child", x.nr[k])
x.log ("call discover", x.nr[k])
        x.mon[k].F(x.me, discover)
      }
    }
    x.Op(x.actVertex)
  } else {
    <-done
  }
}

func (x *distributedGraph) a (a Any, i uint) Any {
  x.awaitAllMonitors()
  s := a.(uint)
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    x.parent = x.nr[j]
// x.log ("parent", x.nr[j])
    for k := uint(0); k < x.n; k++ {
      if k != j {
x.log ("call visit", x.nr[k])
        x.mon[k].F(x.me, visit)
      }
    }
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
// x.log ("child", x.nr[k])
x.log ("call discover", x.nr[k])
        x.mon[k].F(x.me, discover)
      }
    }
    x.Op(x.actVertex)
    done <- 0
  }
  return x.me
}
