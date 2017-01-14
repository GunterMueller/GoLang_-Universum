package rw

// (c) murus.org  v. 150304 - license see murus.go

// >>> 2nd readers/writers problem: implementation with critical sections

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSection2 struct {
                   nR, nW uint
                          cs.CriticalSection
                          }


func NewCriticalSection2() ReaderWriter {
  x:= new (criticalSection2)
  c:= func (i uint) bool {
        if i == reader {
          return x.nW == 0 && ! x.Blocked (writer)
        }
        return x.nR == 0 && x.nW == 0 // writer
      }
/*
  p:= func (a Any, i uint) bool {
        if a.(uint) == reader {
          return x.nW == 0 && ! x.Blocked (writer)
        }
        return x.nR == 0 && x.nW == 0 // writer
      }
*/
  e:= func (a Any, k uint) {
        if k == reader {
//        if a.(uint) == reader {
          x.nR++
        } else { // writer
          x.nW = 1
        }
      }
  l:= func (a Any, k uint) {
        if k == reader {
//        if a.(uint) == reader {
          x.nR--
        } else { // writer
          x.nW = 0
        }
      }
  x.CriticalSection = cs.New (2, c, e, l)
//  x.CriticalSection = cs.New (1, p, e, l)
  return x
}

func (x *criticalSection2) ReaderIn() {
  x.Enter (reader, nil)
//  x.Enter (0, reader)
}

func (x *criticalSection2) ReaderOut() {
  x.Leave (reader, nil)
//  x.Leave (0, reader)
}

func (x *criticalSection2) WriterIn() {
  x.Enter (writer, nil)
//  x.Enter (0, writer)
}

func (x *criticalSection2) WriterOut() {
  x.Leave (writer, nil)
//  x.Leave (0, writer)
}
