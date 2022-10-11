package achan

// (c) Christian Maurer   v. 220801 - license see nU.go

import ("sync"; "nU/buf")

type asynchronousChannel struct {
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
  if a == x.any { panic("fatal error: all goroutines are asleep - deadlock!") }
  return a
}
