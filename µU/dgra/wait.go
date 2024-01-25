package dgra

// (c) Christian Maurer   v. 171226 - license see µU.go

import (
  "sync"
  "time"
)
var
  lock = make(chan int, 1)
var
  mutex sync.Mutex

func (x *distributedGraph) awaitAllMonitors() {
  for k := uint(0); k < x.n; k++ {
    for x.mon[k] == nil {
      time.Sleep (1e8)
    }
  }
}

func pause() {
  time.Sleep (1e9)
  return
}
