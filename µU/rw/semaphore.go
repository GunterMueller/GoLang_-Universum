package rw

// (c) Christian Maurer   v. 170411 - license see µU.go

// >>> 1st readers/writers problem

import
  "µU/sem"
type
  semaphore struct {
                nR uint
             m, rw sem.Semaphore
                   }

func newS() ReaderWriter {
  x := new(semaphore)
  x.m = sem.New (1)
  x.rw = sem.New (1)
  return x
}

func (x *semaphore) ReaderIn() {
  x.m.P()
  x.nR++
  if x.nR == 1 {
    x.rw.P()
  }
  x.m.V()
}

func (x *semaphore) ReaderOut() {
  x.m.P()
  x.nR--
  if x.nR == 0 {
    x.rw.V()
  }
  x.m.V()
}

func (x *semaphore) WriterIn() {
  x.rw.P()
}

func (x *semaphore) WriterOut() {
  x.rw.V()
}

func (x *semaphore) Fin() {
}
