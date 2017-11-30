package rw

// (c) Christian Maurer   v. 171016 - license see nU.go

import "nU/cmon"

type conditionedMonitor struct {
  cmon.Monitor
}

func newCM() ReaderWriter {
  x := new(conditionedMonitor)
  var nR, nW uint
  p := func (i uint) bool {
         switch i {
         case readerIn:
           return nW == 0
         case writerIn:
           return nR == 0 && nW == 0
         }
         return true
       }
  f := func (i uint) uint {
         switch i {
         case readerIn:
           nR++
           return nR
         case readerOut:
           nR--
           return nR
         case writerIn:
           nW = 1
         case writerOut:
           nW = 0
         }
         return nW
       }
  x.Monitor = cmon.New (4, f, p)
  return x
}

func (x *conditionedMonitor) ReaderIn() { x.F (readerIn) }
func (x *conditionedMonitor) ReaderOut() { x.F (readerOut) }
func (x *conditionedMonitor) WriterIn() { x.F (writerIn) }
func (x *conditionedMonitor) WriterOut() { x.F (writerOut) }
