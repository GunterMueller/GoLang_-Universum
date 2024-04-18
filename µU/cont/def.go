package cont

// (c) Christian Maurer   v. 240408 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  N = 3 // number of components of a TelMail
type
  Contact interface { // telephone-Number, cell-nummber and E-mail-address

  Editor
  Stringer
  col.Colourer
  Printer
  TeXer
}

// Returns a new empty Contact.
func New() Contact { return new_() }
