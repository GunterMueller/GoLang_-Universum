package dgra

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/fmon")

func (x *distributedGraph) bfsfm (o Op) {
  go func() {
    fmon.New (uint(0), 1, x.b, AllTrueSp,
              x.actHost, p0 + uint16(x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 1, x.b, AllTrueSp,
                         x.host[i], p0 + uint16(x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  x.parent = inf
  if x.me == x.root {
    x.parent = x.root
    for {
      c := uint(0)
      for k := uint(0); k < x.n; k++ {
        if ! x.visited[k] {
          if x.mon[k].F(x.me + inf * x.distance, 0).(uint) == inf {
            x.visited[k] = true
          } else {
            x.child[k] = true
            c++
          }
        }
      }
      if c == 0 {
        break
      }
      x.distance++
    }
    x.Op (x.me)
  } else {
    <-done // auf root warten
  }
}

func (x *distributedGraph) b (a any, i uint) any {
  x.awaitAllMonitors()
  s := a.(uint) % inf
  j := x.channel(s)
  x.distance = a.(uint) / inf
  x.visited[j] = true
  if x.distance == 0 {
    if x.parent < inf {
      return inf
    }
    x.parent = s // == x.nr[j]
    x.Op (x.me)
    return x.me
  }
  c := uint(0)
  for k := uint(0); k < x.n; k++ {
    if k != j && ! x.visited[k] {
      if x.mon[k].F(x.me + (x.distance - 1) * inf, 0).(uint) == inf {
        x.visited[k] = true
      } else {
        x.child[k] = true
        c++
      }
    }
  }
  if c == 0 { // keine weiteren Kinder mehr erreichbar
    done <- 0
    return inf
  }
  return x.me
}
