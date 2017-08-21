package phil

// (c) murus.org  v. 170629 - license see murus.go

// >>> Solution with far monitor

import (
  . "murus/obj"
  . "murus/lockp"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h host.Host, port uint16, s bool) LockerP {
  nForks := make ([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  p := func (a Any, i uint) bool {
         if i == 0 { // lock
           p := a.(uint) // p-th philosopher
           return nForks[p] == 2
         }
         return true // unlock
       }
  f := func (a Any, i uint) Any {
        p := a.(uint) // p-th philosopher
        if i == 0 { // lock
          nForks[left(p)] --
          nForks[right(p)] --
        } else { // unlock
          nForks[left(p)] ++
          nForks[right(p)] ++
        }
        return p
      }
  return &farMonitor { fmon.New (nil, NPhilos, f, p, h, port, s) }
}

func (x *farMonitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, 0)
  changeStatus (p, dining)
}

func (x *farMonitor) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, 1)
}
