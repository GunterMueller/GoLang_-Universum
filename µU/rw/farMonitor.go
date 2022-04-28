package rw

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> 1st readers/writers problem

import (
  "µU/fmon"
)
type
  farMonitor struct {
                    fmon.FarMonitor
                    }

func newFM (h string, p uint16, s bool) ReaderWriter {
  var nR, nW uint
  x := new(farMonitor)
  ps := func (a any, i uint) bool {
          switch i {
          case readerIn:
            return nW == 0
          case writerIn:
            return nR == 0 && nW == 0
          }
          return true // readerOut, writerOut
        }
  fs := func (a any, i uint) any {
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
  x.FarMonitor = fmon.New (uint(0), 4, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) ReaderIn() {
  x.F (uint(0), readerIn)
}

func (x *farMonitor) ReaderOut() {
  x.F (uint(0), readerOut)
}

func (x *farMonitor) WriterIn() {
  x.F (uint(0), writerIn)
}

func (x *farMonitor) WriterOut() {
  x.F (uint(0), writerOut)
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}
