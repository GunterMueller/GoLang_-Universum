package phone

// (c) Christian Maurer   v. 210511 - license see µU.go

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
  TeXer
  Printer
}

// Returns a new empty phone number.
func New() PhoneNumber { return new_() }
