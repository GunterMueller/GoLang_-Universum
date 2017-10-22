package bnat

// (c) Christian Maurer   v. 170919 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Natural interface { // natural numbers < some power of 10

  Object
  col.Colourer
  EditorGr
  Stringer
  Valuator
  Printer

  Width() uint
// Adder, Multiplier, further arithmetics ?
}

// Returns a new Object, that can hold natural numbers
// with at most d digits, where d = nat.Len(n).
// String() has always leading zeros, iff n % 10 == 0.
func New (n uint) Natural { return new_(n) }
