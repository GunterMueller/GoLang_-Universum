package mcorn

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "sync"
	"µU/corn"
)
type
  mCornet struct {
                 corn.Cornet
                 sync.Mutex
                 }

func new_(a any) MCornet {
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

func (x *mCornet) Ins (a any) {
  x.Mutex.Lock()
  x.Cornet.Ins (a)
  x.Mutex.Unlock()
}

func (x *mCornet) Get() any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  return x.Cornet.Get()
}
