package rw

// (c) Christian Maurer   v. 171126 - license see nU.go

import "nU/cs"

type criticalSection1 struct {
  cs.CriticalSection
}

func newCS1() ReaderWriter {
  x := new(criticalSection1)
  var nR, nW uint
  c := func (i uint) bool {
         if i == reader {
           return nW == 0
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

func (x *criticalSection1) ReaderIn() {
  x.Enter (reader)
}

func (x *criticalSection1) ReaderOut() {
  x.Leave (reader)
}

func (x *criticalSection1) WriterIn() {
  x.Enter (writer)
}

func (x *criticalSection1) WriterOut() {
  x.Leave (writer)
}
