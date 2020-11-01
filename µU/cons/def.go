package cons

// (c) Christian Maurer   v. 201016 - license see µU.go

// >>> This package only serves the implementations of µU/mouse,
//     mu/kbd and mu/cons; it must not no be used elsewhere.

import (
  . "µU/obj"
  . "µU/linewd"
  . "µU/shape"
  . "µU/ptr"
  . "µU/mode"
  "µU/col"
  "µU/font"
)

func MaxMode() Mode { return maxMode() }
func MaxRes() (uint, uint) { return maxRes() }
func Ok (m Mode) bool { return ok(m) }
func Lock() { lock() }
func Unlock() { unlock() }

////////////////////////////////////////////////////////////////////////

func ActIndex() int { return actIndex() }

func New (x, y uint, m Mode) Console { return newCons(x,y,m) }
func NewWH (x, y, w, h uint) Console { return newConsWH(x,y,w,h) }

func NewMax() Console { return newMax() }

// func StartCols() (col.Colour, col.Colour) { return startCols() }
// func StartColsA() (col.Colour, col.Colour) { return startColsA() }

type
  Console interface {

  Fin()

// colours /////////////////////////////////////////////////////////////

  ScrCols() (col.Colour, col.Colour)
  ScrColF() col.Colour
  ScrColB() col.Colour
  Cols() (col.Colour, col.Colour)
  ColF() col.Colour
  ColB() col.Colour
  ScrColours (f, b col.Colour)
  ScrColourF (f col.Colour)
  ScrColourB (b col.Colour)
  Colours (f, b col.Colour)
  ColourF (f col.Colour)
  ColourB (b col.Colour)
  Colour (x, y uint) col.Colour

// ranges //////////////////////////////////////////////////////////////

  Clr (l, c, w, h uint)
  ClrGr (x, y, x1, y1 int)
  Cls()

  Buf (on bool)
  Buffered() bool

  Save (l, c, w, h uint)
  SaveGr (x, y, x1, y1 int)
  Save1()
  Restore (l, c, w, h uint)
  RestoreGr (x, y, x1, y1 int)
  Restore1()

// cursor //////////////////////////////////////////////////////////////

  Warp (l, c uint, s Shape)
  WarpGr (x, y uint, s Shape)

// font ////////////////////////////////////////////////////////////////

  ActFontsize() font.Size
  SetFontsize (f font.Size)
  Wd1() uint
  Ht1() uint
  NLines() uint
  NColumns() uint

// text ////////////////////////////////////////////////////////////////

  Write1 (b byte, l, c uint)
  Write (s string, l, c uint)
  Write1Gr (b byte, x, y int)
  WriteGr (s string, x, y int)
  WriteNat (n, l, c uint)
  Write1InvGr (b byte, x, y int)
  WriteInvGr (s string, x, y int)

  Transparent() bool
  Transparence (t bool)

// graphics ////////////////////////////////////////////////////////////

  ActLinewidth() Linewidth
  SetLinewidth (w Linewidth)

  Point (x, y int)
  PointInv (x, y int)
  Points (xs, ys []int)
  PointsInv (xs, ys []int)
  OnPoint (x, y, a, b int, d uint) bool

  Line (x, y, x1, y1 int)
  LineInv (x, y, x1, y1 int)
  OnLine (x, y, x1, y1, a, b int, d uint) bool

  Lines (xs, ys, xs1, ys1 []int)
  LinesInv (xs, ys, xs1, ys1 []int)
  OnLines (xs, ys, xs1, ys1 []int, a, b int, d uint) bool

  Segments (xs, ys []int)
  SegmentsInv (xs, ys []int)
  OnSegments (xs, ys []int, a, b int, d uint) bool

  InfLine (x, y, x1, y1 int)
  InfLineInv (x, y, x1, y1 int)
  OnInfLine (x, y, x1, y1, a, b int, d uint) bool

  Triangle (x, y, x1, y1, x2, y2 int)
  TriangleInv (x, y, x1, y1, x2, y2 int)
  TriangleFull (x, y, x1, y1, x2, y2 int)
  TriangleFullInv (x, y, x1, y1, x2, y2 int)

  Rectangle (x, y, x1, y1 int)
  RectangleInv (x, y, x1, y1 int)
  RectangleFull (x, y, x1, y1 int)
  RectangleFullInv (x, y, x1, y1 int)
  OnRectangle (x, y, x1, y1, a, b int, d uint) bool
  InRectangle (x, y, x1, y1, a, b int, d uint) bool

  Polygon (xs, ys []int)
  PolygonInv (xs, ys []int)
  PolygonFull (xs, ys []int)
  PolygonFullInv (xs, ys []int)
  OnPolygon (xs, ys []int, a, b int, d uint) bool

  Circle (x, y int, r uint)
  CircleInv (x, y int, r uint)
  CircleFull (x, y int, r uint)
  CircleFullInv (x, y int, r uint)
  OnCircle (x, y int, r uint, a, b int, d uint) bool
//  InCircle (x, y int, r uint, a, b int, d uint) bool

// not yet implemented TODO
  Arc (x, y int, r uint, a, b float64)
  ArcInv (x, y int, r uint, a, b float64)
// not yet implemented TODO
  ArcFull (x, y int, r uint, a, b float64)
  ArcFullInv (x, y int, r uint, a, b float64)

  Ellipse (x, y int, a, b uint)
  EllipseInv (x, y int, a, b uint)
  EllipseFull (x, y int, a, b uint)
  EllipseFullInv (x, y int, a, b uint)
  OnEllipse (x, y int, a, b uint, A, B int, d uint) bool
//  InEllipse (x, y int, a, b uint, A, B int, d uint) bool
  Curve (xs, ys []int)
  CurveInv (xs, ys []int)
  OnCurve (xs, ys []int, a, b int, d uint) bool

// mouse ///////////////////////////////////////////////////////////////

  MouseEx() bool
  SetPointer (p Pointer)
  MousePos() (uint, uint)
  MousePosGr() (int, int)
  WarpMouse (l, c uint)
  WarpMouseGr (x, y int)

  MousePointer (b bool)
  MousePointerOn() bool
  UnderMouse (l, c, w, h uint) bool
  UnderMouseGr (x, y, x1, y1 int, t uint) bool
  UnderMouse1 (x, y int, d uint) bool

// serialisation ///////////////////////////////////////////////////////

  Codelen (w, h uint) uint
  Encode (x, y, w, h uint) Stream
  Decode (bs Stream)

// cut buffer //////////////////////////////////////////////////////////

// not yet implemented TODO
  Copy (s string)
  Paste() string
}
