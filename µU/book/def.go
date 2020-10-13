package book

// (c) Christian Maurer   v. 201005 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Book interface {

//  Indexer
  Object
  Editor
  col.Colourer
}

func New() Book { return new_() }
