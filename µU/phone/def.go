package phone

// (c) Christian Maurer   v. 220831 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  Width=16 // maximal number of digits of a phonenumber
type
  PhoneNumber interface {

  Editor
  col.Colourer
  Stringer
  TeXer
  Printer
}

// Returns a new empty phone number.
func New() PhoneNumber { return new_() }
