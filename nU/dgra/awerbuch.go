package dgra

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/fmon")

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
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F(x.me, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.mon[k].F(x.me, discover)
      }
    }
    x.Op(x.me)
  } else {
    <-done
  }
}

func (x *distributedGraph) a (a any, i uint) any {
  x.awaitAllMonitors()
  s := a.(uint)
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    x.parent = x.nr[j]
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.mon[k].F(x.me, visit)
      }
    }
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.mon[k].F(x.me, discover)
      }
    }
    x.Op(x.me)
    done <- 0
  }
  return x.me
}
