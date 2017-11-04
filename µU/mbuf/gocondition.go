package mbuf

// (c) Christian Maurer   v. 171103 - license see µU.go

// >>> Implementation with a Go-Monitor

import (
  "sync"
  . "µU/obj"
  "µU/buf"
)
type
  condition struct {
                   buf.Buffer
 notFull, notEmpty *sync.Cond
                   sync.Mutex
                   }

func newgo (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil }
  x := new (condition)
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
