package lr

// (c) murus.org  v. 150304 - license see murus.go

// >>> left/right problem: implementation with monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 161

import (
  . "murus/obj"
  "murus/mon"
)
type
  monitor struct {
          nL, nR uint
                 mon.Monitor
                 }

func NewMonitor() LeftRight {
  x:= new (monitor)
  f:= func (a Any, i uint) Any {
        switch i {
        case lIn:
          for x.nR > 0 { x.Monitor.Wait(rIn) }
          x.nL++
          x.Monitor.Signal(lIn)
        case rIn:
          for x.nL > 0 { x.Monitor.Wait(lIn) }
          x.nR++
          x.Monitor.Signal(rIn)
        case lOut:
          x.nL--
          if x.nL == 0 {
            x.Monitor.Signal(rIn)
          } else {
            x.Monitor.Signal(lIn)
          }
        case rOut:
          x.nR--
          if x.nR == 0 {
            x.Monitor.Signal(lIn)
          } else {
            x.Monitor.Signal(rIn)
          }
        }
        return nil
      }
  x.Monitor = mon.New (nFuncs, f, nil)
  return x
}

func (x *monitor) LeftIn() {
  x.F (true, lIn)
}

func (x *monitor) LeftOut() {
  x.F (true, lOut)
}

func (x *monitor) RightIn() {
  x.F (true, rIn)
}

func (x *monitor) RightOut() {
  x.F (true, rOut)
}
