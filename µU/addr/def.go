package addr

// (c) Christian Maurer   v. 210511 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Address interface {

  Object
  col.Colourer
  TeXer
  Editor
  Printer

// Pre: y is of type Address.
// Returns true, iff y has the same postcode as x.
  Equiv (y Any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
