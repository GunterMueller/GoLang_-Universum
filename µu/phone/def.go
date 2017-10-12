package phone

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  PhoneNumber interface {

  Object
  col.Colourer
  Editor
  Stringer
  Printer
}

// Returns a new empty phone number.
func New() PhoneNumber { return new_() }
