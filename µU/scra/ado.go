package scr

// (c) Christian Maurer   v. 211125 - license see µU.go

import (
  "µU/ker"
  "µU/obj"
  "µU/mode"
  "µU/scr/shape"
  "µU/linewd"
  "µU/scr/ptr"
  "µU/font"
  "µU/col"
)
var
  under_C, under_X, under_S bool

func init() {
  under_C = ker.UnderC()
  under_X = ker.UnderX()
  under_S = ker.UnderS()
}

func underC() bool {
  return under_C
}

func underX() bool {
  return under_X
}

func underS() bool {
  return under_S
}

func s() Screen {
  if under_S {
    return actualS
  }
  if under_X {
    return actualW
  }
  return actualC
}

func Fin() { s().Fin() }
func Flush() { s().Flush() }
func Name (n string) { s().Name(n) }

func ActMode() mode.Mode { return s().ActMode() }
func X() uint { return s().X() }
func Y() uint { return s().Y() }
func Wd() uint { return s().Wd() }
func Ht() uint { return s().Ht() }
func Wd1() uint { return s().Wd1() }
func Ht1() uint { return s().Ht1() }
func NLines() uint { return s().NLines() }
func NColumns() uint { return s().NColumns() }
func Proportion() float64 { return s().Proportion() }

func ScrColours (f, b col.Colour) { s().ScrColours(f,b) }
func ScrColourF (f col.Colour) { s().ScrColourF(f) }
func ScrColourB (b col.Colour) { s().ScrColourB(b) }
func ScrCols() (col.Colour, col.Colour) { return s().ScrCols() }
func ScrColF() col.Colour { return s().ScrColF() }
func ScrColB() col.Colour { return s().ScrColB() }
func Colours (f, b col.Colour) { s().Colours(f,b) }
func ColourF (f col.Colour) { s().ColourF(f) }
func ColourB (b col.Colour) { s().ColourB(b) }
func Cols() (col.Colour, col.Colour) { return s().Cols() }
func ColF() col.Colour { return s().ColF() }
func ColB() col.Colour { return s().ColB() }
func Colour (x, y uint) col.Colour { return s().Colour(x,y) }

func Clr (l, c, w, h uint) { s().Clr(l,c,w,h) }
func ClrGr (x, y, x1, y1 int) { s().ClrGr(x,y,x1,y1) }
func Cls() { s().Cls() }
func Buf (on bool) { s().Buf(on) }
func Buffered() bool { return s().Buffered() }
func Save (l, c, w, h uint) { s().Save(l,c,w,h) }
func SaveGr (x, y, x1, y1 int) { s().SaveGr(x,y,x1,y1) }
func Save1() { s().Save1() }
func Restore (l, c, w, h uint) { s().Restore(l,c,w,h) }
func RestoreGr (x, y, x1, y1 int) { s().RestoreGr(x,y,x1,y1) }
func Restore1() { s().Restore1() }

func Warp (l, c uint, h shape.Shape) { s().Warp(l,c,h) }
func WarpGr (x, y uint, h shape.Shape) { s().WarpGr(x,y,h) }

func Write1 (b byte, l, c uint) { s().Write1(b,l,c) }
func Write (t string, l, c uint ) { s().Write(t,l,c) }
func Write1Gr (b byte, x, y int) { s().Write1Gr(b,x,y) }
func WriteGr (t string, x, y int) { s().WriteGr(t,x,y) }
func WriteNat (n, l, c uint) { s().WriteNat(n,l,c) }
func WriteNatGr (n uint, x, y int) { s().WriteNatGr(n,x,y) }
func Write1InvGr (b byte, x, y int) { s().Write1InvGr(b,x,y) }
func WriteInvGr (t string, x, y int) { s().WriteInvGr(t,x,y) }
func Transparent() bool { return s().Transparent() }
func Transparence (t bool) { s().Transparence(t) }

func ActFontsize() font.Size { return s().ActFontsize() }
func SetFontsize (f font.Size) { s().SetFontsize(f) }

func ActLinewidth() linewd.Linewidth { return s().ActLinewidth() }
func SetLinewidth (w linewd.Linewidth) { s().SetLinewidth(w) }
func Point (x, y int) { s().Point(x,y) }
func PointInv (x, y int) { s().PointInv(x,y) }
func Points (xs, ys []int) { s().Points (xs,ys) }
func PointsInv (xs, ys []int) { s().PointsInv(xs,ys) }
func OnPoint (x, y, a, b int, d uint) bool { return s().OnPoint(x,y,a,b,d) }
func Line (x, y, x1, y1 int) { s().Line(x,y,x1,y1) }
func LineInv (x, y, x1, y1 int) { s().LineInv(x,y,x1,y1) }
func OnLine (x, y, x1, y1, a, b int, t uint) bool { return s().OnLine(x,y,x1,y1,a,b,t) }
func Lines (xs, ys, xs1, ys1 []int) { s().Lines(xs,ys,xs1,ys1) }
func LinesInv (xs, ys, xs1, ys1 []int) { s().LinesInv(xs,ys,xs1,ys1) }
func OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool { return s().OnLines(xs,ys,xs1,ys1,a,b,t) }
func Segments (xs, ys []int) { s().Segments(xs,ys) }
func SegmentsInv (xs, ys []int) { s().SegmentsInv(xs,ys) }
func OnSegments (xs, ys []int, a, b int, t uint) bool { return s().OnSegments(xs,ys,a,b,t) }
func InfLine (x, y, x1, y1 int) { s().InfLine (x,y,x1,y1) }
func InfLineInv (x, y, x1, y1 int) { s().InfLineInv (x,y,x1,y1) }
func OnInfLine (x, y, x1, y1, a, b int, t uint) bool { return s().OnInfLine(x,y,x1,y1,a,b,t) }
func Triangle (x, y, x1, y1, x2, y2 int) { s().Triangle(x,y,x1,y1,x2,y2) }
func TriangleInv (x, y, x1, y1, x2, y2 int) { s().TriangleInv(x,y,x1,y1,x2,y2) }
func TriangleFull (x, y, x1, y1, x2, y2 int) { s().TriangleFull(x,y,x1,y1,x2,y2) }
func TriangleFullInv (x, y, x1, y1, x2, y2 int) { s().TriangleFullInv(x,y,x1,y1,x2,y2) }
func Rectangle (x, y, x1, y1 int) { s().Rectangle(x,y,x1,y1) }
func RectangleInv (x, y, x1, y1 int) { s().RectangleInv(x,y,x1,y1) }
func RectangleFull (x, y, x1, y1 int) { s().RectangleFull(x,y,x1,y1) }
func RectangleFullInv (x, y, x1, y1 int) { s().RectangleFullInv(x,y,x1,y1) }
func OnRectangle (x, y, x1, y1, a, b int, t uint) bool { return s().OnRectangle(x,y,x1,y1,a,b,t) }
func InRectangle (x, y, x1, y1, a, b int, t uint) bool { return s().InRectangle(x,y,x1,y1,a,b,t) }
func Polygon (xs, ys []int) { s().Polygon(xs,ys) }
func PolygonInv (xs, ys []int) { s().PolygonInv(xs,ys) }
func PolygonFull (xs, ys []int) { s().PolygonFull(xs,ys) }
func PolygonFullInv (xs, ys []int) { s().PolygonFullInv(xs,ys) }
func OnPolygon (xs, ys []int, a, b int, t uint) bool { return s().OnPolygon(xs,ys,a,b,t) }
func Circle (x, y int, r uint) { s().Circle(x,y,r) }
func CircleInv (x, y int, r uint) { s().CircleInv(x,y,r) }
func CircleFull (x, y int, r uint) { s().CircleFull(x,y,r) }
func CircleFullInv (x, y int, r uint) { s().CircleFullInv(x,y,r) }
func OnCircle (x, y int, r uint, a, b int, t uint) bool { return s().OnCircle(x,y,r,a,b,t) }
// func InCircle (x, y int, r uint, a, b int) bool { return s().InCircle(x,y,r,a,b) } // TODO
func Arc (x, y int, r uint, a, b float64) { s().Arc(x,y,r,a,b) }
func ArcInv (x, y int, r uint, a, b float64) { s().ArcInv(x,y,r,a,b) }
func ArcFull (x, y int, r uint, a, b float64) { s().ArcFull(x,y,r,a,b) }
func ArcFullInv (x, y int, r uint, a, b float64) { s().ArcFullInv(x,y,r,a,b) }
func Ellipse (x, y int, a, b uint) { s().Ellipse(x,y,a,b) }
func EllipseInv (x, y int, a, b uint) { s().EllipseInv(x,y,a,b) }
func EllipseFull (x, y int, a, b uint) { s().EllipseFull(x,y,a,b) }
func EllipseFullInv (x, y int, a, b uint) { s().EllipseFullInv(x,y,a,b) }
func OnEllipse (x, y int, a, b uint, A, B int, t uint) bool { return s().OnEllipse(x,y,a,b,A,B,t) }
// func InEllipse (x, y int, a, b uint, A, B int) bool { return s().InEllipse(x,y,a,b,A,B) } // TODO
func Curve (xs, ys []int) { s().Curve(xs,ys) }
func CurveInv (xs, ys []int) { s().CurveInv(xs,ys) }
func OnCurve (xs, ys []int, a, b int, t uint) bool { return s().OnCurve(xs,ys,a,b,t) }

func MouseEx() bool { return s().MouseEx() }
func SetPointer (p ptr.Pointer) { s().SetPointer(p) }
func MousePos() (uint, uint) { return s().MousePos() }
func MousePosGr() (int, int) { return s().MousePosGr() }
func WarpMouse (l, c uint) { s().WarpMouse(l,c) }
func WarpMouseGr (x, y int) { s().WarpMouseGr(x,y) }
func MousePointer (b bool) { s().MousePointer(b) }
func MousePointerOn() bool { return s().MousePointerOn() }
func UnderMouse (l, c, w, h uint) bool { return s().UnderMouse(l,c,w,h) }
func UnderMouseGr (x, y, x1, y1 int, d uint) bool { return s().UnderMouseGr(x,y,x1,y1,d) }
func UnderMouse1 (x, y int, d uint) bool { return s().UnderMouse1(x,y,d) }

func Codelen (w, h uint) uint { return s().Codelen(w,h) }
func Encode (x, y, w, h uint) obj.Stream { return s().Encode(x,y,w,h) }
func Decode (bs obj.Stream) { s().Decode(bs) }

func PPMHeader (w, h uint) string { return s().PPMHeader(w,h) }
func PPMCodelen (w, h uint) uint { return s().PPMCodelen(w,h) }
func PPMEncode (x, y, w, h uint) obj.Stream { return s().PPMEncode(x,y,w,h) }
func PPMDecode (b obj.Stream, x, y uint) { s().PPMDecode(b,x,y) }
func PPMSize (b obj.Stream) (uint, uint) { return s().PPMSize(b) }
