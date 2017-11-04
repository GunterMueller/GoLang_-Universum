package smok

// (c) Christian Maurer   v. 171102 - license see µU.go

// >>> Solution with a conditioned monitor

import
  "µU/cmon"
type
  condMonitor struct {
                     cmon.Monitor
                     }

func newCM() Smokers {
  var avail [3]bool
  x := new (condMonitor)
  var smoking bool
  p := func (i uint) bool {
         if i < 3 { // Agent
           return ! smoking
         }
         if i == 6 { // SmokerOut
           return true
         }
         u1, u2 := others (i - 3) // SmokerIn
         return avail[u1] && avail[u2]
       }
  f := func (i uint) uint {
         if i < 3 { // Agent
           u1, u2 := others (i)
           avail[u1], avail[u2] = true, true
         } else if i == 6 { // SmokerOut
           smoking = false
         } else { // SmokerIn
           smoking = true
           u1, u2 := others (i - 3)
           avail[u1], avail[u2] = false, false
         }
         return 0
       }
  x.Monitor = cmon.New (7, f, p)
  return x
}

func (x *condMonitor) Agent (u uint) {
  x.F (u)
}

func (x *condMonitor) SmokerIn (u uint) {
  x.F (3 + u)
}

func (x *condMonitor) SmokerOut() {
  x.F (uint(6))
}
