package smok

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Solution with a universal monitor

import
  "µU/mon"
type
  monitor struct {
                 mon.Monitor
                 }

func newM() Smokers {
  var avail [3]bool
  smoking := false
  x := new(monitor)
  f := func (a any, i uint) any {
         u := a.(uint)
         u1, u2 := others (u)
         switch i {
         case agent:
           if smoking {
             x.Wait (3)
           }
           avail[u1], avail[u2] = true, true
           x.Signal (u)
         case smokerOut:
           smoking = false
           x.Signal (3)
           return uint(3)
         case smokerIn:
           if ! (avail[u1] && avail[u2]) {
             x.Wait (u)
           }
           smoking = true
           avail[u1], avail[u2] = false, false
         }
         return u
       }
  x.Monitor = mon.New (4, f)
  return x
}

func (x *monitor) Agent (u uint) {
  x.F (u, agent)
}

func (x *monitor) SmokerIn (u uint) {
  x.F (u, smokerIn)
}

func (x *monitor) SmokerOut() {
  x.F (uint(3), smokerOut)
}
