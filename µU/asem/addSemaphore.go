package asem

// (c) Christian Maurer   v. 190823 - license see µU.go

import (
  . "sync"
  "µU/sem"
)
const
  M = 20
type
  addSemaphore struct {
                 uint "number of processes allowed to use the critical section, that shall be protected by the semaphore"
                   me Mutex
                    b [M]sem.Semaphore
                   nB [M]uint
                      }

func new_(n uint) *addSemaphore {
  x := new(addSemaphore)
  x.uint = n
  for i := 0; i < M; i++ {
    x.b[i] = sem.New (0)
  }
  return x
}

func (x *addSemaphore) P (n uint) {
  x.me.Lock()
  if x.uint >= n {
    x.uint -= n
    x.me.Unlock()
  } else {
    x.nB[n]++
    x.me.Unlock()
    x.b[n].P()
  }
}

func (x *addSemaphore) V (n uint) {
  x.me.Lock()
  x.uint += n
  i := x.uint
  for i > 0 {
    for x.nB[i] > 0 && i < x.uint {
      x.uint -= i // x.uint--
      x.nB[i]--
      x.b[i].V()
    }
    i--
  }
  x.me.Unlock()
}
