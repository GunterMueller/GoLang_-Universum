package lr

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> bounded left/right problem

import
  "µU/fmon"
type
  farMonitorBounded struct {
                           fmon.FarMonitor
                           }

func newFMonB (mL, mR uint, h string, port uint16, s bool) LeftRight {
  if mL * mR == 0 { mL, mR = inf, inf } // unbounded case
  var nL, nR uint
  var tL, tR uint // number of actives within one turn
  x := new(farMonitorBounded)
  p := func (a any, i uint) bool {
         switch i {
         case leftIn:
           return nR == 0 && tL < mL
         case rightIn:
           return nL == 0 && tR < mR
         }
         return true // leftOut, rightOut
       }
  f := func (a any, i uint) any {
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
         return 0
       }
  x.FarMonitor = fmon.New (0, 4, f, p, h, port, s)
  return x
}

func (x *farMonitorBounded) LeftIn() {
  x.F(0, leftIn)
}

func (x *farMonitorBounded) LeftOut() {
  x.F(0, leftOut)
}

func (x *farMonitorBounded) RightIn() {
  x.F(0, rightIn)
}

func (x *farMonitorBounded) RightOut() {
  x.F(0, rightOut)
}

func (x *farMonitorBounded) Fin() {
  x.FarMonitor.Fin()
}
