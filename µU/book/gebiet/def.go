package gebiet

// (c) Christian Maurer   v. 210509 - license see µU.go

import
  . "µU/obj"

type
  Gebiet interface {
    Object
    Editor
    Stringer
    TeXer
  }

func New() Gebiet { return new_() }
