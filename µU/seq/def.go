package seq

// (c) Christian Maurer   v. 221118 - license see µU.go

import
  . "µU/obj"
type
  Sequence interface {

//  Object
  Clearer
  Equaler
//  Comparer
  Coder
  Seeker // hence Collector, hence Clearer
  Predicator

// Pre: x is not ordered.
// The order of the elements in x is reversed.
  Reverse()

// Pre: x is not ordered.
// If x contains at most one element, nothing has happened.
// Otherwise, for b == true, the former last element of x is now the first,
// for b == false, the former first element is now the last.
// The order of the other elements has not changed.
  Rotate (b bool)
}

// Pre: a is atomic or of a type implementing Object.
// If x contains at most one element, nothing has happened.
// Returns otherwise a new empty sequence with pattern object a,
// i.e., for objects of the type of a.
func New (a any) Sequence { return new_(a) }
