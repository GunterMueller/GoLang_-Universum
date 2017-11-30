package sem

// (c) Christian Maurer   v. 170121 - license see nU.go

import "sync"

type correct struct {
  int "value"
  binsem, mutex sync.Mutex
}

func newCorrect (n uint) Semaphore {
  x:= new (correct)
  x.int = int(n)
  x.binsem.Lock()
  return x
}

func (x *correct) P() {
  x.mutex.Lock()
  x.int--
  if x.int < 0 {
    x.mutex.Unlock()
    x.binsem.Lock()
  }
  x.mutex.Unlock()
}

func (x *correct) V() {
  x.mutex.Lock()
  x.int++
  if x.int <= 0 {
    x.binsem.Unlock()
  } else {
    x.mutex.Unlock()
  }
}
