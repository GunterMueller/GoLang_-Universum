package phil

// (c) Christian Maurer   v. 170627 - license see µU.go

// >>> Fair Algorithm of Dijkstra

import (
  "sync"
  . "µU/lockp"
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
  x := new(semaphoreFair)
  x.plate = make([]sync.Mutex, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}

func (x *semaphoreFair) Lock (p uint) {
  changeStatus (p, hungry)
  m.Lock()
  x.test (p)
  m.Unlock()
  x.plate[p].Lock()
  changeStatus (p, dining)
}

func (x *semaphoreFair) Unlock (p uint) {
  changeStatus (p, satisfied)
  m.Lock()
  x.test (left(p))
  x.test (right(p))
  m.Unlock()
}
