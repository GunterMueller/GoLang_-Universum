package edge

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)
type
  Edge interface { // Edges with natural numbers < 10^... as values,
                   // represented as line segments on the screen.

  Object // Empty edges have value 1.

  Valuator

// f, b are the normal colours of x.
  Colours (f, b col.Colour)

// f, b are the actual colours of x.
  ColoursA (f, b col.Colour)

// x is written at its position to the screen,
// for a in its actual, otherwise in its normal colour.
  Write (x, y, x1, y1 int, a bool)

// x has the name and the value edited by the user.
  Edit (x, y, x1, y1 int)
}
// Pre: v == nil or v is of type uint or of type Valuator.
// Returns a new empty edge.
// If v == nil, its value is 1, else it is determined by v.
func New (v Any) Edge { return newEdge(v) }
