package phone

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
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
