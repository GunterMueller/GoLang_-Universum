package rw

// (c) Christian Maurer   v. 171016 - license see µU.go

// >>> 1st readers/writers problem

import
  "µU/asem"
type
  addSemaphore struct {
                      uint "maximal number of concurrent readers"
                      asem.AddSemaphore
                      }

func newAS (m uint) ReaderWriter {
  x := new(addSemaphore)
  x.uint = m
  x.AddSemaphore = asem.New(x.uint)
  return x
}

func (x *addSemaphore) ReaderIn() {
  x.AddSemaphore.P (1)
}

func (x *addSemaphore) ReaderOut() {
  x.AddSemaphore.V (1)
}

func (x *addSemaphore) WriterIn() {
  x.AddSemaphore.P (x.uint)
}

func (x *addSemaphore) WriterOut() {
  x.AddSemaphore.V (x.uint)
}

func (x *addSemaphore) Fin() {
}
