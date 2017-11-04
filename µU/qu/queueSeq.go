package qu

// (c) Christian Maurer   v. 171104 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  squeue struct {
                Any
                seq.Sequence
                }

func newS (a Any) Queue {
  CheckAtomicOrObject (a)
  x := new(squeue)
  x.Any = Clone (a)
  x.Sequence = seq.New(a)
  return x
}

func (x *squeue) Ins (a Any) {
  x.Seek(x.Num())
  x.Sequence.Ins(a)
}

func (x *squeue) Get() Any {
  if x.Empty() {
    return x.Any
  }
  x.Seek(0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
