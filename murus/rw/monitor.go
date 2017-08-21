package rw

// (c) murus.org  v. 170121 - license see murus.go

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


func newMon() ReaderWriter {
  x := new (monitor)
  fs := func (a Any, k uint) Any {
          switch k {
          case readerIn:
            for x.nW > 0 { x.Wait (readerIn) }
            x.nR++
            x.Signal (readerIn)
          case readerOut:
            x.nR--
            if x.nR == 0 {
              x.Signal (writerIn)
            }
          case writerIn:
            for x.nR > 0 || x.nW > 0 { x.Wait (writerIn) }
            x.nW = 1
          case writerOut:
            x.nW = 0
            if x.Awaited (readerIn) {
              x.Signal (readerIn)
            } else {
              x.Signal (writerIn)
            }
          }
          return nil
        }
  x.Monitor = mon.New (nFuncs, fs, nil)
  return x
}

func (x *monitor) ReaderIn() {
  x.F (nil, readerIn)
}

func (x *monitor) ReaderOut() {
  x.F (nil, readerOut)
}

func (x *monitor) WriterIn() {
  x.F (nil, writerIn)
}

func (x *monitor) WriterOut() {
  x.F (nil, writerOut)
}

func (x *monitor) Fin() {
}
