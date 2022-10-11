package achan

// (c) Christian Maurer   v. 220801 - license see nU.go

type AsynchronousChannel interface {

  Send (a any)
  Recv() any
}

// Liefert einen neuen asynchronen Kanal.
func New (a any) AsynchronousChannel { return new_(a) }
