package phone

// (c) Christian Maurer   v. 161216 - license see murus.go

import
  . "murus/obj"
type
  PhoneNumber interface {

  Editor
  Stringer
  Printer
}

// Returns a new empty phone number.
func New() PhoneNumber { return new_() }
