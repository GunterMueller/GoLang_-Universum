package book

// (c) Christian Maurer   v. 210524 - license see µU.go

import
  . "µU/obj"
type
  Book interface {

  TeXer
  Indexer
  Rotator

  Sub (y Any) bool
}

func New() Book { return new_() }
