package mbuf

// (c) murus.org  v. 140330 - license see murus.go

// >>> buffer implementation with asynchronous message passing

import
  . "murus/obj"
type
  channel struct {
               c chan Any
                 }

func NewChannel (a Any, n uint) MBuffer {
  if a == nil || n == 0 { return nil } // panic
  return &channel { make (chan Any, n) }
}

func (x *channel) Ins (a Any) {
  x.c <- a
}

func (x *channel) Get() Any {
  return Clone (<-x.c)
}
