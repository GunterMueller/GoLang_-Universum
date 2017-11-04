package achan

// (c) Christian Maurer   v. 171104 - license see µU.go

import
  . "µU/obj"
type
  AsynchronousChannel interface {

// a is contained in x.
  Send (a Any)

// Returns the message, that was sent to x; the message is removed from x.
  Recv() Any
}

// Returns an new empty asynchronous channel.
func New (a Any) AsynchronousChannel { return new_(a) }
