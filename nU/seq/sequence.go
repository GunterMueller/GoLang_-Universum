package seq

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type sequence struct {
  Any "Musterobjekt"
  ss []Any
}

func new_(a Any) Sequence {
  x := new(sequence)
  x.Any = Clone(a)
  x.ss = make([]Any, 0)
  return x
}

func (x *sequence) Empty() bool {
  return len(x.ss) == 0
}

func (x *sequence) Num() int {
  return len(x.ss)
}

func (x *sequence) InsLast (a Any) {
  x.ss = append(x.ss, a)
}

func (x *sequence) DelFirst() {
  x.ss = x.ss[1:]
}

func (x *sequence) GetFirst() Any {
  if x.Empty() {
    return x.Any
  }
  return x.ss[0]
}
