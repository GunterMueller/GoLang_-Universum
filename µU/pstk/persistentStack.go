package pstk

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/pseq"
)
type
  pstack struct {
                pseq.PersistentSequence
                }

func new_(a any) PersistentStack {
  CheckAtomicOrObject (a)
  return &pstack { pseq.New(a) }
}

func (x *pstack) Name (n string) {
  x.PersistentSequence.Name (n)
}

func (x *pstack) Rename (n string) {
  x.PersistentSequence.Rename (n)
}

func (x *pstack) Push (a any) {
  x.PersistentSequence.Seek (0)
  x.PersistentSequence.Ins (a)
}

func (x *pstack) Pop() any {
  if x.PersistentSequence.Empty() { return nil }
  x.PersistentSequence.Seek (0)
  defer x.PersistentSequence.Del()
  return x.PersistentSequence.Get()
}

func (x *pstack) Fin() {
  x.PersistentSequence.Fin()
}
