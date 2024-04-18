package addr

// (c) Christian Maurer   v. 240407 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  N = 4 // number of components of an address
type
  Address interface {

  Editor
  Stringer
  col.Colourer
  Printer
  TeXer

// Pre: y is of type Address.
// Returns true, iff y has the same postcode as x.
  Equiv (y any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
