package dgra

// (c) Christian Maurer   v. 171123 - license see nU.go

import ("time"; "sync")

var (
  lock = make(chan int, 1)
  mutex sync.Mutex
)

func (x *distributedGraph) awaitAllMonitors() {
  for k := uint(0); k < x.n; k++ {
    for x.mon[k] == nil {
      time.Sleep (100 * 1e6)
    }
  }
}

func pause() {
  time.Sleep (1e9)
  return
}
