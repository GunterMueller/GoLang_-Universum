package asem

// (c) Christian Maurer   v. 201117 - license see nU.go

import ("sync"; "nU/sem")

const
  M = 20
type addSemaphore struct {
  uint "number of processes allowed to use the critical section"
  sync.Mutex
  b [M]sem.Semaphore
  nB [M]uint // number of processes blocked by b
}

func new_(n uint) *addSemaphore {
  x := new(addSemaphore)
  x.uint = n
  for i := 0; i < M; i++ {
    x.b[i] = sem.New(0)
  }
  return x
}

func (x *addSemaphore) P (n uint) {
  x.Mutex.Lock()
  if x.uint >= n {
    x.uint -= n
    x.Mutex.Unlock()
  } else {
    x.nB[n]++
    x.Mutex.Unlock()
    x.b[n].P()
  }
}

func (x *addSemaphore) V (n uint) {
  x.Mutex.Lock()
  for i := uint(1); i < x.uint; i++ {
    for j := x.nB[i]; j > 1; j-- {
      if x.uint >= i {
        x.nB[i]--
        x.uint--
        x.b[i].V()
      }
    }
  }
  x.Mutex.Unlock()
}
