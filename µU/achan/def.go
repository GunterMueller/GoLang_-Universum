package achan

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  AsynchronousChannel interface {

// a is contained in x.
  Send (a any)

// Returns the message, that was sent to x; the message is removed from x.
  Recv() any
}

// Returns an new empty asynchronous channel.
func New (a any) AsynchronousChannel { return new_(a) }
