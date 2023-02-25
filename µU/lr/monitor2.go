package lr

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> 2nd left/right problem

import
  "µU/mon"
type
  monitor2 struct {
                  mon.Monitor
                  }

func newM2() LeftRight {
  var nL, nR uint
  x := new(monitor2)
  f := func (a any, i uint) any {
         switch i {
         case leftIn:
           if nR > 0 || x.Awaited (rightIn) && nL > 0 {
             x.Monitor.Wait (leftIn)
           }
           nL++
           x.Monitor.Signal (leftIn)
         case rightIn:
           if nL > 0 || x.Awaited (leftIn) && nR > 0 {
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
  x.F (nil, leftIn)
}

func (x *monitor2) LeftOut() {
  x.F (nil, leftOut)
}

func (x *monitor2) RightIn() {
  x.F (nil, rightIn)
}

func (x *monitor2) RightOut() {
  x.F (nil, rightOut)
}

func (x *monitor2) Fin() {
}
