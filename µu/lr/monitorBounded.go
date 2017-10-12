package lr

// (c) Christian Maurer   v. 170731 - license see µu.go
//
// >>> bounded left/right problem: implementation with monitors

import (
  . "µu/obj"
  "µu/mon"
)
type
  monitorBounded struct {
                 nL, nR, // number of active lefties/righties
                 tL, tR uint // number of lefties/righties within one turn
                        mon.Monitor
                        }

func newMB (mL, mR uint) LeftRight {
  x := new(monitorBounded)
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            if x.nR > 0 || x.Awaited (right) && x.tL >= mL { x.Monitor.Wait (leftIn) }
//            if x.nR > 0 || x.tL >= mL { x.Monitor.Wait (leftIn) }
            x.nL++
            writeL (x.nL)
            x.tL++
            x.tR = 0
            if x.tL < mL || ! x.Awaited (right) { x.Monitor.Signal (leftIn) }
//            if x.tL < mL { x.Monitor.Signal (leftIn) }
          case rightIn:
            if x.nL > 0 || x.Awaited (left) && x.tR >= mR { x.Monitor.Wait (rightIn) }
//            if x.nL > 0 || x.tR >= mR { x.Monitor.Wait (rightIn) }
            x.nR++
            writeR (x.nR)
            x.tR++
            x.tL = 0
            if x.tR < mR || ! x.Awaited(left) { x.Monitor.Signal (rightIn) }
//            if x.tR < mR { x.Monitor.Signal (rightIn) }
          case leftOut:
            x.nL--
            writeL (x.nL)
            if x.nL == 0 {
              x.Monitor.Signal (rightIn)
            } else {
              if x.tL < mL && ! x.Awaited (right) { x.Monitor.Signal (leftIn) }
//              if x.tL < mL { x.Monitor.Signal (leftIn) }
            }
          case rightOut:
            x.nR--
            writeR (x.nR)
            if x.nR == 0 {
              x.Monitor.Signal (leftIn)
            } else {
              if x.tR < mR && ! x.Awaited (left) { x.Monitor.Signal (rightIn) }
//              if x.tR < mR { x.Monitor.Signal (rightIn) }
            }
          }
          return nil
        }
  x.Monitor = mon.New (nFuncs, fs, nil)
  return x
}

func (x *monitorBounded) LeftIn() {
  x.F (true, leftIn)
}

func (x *monitorBounded) LeftOut() {
  x.F (true, leftOut)
}

func (x *monitorBounded) RightIn() {
  x.F (true, rightIn)
}

func (x *monitorBounded) RightOut() {
  x.F (true, rightOut)
}

func (x *monitorBounded) Fin() {
}
