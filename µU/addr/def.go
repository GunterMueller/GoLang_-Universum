package addr

// (c) Christian Maurer   v. 210409 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Address interface {

  Object
  col.Colourer
  Stringer
  Editor
  Printer

  Equiv (Y Any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
