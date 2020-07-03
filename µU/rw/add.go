package rw

// (c) Christian Maurer   v. 200103 - license see µU.go

// >>> 1st readers/writers problem

import
  "µU/asem"
type
  addS struct {
              uint "maximal number of concurrent readers"
              asem.AddSemaphore
              }

func newAdd (m uint) ReaderWriter {
  x := new(addSemaphore)
  x.uint = m
  x.AddSemaphore = asem.NewNaive(x.uint)
  return x
}

func (x *addS) ReaderIn() {
  x.AddSemaphore.P (1)
}

func (x *addS) ReaderOut() {
  x.AddSemaphore.V (1)
}

func (x *addS) WriterIn() {
  x.AddSemaphore.P (x.uint)
}

func (x *addS) WriterOut() {
  x.AddSemaphore.V (x.uint)
}

func (x *addS) Fin() {
}
