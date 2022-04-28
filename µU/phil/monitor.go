package phil

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Implementation with a universal monitor

import (
  "µU/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM() Philos {
  var m mon.Monitor
  nForks := make([]uint, NPhilos)
  for i := uint(0); i < NPhilos; i++ {
    nForks[i] = 2
  }
  f := func (a any, i uint) any {
         p := a.(uint)
         if i == lock {
           changeStatus (p, starving)
           for nForks[p] < 2 {
             m.Wait (p)
           }
           nForks[left(p)]--
           nForks[right(p)]--
         } else {
           nForks[left(p)]++
           nForks[right(p)]++
           if nForks[left(p)] == 2 {
             m.Signal(left(p))
           }
           if nForks[right(p)] == 2 {
             m.Signal(right(p))
           }
         }
         return p
       }
  m = mon.New (NPhilos, f)
  return &monitor { m }
}

func (x *monitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitor) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
