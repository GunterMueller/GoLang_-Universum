package bbuf

// (c) Christian Maurer   v. 171106 - license see nU.go

import . "nU/obj"

type boundedBuffer0 struct {
  Any
  num int
  cap, in, out uint
  content AnyStream
}

func new0(a Any, n uint) BoundedBuffer {
  x := new(boundedBuffer0)
  x.Any = Clone(a)
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

func (x *boundedBuffer0) Ins (a Any) {
  if x.Full() { return }
  CheckTypeEq (a, x.Any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.num++
}

func (x *boundedBuffer0) Get() Any {
  if x.Empty() {
    return x.Any
  }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.Any)
  x.out = (x.out + 1) % x.cap
  x.num--
  return a
}
