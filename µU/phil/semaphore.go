package phil

// (c) Christian Maurer   v. 171104 - license see µU.go

// >>> naive solution with semaphores - deadlock

import (
  "µU/sem"
  . "µU/lockn"
)
type
  semaphore struct {
              fork []sem.Semaphore
                   }

func newS() LockerN {
  x := new(semaphore)
  x.fork = make([]sem.Semaphore, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.fork[p] = sem.New (p)
  }
  return x
}

func (x *semaphore) Lock (p uint) {
  changeStatus (p, hungry)
  x.fork[left (p)].P()
  changeStatus (p, hasLeftFork)
  x.fork[p].P()
  changeStatus (p, dining)
}

func (x *semaphore) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.fork[left(p)].V()
  x.fork[p].V()
}
