package buf

// (c) Christian Maurer   v. 171106 - license see µU.go

import (
  . "µU/obj"
  "µU/seq"
)
type
  bufferSeq struct {
                   any
                   seq.Sequence
                   }

func newS (a any) Buffer {
  CheckAtomicOrObject (a)
  x := new(bufferSeq)
  x.any = Clone (a)
  x.Sequence = seq.New(a)
  return x
}

func (x *bufferSeq) Ins (a any) {
  x.Seek(x.Num())
  x.Sequence.Ins(a)
}

func (x *bufferSeq) Get() any {
  if x.Empty() {
    return x.any
  }
  x.Seek (0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
