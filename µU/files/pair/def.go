package pair

// (c) Christian Maurer   v. 210419 - license see µU.go

import
  . "µU/obj"
type
  Pair interface {

  Object

  Name () string
  Set (s string, b byte)
  Typ () byte
}

func New() Pair { return new_() }
