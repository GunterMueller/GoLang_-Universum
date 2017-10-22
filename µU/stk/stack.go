package stk

// (c) Christian Maurer   v. 170316 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  stack struct {
               seq.Sequence
               }

func new_ (a Any) Stack {
  return &stack { seq.New(a) }
}

func (x *stack) Push (a Any) {
  x.Sequence.Seek (0)
  x.Sequence.Ins (a)
}

func (x *stack) Pop() {
  if x.Sequence.Empty() { return }
  x.Sequence.Seek (0)
  x.Sequence.Del()
}

func (x *stack) Top() Any {
  if x.Sequence.Empty() { return nil }
  x.Sequence.Seek (0)
  return x.Sequence.Get()
}
