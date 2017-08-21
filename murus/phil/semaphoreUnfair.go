package phil

// (c) murus.org  v. 170627 - license see murus.go

// >>> Unfair solution with semaphores, aushungerungsgefährdet
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 95 ff.

import (
  "sync"
  . "murus/lockp"
)
type
  semaphoreUnfair struct {
                   plate []sync.Mutex
                         }

func (x *semaphoreUnfair) test (p uint) {
  if status[p] == hungry &&
     status[left(p)] != dining && status[right(p)] != dining {
    changeStatus (p, dining)
    x.plate[p].Unlock()
  }
}

func newSU() LockerP {
  x := new (semaphoreUnfair)
  x.plate = make ([]sync.Mutex, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    x.plate[p].Lock()
  }
  return x
}

func (x *semaphoreUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  mutex.Lock()
  x.test (p)
  mutex.Unlock()
  x.plate[p].Lock()
  changeStatus (p, dining)
}

func (x *semaphoreUnfair) Unlock (p uint) {
  changeStatus (p, satisfied)
  mutex.Lock()
  x.test (left (p))
  x.test (right (p))
  mutex.Unlock()
}
