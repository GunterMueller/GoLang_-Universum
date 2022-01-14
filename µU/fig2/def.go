package fig2

// (c) Christian Maurer   v. 220111 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/psp"
)
type
  typ byte; const (
  Points = iota // sequence of points
  Segments // line segment[s]
  Polygon
  Curve // Bezier curve
  InfLine // given by two different points
  Rectangle // borders parallel to the screen borders
  Circle
  Ellipse // main axes parallel to the screen borders
  Text // of almost 40 characters
  Image // in ppm-format
  Ntypes
)
type
  Figure interface {

  Object
  Stringer
  Marker

// x is of typ t.
  SetTyp (t typ)

// Returns the typ of x.
  Typ() typ

// x is of the Type, that was selected interactively by the user.
  Select()

// The defining points of x are shown, iff b.
  ShowPoints (b bool)

// Returns the position of the
// - first point (of the first line), if x has a typ <= Line,
// - top left corner of x, if x is of typ Rectangle or Image,
// - middle point of x, if x is of typ Circle or Ellipse,
// - bottom left corner of first characer, if is of typ Text.
  Pos() (int, int)

// x has Position (x, y)
  SetPos (x, y int)

// Returns true, iff the point at (a, b) has a distance
// of at most t pixels from x.
  On (a, b int, t uint) bool

// x is moved by (a, b).
  Move (a, b int)

// Returns true, iff the the mouse cursor is in the interior of x
// or has a distance of not more than t pixels from its boundary.
  UnderMouse (t uint) bool

// x has the colour c.
  SetColour (c col.Colour)

// Returns the colour of x.
  Colour() col.Colour

// x is drawn at its position in its colour to the screen.
  Write()

// x is drawn at its position in its inverted colour to the screen.
  WriteInv()

// Pre: x has a typ != Image.
// x is now the figure interactively generated by the user-
  Edit()

// Pre: x has the typ Image.
// x is now the image interactively generated by the user-
  EditImage (n string)

// Pre: f is not empty. f != Points and f != Image.
// x is interactively changed by the user.
  Change()

// If x is a text, it has the font size f. // TODO
//  SetFont (font.Size)

// x is printed (see package µU/psp).
  Print (psp.PostscriptPage)
}

// Returns a new empty figure of undefined Type.
func New() Figure { return new_() }
