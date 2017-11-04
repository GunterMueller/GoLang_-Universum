package mbuf

// (c) Christian Maurer   v. 171103 - license see µU.go

// >>> buffer implementation with asynchronous message passing

import
  . "µU/obj"
type
  channel struct {
                 Any
               c chan Any
                 }

func newCh (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil }
  x := new(channel)
  x.Any = Clone (a)
  x.c = make(chan Any, n)
  return x
}

func (x *channel) Ins (a Any) {
  x.c <- a
}

func (x *channel) Get() Any {
  return Clone (<-x.c)
}
