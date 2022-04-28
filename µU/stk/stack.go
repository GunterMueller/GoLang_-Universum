package stk

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  stack struct {
               seq.Sequence
               }

func new_(a any) Stack {
  CheckAtomicOrObject (a)
  return &stack { seq.New(a) }
}

func (x *stack) Push (a any) {
  x.Sequence.Seek (0)
  x.Sequence.Ins (a)
}

func (x *stack) Pop() any {
  if x.Sequence.Empty() { return nil }
  x.Sequence.Seek (0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
