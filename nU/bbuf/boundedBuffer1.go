package bbuf

// (c) Christian Maurer   v. 180101 - license see nU.go

import (. "nU/obj"; "nU/buf")

type boundedBuffer1 struct {
  Any "Musterobjekt"
  buf.Buffer
  num int
  cap uint
}

func new1(a Any, n uint) BoundedBuffer {
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

func (x *boundedBuffer1) Ins (a Any) {
  if x.Full() { return }
  CheckTypeEq (a, x.Any)
  x.Buffer.Ins (a)
  x.num++
}

func (x *boundedBuffer1) Get() Any {
  if x.Empty() {
    return x.Any
  }
  x.num--
  return x.Buffer.Get()
}
