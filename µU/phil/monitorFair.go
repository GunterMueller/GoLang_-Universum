package phil

// (c) Christian Maurer   v. 171127 - license see µU.go

// >>> Fair solution with a monitor due to Dijkstra

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitorFair struct {
                     mon.Monitor
                     }

func newMF() Philos {
  var m mon.Monitor
  f := func (a Any, i uint) Any {
         p := a.(uint)
         if i == lock {
           if status[left(p)] == dining && status[right(p)] == satisfied ||
              status[left(p)] == satisfied && status[right(p)] == dining {
             changeStatus (p, starving)
           }
           for status[left(p)] == dining || status[left(p)] == starving ||
               status[right(p)] == dining || status[right(p)] == starving {
             m.Wait (p)
           }
         } else {
           if status[left(p)] == hungry || status[left(p)] == starving {
             m.Signal (left(p))
           }
           if status[right(p)] == hungry || status[right(p)] == starving {
             m.Signal (right(p))
           }
         }
         return p
       }
  m = mon.New (NPhilos, f)
  return &monitorFair { m }
}

func (x *monitorFair) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitorFair) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
