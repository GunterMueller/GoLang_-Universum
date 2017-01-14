package rw

// (c) murus.org  v. 150304 - license see murus.go

// >>> readers/writers problem: implementation with monitors
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 161

import (
  . "murus/obj"
  "murus/mon"
)
type
  monitor struct {
          nR, nW uint
                 mon.Monitor
                 }


func NewMonitor() ReaderWriter {
  x:= new (monitor)
  f:= func (a Any, k uint) Any {
        switch k { case rIn:
          for x.nW > 0 { x.Wait (rIn) }
          x.nR ++
          x.Signal (rIn)
        case rOut:
          x.nR --
          if x.nR == 0 {
            x.Signal (wIn)
          }
        case wIn:
          for x.nR > 0 || x.nW > 0 { x.Wait (wIn) }
          x.nW = 1
        case wOut:
          x.nW = 0
          if x.Awaited (rIn) {
            x.Signal (rIn)
          } else {
            x.Signal (wIn)
          }
        }
        return nil
      }
  x.Monitor = mon.New (nFuncs, f, nil)
  return x
}

func (x *monitor) ReaderIn() {
  x.F (nil, rIn)
}

func (x *monitor) ReaderOut() {
  x.F (nil, rOut)
}

func (x *monitor) WriterIn() {
  x.F (nil, wIn)
}

func (x *monitor) WriterOut() {
  x.F (nil, wOut)
}
