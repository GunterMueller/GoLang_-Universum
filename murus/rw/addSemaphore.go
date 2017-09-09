package rw

// (c) Christian Maurer   v. 170411 - license see murus.go

// >>> 1st readers/writers problem with additive semaphores
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  "murus/asem"
type
  addSemaphore struct {
                      uint
                      asem.AddSemaphore
                      }

func newASem (m uint) ReaderWriter {
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
