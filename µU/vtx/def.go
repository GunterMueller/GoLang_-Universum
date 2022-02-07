package vtx

// (c) Christian Maurer   v. 220131 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Vertex interface { // Vertices with objects as content and positions on the screen

  Object // Empty means content is empty
  col.Colourer
  Valuator

// Returns the content of x.
  Content() EditorGr

// Returns the width of x.
  Wd() uint

// Returns the height of x.
  Ht() uint

// Returns the width and height of x.
  Size() (uint, uint)

// x has the the center (x, y).
  Set (x, y int)

// Returns the coordinates of x.
  Pos() (int, int)

// TODO Spec
  Contour() (uint, uint)

// The center of x has the position of the mouse.
  Mouse()

// Returns true, iff the mouse has the position of the center of x.
  UnderMouse() bool

// f and b are the actual colours of x.
  ColoursA (f, b col.Colour)

// x is written to the screen at its position in its normal colour.
  Write()

// x is written to the screen at its position;
// if a, in the actual colour, otherwise in its normal colour.
  Write1 (a bool)

// x has the content, that was edited by the user.
  Edit()
}

// Returns a new empty node for content of the type of e
// of size (width, height) = (w, h) (in colums, lines).
func New (e EditorGr, w, h uint) Vertex { return new_(e,w,h) }

func W (v Any, a bool) { v.(Vertex).Write1(a) }
