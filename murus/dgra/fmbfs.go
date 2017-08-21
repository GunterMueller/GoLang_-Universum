package dgra

// (c) murus.org  v. 170510 - license see murus.go

import (
  . "murus/obj"
  "murus/nchan"
  "murus/fmon"
)

func (x *distributedGraph) fmbfs (o Op) {
  go func() { fmon.New (uint(0), 2, x.b, AllTrueSp, x.actHost, nchan.Port0 + uint16(2 * x.me), true) }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 2, x.b, AllTrueSp, x.host[i], nchan.Port0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  x.parent = inf
  if x.me == x.root {
    x.parent = x.me
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
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.me, 1)
      }
    }
    x.Op (x.actVertex)
  } else {
    <-done // wait until root finished
  }
}

func (x *distributedGraph) b (a Any, i uint) Any {
  x.awaitAllMonitors()
  s := a.(uint) % inf
  j := x.channel(s)
  x.distance = a.(uint) / inf
  switch i {
  case 0:
    x.visited[j] = true
    if x.distance == 0 {
      if x.parent < inf {
        return inf
      }
      x.parent = s // == x.nr[j]
      x.Op (x.actVertex)
      return x.me
    }
    c := uint(0) // r > 0
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
    if c == 0 {
      return inf
    }
  case 1:
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.me, 1)
      }
    }
    done <- 0
    return inf
  }
  return x.me
}
