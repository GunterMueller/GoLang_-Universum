package mbuf

// (c) Christian Maurer   v. 171103 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/buf"
)
type
  mbuffer struct {
                 buf.Buffer
                 sync.Mutex
                 }

func new_(a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbuffer)
  x.Buffer = buf.New (a, n)
  return x
}

func (x *mbuffer) Ins (a Any) {
  x.Mutex.Lock()
  x.Buffer.Ins (a)
  x.Mutex.Unlock()
}

func (x *mbuffer) Get() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Buffer.Get()
}
