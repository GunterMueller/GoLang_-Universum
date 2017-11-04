package schan

// (c) Christian Maurer   v. 171104 - license see µU.go

import
  . "µU/obj"
type
  SynchronousChannel interface {

// a is contained in x.
  Send (a Any)

// Returns the message, that was sent to x; the message is removed from x.
// The calling process might have been blocked, until x contained a message.
  Recv() Any
}

// Returns an new empty synchronous channel.
func New (a Any) SynchronousChannel { return new_(a) }
