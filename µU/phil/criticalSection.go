package phil

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> Solution with critical sections

import (
  . "µU/lockp"
  . "µU/cs"
)
type
  criticalSection struct {
                         CriticalSection
                         }

func newCS() LockerP {
  nForks := make([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  c := func (i uint) bool {
         return nForks[i] == 2
       }
  f := func (i uint) uint {
         nForks[left(i)]--
         nForks[right(i)]--
         return uint(0)
       }
  l := func (i uint) {
         nForks[left(i)]++
         nForks[right(i)]++
       }
  return &criticalSection { New (NPhilos, c, f, l) }
}

func (x *criticalSection) Lock (i uint) {
  changeStatus (i, hungry)
  x.Enter (i)
  changeStatus (i, dining)
}

func (x *criticalSection) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.Leave (i)
}
