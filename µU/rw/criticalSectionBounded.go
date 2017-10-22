package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> bounded readers/writers problem

import
  "µU/cs"
type
  criticalSectionBounded struct {
                                cs.CriticalSection
                                }

func newCSB (m uint) ReaderWriter {
  var nR, nW uint
  var tR uint // number of readers within one turn
  x := new(criticalSectionBounded)
  if m < 1 { m = 1 }
  c := func (i uint) bool {
         if i == reader {
           return nW == 0 && (! x.Blocked (writer) || tR < m)
         }
         return nR == 0 && nW == 0 && ! x.Blocked (reader)
       }
  e := func (i uint) uint {
         if i == reader {
           nR++
           tR++
           return nR
         }
         nW = 1
         tR = 0
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

func (x *criticalSectionBounded) ReaderIn() {
  x.Enter (reader)
}

func (x *criticalSectionBounded) ReaderOut() {
  x.Leave (reader)
}

func (x *criticalSectionBounded) WriterIn() {
  x.Enter (writer)
}

func (x *criticalSectionBounded) WriterOut() {
  x.Leave (writer)
}

func (x *criticalSectionBounded) Fin() {
}
