package phil

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> Solution with conditioned monitor

import (
  . "µU/lockn"
  "µU/cmon"
)
type
  conditionedMonitor struct {
                            cmon.Monitor
                            }

func newCM() LockerN {
  nForks := make([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  c := func (i uint) bool {
         if i < NPhilos { // Lock
           return nForks[i] == 2
         }
         return true // Unlock
       }
  f := func (i uint) uint {
         if i < NPhilos {
           nForks[left(i)]--
           nForks[right(i)]--
           return i
         }
         i -= NPhilos
         nForks[left(i)]++
         nForks[right(i)]++
         return i
       }
  return &conditionedMonitor { cmon.New (NPhilos, f, c) }
}

func (x *conditionedMonitor) Lock (i uint) {
  changeStatus (i, hungry)
  x.F (lock + i)
  changeStatus (i, dining)
}

func (x *conditionedMonitor) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.F (NPhilos + i)
}
