package rw

// (c) Christian Maurer   v. 171101 - license see µU.go

// >>> bounded readers/writers problem

import
  "µU/cmon"
type
  condMonitorBounded struct {
                            cmon.Monitor
                            }

func newCMB (m uint) ReaderWriter {
  var nR, nW uint
//  var tR uint // number of readers within one turn XXX
  x := new(condMonitorBounded)
  if m == 0 { m = 1 }
  c := func (i uint) bool {
         switch i {
         case readerIn:
//           return nW == 0 && (x.Blocked (writerIn) == 0 || tR < m)
           return nW == 0 && (x.Blocked (writerIn) == 0 || nR < m)
         case writerIn:
           return nR == 0 && nW == 0 && x.Blocked (readerIn) == 0
         }
         return true
       }
  f := func (i uint) uint {
         switch i {
         case readerIn:
           nR++
//           tR++
           return nR
         case readerOut:
           nR--
           return nR
         case writerIn:
           nW = 1
//           tR = 0
         case writerOut:
           nW = 0
         }
         return nW
       }
  x.Monitor = cmon.New (4, f, c)
  return x
}

func (x *condMonitorBounded) ReaderIn() {
  x.F (readerIn)
}

func (x *condMonitorBounded) ReaderOut() {
  x.F (readerOut)
}

func (x *condMonitorBounded) WriterIn() {
  x.F (writerIn)
}

func (x *condMonitorBounded) WriterOut() {
  x.F (writerOut)
}

func (x *condMonitorBounded) Fin() {
}
