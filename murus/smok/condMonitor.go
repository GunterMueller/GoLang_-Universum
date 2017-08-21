package smok

// (c) murus.org  v. 170627 - license see murus.go

// >>> Solution with a conditioned monitor

import (
  . "murus/obj"
  "murus/mon"
  . "murus/smok/utensil"
)
type
  condMonitor struct {
           smokerOut bool
               avail [NUtensils]bool
                     mon.Monitor
                     }

func newCM() Smokers {
  x := new (condMonitor)
  x.smokerOut = true
  p := func (a Any, k uint) bool {
         if k == NUtensils { // AgentIn
           return x.smokerOut
         }
         if k < NUtensils { // SmokerIn
           u1, u2 := Others(k)
           return x.avail[u1] && x.avail[u2]
         }
         return true // SmokerOut
       }
  f := func (a Any, k uint) Any {
         if k == NUtensils { // AgentIn
           u := a.(uint)
           u1, u2 := Others(u)
           x.avail[u1], x.avail[u2] = true, true
         } else if k < NUtensils { // SmokerIn
           x.smokerOut = false
           u1, u2 := Others(k)
           x.avail[u1], x.avail[u2] = false, false
         } else { // SmokerOut
           x.smokerOut = true
         }
         return a
       }
  x.Monitor = mon.New (NUtensils + 2, f, p)
  return x
}

func (x *condMonitor) Agent (u uint) {
  x.F (u, NUtensils)
}

func (x *condMonitor) SmokerIn (u uint) {
  x.F (nil, u)
}

func (x *condMonitor) SmokerOut() {
  x.F (nil, NUtensils + 1)
}
