package field

// (c) Christian Maurer   v. 210512 - license see µU.go

import
  . "µU/obj"

type
  Field interface {
    Object
    Editor
    TeXer
  }

func New() Field { return new_() }
