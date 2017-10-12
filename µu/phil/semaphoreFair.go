package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Fair Algorithm of Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 98

import (
  "sync"
  . "µu/lockp"
)
type
  semaphoreFair struct {
                 plate []sync.Mutex
                       }

func (x *semaphoreFair) test (p uint) {
  if status[p] == hungry &&
     (status[left(p)] == dining && status[right(p)] == satisfied ||
      status[left(p)] == satisfied && status[right(p)] == dining) {
//    changeStatusUnlocked (p, starving)
    changeStatus (p, starving)
  }
  if (status[p] == hungry || status[p] == starving) &&
   ! (status[left(p)] == dining || status[left(p)] == starving) &&
   ! (status[right(p)] == dining || status[right(p)] == starving) {
//    changeStatusUnlocked (p, dining)
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}

func newSF() LockerP {
  x := new (semaphoreFair)
  x.plate = make ([]sync.Mutex, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}

func (x *semaphoreFair) Lock (p uint) {
  changeStatus (p, hungry)
  mutex.Lock()
  x.test (p)
  mutex.Unlock()
  x.plate[p].Lock()
  changeStatus (p, dining)
}

func (x *semaphoreFair) Unlock (p uint) {
  changeStatus (p, satisfied)
  mutex.Lock()
  x.test (left(p))
  x.test (right(p))
  mutex.Unlock()
}
