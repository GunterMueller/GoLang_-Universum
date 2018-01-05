package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Fair Algorithm of Dijkstra

import "sync"

type semaphoreFair struct {
  plate []sync.Mutex
}

func (x *semaphoreFair) test (p uint) {
  if status[p] == hungry &&
     (status[left(p)] == dining && status[right(p)] == thinking ||
      status[left(p)] == thinking && status[right(p)] == dining) {
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

func newSF() Philos {
  x := new(semaphoreFair)
  x.plate = make([]sync.Mutex, 5)
  for p := uint(0); p < 5; p++ {
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
  changeStatus (p, thinking)
  m.Lock()
  x.test (left(p))
  x.test (right(p))
  m.Unlock()
}
