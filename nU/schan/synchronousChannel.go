package schan

// (c) Christian Maurer   v. 220801 - license see nU.go

import ("sync"; . "nU/obj")

type synchronousChannel struct {
  pattern any
  any "object in channel"
  bool "first at rendezvous"
  mutex, s, r, rendezvous sync.Mutex
}

func new_(a any) SynchronousChannel {
  x := new(synchronousChannel)
  x.pattern = Clone(a)
  x.any = nil // Clone(a)
  x.bool = true
  x.rendezvous.Lock()
  return x
}

func (x *synchronousChannel) Send (a any) {
  x.s.Lock()
  x.mutex.Lock()
  x.any = Clone(a)
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

func (x *synchronousChannel) Recv() any {
  var a any
  x.r.Lock()
  x.mutex.Lock()
  if x.bool {
    x.bool = false
    x.mutex.Unlock()
    x.rendezvous.Lock()
    a = Clone(x.any)
    x.any = x.pattern
    x.mutex.Unlock()
  } else {
    x.bool = true
    a = Clone (x.any)
    x.any = x.pattern
    x.mutex.Unlock()
    x.rendezvous.Unlock()
  }
  x.r.Unlock()
  return a
}
