package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Bounded case:
//     At most m - 1 philosophers are allowed to take place at the table
//     at the same time, where m is the number of participating philsophers.
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 96

import (
  "sync"
  . "µu/lockp"
  "µu/sem"
)
type
  bounded struct {
                 sem.Semaphore "takeSeat"
            fork []sync.Mutex
                 }

func newB() LockerP {
  x := new (bounded)
  x.Semaphore = sem.New (NPhilos - 1)
  x.fork = make ([]sync.Mutex, NPhilos)
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
