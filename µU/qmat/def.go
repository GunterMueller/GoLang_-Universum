package qmat

// (c) Christian Maurer   v. 211106 - license see µU.go

import (
  . "µU/obj"
  "µU/q"
)
type
  QMatrix interface { // matrices of fractions of rationsls

  Object
  Editor
//  Stringer
  TeXer
//  Printer
  Adder
  Multiplier
  Invert()

  Set1 (n ...int)
  Set (n ...int)

// Pre: l < number of lines of x, c < number of columns of x.
// Returns the values of the fraction in position (l, c) of x
// (sign, numerator and denominator).
  Vals (l, c uint) (bool, uint, uint)

// Returns the determinant of x.
  Det() q.Rational
}

// Pre: d <= 3.
// Returns a new empty matrix with m lines and n colums with fractions as entries,
// whose numerators and denominators of it have at most d digits.
func New (m, n, d uint) QMatrix { return new_(m, n, d) }

// Pre: d <= 3.
// Returns a new matrix with m lines and n colums,
// in which all diagonal entries are 1 and all others 0.
func Unit (m, n, d uint) QMatrix { return unit(m, n, d) }
