package edg

// (c) Christian Maurer   v. 170424 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)
type
  Edge interface { // Edges with natural numbers < 10^... as values,
                   // represented as line segments on the screen.

  Object // Empty edges have value 1.

  Valuator

// Returns true, if x is directed.
  Directed() bool

// x is directed iff b == true.
  Direct (b bool)

// f, b are the normal colours of x.
  Colours (f, b col.Colour)

// f, b are the actual colours of x.
  ColoursA (f, b col.Colour)

// If x is directed, the direction is 0 -> 1.
  SetPos0 (int, int)
  SetPos1 (int, int)
  Pos0() (int, int)
  Pos1() (int, int)

// For b == true the screen output of the edge-Values are suppressed.
  Label (b bool)

// x is written at its position to the screen in its normal colour.
  Write()

// x is written at its position to the screen,
// for a in its actual, otherwise in its normal colour.
  Write1 (a bool)

// x has the name and the value edited by the user.
  Edit (/* x, y, x1, y1 int */)
}

// Pre: a == nil or a is of type uint or of type Valuator.
// Returns a new empty edge, that is directed, iff d = true.
// If a == nil, its value is 1, else it is determined by a.
func New (d bool, a Any) Edge { return new_(d,a) }

func W (e Any, a bool) { e.(Edge).Write1(a) }
