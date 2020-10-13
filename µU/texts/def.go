package texts

// (c) Christian Maurer   v. 201005 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Texts interface {

  Object
  Editor
  col.Colourer
  Len() []uint
}

func New (n []uint) Texts { return new_(n) }
