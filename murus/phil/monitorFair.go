package phil

// (c) Christian Maurer   v. 170627 - license see murus.go

// >>> Fair solution with Monitor due to Dijkstra
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 163

import (
  . "murus/obj"; . "murus/lockp"; "murus/mon"
)
type
  monitorFair struct {
                     mon.Monitor
                     }

func newMF() LockerP {
  var m mon.Monitor
  f := func (a Any, i uint) Any {
         p := a.(uint)
         if i == 0 { // lock
           if status[left(p)] == dining && status[right(p)] == satisfied ||
              status[left(p)] == satisfied && status[right(p)] == dining {
             changeStatus (p, starving)
           }
           for status[left(p)] == dining || status[left(p)] == starving ||
             status[right(p)] == dining || status[right(p)] == starving {
             m.Wait (p)
           }
         } else { // unlock
           if status[left(p)] == hungry || status[left(p)] == starving {
             m.Signal (left(p))
           }
           if status[right(p)] == hungry || status[right(p)] == starving {
             m.Signal (right(p))
           }
         }
         return nil
       }
  m = mon.New (NPhilos, f, nil)
  return &monitorFair { m }
}

func (x *monitorFair) Lock (i uint) {
  changeStatus (i, hungry)
  x.F (i, 0)
  changeStatus (i, dining)
}

func (x *monitorFair) Unlock (i uint) {
  changeStatus (i, satisfied)
  x.F (i, 1)
}
