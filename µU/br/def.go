package br

// (c) Christian Maurer   v. 220811 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Real interface { // real numbers < some power of 10

// Returns 4 + the number of digits given by New.
  Width() uint

  Editor
  col.Colourer
  Stringer
// Pre for SetRealVal: x has enough digits.
  RealValuator
  Printer

  Adder
  Multiplier
// further arithmetics ?
}

// Pre: d <= .
// Returns a new empty object with n digits before the dot.
func New (n uint) Real { return new_(n) }
