package rw

// (c) murus.org  v. 170411 - license see murus.go

// >>> readers/writers problem: implementation with conditioned monitors

import (
  . "murus/obj"
  "murus/mon"
)
type
  conditionedMonitor struct {
                     nR, nW uint
                            mon.Monitor
                            }


func newCondMon() ReaderWriter {
  x := new (conditionedMonitor)
  fs := func (a Any, i uint) Any {
          switch i {
          case readerIn:
            x.nR++
          case readerOut:
            x.nR--
          case writerIn:
            x.nW++
          case writerOut:
            x.nW--
          }
          return nil
        }
  ps := func (a Any, i uint) bool {
          switch i {
          case readerIn:
            return x.nW == 0
          case writerIn:
            return x.nR == 0 && x.nW == 0
          }
          return true
        }
  x.Monitor = mon.New (nFuncs, fs, ps)
  return x
}

func (x *conditionedMonitor) ReaderIn() {
  x.F (nil, readerIn)
}

func (x *conditionedMonitor) ReaderOut() {
  x.F (nil, readerOut)
}

func (x *conditionedMonitor) WriterIn() {
  x.F (nil, writerIn)
}

func (x *conditionedMonitor) WriterOut() {
  x.F (nil, writerOut)
}

func (x *conditionedMonitor) Fin() {
}
