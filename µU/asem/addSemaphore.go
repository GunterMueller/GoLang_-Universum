package asem

// (c) Christian Maurer   v. 200421 - license see µU.go

import (
  "sync"
  "µU/sem"
)
const
  M = 20
type
  addSemaphore struct {
                      uint "value of the semaphore"
                      sync.Mutex
                    s [M]sem.Semaphore
                    n [M]uint // n[i] = number of processes blocked on s[i] (i < M)
                      }

func new_(n uint) *addSemaphore {
  x := new(addSemaphore)
  x.uint = n
  for i := 0; i < M; i++ {
    x.s[i] = sem.New (0)
  }
  return x
}

func (x *addSemaphore) P (n uint) {
  x.Mutex.Lock()
  if x.uint >= n {
    x.uint -= n
    x.Mutex.Unlock()
  } else {
    x.n[n] += 1
    x.Mutex.Unlock()
    x.s[n].P()
  }
}

func (x *addSemaphore) V (n uint) {
  x.Mutex.Lock()
  x.uint += n
  for i := x.uint; i > 0; i-- {
    for x.n[i] > 0 && i <= x.uint {
      x.uint -= i
      x.n[i]--
      x.s[i].V()
    }
  }
  x.Mutex.Unlock()
}
