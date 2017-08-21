package vect

// (c) murus.org  v. 170820 - license see murus.go

import (
  . "murus/obj"
  . "murus/spc"
)
const // For float64's a, b "a quasi equals b" means |a - b| < epsilon.
  epsilon = 1.0E-6
type
  Vector interface { // Triples (x0, x1, x2) of float64's.

// A vector is "Empty", iff it is quasi the null vector in this sense;
// Clear sets a vector to the null vector.
  Editor
  Stringer
  Printer
  Adder

// x = (x0, x1, x2).
  Set3 (x0, x1, x2 float64)

// x = (c[Right], c[Top], c[Front]).
  Set (c Coord)

// Returns the coordinates of x (in direction d).
  Coord3 () (float64, float64, float64)
  Coord (d Direction) float64

// x = (x0 + r * cos(phi) * sin(theta),
//      x1 + r * sin(phi) * sin(theta),
//      x2 + r * cos(theta)),
// where the coordinates are those of x before.
  SetPolar (x, y, z, r, phi, theta float64)

//  x = ((a0, 0, 0), (0, b1, 0), (0, 0, c2)).
  Project (a, b, c Vector)

// Returns the inner product <x, y>.
  Int (y Vector) float64

// x = [x0, y] (vector product), where x0 denotes x before.
  Cross (y Vector)

// x = [y, z].
  Ext (y, z Vector) // -> Cross ?  [Ext]Prod ?

// Returns true, if x and y are quasi linearly dependent.
  Collinear (y Vector) bool

// x = a * y.
  Scale (a float64, y Vector)

// x = (x0[0] + a, x0[1] + a, x0[2] + a), where x0 denotes x before.
  Translate (a float64)

// x = a * x0, where x0 denotes x before.
  Dilate (a float64) // name ?

// x = y + t * (z - y).
  Parametrize (y, z Vector, t float64)

// Returns |x| = sqrt (<x, x>),
  Len () float64

// Returns |x - y|.
  Distance (y Vector) float64

// x = 1/2 * (x + y).
// Returns 1/2 * |x0 - y|, where x0 denotes x before.
  Centre (x, y Vector) float64

// Returns true, iff x2 and y2 are quasi equal.
  Flat (y Vector) bool

// If |x| < epsilon, nothing had happened. Otherwise:
// x = 1/|x0| * x0, where x0 denotes x before.
  Norm ()

// Returns true, if |x| is quasi 1.
  Normed() bool

// If x and y are collinear, nothing had happened. Otherwise: 
// x = cos(a) * x0 + <x0, y> * (1 - cos(a)) * y + sin(a) * [y, x0]
// where x0 denotes x before.
  Rot (W Vector, a float64)

// Spec TODO
  Minimax (Min, Max Vector)
}

// Returns a new vector with coords (0, 0, 0).
func New() Vector { return new_() }

// Returns a new vector with coords (x1, x2, x3).
func New3 (x0, x1, x2 float64) Vector { return new3 (x0, x1, x2) }
