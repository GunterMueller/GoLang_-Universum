package schan

// (c) Christian Maurer   v. 171104 - license see µU.go

import (
  "sync"
  . "µU/obj"
)
type
  synchronousChannel struct {
             pattern Any
                     Any "object in channel"
                     bool "first at rendezvous"
               mutex,
    s, r, rendezvous sync.Mutex
                     }

func new_(a Any) SynchronousChannel {
  x := new(synchronousChannel)
  x.pattern = Clone(a)
  x.Any = nil // Clone(a)
  x.bool = true
  x.rendezvous.Lock()
  return x
}

func (x *synchronousChannel) Send (a Any) {
  x.s.Lock()
  x.mutex.Lock()
  x.Any = Clone(a)
  if x.bool {
    x.bool = false
    x.mutex.Unlock()
    x.rendezvous.Lock()
    x.mutex.Unlock()
  } else {
    x.bool = true
    x.rendezvous.Unlock()
  }
  x.s.Unlock()
}

func (x *synchronousChannel) Recv() Any {
  var a Any
  x.r.Lock()
  x.mutex.Lock()
  if x.bool {
    x.bool = false
    x.mutex.Unlock()
    x.rendezvous.Lock()
    a = Clone(x.Any)
    x.Any = x.pattern
    x.mutex.Unlock()
  } else {
    x.bool = true
    a = Clone (x.Any)
    x.Any = x.pattern
    x.mutex.Unlock()
    x.rendezvous.Unlock()
  }
  x.r.Unlock()
  return a
}
