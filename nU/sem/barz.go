package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

import
  "sync"
type
  barz struct {
    cs, mutex sync.Mutex
              int "value"
              }

func newBarz (n uint) Semaphore {
  x:= new (barz)
  x.int = int(n)
  if x.int == 0 {
    x.cs.Lock()
  }
  return x
}

func (x *barz) P() {
  x.cs.Lock()
  x.mutex.Lock()
  x.int--
  if x.int > 0 {
    x.cs.Unlock()
  }
  x.mutex.Unlock()
}

func (x *barz) V() {
  x.mutex.Lock()
  x.int++
  if x.int == 1 {
    x.cs.Unlock()
  }
  x.mutex.Unlock()
}
