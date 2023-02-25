package rw

// (c) Christian Maurer   v. 200421 - license see µU.go

// >>> bounded readers/writers problem

import
  "µU/cr"
type
  criticalResource struct {
                          cr.CriticalResource
                          }

func newCR (m uint) ReaderWriter {
  const nc = 2
  x := &criticalResource { cr.New (nc, 1) }
  max := make([][]uint, nc)
  for i := uint(0); i < nc; i++ { max[i] = make([]uint, 1) }
  max[reader][0], max[writer][0] = m, 1
  x.Limit (max)
  return x
}

func (x *criticalResource) ReaderIn() {
  x.Enter (reader)
}

func (x *criticalResource) ReaderOut() {
  x.Leave (reader)
}

func (x *criticalResource) WriterIn() {
  x.Enter (writer)
}

func (x *criticalResource) WriterOut() {
  x.Leave (writer)
}

func (x *criticalResource) Fin() {
}
