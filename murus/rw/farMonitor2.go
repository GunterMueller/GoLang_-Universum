package rw

// (c) Christian Maurer   v. 170411 - license see murus.go

// >>> readers/writers problem: implementation with far monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, Abschnitt 8.3

import (
  . "murus/obj"
  "murus/host"
  "murus/fmon"
)
type
  farMonitor2 struct {
              nR, nW,
           nWblocked uint
                     fmon.FarMonitor
                     }

func newFMon2 (h host.Host, p uint16, s bool) ReaderWriter {
  x := new(farMonitor2)
  ps := func (a Any, k uint) bool {
          switch k {
          case readerIn:
            return x.nW == 0 && x.nWblocked == 0
          case writerIn:
            ok :=  x.nR == 0 && x.nW == 0
            if ! ok {
              x.nWblocked++
            }
            return ok
          case readerOut:
            if x.nWblocked > 0 && x.nR == 1 && x.nW == 0 {
              x.nWblocked--
            }
            return true
          case writerOut:
            if x.nWblocked > 0 && x.nR == 0 && x.nW == 0 {
              x.nWblocked--
            }
            return true
          }
          panic("unreachable")
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

func (x *farMonitor2) ReaderIn() {
  x.F(true, readerIn)
}

func (x *farMonitor2) ReaderOut() {
  x.F(true, readerOut)
}

func (x *farMonitor2) WriterIn() {
  x.F(true, writerIn)
}

func (x *farMonitor2) WriterOut() {
  x.F(true, writerOut)
}

func (x *farMonitor2) Fin() {
  x.FarMonitor.Fin()
}
