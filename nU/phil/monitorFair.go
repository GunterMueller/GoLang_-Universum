package phil

// (c) Christian Maurer   v. 220702 - license see nU.go

// >>> Fair solution with a monitor due to Dijkstra

import
  "nU/mon"
type
  monitorFair struct {
                     mon.Monitor
                     }

func newMF() Philos {
  var m mon.Monitor
  f := func (a any, i uint) any {
         p := a.(uint)
         if i == lock {
           if status[left(p)] == dining && status[right(p)] == thinking ||
              status[left(p)] == thinking && status[right(p)] == dining {
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
  m = mon.New (5, f)
  return &monitorFair { m }
}

func (x *monitorFair) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitorFair) Unlock (p uint) {
  changeStatus (p, thinking)
  x.F (p, unlock)
}
