package qmat

// (c) Christian Maurer   v. 220831 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/q"
)
type
  QMatrix interface { // matrices of fractions of rationals

  Editor
  col.Colourer
//  Stringer
  TeXer
//  Printer
  Adder
  Multiplier
  Invert()

// Pre: n is the product of the number of lines
// and the number of columns of x.
// x has []n/[]n+1 as entry in row 1 and column1,
// []n+2/[]n+3 as entry in row 1 and column 2,
// and so on.
  Set (n ...int)
// x has the entries n/1.
  Set1 (n ...int)

// Pre: l < number of lines of x, c < number of columns of x.
// Returns the values of the fraction in position (l, c) of x
// (sign, numerator and denominator).
  Vals (l, c uint) (bool, uint, uint)

// Returns the determinant of x.
  Det() q.Rational
}

// Returns a new empty matrix with m lines and n colums
// for rationals as entries.
func New (m, n uint) QMatrix { return new_(m, n) }

// Returns a new matrix with m lines and n colums,
// in which all diagonal entries are 1 and all others 0.
func Unit (m, n uint) QMatrix { return unit(m, n) }
