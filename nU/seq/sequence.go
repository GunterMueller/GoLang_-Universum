package seq

// (c) Christian Maurer   v. 220702 - license see nU.go

import
  . "nU/obj"
type
  sequence struct {
                  any "Musterobjekt"
               ss []any
                  }

func new_(a any) Sequence {
  x := new(sequence)
  x.any = Clone(a)
  x.ss = make([]any, 0)
  return x
}

func (x *sequence) Empty() bool {
  return len(x.ss) == 0
}

func (x *sequence) Num() int {
  return len(x.ss)
}

func (x *sequence) InsLast (a any) {
  x.ss = append(x.ss, a)
}

func (x *sequence) DelFirst() {
  x.ss = x.ss[1:]
}

func (x *sequence) GetFirst() any {
  if x.Empty() {
    return x.any
  }
  return x.ss[0]
}
