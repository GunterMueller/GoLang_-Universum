package sem

// (c) Christian Maurer   v. 170121 - license see murus.go

// >>> Implementation with a far Monitor
//     Nichtsequentielle Programmierung mit Go 1 kompakt, S. 208

import (
  . "murus/obj"
  "murus/host"
  "murus/fmon"
)
const (
  pOp = uint(iota)
  vOp
  nFuncs
)
type
  farMonitor struct {
                    uint "value"
                    fmon.FarMonitor
                    }

func newFMon (n uint, h host.Host, p uint16, s bool) Semaphore {
  x := new(farMonitor)
  x.uint = n
  c := func (a Any, i uint) bool {
         if i == pOp {
           return x.uint > 0
         }
         return true // vOp
       }
  f := func (a Any, i uint) Any {
         switch i {
         case pOp:
           x.uint--
         case vOp:
           x.uint++
         }
         return true
       }
  x.FarMonitor = fmon.New (false, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) P() {
  x.F(true, pOp)
}

func (x *farMonitor) V() {
  x.F(true, vOp)
}
