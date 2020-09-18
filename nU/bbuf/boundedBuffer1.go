package bbuf

// (c) Christian Maurer   v. 200908 - license see nU.go

import (. "nU/obj"; "nU/buf")

type boundedBuffer1 struct {
  Any "Musterobjekt"
  int "Anzahl der Objekte im Puffer"
  cap uint
  buf.Buffer
}

func new1(a Any, n uint) BoundedBuffer {
  x := new(boundedBuffer1)
  x.Any = Clone(a)
  x.cap = n
  x.Buffer = buf.New (a)
  return x
}

func (x *boundedBuffer1) Full() bool {
  return x.Num() == int(x.cap - 1)
}
