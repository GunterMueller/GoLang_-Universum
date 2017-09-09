package rw

// (c) Christian Maurer   v. 170411 - license see murus.go

// >>> readers/writers problem: implementation with critical resources

import
  "murus/cr"
type
  criticalResource struct {
                          cr.CriticalResource
                          }

func newCR() ReaderWriter {
  const nc = 2
  x := &criticalResource { cr.New (nc, 1) }
  m := make ([][]uint, nc)
  for i := uint(0); i < nc; i++ { m[i] = make ([]uint, 1) }
  m[reader][0], m[writer][0] = 100, 1
  x.Limit (m)
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
