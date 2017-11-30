package dgra

// (c) Christian Maurer   v. 171123 - license see µU.go

import (
  "sync"
  "µU/ker"
)
var
  lock = make(chan int, 1)
var
  mutex sync.Mutex

func (x *distributedGraph) awaitAllMonitors() {
  for k := uint(0); k < x.n; k++ {
    for x.mon[k] == nil {
      ker.Msleep (100)
    }
  }
}
