package mbbuf

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "sync"
  "µU/bbuf"
)
type
  mbuffer1 struct {
                  bbuf.BoundedBuffer
         notEmpty,
          notFull,
         ins, get sync.Mutex
                  }

func new1 (a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbuffer1)
  x.BoundedBuffer = bbuf.New (a, n)
  x.notEmpty.Lock()
  return x
}

func (x *mbuffer1) Ins (a any) {
  x.notFull.Lock()
  defer x.notEmpty.Unlock()
  x.ins.Lock()
  defer x.ins.Unlock()
  x.BoundedBuffer.Ins (a)
}

func (x *mbuffer1) Get() any {
  x.notEmpty.Lock()
  defer x.notFull.Unlock()
  x.get.Lock()
  defer x.get.Unlock()
  return x.BoundedBuffer.Get()
}
