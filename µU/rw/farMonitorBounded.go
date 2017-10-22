package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> bounded readers/writers problem

import (
  . "µU/obj"
  "µU/host"
  "µU/fmon"
)
type
  farMonitorBounded struct {
                           fmon.FarMonitor
                           }

func newFMB (m uint, h host.Host, port uint16, s bool) ReaderWriter {
  if m == 0 { m = inf } // unbounded case
  var nR, nW uint
  var tR uint // number of active readers within one turn
  x := new(farMonitorBounded)
  p := func (a Any, i uint) bool {
         switch i {
         case readerIn:
           return nW == 0 && tR < m
         case writerIn:
           return nR == 0 && nW == 0
         }
         return true // readerOut, writerOut
       }
  f := func (a Any, i uint) Any {
         switch i {
         case readerIn:
           nR++
           tR++
           return nR
         case readerOut:
           nR--
           return nR
         case writerIn:
           nW = 1
           tR = 0
         case writerOut:
           nW = 0
         }
         return nW
       }
  x.FarMonitor = fmon.New (true, nFuncs, f, p, h, port, s)
  return x
}

func (x *farMonitorBounded) ReaderIn() {
 x.F(true, readerIn)
}

func (x *farMonitorBounded) ReaderOut() {
  x.F(true, readerOut)
}

func (x *farMonitorBounded) WriterIn() {
  x.F(true, writerIn)
}

func (x *farMonitorBounded) WriterOut() {
  x.F(true, writerOut)
}

func (x *farMonitorBounded) Fin() {
  x.FarMonitor.Fin()
}
