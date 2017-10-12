package brat

// (c) Christian Maurer   v. 170919 - license see µu.go

import
  . "µu/obj"
type
  Rational interface {

  Object
  Editor
  Stringer
  Printer
  Adder
  Multiplier

// x = 1/x0, where x0 denotes x before.
  Invert ()

  RealVal () float64
  Set (a, b int) bool
  Integer () bool
  GeqNull () bool
}

// Returns a new empty rational.
func New() Rational { return new_() }
