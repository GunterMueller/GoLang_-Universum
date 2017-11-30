package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/fmon")

type farMonitor struct {
  fmon.FarMonitor
}

func newFMon (h string, port uint16, s bool) LeftRight {
  var nL, nR uint
  x := new(farMonitor)
  p := func (a Any, i uint) bool {
         switch i {
         case leftIn:
           return nR == 0
         case rightIn:
           return nL == 0
         }
         return true // leftOut, rightOut
       }
  f := func (a Any, i uint) Any {
         switch i {
         case leftIn:
           nL++
         case rightIn:
           nR++
         case leftOut:
           nL--
         case rightOut:
           nR--
         }
         return 0
       }
  x.FarMonitor = fmon.New (0, 4, f, p, h, port, s)
  return x
}

func (x *farMonitor) LeftIn() {
  x.F (0, leftIn)
}

func (x *farMonitor) LeftOut() {
  x.F (0, leftOut)
}

func (x *farMonitor) RightIn() {
  x.F (0, rightIn)
}

func (x *farMonitor) RightOut() {
  x.F (0, rightOut)
}
