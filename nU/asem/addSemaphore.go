package asem

// (c) Christian Maurer   v. 171227 - license see nU.go

import (. "sync"; "nU/sem")

const
  M = 20
type addSemaphore struct {
  int "number of processes allowed to use the critical section, that shall be protected by the semaphore"
  me Mutex
  b [M]sem.Semaphore
  nB [M]int
}

func new_(n uint) *addSemaphore {
  x := new(addSemaphore)
  x.int = int(n)
  for i := 0; i < M; i++ {
    x.b[i] = sem.New (0)
  }
  return x
}

func (x *addSemaphore) P (n uint) {
  x.me.Lock()
  if x.int >= int(n) {
    x.int -= int(n)
    x.me.Unlock()
  } else {
    x.nB[n]++
    x.me.Unlock()
    x.b[n].P()
  }
}

func (x *addSemaphore) V (n uint) {
  x.me.Lock()
  x.int += int(n)
  i:= x.int
  for i > 0 {
    for x.nB[i] > 0 && i < x.int {
      x.int -= i // x.int--
      x.nB[i]--
      x.b[i].V()
    }
    i--
  }
  x.me.Unlock()
}
