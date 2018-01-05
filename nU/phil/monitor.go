package phil

// (c) Christian Maurer   v. 171229 - license see nU.go

// >>> Implementation with a universal monitor

import (. "nU/obj"; "nU/mon")

type monitor struct {
  mon.Monitor
}

func newM() Philos {
  var m mon.Monitor
  nForks := make([]uint, 5)
  for i := uint(0); i < 5; i++ {
    nForks[i] = 2
  }
  f := func (a Any, i uint) Any {
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
  m = mon.New (5, f)
  return &monitor { m }
}

func (x *monitor) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitor) Unlock (p uint) {
  changeStatus (p, thinking)
  x.F (p, unlock)
}
