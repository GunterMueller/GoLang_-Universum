package smok

// (c) murus.org  v. 170629 - license see murus.go

// >>> Solution with a critical section

import (
  . "murus/obj"
  "murus/cs"
  . "murus/smok/utensil"
)
type
  criticalSection struct {
                   avail [NUtensils]bool
               smokerOut bool
                         cs.CriticalSection
                         }


func newCS() Smokers {
  x := new(criticalSection)
  c := func (k uint) bool {
         if k == NUtensils { // Agent
           return x.smokerOut
         }
         u1, u2 := Others(k)
         return x.avail[u1] && x.avail[u2] // SmokerIn
       }
  e := func (a Any, k uint) {
         if k == NUtensils { // Agent
           u1, u2 := Others(a.(uint))
           x.avail[u1], x.avail[u2] = true, true
         } else { // SmokerIn
           x.smokerOut = false
           u1, u2 := Others(k)
           x.avail[u1], x.avail[u2] = false, false
         }
       }
  l := func (a Any, k uint) {
         x.smokerOut = true // SmokerOut
       }
  x.CriticalSection = cs.New (NUtensils + 1, c, e, l)
  return x
}

func (x *criticalSection) Agent (u uint) {
  x.Enter (NUtensils, u)
}

func (x *criticalSection) SmokerIn (u uint) {
  x.Enter (u, 0)
}

func (x *criticalSection) SmokerOut() {
  x.Leave (0, nil)
}
