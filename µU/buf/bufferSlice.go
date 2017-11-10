package buf

// (c) Christian Maurer   v. 171106 - license see µU.go

import
  . "µU/obj"
type
  buffer struct {
                Any "pattern object of the buffer"
              s []Any
                }

func new_(a Any) Buffer {
  if a == nil { return nil }
  x := new(buffer)
  x.Any = Clone(a)
  x.s = make([]Any, 0)
  return x
}

func (x *buffer) Empty() bool {
  return len(x.s) == 0
}

func (x *buffer) Num() uint {
  return uint(len(x.s))
}

func (x *buffer) Ins (a Any) {
  CheckTypeEq (a, x.Any)
  x.s = append (x.s, a)
}

func (x *buffer) Get() Any {
  if x.Empty() {
    return x.Any
  }
  a := x.s[0]
  x.s = x.s[1:]
  return a
}
