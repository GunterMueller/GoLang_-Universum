package addr

// (c) Christian Maurer   v. 221003 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Address interface {

  Editor
  col.Colourer
  Printer
  TeXer

// Pre: y is of type Address.
// Returns true, iff y has the same postcode as x.
  Equiv (y any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
