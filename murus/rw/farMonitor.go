package rw

// (c) murus.org  v. 170520 - license see murus.go

// >>> readers/writers problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, Abschnitt 8.3

import (
  . "murus/obj"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor struct {
             nR, nW uint
                    fmon.FarMonitor
                    }

func newFMon (h host.Host, p uint16, s bool) ReaderWriter {
  x := new(farMonitor)
  ps := func (a Any, k uint) bool {
          switch k {
          case readerIn:
            return x.nW == 0
          case writerIn:
            return x.nR == 0 && x.nW == 0
          }
          return true // rOut, wOut
        }
  fs := func (a Any, k uint) Any {
          switch k {
          case readerIn:
            x.nR++
            writeR (x.nR)
          case readerOut:
            x.nR--
            writeR (x.nR)
          case writerIn:
            x.nW = 1
            writeW (x.nW)
          case writerOut:
            x.nW = 0
            writeW (x.nW)
           }
          return true
        }
  x.FarMonitor = fmon.New (true, nFuncs, fs, ps, h, p, s)
  return x
}

func (x *farMonitor) ReaderIn() {
  x.F(true, readerIn)
}

func (x *farMonitor) ReaderOut() {
  x.F(true, readerOut)
}

func (x *farMonitor) WriterIn() {
  x.F(true, writerIn)
}

func (x *farMonitor) WriterOut() {
  x.F(true, writerOut)
}

func (x *farMonitor) Fin() {
  x.FarMonitor.Fin()
}
