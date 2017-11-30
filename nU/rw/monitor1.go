package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/mon")

type monitor1 struct {
  mon.Monitor
}

func newM1() ReaderWriter {
  x := new(monitor1)
  var nR, nW uint
  f := func (a Any, k uint) Any {
         switch k {
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

func (x *monitor1) ReaderIn() { x.F (nil, readerIn) }
func (x *monitor1) ReaderOut() { x.F (nil, readerOut) }
func (x *monitor1) WriterIn() { x.F (nil, writerIn) }
func (x *monitor1) WriterOut() { x.F (nil, writerOut) }
