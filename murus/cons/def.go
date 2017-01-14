package cons

// (c) murus.org  v. 161216 - license see murus.go

// >>> This package only serves the implementations of murus/mouse, 
//     murus/kbd and murus/cons; it must not no be used elsewhere.

import (
  . "murus/linewd"; . "murus/shape"; . "murus/ptr"; . "murus/mode"
  "murus/col"; "murus/font"
)

func MaxMode() Mode { return maxMode() }
func MaxRes() (uint, uint) { return maxRes() }
func Ok (m Mode) bool { return ok(m) }
func Lock() { lock() }; func Unlock() { unlock() }

////////////////////////////////////////////////////////////////////////

func ActIndex() int { return actIndex() }

func New (x, y uint, m Mode) Console { return newCons(x,y,m) }

func NewMax() Console { return newMax() }

type
  Console interface {

  Fin()

// experimental ////////////////////////////////////////////////////////

  Moved (x, y int) bool

// modes and sizes /////////////////////////////////////////////////////

  X() uint; Y() uint
  Wd() uint; Ht() uint
  Wd1() uint; Ht1() uint
  NLines() uint; NColumns() uint

// colours /////////////////////////////////////////////////////////////

  ScrColours (f, b col.Colour); ScrColourF (f col.Colour); ScrColourB (b col.Colour)
  ScrCols() (col.Colour, col.Colour); ScrColF() col.Colour; ScrColB() col.Colour
  Colours (f, b col.Colour); ColourF (f col.Colour); ColourB (b col.Colour)
  Cols() (col.Colour, col.Colour); ColF() col.Colour; ColB() col.Colour
  Colour (x, y uint) col.Colour

// ranges //////////////////////////////////////////////////////////////

  Clr (l, c, w, h uint); ClrGr (x, y, x1, y1 int); Cls()

  Buf (on bool); Buffered() bool

  Save (l, c, w, h uint); SaveGr (x, y, x1, y1 int); Save1()
  Restore (l, c, w, h uint); RestoreGr (x, y, x1, y1 int); Restore1()

// cursor //////////////////////////////////////////////////////////////

  Warp (l, c uint, s Shape); WarpGr (x, y uint, s Shape)

// font ////////////////////////////////////////////////////////////////

  ActFontsize() font.Size; SetFontsize (f font.Size)

// text ////////////////////////////////////////////////////////////////

  Write1 (b byte, l, c uint); Write (s string, l, c uint)
  Write1Gr (b byte, x, y int); WriteGr (s string, x, y int)
  WriteNat (n, l, c uint)
  Write1InvGr (b byte, x, y int); WriteInvGr (s string, x, y int)

  Transparent() bool; Transparence (t bool)

// graphics ////////////////////////////////////////////////////////////

  ActLinewidth() Linewidth; SetLinewidth (w Linewidth)

  Point (x, y int); PointInv (x, y int)
  Points (xs, ys []int); PointsInv (xs, ys []int)

  Line (x, y, x1, y1 int); LineInv (x, y, x1, y1 int)
  OnLine (x, y, x1, y1, a, b int, t uint) bool

  Lines (xs, ys, xs1, ys1 []int); LinesInv (xs, ys, xs1, ys1 []int)
  OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool

  Segments (xs, ys []int); SegmentsInv (xs, ys []int)
  OnSegments (xs, ys []int, a, b int, t uint) bool

  InfLine (x, y, x1, y1 int); InfLineInv (x, y, x1, y1 int)
  OnInfLine (x, y, x1, y1, a, b int, t uint) bool

  Rectangle (x, y, x1, y1 int); RectangleInv (x, y, x1, y1 int)
  RectangleFull (x, y, x1, y1 int); RectangleFullInv (x, y, x1, y1 int)
  OnRectangle (x, y, x1, y1, a, b int, t uint) bool; InRectangle (x, y, x1, y1, a, b int, t uint) bool

  Polygon (xs, ys []int); PolygonInv (xs, ys []int)
  PolygonFull (xs, ys []int); PolygonFullInv (xs, ys []int)
  OnPolygon (xs, ys []int, a, b int, t uint) bool

  Circle (x, y int, r uint); CircleInv (x, y int, r uint)
  CircleFull (x, y int, r uint); CircleFullInv (x, y int, r uint)
  OnCircle (x, y int, r uint, a, b int, t uint) bool // ; InCircle (x, y int, r uint, a, b int, t uint) bool

// not yet implemented TODO
  Arc (x, y int, r uint, a, b float64); ArcInv (x, y int, r uint, a, b float64)
// not yet implemented TODO
  ArcFull (x, y int, r uint, a, b float64); ArcFullInv (x, y int, r uint, a, b float64)

  Ellipse (x, y int, a, b uint); EllipseInv (x, y int, a, b uint)
  EllipseFull (x, y int, a, b uint); EllipseFullInv (x, y int, a, b uint)
  OnEllipse (x, y int, a, b uint, A, B int, t uint) bool // ; InEllipse (x, y int, a, b uint, A, B int, t uint) bool

  Curve (xs, ys []int); CurveInv (xs, ys []int)
  OnCurve (xs, ys []int, a, b int, t uint) bool

// mouse ///////////////////////////////////////////////////////////////

  MouseEx() bool
  SetPointer (p Pointer)
  MousePos() (uint, uint); MousePosGr() (int, int)
  WarpMouse (l, c uint); WarpMouseGr (x, y int)

  MousePointer (b bool); MousePointerOn() bool
  UnderMouse (l, c, w, h uint) bool; UnderMouseGr (x, y, x1, y1 int, t uint) bool
  UnderMouse1 (x, y int, d uint) bool

// serialisation ///////////////////////////////////////////////////////

  Codelen (w, h uint) uint
  Encode (x, y, w, h uint) []byte; Decode (bs []byte)

// openGL //////////////////////////////////////////////////////////////

  WriteGlx()

// cut buffer //////////////////////////////////////////////////////////

// not yet implemented TODO
  Cut (s string)
// not yet implemented TODO
  Paste() string
}
