package pair

// (c) Christian Maurer  v. 230303

import
  . "µU/obj"
type
  Symbol = byte
type
  Pair interface {

  Object

  Set (s Symbol, i uint)
  Get() (Symbol, uint)
}

func New() Pair { return new_() }
