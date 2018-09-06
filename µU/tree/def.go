package tree

// (c) Christian Maurer   v. 180817 - license see µU.go

import
  . "µU/obj"
type
  Tree interface {

  Object
//  Iterator // TODO
  Root() Any
  // Returns the father of x, if that exists; returns otherwise nil.
  Father() Any
  // Returns the number of the sons of x.
  NumSons() uint
  // Pre: i < NumSons of x.
  Son (i uint) Tree
  Depth() uint
  // Pre: d < Depth of x; n < Num
//  Ins (a Any, d, n uint) // TODO
  // Pre: d < Depth of x.
//  Del (a Any, d, n uint) // TODO
}

// Pre: a is atomic or of a type implementing Object.
// Returns a new empty tree for objects of the type of a.
func New(a Any) Tree { return new_(a) }
