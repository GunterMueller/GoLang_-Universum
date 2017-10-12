package addr

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  Address interface {

  Object
  col.Colourer
  Editor
  Printer

  Equiv (Y Any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
