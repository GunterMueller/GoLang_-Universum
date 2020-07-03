package buf

// (c) Christian Maurer   v. 171106 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  bufferSeq struct {
                   Any
                   seq.Sequence
                   }

func newS (a Any) Buffer {
  CheckAtomicOrObject (a)
  x := new(bufferSeq)
  x.Any = Clone (a)
  x.Sequence = seq.New(a)
  return x
}

func (x *bufferSeq) Ins (a Any) {
  x.Seek(x.Num())
  x.Sequence.Ins(a)
}

func (x *bufferSeq) Get() Any {
  if x.Empty() {
    return x.Any
  }
  x.Seek (0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
