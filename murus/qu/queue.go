package qu

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/seq"
)
type
  queue struct {
               seq.Sequence
               }

func newQu (a Any) Queue {
  return &queue { seq.New(a) }
}

func (x *queue) Ins (a Any) {
  x.Seek(x.Num ())
  x.Sequence.Ins(a)
}

func (x *queue) Get () Any {
  if x.Empty() { return nil }
  x.Seek(0)
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
