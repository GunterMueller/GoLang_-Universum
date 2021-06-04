package book

// (c) Christian Maurer   v. 210525 - license see µU.go

import
  . "µU/obj"
type
  Book interface {

  TeXer
  Indexer
  Rotator

// Pre: y is of type Book.
// Returns true, iff x is a part of y.
  Sub (y Any) bool
}

func New() Book { return new_() }
