package buf

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  buffer struct {
                any "pattern object"
              s []any
                }

func new_(a any) Buffer {
  if a == nil { return nil }
  x := new(buffer)
  x.any = Clone(a)
  x.s = make([]any, 0)
  return x
}

func (x *buffer) Empty() bool {
  return len(x.s) == 0
}

func (x *buffer) Num() uint {
  return uint(len(x.s))
}

func (x *buffer) Ins (a any) {
  CheckTypeEq (a, x.any)
  x.s = append (x.s, a)
}

func (x *buffer) Get() any {
  if x.Empty() {
    return x.any
  }
  a := x.s[0]
  x.s = x.s[1:]
  return a
}
