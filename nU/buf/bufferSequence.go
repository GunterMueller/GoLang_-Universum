package buf

// (c) Christian Maurer   v. 220702 - license see nU.go

import (
  . "nU/obj"
  "nU/seq"
)
type
  bufferSeq struct {
                   any "Musterobjekt"
                   seq.Sequence
                   }

func newS (a any) Buffer {
  x := new(bufferSeq)
  x.any = Clone(a)
  x.Sequence = seq.New(a)
  return x
}

func (x *bufferSeq) Ins (a any) {
  x.InsLast (a)
}

func (x *bufferSeq) Get() any {
  if x.Empty() {
    return x.any
  }
  defer x.DelFirst()
  return x.GetFirst()
}
