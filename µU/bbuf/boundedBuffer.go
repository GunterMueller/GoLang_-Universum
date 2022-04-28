package bbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  boundedBuffer struct {
                       any "pattern object"
     cap, num, in, out uint
               content AnyStream
                       }

func new_(a any, n uint) BoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(boundedBuffer)
  x.any = Clone(a)
  x.cap = n
  x.content = make (AnyStream, x.cap)
  return x
}

func (x *boundedBuffer) Empty() bool {
  return x.num == 0
}

func (x *boundedBuffer) Num() uint {
  return x.num
}

func (x *boundedBuffer) Full() bool {
  return x.num == x.cap - 1
}

func (x *boundedBuffer) Ins (a any) {
  if x.Full() { return }
  CheckTypeEq (a, x.any)
  x.content[x.in] = Clone (a)
  x.in = (x.in + 1) % x.cap
  x.num++
}

func (x *boundedBuffer) Get() any {
  if x.Empty() {
    return x.any
  }
  a := Clone (x.content[x.out])
  x.content[x.out] = Clone (x.any)
  x.out = (x.out + 1) % x.cap
  x.num--
  return a
}
