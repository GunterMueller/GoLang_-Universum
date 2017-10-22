package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st readers/writers problem

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h host.Host, port uint16, s bool) ReaderWriter {
  var nR, nW uint
  x := new(farMonitor)
  p := func (a Any, i uint) bool {
         switch i {
         case readerIn:
           return nW == 0
         case writerIn:
           return nR == 0 && nW == 0
         }
         return true // rOut, wOut
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
  x.FarMonitor = fmon.New (true, nFuncs, f, p, h, port, s)
  return x
}

func (x *farMonitor) ReaderIn() {
  x.F (true, readerIn)
}

func (x *farMonitor) ReaderOut() {
  x.F (true, readerOut)
}

func (x *farMonitor) WriterIn() {
  x.F (true, writerIn)
}

func (x *farMonitor) WriterOut() {
  x.F (true, writerOut)
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}
