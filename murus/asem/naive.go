package asem

// (c) murus.org  v. 170410 - license see murus.go

// >>> incorrect naive representation
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 100

import
  "murus/sem"
type
  naive struct {
               uint32 "number of processes allowed to use the critical section, that shall be protected by the semaphore"
               sem.Semaphore
               }


func newNaive (n uint) AddSemaphore {
  return &naive { uint32(n), sem.New (n) }
}

func (x *naive) P (n uint) {
  for n > 0 {
    x.Semaphore.P()
    n --
  }
}

func (x *naive) V (n uint) {
  for n > 0 {
    x.Semaphore.V()
    n --
  }
}
