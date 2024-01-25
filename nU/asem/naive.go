package asem

// (c) Christian Maurer   v. 170410 - license see nU.go

import
  "nU/sem"
type
  naive struct {
               uint32 "Anzahl der Prozesse, die den k.A. betreten dürfen, der durch das Semaphor geschützt werden soll"
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
