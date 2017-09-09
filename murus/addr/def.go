package addr

// (c) Christian Maurer   v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Address interface {

  Editor
  Printer

  Equiv (Y Any) bool
}

// Returns a new empty address.
func New() Address { return new_() }
