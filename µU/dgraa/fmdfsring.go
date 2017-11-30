package dgra

// (c) Christian Maurer   v. 170507 - license see µU.go
//
// >>> Construction of a directed ring using the idea pf Awerbuch's DFS-algorithm
//     The sequence of the vertices in the ring is given by the discover-times x.time.

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) fmdfsring() {
//  go func() { fmon.New (uint(0), 2, x.dr, AllTrueSp, x.actHost, p0 + uint16(2 * x.me), true) }()
  go func() { fmon.New (uint(0), 2, x.dr, AllTrueSp, x.actHost, uint16(2 * x.me), true) }()
  for i := uint(0); i < x.n; i++ {
//    x.mon[i] = fmon.New (uint(0), 2, x.dr, AllTrueSp, x.host[i], p0 + uint16(2 * x.nr[i]), false)
    x.mon[i] = fmon.New (uint(0), 2, x.dr, AllTrueSp, x.host[i], uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  if x.me == x.root {
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F(x.me, visit)
    }
    x.time = 0
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.mon[k].F(x.me + x.time * inf, discover)
      }
    }
  } else {
    <-done
  }
}

func (x *distributedGraph) dr (a Any, i uint) Any {
  x.awaitAllMonitors()
  s := a.(uint) % inf
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.mon[k].F(x.me, visit)
      }
    }
    t := a.(uint) / inf
    t++
    x.time = t
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        t = x.mon[k].F(x.me + t * inf, discover).(uint) / inf
      }
    }
    done <- 0
    return x.me + inf * t
  }
  return 0
}
