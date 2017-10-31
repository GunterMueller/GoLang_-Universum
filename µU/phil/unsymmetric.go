package phil

// (c) Christian Maurer   v. 170627 - license see µU.go

// >>> Unsymmetric case:
//     Störung der Symmetrie dadurch, dass manche (aber nicht alle)
//     Philosophen zuerst die linke Gabel aufnehmen

import (
  "sync"
  . "µU/lockn"
)
type
  unsymmetric struct {
                fork []sync.Mutex
                     }

func newU() LockerN {
  x := new(unsymmetric)
  x.fork = make([]sync.Mutex, NPhilos)
  return x
}

func (x *unsymmetric) Lock (p uint) {
  changeStatus (p, hungry)
  if p % 2 == 1 {
//  if p == 0 {
    x.fork [left (p)].Lock()
    changeStatus (p, hasLeftFork)
    x.fork [p].Lock()
  } else {
    x.fork [p].Lock()
    changeStatus (p, hasRightFork)
    x.fork [left (p)].Lock()
  }
  changeStatus (p, dining)
}

func (x *unsymmetric) Unlock (p uint) {
  x.fork[p].Unlock()
  x.fork[left (p)].Unlock()
  changeStatus (p, satisfied)
}
