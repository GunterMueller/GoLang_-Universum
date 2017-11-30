package mbbuf

// (c) Christian Maurer   v. 171106 - license see nU.go

import ("sync"; . "nU/obj"; "nU/bbuf")

type mbuffer struct {
  bbuf.BoundedBuffer
  sync.Mutex
}

func new_(a Any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbuffer)
  x.BoundedBuffer = bbuf.New (a, n)
  return x
}

func (x *mbuffer) Ins (a Any) {
  x.Mutex.Lock()
  x.BoundedBuffer.Ins (a)
  x.Mutex.Unlock()
}

func (x *mbuffer) Get() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.BoundedBuffer.Get()
}
