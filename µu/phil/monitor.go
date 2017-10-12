package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> monitor solution

import (
  . "µu/obj"
  . "µu/lockp"
  "µu/mon"
)
type
  monitor struct {
                 mon.Monitor
                 }

func newM() LockerP {
  var m mon.Monitor
  nForks := make([]uint, NPhilos)
  for p := uint(0); p < NPhilos; p++ {
    nForks[p] = 2
  }
  f := func (a Any, i uint) Any {
        p := a.(uint)
        if i == 0 { // lock
          changeStatus (p, starving)
          for nForks[p] < 2 {
            m.Wait (p)
          }
          nForks[left(p)] --
          nForks[right(p)] --
        } else { // i == unlock
          nForks[left(p)] ++
          nForks[right(p)] ++
          if nForks[left(p)] == 2 {
            m.Signal(left(p))
          }
          if nForks[right(p)] == 2 {
            m.Signal(right(p))
          }
        }
        return nil
      }
  m = mon.New (NPhilos, f, nil)
  return &monitor { m }
}

func (x *monitor) Lock (i uint) {
  changeStatus (i, hungry)
  x.F (i, 0)
  changeStatus (i, dining)
}

func (x *monitor) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.F (i, 1)
}
