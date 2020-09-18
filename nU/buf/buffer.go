package buf

// (c) Christian Maurer   v. 200908 - license see nU.go

import . "nU/obj"

type buffer struct {
  Any "Musterobjekt"
  s []Any
}

func new_(a Any) Buffer {
  x := new(buffer)
  x.Any = Clone(a)
  x.s = make([]Any, 0)
  return x
}

func (x *buffer) Empty() bool {
  return len(x.s) == 0
}

func (x *buffer) Num() int {
  return len(x.s)
}

func (x *buffer) Ins (a Any) {
  x.s = append(x.s, a)
}

func (x *buffer) Get() Any {
  if x.Empty() {
    return x.Any
  }
  a := x.s[0]
  x.s = x.s[1:]
  return a
}
