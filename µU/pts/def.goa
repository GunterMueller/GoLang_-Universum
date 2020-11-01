package pts

// (c) Christian Maurer   v. 201025 - license see µU.go

// >>> experimental package, probably superfluous

import (
  . "µU/obj"
  "µU/col"
  "µU/vect"
  "µU/pt"
)
type
  Points interface { // persistent sequences of objects of type Point

  Clearer
  Persistor

// x is defined by the user's selection.
  Select ()

// x is defined by the 1st parameter of the program call.
  NameCall ()

// Pre: n > 0; pt.Class <= c < pt. nPolygon.
//      x does not contain a point of type pt.Start.
// The points defined by c and v with normal (0, 0, 1) are appended in x;
// the i-th point of x has the number len(v) - 1 - i.
  Ins1 (c pt.Class, v []vect.Vector, f col.Colour)

// Pre: len(v) == len(n) > 0.
//      x does not contain a point of type pt.Start.
// The points defined by c, v and n are appended in x;
// the i-th point of x has the number len(v) - 1 - i.
  Ins (c pt.Class, v, n []vect.Vector, f col.Colour)

// TODO Spec
  InsLight (l uint, v, n []vect.Vector, f col.Colour)

// Pre: (x, y, z) != (x1, y1, z1).
//      x does not contain a point of type pt.Start.
// The point of class pt.Start with number 1, coordinates (x, y, z), normal (x1, y1, z1)
// and colour Black is appended in x;
// same effect as "Ins (pt.Start, 1, v, n, col.Black)" with v = (x, y, z) and n = (x1, y1, z1).
  Start (x, y, z, x1, y1, z1 float64)

// Returns the coordinates of the last point in x and its normal, if
// that point is of class Start; returns otherwise (0, 0, 1, 0, 0, 0).
  StartCoord () (float64, float64, float64, float64, float64, float64)

// Pre: x is defined and its last point is of class pt.Start.
//      eye is defined.
// All points of x are given by openGL to the screen.
// TODO details.
  Write (d chan bool)
}

// Returns a new empty sequence of points with no name.
// (See also murus/pt.)
func New() Points { return new_() }
