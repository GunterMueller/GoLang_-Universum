package corn

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
	"µU/rand"
	"µU/seq"
)
type
  cornet struct {
                seq.Sequence
                }

func new_(a any) Cornet {
  return &cornet { seq.New(a) }
}

func (x *cornet) Ins (a any) {
  x.Seek (x.Num())
	x.Sequence.Ins (a)
}

func (x *cornet) Get() any {
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

func (x *cornet) All() []any {
  all := make([]any, 0)
  n := x.Num()
  if n == 0 {
    return all
  }
  for i := uint(0); i < n; i++ {
    x.Sequence.Seek (i)
    all = append (all, x.Sequence.Get().(any))
  }
  return all
}
