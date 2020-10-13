package tval

// (c) Christian Maurer   v. 200906 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  TruthValue interface { // values "indetermined", "false" and "true"

  Object // Clear sets the value of x to "indetermined".
  col.Colourer
  Editor
  Stringer
  Printer

// Pre: len(i) == len(f) == len(t) > 0.
// indetermined/false/true as strings are represented by i/f/t.
  SetFormat (i, f, t string)

// The value of x is set to b.
  Set (b bool)
}

// Returns a new truth value with value undecidable.
func New() TruthValue { return new_() }
