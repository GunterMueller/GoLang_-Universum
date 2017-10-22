package mcorn

// (c) Christian Maurer   v. 170320

import (
  "sync"
  . "µU/obj"
	"µU/corn"
)
type
  mCornet struct {
                  corn.Cornet
        notEmpty,
            mutex sync.Mutex
                 }

func new_(a Any) MCornet {
  if a == nil { return nil }
  x := new(mCornet)
  x.Cornet = corn.New(a)
  x.notEmpty.Lock()
  return x
}

func (x *mCornet) Empty() bool {
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Cornet.Empty()
}

func (x *mCornet) Num() uint {
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Cornet.Num()
}

func (x *mCornet) Ins (a Any) {
  x.mutex.Lock()
  x.Cornet.Ins (a)
  x.mutex.Unlock()
  x.notEmpty.Unlock()
}

func (x *mCornet) Get() Any {
  x.notEmpty.Lock()
  x.mutex.Lock()
  defer x.mutex.Unlock()
  return x.Cornet.Get()
}
