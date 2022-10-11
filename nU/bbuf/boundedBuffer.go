package bbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

type boundedBuffer struct {
  any "Musterobjekt"
  int "Anzahl der Objekte im Puffer"
  cap, in, out uint
  content AnyStream
}

func new_(a any, n uint) BoundedBuffer {
  x := new(boundedBuffer)
  x.any = Clone(a)
  x.cap = n
  x.content = make(AnyStream, x.cap)
  return x
}

func (x *boundedBuffer) Empty() bool {
  return x.int == 0
}

func (x *boundedBuffer) Num() int {
  return x.int
}

func (x *boundedBuffer) Full() bool {
  return x.int == int(x.cap - 1)
}

func (x *boundedBuffer) Ins (a any) {
  if x.Full() { return }
  CheckTypeEq (a, x.any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.int++
}

func (x *boundedBuffer) Get() any {
  if x.Empty() {
    return x.any
  }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.any)
  x.out = (x.out + 1) % x.cap
  x.int--
  return a
}
