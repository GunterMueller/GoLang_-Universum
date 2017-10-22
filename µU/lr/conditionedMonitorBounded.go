package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> bounded left/right problem

import
  "µU/cmon"
type
  condMonitorBounded struct {
                            cmon.Monitor
                            }

func newCMB (mL, mR uint) LeftRight {
  var nL, nR uint
  var tL, tR uint // number of actives within one turn
  x := new(condMonitorBounded)
  c := func (i uint) bool {
         switch i {
         case leftIn:
           return nR == 0 && (x.Blocked (rightIn) == 0 || tL < mL)
         case rightIn:
           return nL == 0 && (x.Blocked (leftIn) == 0 || tR < mR)
         }
         return true
       }
  f := func (i uint) uint {
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
         return uint(0)
       }
  x.Monitor = cmon.New (nFuncs, f, c)
  return x
}

func (x *condMonitorBounded) LeftIn() {
  x.F (leftIn)
}

func (x *condMonitorBounded) LeftOut() {
  x.F (leftOut)
}

func (x *condMonitorBounded) RightIn() {
  x.F (rightIn)
}

func (x *condMonitorBounded) RightOut() {
  x.F (rightOut)
}

func (x *condMonitorBounded) Fin() {
}
