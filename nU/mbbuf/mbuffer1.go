package mbbuf

// (c) Christian Maurer   v. 171106 - license see nU.go

import (. "nU/obj"; "nU/bbuf"; "nU/sem")

type mBuffer1 struct {
  bbuf.BoundedBuffer
  notEmpty, notFull, ins, get sem.Semaphore
}

func new1 (a Any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mBuffer1)
  x.BoundedBuffer = bbuf.New (a, n)
  x.notEmpty, x.notFull = sem.New(0), sem.New(n)
  x.ins, x.get = sem.New(1), sem.New(1)
  return x
}

func (x *mBuffer1) Ins (a Any) {
  x.notFull.P()
  x.ins.P()
  x.BoundedBuffer.Ins (a)
  x.ins.V()
  x.notEmpty.V()
}

func (x *mBuffer1) Get() Any {
  x.notEmpty.P()
  x.get.P()
  a := x.BoundedBuffer.Get()
  x.get.V()
  x.notFull.V()
  return a
}
