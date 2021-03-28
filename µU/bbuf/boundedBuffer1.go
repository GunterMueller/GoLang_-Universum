package bbuf

// (c) Christian Maurer   v. 210314 - license see µU.go

import (
  . "µU/obj"
  "µU/buf"
)

type
  boundedBuffer1 struct {
                        Any "pattern object"
                        cap uint
                        buf.Buffer
}

func new1 (a Any, n uint) BoundedBuffer {
  x := new(boundedBuffer1)
  x.Any = Clone(a)
  x.cap = n
  x.Buffer = buf.New (a)
  return x
}

func (x *boundedBuffer1) Full() bool {
  return x.Num() == x.cap - 1
}
