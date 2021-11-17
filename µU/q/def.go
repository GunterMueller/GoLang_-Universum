package q

// (c) Christian Maurer   v. 211106 - license see µU.go

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

  RealVal() float64
  Set (a, b int) bool
  Set1 (a int) bool
  Vals() (bool, uint, uint)
  Integer() bool
  GeqNull() bool
}

// Returns a new empty rational.
func New() Rational { return new_() }
