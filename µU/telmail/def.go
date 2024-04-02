package telmail

// (c) Christian Maurer   v. 221003 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  TelMail interface { // Telephone-Numbers and E-Mail-Address

  Editor
  col.Colourer
  Printer
  TeXer
}

// Returns a new empty TelMail.
func New() TelMail { return new_() }
