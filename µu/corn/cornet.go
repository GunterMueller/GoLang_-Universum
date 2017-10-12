package corn

// (c) Christian Maurer   v. 170318 - license see µu.go

import (
  . "µu/obj"
	"µu/rand"
	"µu/seq"
)
type
  cornet struct {
                seq.Sequence
                }

func new_(a Any) Cornet {
  return &cornet { seq.New(a) }
}

func (x *cornet) Ins (a Any) {
  x.Seek(x.Num())
	x.Sequence.Ins(a)
}

func (x *cornet) Get() Any {
  n := x.Num()
  switch n {
  case 0:
    return nil
  case 1:
    x.Seek(0)
  default:
    x.Seek(rand.Natural(x.Num() - 1) + 1)
  }
  defer x.Sequence.Del()
  return x.Sequence.Get()
}
