package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  "nU/cmon"
type
  condMonitorBounded struct {
                            cmon.Monitor
                            }

func newCMB (mR uint) ReaderWriter {
  x := new(condMonitorBounded)
  var nR, nW,
      tR uint // number of readerss within one turn
  p := func (i uint) bool {
         switch i {
         case readerIn:
           return nW == 0 && (x.Blocked (writerIn) == 0 || tR < mR)
         case writerIn:
           return nR == 0 && nW == 0 && x.Blocked (readerIn) == 0
         }
         return true
       }
  f := func (i uint) uint {
         switch i {
         case readerIn:
           nR++
           tR++
           return nR
         case readerOut:
           nR--
           return nR
         case writerIn:
           nW = 1
           tR = 0
         case writerOut:
           nW = 0
         }
         return nW
       }
  x.Monitor = cmon.New (4, f, p)
  return x
}

func (x *condMonitorBounded) ReaderIn() { x.F (readerIn) }
func (x *condMonitorBounded) ReaderOut() { x.F (readerOut) }
func (x *condMonitorBounded) WriterIn() { x.F (writerIn) }
func (x *condMonitorBounded) WriterOut() { x.F (writerOut) }
