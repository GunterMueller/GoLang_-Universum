package sem

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> Implementation with a conditioned universal monitor

import
  "µU/cmon"
type
  conditionedMonitor struct {
                            cmon.Monitor
                            }

func newCM (n uint) Semaphore {
  val := n
  x := new(conditionedMonitor)
  c := func (i uint) bool {
         if i == p {
           return val > 0
         }
         return true
       }
  f := func (i uint) uint {
         if i == p {
           val--
         } else {
           val++
         }
         return val
       }
  x.Monitor = cmon.New (2, f, c)
  return x
}

func (x *conditionedMonitor) P() {
  x.F (p)
}

func (x *conditionedMonitor) V() {
  x.F (v)
}
