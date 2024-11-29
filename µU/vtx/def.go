package vtx

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Vertex interface { // Vertices with objects as content and positions on the screen

  Object // Empty means content is empty
  Valuator
//  Marker
  Stringer

// Returns the content of x.
  Content() EditorGr

// Returns the width of x.
  Wd() uint

// Returns the height of x.
  Ht() uint

// Returns the width and height of x.
  Size() (uint, uint)

// x has the the center coordinates (x, y).
  Set (x, y int)

// Returns the center coordinates of x.
  Pos() (int, int)

// The center of x has the position of the mouse.
  Mouse()

// Returns true, iff the mouse has the position of the center of x.
  UnderMouse() bool

// f and b are the normal colours of x
// and fm and bm the mark colours of x.
  Colours (f, b, fm, bm col.Colour)

// Returns the normal fore- and backgroundcolour of x.
  Cols() (col.Colour, col.Colour)

// Returns the mark fore- and backgroundcolour of x.
  ColsM() (col.Colour, col.Colour)

// x is written to the screen at its position;
// if x is marked, in the mark colour, otherwise in its normal colour.
  Write()

// x has the content, that was edited by the user.
  Edit()

  Mark (m bool)
  Marked() bool
}

// Returns a new empty node for content of the type of e
// of size (width, height) = (w, h) (in colums, lines).
func New (e EditorGr, w, h uint) Vertex { return new_(e,w,h) }

// func W (v any, b bool) { v.(Vertex).Write1(b) }
func W (v any) { v.(Vertex).Write() }
