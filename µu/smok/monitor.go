package smok

// (c) Christian Maurer   v. 170629 - license see µu.go

// >>> Solution with a monitor

import (
  "µu/ker"
  . "µu/obj"
  "µu/mon"
  . "µu/smok/utensil"
)
type
  monitor struct {
           avail [NUtensils]bool
       smokerOut bool
                 mon.Monitor
                 }

func newM() Smokers {
  x := new (monitor)
  f := func (a Any, k uint) Any {
         if k < NUtensils {
           if a == nil { // AgentOut
             ker.Panic ("oops")
             x.Wait (NUtensils)
           } else { // AgentIn
             for ! x.smokerOut {
               x.Wait (NUtensils)
             }
             u1, u2 := Others(k)
             x.avail[u1], x.avail[u2] = true, true
             x.Signal (k)
           }
         } else { // k == NUtensils
           if a == nil { // SmokerOut
             x.smokerOut = true
             x.Signal (NUtensils)
           } else { // SmokerIn
             u1, u2 := Others(a.(uint))
             for ! x.avail[u1] || ! x.avail[u2] {
               x.Wait (a.(uint))
             }
             x.smokerOut = false
             x.avail[u1], x.avail[u2] = false, false
           }
         }
         return a
       }
  x.Monitor = mon.New (NUtensils + 1, f, nil)
  return x
}

func (x *monitor) Agent (u uint) {
  x.F (0, u)
}

func (x *monitor) SmokerIn (u uint) {
  x.F (u, NUtensils)
}

func (x *monitor) SmokerOut() {
  x.F (nil, NUtensils)
}
