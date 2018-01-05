package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/fmon")

type farMonitor struct {
  fmon.FarMonitor
}

func newFM (h string, p uint16, s bool) ReaderWriter {
  var nR, nW uint
  x := new(farMonitor)
  c := func (a Any, i uint) bool {
         switch i {
         case readerIn:
           return nW == 0
         case writerIn:
           return nR == 0 && nW == 0
         }
         return true // readerOut, writerOut
       }
  f := func (a Any, i uint) Any {
         switch i {
         case readerIn:
           nR++
           return nR
         case readerOut:
           nR--
           return nR
         case writerIn:
           nW = 1
         case writerOut:
           nW = 0
         }
         return nW
       }
  x.FarMonitor = fmon.New (uint(0), 4, f, c, h, p, s)
  return x
}

func (x *farMonitor) ReaderIn() {
  x.F (0, readerIn)
}

func (x *farMonitor) ReaderOut() {
  x.F (0, readerOut)
}

func (x *farMonitor) WriterIn() {
  x.F (0, writerIn)
}

func (x *farMonitor) WriterOut() {
  x.F (0, writerOut)
}
