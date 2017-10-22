package lr

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st left/right problem

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitor1 struct {
                  mon.Monitor
                  }

func newM1() LeftRight {
  x := new(monitor1)
  var nL, nR uint
  f := func (a Any, i uint) Any {
         switch i {
         case leftIn:
           if nR > 0 {
             x.Monitor.Wait (leftIn)
           }
           nL++
           x.Monitor.Signal (leftIn)
         case rightIn:
           if nL > 0 {
             x.Monitor.Wait (rightIn)
           }
           nR++
           x.Monitor.Signal (rightIn)
         case leftOut:
           nL--
           if nL == 0 {
             x.Monitor.Signal (rightIn)
           } else {
             x.Monitor.Signal (leftIn)
           }
         case rightOut:
           nR--
           if nR == 0 {
             x.Monitor.Signal (leftIn)
           } else {
             x.Monitor.Signal (rightIn)
           }
         }
         return nil
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *monitor1) LeftIn() {
  x.F (nil, leftIn)
}

func (x *monitor1) LeftOut() {
  x.F (nil, leftOut)
}

func (x *monitor1) RightIn() {
  x.F (nil, rightIn)
}

func (x *monitor1) RightOut() {
  x.F (nil, rightOut)
}

func (x *monitor1) Fin() {
}
