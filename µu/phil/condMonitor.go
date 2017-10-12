package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Solution with conditioned monitor
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 164

import (
  . "µu/obj"
  . "µu/lockp"
  "µu/mon"
)
type
  condMonitor struct {
                     mon.Monitor
                     }

func newCM() LockerP {
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
  return &condMonitor { mon.New (NPhilos, f, p) }
}

func (x *condMonitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, 0)
  changeStatus (p, dining)
}

func (x *condMonitor) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, 1)
}
