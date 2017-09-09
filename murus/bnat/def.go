package bnat

// (c) Christian Maurer   v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Natural interface { // natural numbers < some power of 10

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
