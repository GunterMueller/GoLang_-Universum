package texts

// (c) Christian Maurer   v. 220831 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Texts interface {

  Editor
  col.Colourer
  Len() []uint
}

func New (n []uint) Texts { return new_(n) }
