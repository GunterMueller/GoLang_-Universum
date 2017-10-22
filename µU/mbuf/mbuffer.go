package mbuf

// (c) Christian Maurer   v. 170218 - license see µU.go

import (
  "sync"
  "µU/ker"
  . "µU/obj"
  "µU/buf"
)
type
  mbuffer struct {
                 buf.Buffer
        notEmpty,
         notFull,
        ins, get sync.Mutex
                 }

func new_(a Any, n uint) MBuffer {
  if a == nil || n == 0 { ker.Panic ("mbuf.New with param nil or 0") }
  x := new (mbuffer)
  x.Buffer = buf.New (a, n)
  x.notEmpty.Lock()
  return x
}

func (x *mbuffer) Ins (a Any) {
  x.notFull.Lock()
  defer x.notEmpty.Unlock()
  x.ins.Lock()
  defer x.ins.Unlock()
  x.Buffer.Ins (a)
}

func (x *mbuffer) Get() Any {
  x.notEmpty.Lock()
  defer x.notFull.Unlock()
  x.get.Lock()
  defer x.get.Unlock()
  return x.Buffer.Get()
}
