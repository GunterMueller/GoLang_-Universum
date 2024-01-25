package phil

// (c) Christian Maurer   v. 220702 - license see nU.go

import
  "nU/mon"
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
  m = mon.New (5, f)
  return &monitorUnfair { m }
}

func (x *monitorUnfair) Lock (p uint) {
  changeStatus (p, hungry)
  x.F (p, lock)
  changeStatus (p, dining)
}

func (x *monitorUnfair) Unlock (p uint) {
  changeStatus (p, thinking)
  x.F (p, unlock)
}
