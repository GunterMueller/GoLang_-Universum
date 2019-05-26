package achan

// (c) Christian Maurer   v. 171104 - license see nU.go

import . "nU/obj"

type AsynchronousChannel interface {

  Send (a Any)

  Recv() Any
}

// Liefert einen neuen asynchronen Kanal.
func New (a Any) AsynchronousChannel { return new_(a) }
