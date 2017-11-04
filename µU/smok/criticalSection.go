package smok

// (c) Christian Maurer   v. 171102 - license see µU.go

// >>> Solution with a critical section

import
  "µU/cs"
type
  criticalSection struct {
                         cs.CriticalSection
                         }


func newCS() Smokers {
  var avail [3]bool
  var smoking bool
  x := new(criticalSection)
  c := func (i uint) bool {
         if i < 3 { // Agent
           return ! smoking
         } else if i < 6 { // SmokerIn
           u1, u2 := others (i)
           return avail[u1] && avail[u2]
         }
         return true // SmokerOut
       }
  f := func (i uint) uint {
         u1, u2 := others (i)
         if i < 3 { // Agent
           avail[u1], avail[u2] = true, true
         } else if i < 6 { // SmokerIn
           smoking = true
           avail[u1], avail[u2] = false, false
         }
         return uint(0)
       }
  l := func (i uint) {
         smoking = false
       }
  x.CriticalSection = cs.New (6, c, f, l)
  return x
}

func (x *criticalSection) Agent (u uint) {
  x.Enter (u)
}

func (x *criticalSection) SmokerIn (u uint) {
  x.Enter (3 + u)
}

func (x *criticalSection) SmokerOut() {
  x.Leave (0)
}
