package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import "nU/cmon"

type condMonitorBounded struct {
  cmon.Monitor
}

func newCMB (mL, mR uint) LeftRight {
  x := new(condMonitorBounded)
  var nL, nR,
      tL, tR uint // number of lefties/righties within one turn
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
           return nL
         case rightIn:
           nR++
           tR++
           tL = 0
           return nL
         case leftOut:
           nL--
         case rightOut:
           nR--
         }
         return nR
       }
  x.Monitor = cmon.New (4, f, c)
  return x
}

func (x *condMonitorBounded) LeftIn() { x.F (leftIn) }
func (x *condMonitorBounded) LeftOut() { x.F (leftOut) }
func (x *condMonitorBounded) RightIn() { x.F (rightIn) }
func (x *condMonitorBounded) RightOut() { x.F (rightOut) }
