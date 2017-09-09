package rw

// (c) Christian Maurer   v. 170411 - license see murus.go

// >>> 3rd readers/writers problem (number of concurrent readers bounded)

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSectionBounded struct {
                     nR, nW, rR uint
                                cs.CriticalSection
                                }


func newCSB (m uint) ReaderWriter {
  x := new (criticalSectionBounded)
  if m < 1 { m = 1 }
  c := func (i uint) bool {
         if i == reader {
           return x.nW == 0 && (x.Blocked (writer) || x.rR < m)
         }
         return x.nR == 0 && x.nW == 0 /* && ! x.Blocked (reader) */ // writer
       }
  es := func (a Any, i uint) {
          if i == reader {
            x.nR++
            x.rR++
          } else { // writer
            x.nW++
            x.rR = 0
          }
        }
  ls := func (a Any, i uint) {
          if i == reader {
            x.nR--
          } else { // writer
            x.nW--
          }
        }
  x.CriticalSection = cs.New (2, c, es, ls)
  return x
}

func (x *criticalSectionBounded) ReaderIn() {
  x.Enter (reader, nil)
}

func (x *criticalSectionBounded) ReaderOut() {
  x.Leave (reader, nil)
}

func (x *criticalSectionBounded) WriterIn() {
  x.Enter (writer, nil)
}

func (x *criticalSectionBounded) WriterOut() {
  x.Leave (writer, nil)
}

func (x *criticalSectionBounded) Fin() {
}
