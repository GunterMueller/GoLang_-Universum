package seq

// (c) Christian Maurer   v. 201011 - license see µU.go

import
  . "µU/obj"
type
  Sequence interface {

  Object
  Sorter
  Seeker

  Reverse() // destroys the order, if x is ordered
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty sequence for objects of the type of a.
func New (a Any) Sequence { return new_(a) }
