package q

// (c) Christian Maurer   v. 201128 - license see µU.go

import
  . "µU/obj"
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
