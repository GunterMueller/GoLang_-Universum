package bbuf

// (c) Christian Maurer   v. 171106 - license see nU.go

import . "nU/obj"

type boundedBuffer struct {
  Any
  num int
  cap, in, out uint
  content AnyStream
}

func new_(a Any, n uint) BoundedBuffer {
  x := new(boundedBuffer)
  x.Any = Clone(a)
  x.cap = n
  x.content = make (AnyStream, x.cap)
  return x
}

func (x *boundedBuffer) Empty() bool {
  return x.num == 0
}

func (x *boundedBuffer) Num() int {
  return x.num
}

func (x *boundedBuffer) Full() bool {
  return x.num == int(x.cap - 1)
}

func (x *boundedBuffer) Ins (a Any) {
  if x.Full() { return }
  CheckTypeEq (a, x.Any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.num++
}

func (x *boundedBuffer) Get() Any {
  if x.Empty() {
    return x.Any
  }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.Any)
  x.out = (x.out + 1) % x.cap
  x.num--
  return a
}
