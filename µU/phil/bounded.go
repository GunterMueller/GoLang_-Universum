package phil

// (c) Christian Maurer   v. 171127 - license see µU.go

// >>> Bounded case:
//     At most m - 1 philosophers are allowed to take place at the table
//     at the same time, where m is the number of participating philsophers.

import (
  "sync"
  "µU/sem"
)
type
  bounded struct {
                 sem.Semaphore "takeSeat"
            fork []sync.Mutex
                 }

func newB() Philos {
  x := new(bounded)
  x.Semaphore = sem.New (NPhilos - 1)
  x.fork = make([]sync.Mutex, NPhilos)
  return x
}

func (x *bounded) Lock (p uint) {
  x.Semaphore.P()
  changeStatus (p, hungry)
  x.fork[left (p)].Lock()
//  changeStatus (p, hasRightFork)
  changeStatus (p, hasLeftFork)
  x.fork[p].Lock()
  changeStatus (p, dining)
}

func (x *bounded) Unlock (p uint) {
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
  changeStatus (p, satisfied)
  x.Semaphore.V()
}
