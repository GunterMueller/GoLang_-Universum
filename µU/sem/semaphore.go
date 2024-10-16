package sem

// (c) Christian Maurer   v. 240930 - license see µU.go

// >>> corrected naive solution

import
  "sync"
type
  semaphore struct {
               int "value"
             block,
             mutex sync.Mutex
                   }

func new_(n uint) Semaphore {
  x := new(semaphore)
  x.int = int(n)
  x.block.Lock()
  return x
}

func (x *semaphore) P() {
  x.mutex.Lock()
  x.int--
  if x.int < 0 {
    x.mutex.Unlock()
    x.block.Lock()
  }
  x.mutex.Unlock()
}

func (x *semaphore) V() {
  x.mutex.Lock()
  x.int++
  if x.int <= 0 {
    x.block.Unlock()
  } else {
    x.mutex.Unlock()
  }
}
