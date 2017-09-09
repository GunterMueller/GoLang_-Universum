package mbuf

// (c) Christian Maurer   v. 170218 - license see murus.go

// >>> buffer implementation with asynchronous message passing

import (
  "murus/ker"
  . "murus/obj"
)
type
  channel struct {
               c chan Any
                 }

func newch (a Any, n uint) MBuffer {
  if a == nil || n == 0 { ker.Panic ("mbuf.NewCh with param nil or 0") }
  return &channel { make (chan Any, n) }
}

func (x *channel) Ins (a Any) {
  x.c <- a
}

func (x *channel) Get() Any {
  return Clone (<-x.c)
}
