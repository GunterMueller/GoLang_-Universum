package stk

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/seq"
)
type
  stack struct {
               seq.Sequence
               }

func newStk (a Any) Stack {
  return &stack { seq.New(a) }
}

func (x *stack) Push (a Any) {
  x.Sequence.Seek (0)
  x.Sequence.Ins (a)
}

/*
func (x *stack) Empty() bool {
  return x.Sequence.Empty ()
}
*/

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
