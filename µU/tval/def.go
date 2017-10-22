package tval

// (c) Christian Maurer   v. 170919 - license see µU.go

// >>> TODO logical operations

import (
  . "µU/obj"
  "µU/col"
)
type
  TruthValue interface { // truth values "true", "false" and "undecidable"

  Object
  col.Colourer
  Editor
  Stringer
  Printer

// Pre: len(f) == len(t) > 0.
// false/true as strings are represented by f/t;
// undecidable by an empty string of the same length.
  SetFormat (f, t string)

// The value of x is set to b.
  Set (b bool)
}

// Returns a new truth value with value undecidable.
func New() TruthValue { return new_() }
