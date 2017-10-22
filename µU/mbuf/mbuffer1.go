package mbuf

// (c) Christian Maurer   v. 170218 - license see µU.go

import (
  "sync"
  "µU/ker"
  . "µU/obj"
  "µU/buf"
)
type
  mbuffer1 struct {
                  buf.Buffer
                  sync.Mutex
                  }

func new1 (a Any, n uint) MBuffer {
  if a == nil || n == 0 { ker.Panic ("mbuf.New1 with param nil or 0") }
  x := new (mbuffer1)
  x.Buffer = buf.New (a, n)
  return x
}

func (x *mbuffer1) Ins (a Any) {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  x.Buffer.Ins (a)
}

func (x *mbuffer1) Get() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Buffer.Get()
}
