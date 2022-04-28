package sem

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> Implementation with a far monitor

import (
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (n uint, h string, port uint16, s bool) Semaphore {
  x := new(farMonitor)
  val := n
  c := func (a any, i uint) bool {
         if i == p {
           return val > 0
         }
         return true // vOp
       }
  f := func (a any, i uint) any {
         switch i {
         case p:
           val--
         case v:
           val++
         }
         return true
       }
  x.FarMonitor = fmon.New (false, 2, f, c, h, port, s)
  return x
}

func (x *farMonitor) P() {
  x.F (true, p)
}

func (x *farMonitor) V() {
  x.F (true, v)
}
