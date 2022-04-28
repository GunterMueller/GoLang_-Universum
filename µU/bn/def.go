package bn

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const
  M = 20 // 1<<64 - 2 = 18446744073709551614 has 20 digits
type
  Natural interface { // natural numbers < 2^64 - 1.

  Object
  col.Colourer
  EditorGr
  Stringer
  Valuator
  Printer
  Adder
//  Multiplier ?

// Returns the width of x given by New.
  Width() uint

// Pre: s contains only the digits 0 and 1.
// x is the natural number with the binary represenation s.
  Decimal (s string)

// Returns the binary representation of x.
  Dual() string
}

// Pre: 0 < n <= M.
// Returns a new Natural with value 0 for numbers with at most n digits.
func New (n uint) Natural { return new_(n) }
