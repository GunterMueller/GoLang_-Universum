package phil

// (c) Christian Maurer   v. 170627 - license see µu.go

// >>> Unfair monitor solution due to Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 97

import (
  . "µu/obj"
  . "µu/lockp"
  "µu/mon"
)
type
  monitorUnfair struct {
                       mon.Monitor
                       }

func newMU() LockerP {
  var m mon.Monitor
  f := func (a Any, i uint) Any {
        p := a.(uint)
        if i == 0 { // lock
          changeStatus (p, starving)
          for status[left(p)] == dining || status[right(p)] == dining {
            m.Wait (p)
          }
        } else { // unlock
          m.Signal (left(p))
          m.Signal (right(p))
        }
        return nil
      }
  m = mon.New (NPhilos, f, nil)
  return &monitorUnfair { m }
}

func (x *monitorUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, 0)
  changeStatus (p, dining)
}

func (x *monitorUnfair) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, 1)
}
