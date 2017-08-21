package rw

// (c) murus.org  v. 170411 - license see murus.go

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


func newCS2() ReaderWriter {
  x := new (criticalSection2)
  c := func (k uint) bool {
         if k == reader {
           return x.nW == 0 && ! x.Blocked (writer)
         }
         return x.nR == 0 && x.nW == 0 // writer
       }
  es := func (a Any, k uint) {
          if k == reader {
            x.nR++
          } else { // writer
            x.nW = 1
          }
        }
  ls := func (a Any, k uint) {
          if k == reader {
            x.nR--
          } else { // writer
            x.nW = 0
          }
        }
  x.CriticalSection = cs.New (2, c, es, ls)
  return x
}

func (x *criticalSection2) ReaderIn() {
  x.Enter (reader, nil)
}

func (x *criticalSection2) ReaderOut() {
  x.Leave (reader, nil)
}

func (x *criticalSection2) WriterIn() {
  x.Enter (writer, nil)
}

func (x *criticalSection2) WriterOut() {
  x.Leave (writer, nil)
}

func (x *criticalSection2) Fin() {
}
