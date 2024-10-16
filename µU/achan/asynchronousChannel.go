package achan

// (c) Christian Maurer   v. 220420 - license see µU.go

// >>> simulation of asynchronous message passing sich synchronous message passing

import (
  "sync"
  "µU/buf"
)
type
  asynchronousChannel struct {
                             any
                             buf.Buffer
                          ch chan any
                             sync.Mutex
                             }

func new_(a any) AsynchronousChannel {
  x := new(asynchronousChannel)
  x.Buffer = buf.New(a)
  x.ch = make(chan any) // synchronous !
  return x
}

func (x *asynchronousChannel) Send (a any) {
  x.Mutex.Lock()
  x.Buffer.Ins (a)
  x.Mutex.Unlock()
}

func (x *asynchronousChannel) Recv() any {
  x.Mutex.Lock()
  defer x.Mutex.Unlock()
  a := x.Buffer.Get()
  if a == x.any { panic ("fatal error: all processes are asleep - deadlock!") }
  return a
}
