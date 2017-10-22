package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st left/right problem

import
  "µU/cmon"
type
  cmonitor1 struct {
                   cmon.Monitor
                   }

func newcM1() LeftRight {
  var nL, nR uint
  x := new(cmonitor1)
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
  x.Monitor = cmon.New (nFuncs, f, c)
  return x
}

func (x *cmonitor1) LeftIn() {
  x.F (leftIn)
}

func (x *cmonitor1) LeftOut() {
  x.F (leftOut)
}

func (x *cmonitor1) RightIn() {
  x.F (rightIn)
}

func (x *cmonitor1) RightOut() {
  x.F (rightOut)
}

func (x *cmonitor1) Fin() {
}
