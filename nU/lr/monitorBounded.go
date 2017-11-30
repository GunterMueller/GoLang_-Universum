package lr

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/mon")

type monitorBounded struct {
  mon.Monitor
}

func newMB (mL, mR uint) LeftRight {
  x := new(monitorBounded)
  var nL, nR, // number of active lefties/righties
      tL, tR uint // number of lefties/righties within one turn
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            if nR > 0 || x.Awaited (rightIn) && tL >= mL {
              x.Monitor.Wait (leftIn)
            }
            nL++
            tL++
            tR = 0
            if tL < mL || ! x.Awaited (rightIn) { x.Monitor.Signal (leftIn) }
          case rightIn:
            if nL > 0 || x.Awaited (leftIn) && tR >= mR {
              x.Monitor.Wait (rightIn)
            }
            nR++
            tR++
            tL = 0
            if tR < mR || ! x.Awaited(leftIn) {
              x.Monitor.Signal (rightIn)
            }
          case leftOut:
            nL--
            if nL == 0 {
              x.Monitor.Signal (rightIn)
            } else {
              if tL < mL && ! x.Awaited (rightIn) {
                x.Monitor.Signal (leftIn)
              }
            }
          case rightOut:
            nR--
            if nR == 0 {
              x.Monitor.Signal (leftIn)
            } else {
              if tR < mR && ! x.Awaited (leftIn) {
                x.Monitor.Signal (rightIn)
              }
            }
          }
          return nil
        }
  x.Monitor = mon.New (4, fs)
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
