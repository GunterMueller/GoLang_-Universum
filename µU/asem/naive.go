package asem

// (c) Christian Maurer   v. 190823 - license see µU.go

// >>> incorrect naive representation

import
  "µU/sem"
type
  naive struct {
               uint "number of processes allowed to use the critical section, that shall be protected by the semaphore"
               sem.Semaphore
               }


func newNaive (n uint) AddSemaphore {
  return &naive { n, sem.New (n) }
}

func (x *naive) P (n uint) {
  for i := uint(0); i < n; i++ {
    x.Semaphore.P()
  }
}

func (x *naive) V (n uint) {
  for i := uint(0); i < n; i++ {
    x.Semaphore.V()
  }
}
