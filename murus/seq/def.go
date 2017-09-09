package seq

// (c) Christian Maurer   v. 170424 - license see murus.go

import
  . "murus/obj"
type
  Sequence interface {

  Object
  Sorter
  Iterator
  Seeker

  Reverse() // destroys the order, if x is ordered
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty sequence for objects of the type of a.
func New (a Any) Sequence { return new_(a) }
