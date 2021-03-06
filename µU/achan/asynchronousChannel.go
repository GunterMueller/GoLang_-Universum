package achan

// (c) Christian Maurer   v. 171106 - license see µU.go

// >>> simulation of asynchronous message passing sich synchronous message passing

import (
  "sync"
  . "µU/obj"
  "µU/buf"
)
type
  asynchronousChannel struct {
                             Any
                             buf.Buffer
                          ch chan Any
                             sync.Mutex
                             }

func new_(a Any) AsynchronousChannel {
  x := new(asynchronousChannel)
  x.Buffer = buf.New(a)
  x.ch = make(chan Any) // synchronous !
  return x
}

func (x *asynchronousChannel) Send (a Any) {
  x.Mutex.Lock()
  x.Buffer.Ins (a)
  x.Mutex.Unlock()
}

func (x *asynchronousChannel) Recv() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  a := x.Buffer.Get()
  if a == x.Any { panic("fatal error: alle goroutines are asleep - deadlock!") }
  return a
}
