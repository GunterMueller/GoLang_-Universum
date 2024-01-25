package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

import
  "sync"
type
  naive struct {
               int "value"
  block, mutex sync.Mutex
}

func newNaive (n uint) Semaphore {
  x:= new(naive)
  x.int = int(n)
  x.block.Lock()
  return x
}

func (x *naive) P() {
  x.mutex.Lock()
  x.int--
  if x.int < 0 {
    x.mutex.Unlock()
    x.block.Lock()
  } else {
    x.mutex.Unlock()
  }
}

func (x *naive) V() {
  x.mutex.Lock()
  x.int++
  if x.int <= 0 {
    x.block.Unlock()
  }
  x.mutex.Unlock()
}
