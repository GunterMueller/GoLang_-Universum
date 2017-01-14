package rw

// (c) murus.org  v. 140330 - license see murus.go

// >>> 1st readers/writers problem (readers' preference)
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  "murus/sem"
type
  semaphore struct {
                nR int
         mutex, rw sem.Semaphore
                   }

func NewSemaphore() ReaderWriter {
  return &semaphore { mutex: sem.New (1), rw: sem.New (1)  }
}

func (x *semaphore) ReaderIn() {
  x.mutex.P()
  x.nR ++
  if x.nR == 1 {
    x.rw.P()
  }
  x.mutex.V()
}

func (x *semaphore) ReaderOut() {
  x.mutex.P()
  x.nR --
  if x.nR == 0 {
    x.rw.V()
  }
  x.mutex.V()
}

func (x *semaphore) WriterIn() {
  x.rw.P()
}

func (x *semaphore) WriterOut() {
  x.rw.V()
}
