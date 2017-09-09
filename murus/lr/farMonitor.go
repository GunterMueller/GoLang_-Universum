package lr

// (c) Christian Maurer   v. 170520 - license see murus.go

// >>> left/right problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 209

import (
  . "murus/obj"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor struct {
             nL, nR uint
                    fmon.FarMonitor
                    }

func newFMon (h host.Host, p uint16, s bool) LeftRight {
  x := new(farMonitor)
  ps := func (a Any, i uint) bool {
          switch i {
          case leftIn:
            return x.nR == 0
          case rightIn:
            return x.nL == 0
          }
          return true // leftOut, rightOut
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            x.nL++
            writeL (x.nL)
          case rightIn:
            x.nR++
            writeR (x.nR)
          case leftOut:
            x.nL--
            writeL (x.nL)
          case rightOut:
            x.nR--
            writeR (x.nR)
          }
          return true
        }
  x.FarMonitor = fmon.New (false, nFuncs, fs, ps, h, p, s)
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
