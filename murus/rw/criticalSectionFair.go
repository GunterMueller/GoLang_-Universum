package rw

// (c) murus.org  v. 140330 - license see murus.go

// >>> readers/writers problem: fair solution
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 84

import (
  . "murus/obj"
  "murus/cs"
)
type
  criticalSectionFair struct {
                      nR, nW uint
                       lastR bool
                             cs.CriticalSection
                             }


func NewCriticalSectionFair() ReaderWriter {
  x:= new (criticalSectionFair)
  c:= func (i uint) bool {
        if i == reader {
          return x.nW == 0 && (! x.Blocked (writer) || ! x.lastR)
        }
        return x.nR == 0 && x.nW == 0 && (! x.Blocked (reader) || x.lastR) // writer
      }
  e:= func (a Any, i uint) {
        if i == reader {
          x.nR++
          x.lastR = true
        } else { // writer
          x.nW++
          x.lastR = false
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

func (x *criticalSectionFair) ReaderIn() {
  x.Enter (reader, nil)
}

func (x *criticalSectionFair) ReaderOut() {
  x.Leave (reader, nil)
}

func (x *criticalSectionFair) WriterIn() {
  x.Enter (writer, nil)
}

func (x *criticalSectionFair) WriterOut() {
  x.Leave (writer, nil)
}
