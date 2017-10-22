package rw

// (c) Christian Maurer   v. 171019 - license see µU.go

// >>> 1st readers/writers problem

import
  "µU/cmon"
type
  conditionedMonitor struct {
                            cmon.Monitor
                            }


func newCM() ReaderWriter {
  var nR, nW uint
  x := new(conditionedMonitor)
  c := func (i uint) bool {
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
  x.Monitor = cmon.New (nFuncs, f, c)
  return x
}

func (x *conditionedMonitor) ReaderIn() {
  x.F (readerIn)
}

func (x *conditionedMonitor) ReaderOut() {
  x.F (readerOut)
}

func (x *conditionedMonitor) WriterIn() {
  x.F (writerIn)
}

func (x *conditionedMonitor) WriterOut() {
  x.F (writerOut)
}

func (x *conditionedMonitor) Fin() {
}
