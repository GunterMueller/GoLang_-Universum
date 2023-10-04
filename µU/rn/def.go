package rn

// (c) Christian Maurer   v. 230924 - license see µU.go

import
  . "µU/obj"

type
  RomanNatural interface {

  Editor
  Stringer
  Undef() bool
  Valuator
  Adder
  Multiplier
}

// Liefert eine neue leere römische Zahl.
func New0() RomanNatural { return new0() }

// Pre: n > 0, n <= 10000.
// Liefert die römische Zahl mit dem Wert n.
func New (n uint) RomanNatural { return new_(n) }
