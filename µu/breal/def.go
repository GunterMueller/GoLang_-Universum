package breal

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  Real interface { // real numbers < some power of 10

  Object
  col.Colourer
  Editor
  Stringer
  Printer

  RealVal () float64
  SetReal (r float64) bool
// Adder, Multiplier, further arithmetics ?
}

// Returns a new Object, that can hold real numbers
// with at most d digits , where d = nat.Len (n).
func New (n uint) Real { return new_(n) }
