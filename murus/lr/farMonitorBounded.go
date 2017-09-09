package lr

// (c) Christian Maurer   v. 170411 - license see murus.go

// >>> left/right problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 209

import (
  . "murus/obj"
  "murus/ker"
  "murus/host"
  "murus/fmon"
)
type
  farMonitorBounded struct {
                    nL, nR, // number of active lefties/righties
                    tL, tR uint // number of lefties/righties within one turn
                           fmon.FarMonitor
                           }

func newFMonB (mL, mR uint, h host.Host, p uint16, s bool) LeftRight {
  if mL * mR == 0 {
    mL, mR = ker.MaxNat(), ker.MaxNat() // unbounded case
  }
  x := new(farMonitorBounded)
  ps := func (a Any, i uint) bool {
          switch i {
          case leftIn:
            return x.nR == 0 && x.tL < mL
          case rightIn:
            return x.nL == 0 && x.tR < mR
          }
          return true // leftOut, rightOut
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            x.nL++
            writeL (x.nL)
            x.tL++
            x.tR = 0
          case rightIn:
            x.nR++
            writeR (x.nR)
            x.tR++
            x.tL = 0
          case leftOut:
            x.nL--
            writeL (x.nL)
          case rightOut:
            x.nR--
            writeR (x.nR)
          }
          return true
        }
  x.FarMonitor = fmon.New (true, nFuncs, fs, ps, h, p, s)
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
