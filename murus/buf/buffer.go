package buf

// (c) murus.org  v. 170621 - license see murus.go

import
  . "murus/obj"
type
  buffer struct {
                Any
       cap, num,
        in, out uint
        content []Any
                }

func new_(a Any, n uint) Buffer {
  if a == nil || n == 0 { return nil } // TODO Panic() ?
  x := new(buffer)
  x.Any = Clone(a)
  x.cap = n
  x.content = make ([]Any, x.cap)
  return x
}

func (x *buffer) Empty() bool {
  return x.num == 0
}

func (x *buffer) Num() uint {
  return x.num
}

func (x *buffer) Full() bool {
  return x.num == x.cap - 1
}

func (x *buffer) Ins (a Any) {
  if x.Full() { return }
  CheckTypeEq (a, x.Any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.num++
}

func (x *buffer) Get() Any {
  if x.Empty() { return nil }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.Any)
  x.out = (x.out + 1) % x.cap
  x.num--
  return a
}
