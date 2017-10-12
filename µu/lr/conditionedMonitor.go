package lr

// (c) Christian Maurer   v. 170731 - license see µu.go

// >>> left/right problem: implementation with conditioned monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 155 ff.

import (
  . "µu/obj"
  "µu/mon"
)
type
  conditionedMonitor struct {
                     nL, nR uint
                            mon.Monitor
                            }

func newCM() LeftRight {
  x := new (conditionedMonitor)
  ps := func (a Any, i uint) bool {
          switch i {
          case leftIn:
            return x.nR == 0
          case rightIn:
            return x.nL == 0
          }
          return true
        }
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            x.nL++
            writeL (x.nL)
          case leftOut:
            x.nL--
            writeL (x.nL)
          case rightIn:
            x.nR++
            writeR (x.nR)
          case rightOut:
            x.nR--
            writeR (x.nR)
          }
          return a
        }
  x.Monitor = mon.New (nFuncs, fs, ps)
  return x
}

func (x *conditionedMonitor) LeftIn() {
  x.F (true, leftIn)
}

func (x *conditionedMonitor) LeftOut() {
  x.F (true, leftOut)
}

func (x *conditionedMonitor) RightIn() {
  x.F (true, rightIn)
}

func (x *conditionedMonitor) RightOut() {
  x.F (true, rightOut)
}

func (x *conditionedMonitor) Fin() {
}
