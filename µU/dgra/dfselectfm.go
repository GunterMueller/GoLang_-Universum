package dgra

// (c) Christian Maurer   v. 220420 - license see µU.go

// Find the leader by depth-first-search:
// Compare the own identity with the received value and return the appropriate value.

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) dfselectfm() {
  go func() {
    fmon.New (uint(0), 3, x.s, AllTrueSp,
              x.actHost, p0 + uint16(3 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (uint(0), 3, x.s, AllTrueSp,
                         x.host[i], p0 + uint16(3 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  if x.me == x.root {
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.mon[k].F (x.me + inf * x.leader, 0)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        v := x.mon[k].F (x.me + inf * x.leader, 1).(uint)
        if v > x.leader {
          x.leader = v
        }
      }
    }
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.me + inf * x.leader, 2)
      }
    }
  } else {
    <-done
  }
}

func (x *distributedGraph) s (a any, i uint) any {
  x.awaitAllMonitors()
  s, v := a.(uint) % inf, a.(uint) / inf
  j := x.channel(s)
  switch i {
  case 0:
    x.visited[j] = true
  case 1:
    x.parent = s
    if v > x.me {
      x.leader = v
    }
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.mon[k].F(x.me + inf * x.leader, 0)
      }
    }
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        v = x.mon[k].F(x.me + inf * x.leader, 1).(uint)
        if v > x.leader {
          x.leader = v
        }
      }
    }
    return x.leader
  case 2:
    x.leader = v
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.me + inf * x.leader, 2)
      }
    }
    done <- 0
  }
  return x.me
}
