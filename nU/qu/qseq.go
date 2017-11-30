package qu

// (c) Christian Maurer   v. 171125 - license see nU.go

import (. "nU/obj"; "nU/seq")

type queueSeq struct {
  Any "Musterobjekt"
  seq.Sequence
}

func news (a Any) Queue {
  x := new(queueSeq)
  x.Any = Clone(a)
  x.Sequence = seq.New(a)
  return x
}

func (x *queueSeq) Enqueue (a Any) {
  x.InsLast (Clone(a))
}

func (x *queueSeq) Num() int {
  return x.Num()
}

func (x *queueSeq) Dequeue() Any {
  if x.Empty() {
    return x.Any
  }
  defer x.DelFirst()
  return x.GetFirst()
}
