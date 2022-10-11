package qu

// (c) Christian Maurer   v. 220801 - license see nU.go

import (. "nU/obj"; "nU/seq")

type queueSeq struct {
  any "Musterobjekt"
  seq.Sequence
}

func news (a any) Queue {
  x := new(queueSeq)
  x.any = Clone(a)
  x.Sequence = seq.New(a)
  return x
}

func (x *queueSeq) Enqueue (a any) {
  x.InsLast (Clone(a))
}

func (x *queueSeq) Num() int {
  return x.Num()
}

func (x *queueSeq) Dequeue() any {
  if x.Empty() {
    return x.any
  }
  defer x.DelFirst()
  return x.GetFirst()
}
