package lr

// (c) Christian Maurer   v. 171017 - license see µU.go

// >>> bounded left/right problem

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
type
  farMonitorBounded struct {
                           fmon.FarMonitor
                           }

func newFMonB (mL, mR uint, h host.Host, port uint16, s bool) LeftRight {
  if mL * mR == 0 { mL, mR = inf, inf } // unbounded case
  var nL, nR uint
  var tL, tR uint // number of actives within one turn
  x := new(farMonitorBounded)
  p := func (a Any, i uint) bool {
         switch i {
         case leftIn:
           return nR == 0 && tL < mL
         case rightIn:
           return nL == 0 && tR < mR
         }
         return true // leftOut, rightOut
       }
  f := func (a Any, i uint) Any {
         switch i {
         case leftIn:
           nL++
           tL++
           tR = 0
         case leftOut:
           nL--
         case rightIn:
           nR++
           tR++
           tL = 0
         case rightOut:
           nR--
         }
         return true
       }
  x.FarMonitor = fmon.New (true, nFuncs, f, p, h, port, s)
  return x
}

func (x *farMonitorBounded) LeftIn() {
  x.F(true, leftIn)
}

func (x *farMonitorBounded) LeftOut() {
  x.F(true, leftOut)
}

func (x *farMonitorBounded) RightIn() {
  x.F(true, rightIn)
}

func (x *farMonitorBounded) RightOut() {
  x.F(true, rightOut)
}

func (x *farMonitorBounded) Fin() {
  x.FarMonitor.Fin()
}
