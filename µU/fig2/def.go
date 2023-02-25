package fig2

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/fontsize"
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
  Figure2 interface {

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
// Returns true, if x is now the image interactively generated by the user.
  ImageEdited (n string) bool

// Pre: f is not empty. f != Points and f != Image.
// x is interactively changed by the user.
  Change()

// If x is a text, it has the font size f. // TODO
SetFontsize (fontsize.Size)

// x is printed (see package µU/psp).
  Print (psp.PostscriptPage)
}

// Returns a new empty figure with undefined typ.
func New() Figure2 { return new_() }

// Return a new empty figure with the corresponding typ.
func NewPoints (xs, ys []int, c col.Colour) Figure2 { return newPoints(xs,ys,c) }
func NewSegments (xs, ys []int, c col.Colour) Figure2 { return newSegments(xs,ys,c) }
func NewPolygon (xs, ys []int, f bool, c col.Colour) Figure2 { return newPolygon(xs,ys,f,c) }
func NewCurve (xs, ys []int, c col.Colour) Figure2 { return newCurve(xs,ys,c) }
func NewInfLine (x, y, x1, y1 int, c col.Colour) Figure2 { return newInfLine(x,y,x1,y1,c) }
func NewRectangle (x, y, x1, y1 int, f bool, c col.Colour) Figure2 { return newRectangle(x,y,x1,y1,f,c) }
func NewCircle (x, y, r int, f bool, c col.Colour) Figure2 { return newCircle(x,y,r,f,c) }
func NewEllipse (x, y, a, b int, f bool, c col.Colour) Figure2 { return newEllipse(x,y,a,b,f,c) }
func NewText (x, y int, s string, c col.Colour) Figure2 { return newText(x,y,s,c) }
// func NewImage (x, y int, n string) Figure2 { return newImage(x,y,n) }
