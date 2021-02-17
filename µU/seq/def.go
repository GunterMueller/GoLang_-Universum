package seq

// (c) Christian Maurer   v. 210213 - license see µU.go

import
  . "µU/obj"
type
  Sequence interface {

  Object
  Seeker
  Predicator

// Pre: x is not ordered.
  Reverse()
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty sequence with pattern object a,
// i.e. for objects of the type of a.
func New (a Any) Sequence { return new_(a) }
