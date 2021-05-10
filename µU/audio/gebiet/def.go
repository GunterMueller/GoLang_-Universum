package gebiet

// (c) Christian Maurer   v. 210510 - license see µU.go

import
  . "µU/obj"

type
  Gebiet interface {
    Object
    Editor
    TeXer
  }

func New() Gebiet { return new_() }
