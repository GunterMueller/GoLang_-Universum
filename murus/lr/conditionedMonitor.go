package lr

// (c) murus.org  v. 140330 - license see murus.go

// >>> left/right problem: implementation with conditioned monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 155 ff.

import (
  . "murus/obj"
  "murus/mon"
)
type
  conditionedMonitor struct {
                     nL, nR uint
                            mon.Monitor
                            }

func NewConditionedMonitor() LeftRight {
  x:= new (conditionedMonitor)
  p:= func (a Any, i uint) bool {
        switch i {
        case lIn:
          return x.nR == 0
        case rIn:
          return x.nL == 0
        }
        return true // lOut, rOut
      }
  f:= func (a Any, i uint) Any {
        switch i {
        case lIn:
          x.nL ++
        case lOut:
          x.nL --
        case rIn:
          x.nR ++
        case rOut:
          x.nR --
        }
        return a
      }
  x.Monitor = mon.NewF (nil, nFuncs, f, p)
  x.Monitor = mon.New (nFuncs, f, p)
  return x
}

func (x *conditionedMonitor) LeftIn() {
  x.F (true, lIn)
}

func (x *conditionedMonitor) LeftOut() {
  x.F (true, lOut)
}

func (x *conditionedMonitor) RightIn() {
  x.F (true, rIn)
}

func (x *conditionedMonitor) RightOut() {
  x.F (true, rOut)
}
