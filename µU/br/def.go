package br

// (c) Christian Maurer   v. 201206 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Real interface { // real numbers < some power of 10

  Object
  col.Colourer
  Editor
  Stringer
  Printer

  Float64() float64
  SetFloat64 (r float64) bool

  Adder
  Multiplier
// further arithmetics ?
}

// Returns a new Object, that can hold real numbers
// with at most d digits , where d = nat.Len (n).
func New (n uint) Real { return new_(n) }
