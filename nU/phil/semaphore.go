package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> naive solution with semaphores - deadlock

import "nU/sem"

type semaphore struct {
  fork []sem.Semaphore
}

func newS() Philos {
  x := new(semaphore)
  x.fork = make([]sem.Semaphore, 5)
  for p := uint(0); p < 5; p++ {
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
  changeStatus (p, thinking)
  x.fork[left(p)].V()
  x.fork[p].V()
}
