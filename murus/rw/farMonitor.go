package rw

// (c) murus.org  v. 161226 - license see murus.go

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

func NewFarMonitor (h host.Host, p uint16, s bool) ReaderWriter {
  x:= new (farMonitor)
  c:= func (a Any, k uint) bool {
        switch k {
        case rIn:
          return x.nW == 0
        case wIn:
          return x.nR == 0 && x.nW == 0
        }
        return true // rOut, wOut
      }
  f:= func (a Any, k uint) Any {
        switch k {
        case rIn:
          x.nR ++
        case rOut:
          x.nR --
        case wIn:
          x.nW = 1
        case wOut:
          x.nW = 0
        }
        return true
      }
  x.FarMonitor = fmon.New (true, nFuncs, f, c, h, p, s)
  return x
}

func (x *farMonitor) ReaderIn() {
  x.F (true, rIn)
}

func (x *farMonitor) ReaderOut() {
  x.F (true, rOut)
}

func (x *farMonitor) WriterIn() {
  x.F (true, wIn)
}

func (x *farMonitor) WriterOut() {
  x.F (true, wOut)
}
