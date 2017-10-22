package rw

// (c) Christian Maurer   v. 171017 - license see µU.go

// >>> 1st readers/writers problem

import (
  . "µU/obj"
  "µU/mon"
)
type
  monitor1 struct {
                  mon.Monitor
                  }

func newM1() ReaderWriter {
  var nR, nW uint
  x := new(monitor1)
  f := func (a Any, i uint) Any {
         switch i {
         case readerIn:
           for nW > 0 {
             x.Wait (readerIn)
           }
           nR++
           x.Signal (readerIn)
         case readerOut:
           nR--
           if nR == 0 {
             x.Signal (writerIn)
           }
         case writerIn:
           for nR > 0 || nW > 0 {
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
         return nR
       }
  x.Monitor = mon.New (nFuncs, f)
  return x
}

func (x *monitor1) ReaderIn() {
  x.F (nil, readerIn)
}

func (x *monitor1) ReaderOut() {
  x.F (nil, readerOut)
}

func (x *monitor1) WriterIn() {
  x.F (nil, writerIn)
}

func (x *monitor1) WriterOut() {
  x.F (nil, writerOut)
}

func (x *monitor1) Fin() {
}
