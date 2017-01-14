package rw

// (c) murus.org  v. 140330 - license see murus.go

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


func NewCriticalSectionBounded (m uint) ReaderWriter {
  x:= new (criticalSectionBounded)
  if m < 1 { m = 1 }
  c:= func (i uint) bool {
        if i == reader {
          return x.nW == 0 && (x.Blocked (writer) || x.rR < m)
        }
        return x.nR == 0 && x.nW == 0 /* && ! x.Blocked (reader) */ // writer
      }
  e:= func (a Any, i uint) {
        if i == reader {
          x.nR++
          x.rR++
        } else { // writer
          x.nW++
          x.rR = 0
        }
      }
  l:= func (a Any, i uint) {
        if i == reader {
          x.nR--
        } else { // writer
          x.nW--
        }
      }
  x.CriticalSection = cs.New (2, c, e, l)
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
