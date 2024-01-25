package rw

// (c) Christian Maurer   v. 171125 - license see nU.go

import
  "nU/asem"
type
  addSemaphore struct {
                 uint "Maximalzahl nebenläufiger Leser"
                      asem.AddSemaphore
                      }

func newAS (m uint) ReaderWriter {
  x := new(addSemaphore)
  x.uint = m
  x.AddSemaphore = asem.New(x.uint)
  return x
}

func (x *addSemaphore) ReaderIn() { x.AddSemaphore.P (1) }
func (x *addSemaphore) ReaderOut() { x.AddSemaphore.V (1) }
func (x *addSemaphore) WriterIn() { x.AddSemaphore.P (x.uint) }
func (x *addSemaphore) WriterOut() { x.AddSemaphore.V (x.uint) }
