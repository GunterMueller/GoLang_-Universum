package buf

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/seq")

type bufferSeq struct {
  Any "Musterobjekt"
  seq.Sequence
}

func newS (a Any) Buffer {
  x := new(bufferSeq)
  x.Any = Clone(a)
  x.Sequence = seq.New(a)
  return x
}

func (x *bufferSeq) Ins (a Any) {
  x.InsLast (a)
}

func (x *bufferSeq) Get() Any {
  if x.Empty() {
    return x.Any
  }
  defer x.DelFirst()
  return x.GetFirst()
}
