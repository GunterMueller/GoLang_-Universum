package schan

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  SynchronousChannel interface {

// a is contained in x.
  Send (a any)

// Returns the message, that was sent to x; the message is removed from x.
// The calling process might have been blocked, until x contained a message.
  Recv() any
}

// Returns an new empty synchronous channel.
func New (a any) SynchronousChannel { return new_(a) }
