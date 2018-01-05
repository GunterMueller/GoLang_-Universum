package barb

// (c) Christian Maurer   v. 170731 - license see nU.go

import "sync"

type barberSem struct {
  waiting, mutex sync.Mutex
  n, k uint
}

func newS() Barber {
  x := new (barberSem)
  x.waiting.Lock()
  return x
}

func (x *barberSem) Customer() {
  x.mutex.Lock()
  x.n++
  if x.n== 1 {
    x.waiting.Unlock()
  }
  x.mutex.Unlock()
}

func (x *barberSem) Barber() {
  if x.k == 0 {
    x.waiting.Lock()
  }
  x.mutex.Lock()
  x.n--
  x.k = x.n
  x.mutex.Unlock()
}
