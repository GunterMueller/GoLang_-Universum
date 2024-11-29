package pair

import
  . "µU/obj"
type
  Pair interface {

  Object

  Uint() uint
  Bool() bool

  Set (u uint, b bool)
}

func New() Pair { return new_() }
