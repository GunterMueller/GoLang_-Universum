package rw

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> bounded readers/writers problem

import (
  "µU/fmon"
)
type
  farMonitorBounded struct {
                           fmon.FarMonitor
                           }

func newFMB (m uint, h string, port uint16, s bool) ReaderWriter {
  if m == 0 { m = inf } // unbounded case
  var nR, nW uint
  var tR uint // number of active readers within one turn
  x := new(farMonitorBounded)
  p := func (a any, i uint) bool {
         switch i {
         case readerIn:
           return nW == 0 && tR < m
         case writerIn:
           return nR == 0 && nW == 0
         }
         return true // readerOut, writerOut
       }
  f := func (a any, i uint) any {
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
  x.FarMonitor = fmon.New (uint(0), 4, f, p, h, port, s)
  return x
}

func (x *farMonitorBounded) ReaderIn() {
 x.F(uint(0), readerIn)
}

func (x *farMonitorBounded) ReaderOut() {
  x.F(uint(0), readerOut)
}

func (x *farMonitorBounded) WriterIn() {
  x.F(uint(0), writerIn)
}

func (x *farMonitorBounded) WriterOut() {
  x.F(uint(0), writerOut)
}

func (x *farMonitorBounded) Fin() {
  x.FarMonitor.Fin()
}
