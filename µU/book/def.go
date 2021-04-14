package book

// (c) Christian Maurer   v. 210409 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Book interface {

  Object
  Editor
  Stringer
  col.Colourer
  Indexer
  Rotator
}

func New() Book { return new_() }
