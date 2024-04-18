package telmail

// (c) Christian Maurer   v. 240407 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  N = 3 // number of components of a TelMail
type
  TelMail interface { // Telephone-Numbers and E-Mail-Address

  Editor
  Stringer
  col.Colourer
  Printer
  TeXer
}

// Returns a new empty TelMail.
func New() TelMail { return new_() }
