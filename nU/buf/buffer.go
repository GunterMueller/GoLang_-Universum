package buf

// (c) Christian Maurer   v. 220702 - license see nU.go

import
  . "nU/obj"
type
  buffer struct {
            any "Musterobjekt"
              s []any
                }

func new_(a any) Buffer {
  x := new(buffer)
  x.any = Clone(a)
  x.s = make([]any, 0)
  return x
}

func (x *buffer) Empty() bool {
  return len(x.s) == 0
}

func (x *buffer) Num() int {
  return len(x.s)
}

func (x *buffer) Ins (a any) {
  x.s = append(x.s, a)
}

func (x *buffer) Get() any {
  if x.Empty() {
    return x.any
  }
  a := x.s[0]
  x.s = x.s[1:]
  return a
}
