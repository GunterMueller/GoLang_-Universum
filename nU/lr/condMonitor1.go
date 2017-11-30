package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import "nU/cmon"

type conditionedMonitor1 struct {
  cmon.Monitor
}

func newCM1() LeftRight {
  x := new(conditionedMonitor1)
  var nL, nR uint
  c := func (i uint) bool {
         switch i {
         case leftIn:
           return nR == 0
         case rightIn:
           return nL == 0
         }
         return true
       }
  f := func (i uint) uint {
         switch i {
         case leftIn:
           nL++
           return nL
         case leftOut:
           nL--
           return nL
         case rightIn:
           nR++
         case rightOut:
           nR--
         }
         return nR
       }
  x.Monitor = cmon.New (4, f, c)
  return x
}

func (x *conditionedMonitor1) LeftIn() { x.F (leftIn) }
func (x *conditionedMonitor1) LeftOut() { x.F (leftOut) }
func (x *conditionedMonitor1) RightIn() { x.F (rightIn) }
func (x *conditionedMonitor1) RightOut() { x.F (rightOut) }
