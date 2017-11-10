package mbbuf

// (c) Christian Maurer   v. 171106 - license see µU.go

// >>> Implementation with a Go-Monitor

import (
  "sync"
  . "µU/obj"
  "µU/bbuf"
)
type
  condition struct {
                   bbuf.BoundedBuffer
 notFull, notEmpty *sync.Cond
                   sync.Mutex
                   }

func newgo (a Any, n uint) MBoundedBuffer {
  if a == nil || n == 0 { return nil }
  x := new (condition)
  x.BoundedBuffer = bbuf.New (a, n)
  x.notFull = sync.NewCond (&x.Mutex)
  x.notEmpty = sync.NewCond (&x.Mutex)
  return x
}

func (x *condition) Ins (a Any) {
  x.Mutex.Lock()
  for x.BoundedBuffer.Full() {
    x.notFull.Wait()
  }
  x.BoundedBuffer.Ins (a)
  x.notEmpty.Signal()
  x.Mutex.Unlock()
}

func (x *condition) Get() Any {
  x.Mutex.Lock()
  for x.BoundedBuffer.Num() == 0 {
    x.notEmpty.Wait()
  }
  defer x.Mutex.Unlock()
  x.notFull.Signal()
  return x.BoundedBuffer.Get()
}
