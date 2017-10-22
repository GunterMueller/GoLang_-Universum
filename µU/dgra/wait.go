package dgra

// (c) Christian Maurer   v. 170506 - license see µU.go

import (
  "sync"
  "µU/ker"
)

func (x *distributedGraph) awaitAllMonitors() {
  for k := uint(0); k < x.n; k++ {
    for x.mon[k] == nil {
      ker.Msleep (100)
    }
  }
}

func (x *distributedGraph) awaitAllMonitorsM() {
  for k := uint(0); k < x.n; k++ {
    for x.monM[k] == nil {
      ker.Msleep (100)
    }
  }
}

var
  lock = make(chan int, 1)
var
  mutex sync.Mutex
