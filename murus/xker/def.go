package xker

// (c) murus.org  v. 161216 - license see murus.go

// >>> This package only serves the implementations of murus/mouse, 
//     murus/kbd and murus/cons; it must not be used elsewhere.

import (
  . "murus/linewd"; . "murus/shape"; . "murus/ptr"; . "murus/mode"
  "murus/col"; "murus/font"
)
type
  Event struct {
             T,     // type
             C,     // xkey.keycode, xbutton.button, xmotion.is_hint
             S uint // state
        }
var
  Eventpipe chan Event = make (chan Event)

func UnderX() bool { return underX() }
//  func Far() bool
func MaxMode() Mode { return maxMode() }
func MaxRes() (uint, uint) { return maxRes() }
func Ok (m Mode) bool { return ok(m) }
func Lock() { lock() }; func Unlock() { unlock() }

////////////////////////////////////////////////////////////////////////

func Act() Window { return actual }

// Returns a new window with left upper Corner (x, y)
// in the size of m (see mode/def.go).
func New (x, y uint, m Mode) Window { return newWin(x,y,m) }

// Returns a new window with left upper Corner (0, 0)
// in the maximal possible size (depending on the actual screen).
func NewMax() Window { return newMax() }

type
  Window interface {

  Fin()

  Flush()

  OnFocus() bool; OffFocus() bool
  Name (s string)

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
  Win2buf()
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

  Arc (x, y int, r uint, a, b float64); ArcInv (x, y int, r uint, a, b float64)
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
  UnderMouse (l, c, w, h uint) bool; UnderMouseGr (x, y, x1, y1 int, d uint) bool
  UnderMouse1 (x, y int, d uint) bool

// serialisation ///////////////////////////////////////////////////////

  Codelen (w, h uint) uint
  Encode (x, y, w, h uint) []byte; Decode (bs []byte)

// openGL //////////////////////////////////////////////////////////////

  WriteGlx()

// cut buffer //////////////////////////////////////////////////////////

  Cut (s string)
  Paste() string

}
