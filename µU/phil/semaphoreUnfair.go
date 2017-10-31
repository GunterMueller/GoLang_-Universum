package phil

// (c) Christian Maurer   v. 170627 - license see µU.go

// >>> Unfair solution with semaphores, danger of starvation

import (
  "sync"
  . "µU/lockn"
)
type
  semaphoreUnfair struct {
                   plate []sync.Mutex
                         }
var
  m sync.Mutex

func (x *semaphoreUnfair) test (p uint) {
  if status[p] == hungry &&
     status[left(p)] != dining && status[right(p)] != dining {
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}

func newSU() LockerN {
  x := new(semaphoreUnfair)
  x.plate = make([]sync.Mutex, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}

func (x *semaphoreUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  m.Lock()
  x.test (p)
  m.Unlock()
  x.plate[p].Lock()
  changeStatus (p, dining)
}

func (x *semaphoreUnfair) Unlock (p uint) {
  changeStatus (p, satisfied)
  m.Lock()
  x.test (left (p))
  x.test (right (p))
  m.Unlock()
}
