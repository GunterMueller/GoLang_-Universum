package book

// (c) Christian Maurer   v. 210509 - license see µU.go

import
  . "µU/obj"
type
  Book interface {

  Stringer
  TeXer
  Indexer
  Rotator
}

func New() Book { return new_() }
