package stk

// (c) Christian Maurer   v. 201103 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  stack struct {
               seq.Sequence
               }

func new_(a Any) Stack {
  CheckAtomicOrObject (a)
  return &stack { seq.New(a) }
}

func (x *stack) Push (a Any) {
  x.Sequence.Seek (0)
  x.Sequence.Ins (a)
}

func (x *stack) Pop() Any {
  if x.Sequence.Empty() { return nil }
  x.Sequence.Seek (0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
