package phil

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Solution with far monitor

import (
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h string, port uint16, s bool) Philos {
  nForks := make([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  p := func (a any, i uint) bool {
         if i == lock {
           j := a.(uint) // j-th philosopher
           return nForks[j] == 2
         }
         return true // unlock
       }
  f := func (a any, i uint) any {
         j := a.(uint) // j-th philosopher
         if i == lock {
           nForks[left(j)]--
           nForks[right(j)]--
         } else { // unlock
           nForks[left(j)]++
           nForks[right(j)]++
         }
         return j
       }
  return &farMonitor { fmon.New (uint(0), NPhilos, f, p, h, port, s) }
}

func (x *farMonitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *farMonitor) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
