package mbuf

// (c) murus.org  v. 140330 - license see murus.go

import (
  "sync"
  . "murus/obj"; "murus/buf"
)
type
  mbuffer1 struct {
                  buf.Buffer
                  sync.Mutex
                  }

func New1 (a Any, n uint) MBuffer {
  if a == nil || n == 0 { panic ("mbuf.New with param nil or 0") }
  x:= new (mbuffer1)
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
