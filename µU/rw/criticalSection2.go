package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 2nd readers/writers problem

import
  "µU/cs"
type
  criticalSection2 struct {
                          cs.CriticalSection
                          }


func newCS2() ReaderWriter {
  var nR, nW uint
  x := new(criticalSection2)
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
         }
         nW = 1
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

func (x *criticalSection2) Fin() {
}
