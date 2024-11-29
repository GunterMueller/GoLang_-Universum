package edg

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Edge interface { // Edges with natural numbers < 10^... as values,
                   // represented as line segments on the screen.

  Object // Empty edges have value 1.
  Valuator
  Marker

// Returns true, if x is directed.
  Directed() bool

// x is directed iff b == true.
  Direct (b bool)

// f, b are the normal colours of x,
// fm, bm are the mark colours of x.
  Colours (f, b, fm, bm col.Colour)

// If x is directed, the direction is 0 -> 1.
  SetPos0 (int, int)
  SetPos1 (int, int)
  Pos0() (int, int)
  Pos1() (int, int)

// For b == true the screen output of the edge-Values are suppressed.
  Label (b bool)

// x is written at its position to the screen in its normal colour.

// x is written at its position to the screen,
// if x is marked, in its marked, otherwise in its normal colour.
  Write ()

// x has the name and the value edited by the user.
  Edit()
}

// Pre: a == nil or a is of an uint-type or of type Valuator.
// Returns a new empty edge, that is directed, iff d = true.
// If a == nil, its value is 1, else it is determined by a.
func New (d bool, a any) Edge { return new_(d,a) }

func W (e any) { e.(Edge).Write() }
