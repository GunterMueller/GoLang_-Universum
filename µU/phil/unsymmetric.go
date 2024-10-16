package phil

// (c) Christian Maurer   v. 170627 - license see µU.go

// >>> Unsymmetric case:
//     Symmetry is disturbed by the fact that some
//     (but not all) philosophers pick up the left fork first.

import
  "sync"
type
  unsymmetric struct {
                fork []sync.Mutex
                     }

func newU() Philos {
  x := new(unsymmetric)
  x.fork = make([]sync.Mutex, NPhilos)
  return x
}

func (x *unsymmetric) Lock (p uint) {
  changeStatus (p, hungry)
  if p % 2 == 1 {
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
