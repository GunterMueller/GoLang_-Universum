package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Solution with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 98

import (
  . "µu/obj"
  . "µu/lockp"
  . "µu/cs"
)
type
  criticalSection struct {
                         CriticalSection
                         }

func newCS() LockerP {
  nForks := make ([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  c := func (i uint) bool {
         return nForks[i] == 2
       }
  l := func (a Any, i uint) {
         nForks[left(i)] --
         nForks[right(i)] --
       }
  u := func (a Any, i uint) {
         nForks[left(i)] ++
         nForks[right(i)] ++
       }
  return &criticalSection { New (NPhilos, c, l, u) }
}

func (x *criticalSection) Lock (i uint) {
  changeStatus (i, hungry)
  x.Enter (i, nil)
  changeStatus (i, dining)
}

func (x *criticalSection) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.Leave (i, nil)
}
