package rw

// (c) murus.org  v. 140330 - license see murus.go

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


func NewConditionedMonitor() ReaderWriter {
  x:= new (conditionedMonitor)
  f:= func (a Any, i uint) Any {
        switch i {
        case rIn:
          x.nR++
        case rOut:
          x.nR--
        case wIn:
          x.nW++
        case wOut:
          x.nW--
        }
        return nil
      }
  p:= func (a Any, i uint) bool {
        switch i {
        case rIn:
          return x.nW == 0
        case wIn:
          return x.nR == 0 && x.nW == 0
        }
        return true
      }
  x.Monitor = mon.New (nFuncs, f, p)
  return x
}

func (x *conditionedMonitor) ReaderIn() {
  x.F (nil, rIn)
}

func (x *conditionedMonitor) ReaderOut() {
  x.F (nil, rOut)
}

func (x *conditionedMonitor) WriterIn() {
  x.F (nil, wIn)
}

func (x *conditionedMonitor) WriterOut() {
  x.F (nil, wOut)
}
