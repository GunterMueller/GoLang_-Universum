package asem

// (c) Christian Maurer   v. 200421 - license see nU.go

import (. "sync"; "nU/sem")

const
  M = 20
type addSemaphore struct {
  int "number of processes allowed to use the critical section"
  Mutex
  sem [M]sem.Semaphore
  n [M]int
}

func new_(n uint) *addSemaphore {
  x := new(addSemaphore)
  x.int = int(n)
  for i := 0; i < M; i++ {
    x.sem[i] = sem.New(0)
  }
  return x
}

func (x *addSemaphore) P (n uint) {
  x.Mutex.Lock()
  if x.int >= int(n) {
    x.int -= int(n)
    x.Mutex.Unlock()
  } else {
    x.n[n]++
    x.Mutex.Unlock()
    x.sem[n].P()
  }
}

func (x *addSemaphore) V (n uint) {
  x.Mutex.Lock()
  x.int += int(n)
  i := x.int
  for i > 0 {
    for x.n[i] > 0 && i < x.int {
      x.int -= i
      x.n[i]--
      x.sem[i].V()
    }
    i--
  }
  x.Mutex.Unlock()
}
