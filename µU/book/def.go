package book

// (c) Christian Maurer   v. 220420 - license see µU.go

import
  . "µU/obj"
type
  Book interface {

  Stringer
  TeXer
  Indexer
  Rotator

// Pre: y is of type Book.
// Returns true, iff x is a part of y.
  Sub (y any) bool
}

func New() Book { return new_() }
