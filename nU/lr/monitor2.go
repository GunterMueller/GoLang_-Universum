package lr

// (c) Christian Maurer   v. 220702 - license see nU.go

import "nU/mon"

type monitor2 struct {
  mon.Monitor
}

func newM2() LeftRight {
  x := new(monitor2)
  var nL, nR uint
  f := func (a any, i uint) any {
         switch i {
         case leftIn:
           for nR > 0 || x.Awaited (rightIn) && nL > 0 {
             x.Monitor.Wait (leftIn)
           }
           nL++
           x.Monitor.Signal (leftIn)
         case rightIn:
           for nL > 0 || x.Awaited (leftIn) && nR > 0 {
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
  x.Monitor = mon.New (4, f)
  return x
}

func (x *monitor2) LeftIn() {
  x.F (true, leftIn)
}

func (x *monitor2) LeftOut() {
  x.F (true, leftOut)
}

func (x *monitor2) RightIn() {
  x.F (true, rightIn)
}

func (x *monitor2) RightOut() {
  x.F (true, rightOut)
}
