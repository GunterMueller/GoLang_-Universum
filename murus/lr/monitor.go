package lr

// (c) murus.org  v. 170731 - license see murus.go

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

func newMon() LeftRight {
  x := new(monitor)
  fs := func (a Any, i uint) Any {
          switch i {
          case leftIn:
            for x.nR > 0 { x.Monitor.Wait (leftIn) }
            x.nL++
            writeL (x.nL)
            x.Monitor.Signal (leftIn)
          case rightIn:
            for x.nL > 0 { x.Monitor.Wait (rightIn) }
            x.nR++
            writeR (x.nR)
            x.Monitor.Signal (rightIn)
          case leftOut:
            x.nL--
            writeL (x.nL)
            if x.nL == 0 {
              x.Monitor.Signal (rightIn)
            } else {
              x.Monitor.Signal (leftIn)
            }
          case rightOut:
            x.nR--
            writeR (x.nR)
            if x.nR == 0 {
              x.Monitor.Signal (leftIn)
            } else {
              x.Monitor.Signal (rightIn)
            }
          }
          return nil
        }
  x.Monitor = mon.New (nFuncs, fs, nil)
  return x
}

func (x *monitor) LeftIn() {
  x.F (true, leftIn)
}

func (x *monitor) LeftOut() {
  x.F (true, leftOut)
}

func (x *monitor) RightIn() {
  x.F (true, rightIn)
}

func (x *monitor) RightOut() {
  x.F (true, rightOut)
}

func (x *monitor) Fin() {
}
