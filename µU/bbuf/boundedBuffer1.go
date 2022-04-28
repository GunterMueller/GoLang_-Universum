package bbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
)

type
  boundedBuffer1 struct {
                        any "pattern object"
                        cap uint
                        buf.Buffer
}

func new1 (a any, n uint) BoundedBuffer {
  x := new(boundedBuffer1)
  x.any = Clone(a)
  x.cap = n
  x.Buffer = buf.New (a)
  return x
}

func (x *boundedBuffer1) Full() bool {
  return x.Num() == x.cap - 1
}
