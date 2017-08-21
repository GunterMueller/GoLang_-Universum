package mqu

// (c) murus.org  v. 161216 - license see murus.go

import (
  "sync"
  . "murus/obj"
  "murus/qu"
)
type
  mQueue struct {
                qu.Queue
       notEmpty,
          mutex sync.Mutex
                }

func new_(a Any) MQueue {
  x := new(mQueue)
  x.Queue = qu.New(a)
  x.notEmpty.Lock()
  return x
}

func (x *mQueue) Ins (a Any) {
  x.mutex.Lock()
  x.Queue.Ins(a)
  x.mutex.Unlock()
  x.notEmpty.Unlock()
}

func (x *mQueue) Get() Any {
  x.notEmpty.Lock()
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Queue.Get()
}
