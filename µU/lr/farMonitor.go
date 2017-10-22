package lr

// (c) Christian Maurer   v. 171017 - license see µU.go

// >>> 1st left/right problem

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFMon (h host.Host, port uint16, s bool) LeftRight {
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
         return true
       }
  x.FarMonitor = fmon.New (false, nFuncs, f, p, h, port, s)
  return x
}

func (x *farMonitor) LeftIn() {
  x.F(true, leftIn)
}

func (x *farMonitor) LeftOut() {
  x.F(true, leftOut)
}

func (x *farMonitor) RightIn() {
  x.F(true, rightIn)
}

func (x *farMonitor) RightOut() {
  x.F(true, rightOut)
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}
