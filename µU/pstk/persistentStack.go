package pstk

// (c) Christian Maurer   v. 210109 - license see µU.go

import (
  . "µU/obj"
  "µU/pseq"
)
type
  pstack struct {
                pseq.PersistentSequence
                }

func new_(a Any) PersistentStack {
  CheckAtomicOrObject (a)
  return &pstack { pseq.New(a) }
}

func (x *pstack) Name (n string) {
  x.PersistentSequence.Name (n)
}

func (x *pstack) Rename (n string) {
  x.PersistentSequence.Rename (n)
}

func (x *pstack) Push (a Any) {
  x.PersistentSequence.Seek (0)
  x.PersistentSequence.Ins (a)
}

func (x *pstack) Pop() Any {
  if x.PersistentSequence.Empty() { return nil }
  x.PersistentSequence.Seek (0)
  defer x.PersistentSequence.Del()
  return x.PersistentSequence.Get()
}

func (x *pstack) Fin() {
  x.PersistentSequence.Fin()
}
