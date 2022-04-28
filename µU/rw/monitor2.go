package rw

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> 2nd readers/writers problem

import (
  "µU/mon"
)
type
  monitor2 struct {
                  mon.Monitor
                  }

func newM2() ReaderWriter {
  var nR, nW uint
  x := new(monitor2)
  f := func (a any, i uint) any {
         switch i {
         case readerIn:
           if nW > 0 || x.Awaited (writerIn) && nR > 0 {
             x.Wait (readerIn)
           }
           nR++
           x.Signal (readerIn)
           return nR
         case readerOut:
           nR--
           if nR == 0 {
             x.Signal (writerIn)
           }
           return nR
         case writerIn:
           if nR > 0 || nW > 0 {
             x.Wait (writerIn)
           }
           nW = 1
           return nW
         case writerOut:
           nW = 0
           if x.Awaited (readerIn) {
             x.Signal (readerIn)
           } else {
             x.Signal (writerIn)
           }
           return nW
         }
         return 0
       }
  x.Monitor = mon.New (4, f)
  return x
}

func (x *monitor2) ReaderIn() {
  x.F (nil, readerIn)
}

func (x *monitor2) ReaderOut() {
  x.F (nil, readerOut)
}

func (x *monitor2) WriterIn() {
  x.F (nil, writerIn)
}

func (x *monitor2) WriterOut() {
  x.F (nil, writerOut)
}

func (x *monitor2) Fin() {
}
