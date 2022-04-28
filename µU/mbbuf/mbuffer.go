package mbbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "sync"
  "µU/bbuf"
)
type
  mbuffer struct {
                 bbuf.BoundedBuffer
                 sync.Mutex
                 }

func new_(a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbuffer)
  x.BoundedBuffer = bbuf.New (a, n)
  return x
}

func (x *mbuffer) Ins (a any) {
  x.Mutex.Lock()
  x.BoundedBuffer.Ins (a)
  x.Mutex.Unlock()
}

func (x *mbuffer) Get() any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.BoundedBuffer.Get()
}
