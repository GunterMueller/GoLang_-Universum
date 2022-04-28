package lr

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> bounded left/right problem

import (
  "µU/mon"
)
type
  monitorBounded struct {
                        mon.Monitor
                        }

func newMB (mL, mR uint) LeftRight {
  var nL, nR uint
  var tL, tR uint // number of actives within one turn
  x := new(monitorBounded)
  f := func (a any, i uint) any {
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
         return uint(0)
       }
  x.Monitor = mon.New (4, f)
  return x
}

func (x *monitorBounded) LeftIn() {
  x.F (nil, leftIn)
}

func (x *monitorBounded) LeftOut() {
  x.F (nil, leftOut)
}

func (x *monitorBounded) RightIn() {
  x.F (nil, rightIn)
}

func (x *monitorBounded) RightOut() {
  x.F (nil, rightOut)
}

func (x *monitorBounded) Fin() {
}
