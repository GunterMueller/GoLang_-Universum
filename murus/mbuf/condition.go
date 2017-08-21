package mbuf

// (c) murus.org  v. 140330 - license see murus.go

// >>> Implementation with a conditioned monitor

import (
  "sync"
  . "murus/obj"
  "murus/buf"
)
type
  condition struct {
                   buf.Buffer
 notFull, notEmpty *sync.Cond
                   sync.Mutex
                   }

func NewCondition (a Any, n uint) MBuffer {
  x:= new (condition)
  x.Buffer = buf.New (a, n)
  x.notFull = sync.NewCond (&x.Mutex)
  x.notEmpty = sync.NewCond (&x.Mutex)
  return x
}

func (x *condition) Ins (a Any) {
  x.Mutex.Lock()
  for x.Buffer.Full() {
    x.notFull.Wait()
  }
  x.Buffer.Ins (a)
  x.notEmpty.Signal()
  x.Mutex.Unlock()
}

func (x *condition) Get() Any {
  x.Mutex.Lock()
  for x.Buffer.Num() == 0 {
    x.notEmpty.Wait()
  }
  defer x.Mutex.Unlock()
  x.notFull.Signal()
  return x.Buffer.Get()
}
