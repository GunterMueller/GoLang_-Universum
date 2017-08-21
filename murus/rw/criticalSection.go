package rw

// (c) murus.org  v. 170411 - license see murus.go

// >>> readers/writers problem: implementation with critical sections
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 92

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSection struct {
                  nR, nW uint
                         cs.CriticalSection
                         }

func newCS() ReaderWriter {
  x := new (criticalSection)
  c := func (i uint) bool {
         if i == reader {
           return x.nW == 0
         }
         return x.nR == 0 && x.nW == 0 // writer
       }
  es := func (a Any, i uint) {
          if i == reader {
            x.nR++
          } else { // writer
            x.nW = 1
          }
        }
  ls := func (a Any, i uint) {
          if i == reader {
            x.nR--
          } else { // writer
            x.nW = 0
          }
        }
  x.CriticalSection = cs.New (2, c, es, ls)
  return x
}

func (x *criticalSection) ReaderIn() {
  x.Enter (reader, nil)
}

func (x *criticalSection) ReaderOut() {
  x.Leave (reader, nil)
}

func (x *criticalSection) WriterIn() {
  x.Enter (writer, nil)
}

func (x *criticalSection) WriterOut() {
  x.Leave (writer, nil)
}

func (x *criticalSection) Fin() {
}
