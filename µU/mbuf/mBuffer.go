package mbuf

// (c) Christian Maurer   v. 201101 - license see µU.go

import (
  "sync"
  . "µU/obj"
  "µU/buf"
)
type
  mBuffer struct {
                 buf.Buffer
       notEmpty,
          mutex sync.Mutex
                }

func new_(a Any) MBuffer {
  CheckAtomicOrObject(a)
  x := new(mBuffer)
  x.Buffer = buf.New(a)
  x.notEmpty.Lock()
  return x
}

func (x *mBuffer) Ins (a Any) {
  x.mutex.Lock()
  x.Buffer.Ins(a)
  x.mutex.Unlock()
  x.notEmpty.Unlock()
}

func (x *mBuffer) Get() Any {
  x.notEmpty.Lock()
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Buffer.Get()
}
