package mbbuf

// (c) Christian Maurer   v. 240929 - license see µU.go

import (
  "sync"
  "µU/bbuf"
)
type
  mbbuffer struct {
                  bbuf.BoundedBuffer
         notEmpty,
          notFull sync.Mutex
                  }

func new_(a any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbbuffer)
  x.BoundedBuffer = bbuf.New (a, n)
  x.notEmpty.Lock()
  return x
}

func (x *mbbuffer) Ins (a any) {
  x.notFull.Lock()
  defer x.notEmpty.Unlock()
  x.BoundedBuffer.Ins (a)
}

func (x *mbbuffer) Get() any {
  x.notEmpty.Lock()
  defer x.notFull.Unlock()
  return x.BoundedBuffer.Get()
}
