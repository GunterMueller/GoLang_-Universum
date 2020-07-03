package mcorn

// (c) Christian Maurer   v. 200120

import (
  "sync"
  . "µU/obj"
	"µU/corn"
)
type
  mCornet struct {
                 corn.Cornet
                 sync.Mutex
                 }

func new_(a Any) MCornet {
  if a == nil { return nil }
  x := new(mCornet)
  x.Cornet = corn.New(a)
  return x
}

func (x *mCornet) Empty() bool {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Cornet.Empty()
}

func (x *mCornet) Num() uint {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Cornet.Num()
}

func (x *mCornet) Ins (a Any) {
  x.Mutex.Lock()
  x.Cornet.Ins (a)
  x.Mutex.Unlock()
}

func (x *mCornet) Get() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Cornet.Get()
}
