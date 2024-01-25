package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

import (
  "sync"
  "nU/sem"
)
type
  bounded struct {
                 sem.Semaphore "takeSeat"
            fork []sync.Mutex
                 }

func newB() Philos {
  x := new(bounded)
  x.Semaphore = sem.New (5 - 1)
  x.fork = make([]sync.Mutex, 5)
  return x
}

func (x *bounded) Lock (p uint) {
  x.Semaphore.P()
  changeStatus (p, hungry)
  x.fork[left (p)].Lock()
  changeStatus (p, hasLeftFork)
  x.fork[p].Lock()
  changeStatus (p, dining)
}

func (x *bounded) Unlock (p uint) {
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
  changeStatus (p, thinking)
  x.Semaphore.V()
}
