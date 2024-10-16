package phil

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Unfair monitor solution due to Dijkstra

import
  "µU/mon"
type
  monitorUnfair struct {
                       mon.Monitor
                       }

func newMU() Philos {
  var m mon.Monitor
  f := func (a any, i uint) any {
         p := a.(uint)
         if i == lock {
           changeStatus (p, starving)
           for status[left(p)] == dining || status[right(p)] == dining {
             m.Wait (p)
           }
         } else {
           m.Signal (left(p))
           m.Signal (right(p))
        }
        return p
      }
  m = mon.New (NPhilos, f)
  return &monitorUnfair { m }
}

func (x *monitorUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitorUnfair) Unlock (p uint) {
  changeStatus (p, satisfied)
  x.F (p, unlock)
}
