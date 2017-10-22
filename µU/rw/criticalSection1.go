package rw

// (c) Christian Maurer   v. 171018 - license see µU.go

// >>> 1st readers/writers problem

import
  "µU/cs"
type
  criticalSection1 struct {
                          cs.CriticalSection
                          }

func newCS1() ReaderWriter {
  var nR, nW uint
  x := new(criticalSection1)
  c := func (i uint) bool {
         if i == reader {
           return nW == 0
         }
         return nR == 0 && nW == 0
       }
  e := func (i uint) uint {
         if i == reader {
           nR++
           return nR
         }
         nW = 1
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

func (x *criticalSection1) Fin() {
}
