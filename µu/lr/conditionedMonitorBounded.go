package lr

// (c) Christian Maurer   v. 170731 - license see µu.go

// >>> bounded left/right problem: implementation with monitors

import (
  . "µu/obj"
  "µu/mon"
)
type
  monitorCondBounded struct {
                     nL, nR, // number of active lefties/righties
                     tL, tR, // number of lefties/righties within one turn
                     bL, bR uint // number of blocked lefties/righties
                            mon.Monitor
                            }

func newCMB (mL, mR uint) LeftRight {
  x := new(monitorCondBounded)
  ps := func (a Any, i uint) bool {
          switch i {
          case leftIn:
            p := x.nR == 0 && (x.bR == 0 || x.tL < mL)
            if ! p { x.bL++ }
            return p
          case rightIn:
            p := x.nL == 0 && (x.bL == 0 || x.tR < mR)
            if ! p { x.bR++ }
            return p
          }
          return true
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            x.nL++
            writeL (x.nL)
            x.tL++
            if x.bL > 0 { x.bL-- }
            x.tR = 0
          case rightIn:
            x.nR++
            writeR (x.nR)
            x.tR++
            if x.bR > 0 { x.bR-- }
            x.tL = 0
          case leftOut:
            x.nL--
            writeL (x.nL)
          case rightOut:
            x.nR--
            writeR (x.nR)
          }
          return nil
        }
  x.Monitor = mon.New (nFuncs, fs, ps)
  return x
}

func (x *monitorCondBounded) LeftIn() {
  x.F (true, leftIn)
}

func (x *monitorCondBounded) LeftOut() {
  x.F (true, leftOut)
}

func (x *monitorCondBounded) RightIn() {
  x.F (true, rightIn)
}

func (x *monitorCondBounded) RightOut() {
  x.F (true, rightOut)
}

func (x *monitorCondBounded) Fin() {
}
