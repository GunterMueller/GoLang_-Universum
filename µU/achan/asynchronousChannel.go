package achan

// (c) Christian Maurer   v. 171104 - license see µU.go

// >>> simulation of asynchronous message passing sich synchronous message passing

import (
  "sync"
  . "µU/obj"
  "µU/qu"
)
type
  asynchronousChannel struct {
                             Any
                             qu.Queue
                          ch chan Any
                             sync.Mutex
                             }

func new_(a Any) AsynchronousChannel {
  x := new(asynchronousChannel)
  x.Queue = qu.New(a)
  x.ch = make(chan Any) // synchronous !
  return x
}

func (x *asynchronousChannel) Send (a Any) {
  x.Mutex.Lock()
  x.Queue.Ins (a)
  x.Mutex.Unlock()
}

func (x *asynchronousChannel) Recv() Any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  a := x.Queue.Get()
  if a == x.Any { panic("fatal error: alle goroutines are asleep - deadlock!") }
  return a
}
