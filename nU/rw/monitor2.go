package rw

// (c) Christian Maurer   v. 220702 - license see nU.go

import "nU/mon"

type monitor2 struct {
  mon.Monitor
}

func newM2() ReaderWriter {
  x := new(monitor2)
  var nR, nW uint
  f := func (a any, k uint) any {
         switch k {
         case readerIn:
           for nW > 0 || x.Awaited (writerIn) && nR > 0 {
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
         case writerOut:
           nW = 0
           if x.Awaited (readerIn) {
             x.Signal (readerIn)
           } else {
             x.Signal (writerIn)
           }
         }
         return nil
       }
  x.Monitor = mon.New (4, f)
  return x
}

func (x *monitor2) ReaderIn() { x.F (nil, readerIn) }
func (x *monitor2) ReaderOut() { x.F (nil, readerOut) }
func (x *monitor2) WriterIn() { x.F (nil, writerIn) }
func (x *monitor2) WriterOut() { x.F (nil, writerOut) }
