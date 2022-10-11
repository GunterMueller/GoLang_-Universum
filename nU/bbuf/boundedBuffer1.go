package bbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/buf")

type boundedBuffer1 struct {
  any "Musterobjekt"
  buf.Buffer
  num int
  cap uint
}

func new1(a any, n uint) BoundedBuffer {
  x := new(boundedBuffer1)
  x.Buffer = buf.New (a)
  x.cap = n
  return x
}

func (x *boundedBuffer1) Empty() bool {
  return x.Buffer.Empty()
}

func (x *boundedBuffer1) Num() int {
  return x.num
}

func (x *boundedBuffer1) Full() bool {
  return x.num == int(x.cap - 1)
}

func (x *boundedBuffer1) Ins (a any) {
  if x.Full() { return }
  CheckTypeEq (a, x.any)
  x.Buffer.Ins (a)
  x.num++
}

func (x *boundedBuffer1) Get() any {
  if x.Empty() {
    return x.any
  }
  x.num--
  return x.Buffer.Get()
}
