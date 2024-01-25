package rw

// (c) Christian Maurer   v. 171126 - license see nU.go

import
  "nU/cs"
type
  criticalSectionFair struct {
                             cs.CriticalSection
                             }

func newCSF() ReaderWriter {
  x := new(criticalSectionFair)
  var nR, nW uint
  var lastR bool
  c := func (i uint) bool {
         if i == reader {
           return nW == 0 && (! x.Blocked (writer) || ! lastR)
         }
         return nR == 0 && nW == 0 && (! x.Blocked (reader) || lastR)
       }
  e := func (i uint) uint {
         if i == reader {
           nR++
           lastR = true
           return nR
         } else { // writer
           nW = 1
           lastR = false
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

func (x *criticalSectionFair) ReaderIn() {
  x.Enter (reader)
}

func (x *criticalSectionFair) ReaderOut() {
  x.Leave (reader)
}

func (x *criticalSectionFair) WriterIn() {
  x.Enter (writer)
}

func (x *criticalSectionFair) WriterOut() {
  x.Leave (writer)
}
