package q

// (c) Christian Maurer   v. 211117 - license see µU.go

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

// Returns the real value of x.
  RealVal() float64

// x is set to a/b.
  Set (a, b int) bool

// x is set to a/1.
  Set1 (a int) bool

// Returns for x = a/b (x>=0, a, b).
  Vals() (bool, uint, uint)

// Returns true, iff x = a/1.
  Integer() bool

// Returns true, iff x >= 0.
  GeqNull() bool

// Returns the length of x.String().
  Wd() uint
}

// Returns a new empty rational.
func New() Rational { return new_() }
