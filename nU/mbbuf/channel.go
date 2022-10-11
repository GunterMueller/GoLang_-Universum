package mbbuf

// (c) Christian Maurer   v. 220702 - license see nU.go

// >>> buffer implementation with asynchronous message passing

import . "nU/obj"

type channel struct {
  any
  c chan any
}

func newCh (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(channel)
  x.any = Clone (a)
  x.c = make(chan any, n)
  return x
}

func (x *channel) Ins (a any) {
  x.c <- a
}

func (x *channel) Get() any {
  return Clone (<-x.c)
}
