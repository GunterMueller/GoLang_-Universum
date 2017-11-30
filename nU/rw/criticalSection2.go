package rw

// (c) Christian Maurer   v. 171126 - license see nU.go

import "nU/cs"

type criticalSection2 struct {
  cs.CriticalSection
}

func newCS2() ReaderWriter {
  x := new(criticalSection2)
  var nR, nW uint
  c := func (i uint) bool {
         if i == reader {
           return nW == 0 && ! x.Blocked (writer)
         }
         return nR == 0 && nW == 0 // writer
       }
  e := func (i uint) uint {
          if i == reader {
            nR++
            return nR
          } else { // writer
            nW = 1
          }
          return nW
        }
  l := func (i uint) {
         if i == reader {
           nR--
         } else { // writer
           nW = 0
         }
       }
  x.CriticalSection = cs.New (2, c, e, l)
  return x
}

func (x *criticalSection2) ReaderIn() {
  x.Enter (reader)
}

func (x *criticalSection2) ReaderOut() {
  x.Leave (reader)
}

func (x *criticalSection2) WriterIn() {
  x.Enter (writer)
}

func (x *criticalSection2) WriterOut() {
  x.Leave (writer)
}
