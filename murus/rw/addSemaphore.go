package rw

// (c) murus.org  v. 140330 - license see murus.go

// >>> 1st readers/writers problem with additive semaphores
//     s. Nichtsequentielle Programmierung mit Go 1 kompakt, S. 75

import
  "murus/asem"
const
  m = 19
type
  addSemaphore struct {
                      asem.AddSemaphore
                      }

func NewAddSemaphore() ReaderWriter {
  return &addSemaphore { asem.New (m) }
}

func (x *addSemaphore) ReaderIn() {
  x.AddSemaphore.P (1)
}

func (x *addSemaphore) ReaderOut() {
  x.AddSemaphore.V (1)
}

func (x *addSemaphore) WriterIn() {
  x.AddSemaphore.P (m)
}

func (x *addSemaphore) WriterOut() {
  x.AddSemaphore.V (m)
}
