package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  "nU/cmon"
type
  conditionedMonitor2 struct {
                             cmon.Monitor
                             }

func newCM2() LeftRight {
  x := new(conditionedMonitor2)
  var nL, nR uint
  c := func (i uint) bool {
         switch i {
         case leftIn:
           return nR == 0 && (x.Blocked (rightIn) == 0 || nL == 0)
         case rightIn:
           return nL == 0 && (x.Blocked (leftIn) == 0 || nR == 0)
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

func (x *conditionedMonitor2) LeftIn() { x.F (leftIn) }
func (x *conditionedMonitor2) LeftOut() { x.F (leftOut) }
func (x *conditionedMonitor2) RightIn() { x.F (rightIn) }
func (x *conditionedMonitor2) RightOut() { x.F (rightOut) }
