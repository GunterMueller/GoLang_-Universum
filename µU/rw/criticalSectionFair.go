package rw

// (c) Christian Maurer   v. 171018 - license see µU.go

// >>> readers/writers problem: fair solution

import
  "µU/cs"
type
  criticalSectionFair struct {
                             cs.CriticalSection
                             }

func newCSF() ReaderWriter {
  var nR, nW uint
  var lastR bool
  x := new(criticalSectionFair)
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
         }
         nW = 1
         lastR = false
         return nW
       }
  l := func (i uint) {
         if i == reader {
           nR--
         } else {
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

func (x *criticalSectionFair) Fin() {
}
