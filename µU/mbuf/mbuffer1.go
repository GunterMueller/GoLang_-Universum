package mbuf

// (c) Christian Maurer   v. 171103 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/buf"
)
type
  mbuffer1 struct {
                  buf.Buffer
         notEmpty,
          notFull,
         ins, get sync.Mutex
                  }

func new1 (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil }
  x := new(mbuffer1)
  x.Buffer = buf.New (a, n)
  x.notEmpty.Lock()
  return x
}

func (x *mbuffer1) Ins (a Any) {
  x.notFull.Lock()
  defer x.notEmpty.Unlock()
  x.ins.Lock()
  defer x.ins.Unlock()
  x.Buffer.Ins (a)
}

func (x *mbuffer1) Get() Any {
  x.notEmpty.Lock()
  defer x.notFull.Unlock()
  x.get.Lock()
  defer x.get.Unlock()
  return x.Buffer.Get()
}
