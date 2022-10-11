package bbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

type boundedBuffer0 struct {
  any
  num int
  cap, in, out uint
  content AnyStream
}

func new0(a any, n uint) BoundedBuffer {
  x := new(boundedBuffer0)
  x.any = Clone(a)
  x.cap = n
  x.content = make (AnyStream, x.cap)
  return x
}

func (x *boundedBuffer0) Empty() bool {
  return x.num == 0
}

func (x *boundedBuffer0) Num() int {
  return x.num
}

func (x *boundedBuffer0) Full() bool {
  return x.num == int(x.cap - 1)
}

func (x *boundedBuffer0) Ins (a any) {
  if x.Full() { return }
  CheckTypeEq (a, x.any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.num++
}

func (x *boundedBuffer0) Get() any {
  if x.Empty() {
    return x.any
  }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.any)
  x.out = (x.out + 1) % x.cap
  x.num--
  return a
}
