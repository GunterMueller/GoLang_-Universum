package scr

// (c) Christian Maurer   v. 210314 - license see µU.go

import (
  . "µU/obj"
  "µU/env"
  "µU/mode"
  "µU/scr/shape"
  "µU/linewd"
  "µU/scr/ptr"
  "µU/font"
  "µU/col"
)

func sc() Screen {
  if env.UnderX() {
    return actualW
  }
  return actualC
}

func Fin() { sc().Fin() }
func Flush() { sc().Flush() }
func Name (s string) { sc().Name(s) }

func ActMode() mode.Mode { return sc().ActMode() }
func X() uint { return sc().X() }
func Y() uint { return sc().Y() }
func Wd() uint { return sc().Wd() }
func Ht() uint { return sc().Ht() }
func Wd1() uint { return sc().Wd1() }
func Ht1() uint { return sc().Ht1() }
func NLines() uint { return sc().NLines() }
func NColumns() uint { return sc().NColumns() }
func Proportion() float64 { return sc().Proportion() }

func ScrColours (f, b col.Colour) { sc().ScrColours(f,b) }
func ScrColourF (f col.Colour) { sc().ScrColourF(f) }
func ScrColourB (b col.Colour) { sc().ScrColourB(b) }
func ScrCols() (col.Colour, col.Colour) { return sc().ScrCols() }
func ScrColF() col.Colour { return sc().ScrColF() }
func ScrColB() col.Colour { return sc().ScrColB() }
func Colours (f, b col.Colour) { sc().Colours(f,b) }
func ColourF (f col.Colour) { sc().ColourF(f) }
func ColourB (b col.Colour) { sc().ColourB(b) }
func Cols() (col.Colour, col.Colour) { return sc().Cols() }
func ColF() col.Colour { return sc().ColF() }
func ColB() col.Colour { return sc().ColB() }
func Colour (x, y uint) col.Colour { return sc().Colour(x,y) }

func Clr (l, c, w, h uint) { sc().Clr(l,c,w,h) }
func ClrGr (x, y, x1, y1 int) { sc().ClrGr(x,y,x1,y1) }
func Cls() { sc().Cls() }
func Buf (on bool) { sc().Buf(on) }
func Buffered() bool { return sc().Buffered() }
func Save (l, c, w, h uint) { sc().Save(l,c,w,h) }
func SaveGr (x, y, x1, y1 int) { sc().SaveGr(x,y,x1,y1) }
func Save1() { sc().Save1() }
func Restore (l, c, w, h uint) { sc().Restore(l,c,w,h) }
func RestoreGr (x, y, x1, y1 int) { sc().RestoreGr(x,y,x1,y1) }
func Restore1() { sc().Restore1() }

func Warp (l, c uint, s shape.Shape) { sc().Warp(l,c,s) }
func WarpGr (x, y uint, s shape.Shape) { sc().WarpGr(x,y,s) }

func Write1 (b byte, l, c uint) { sc().Write1(b,l,c) }
func Write (s string, l, c uint ) { sc().Write(s,l,c) }
func Write1Gr (b byte, x, y int) { sc().Write1Gr(b,x,y) }
func WriteGr (s string, x, y int) { sc().WriteGr(s,x,y) }
func WriteNat (n, l, c uint) { sc().WriteNat(n,l,c) }
func WriteNatGr (n uint, x, y int) { sc().WriteNatGr(n,x,y) }
func Write1InvGr (b byte, x, y int) { sc().Write1InvGr(b,x,y) }
func WriteInvGr (s string, x, y int) { sc().WriteInvGr(s,x,y) }
func Transparent() bool { return sc().Transparent() }
func Transparence (t bool) { sc().Transparence(t) }

func ActFontsize() font.Size { return sc().ActFontsize() }
func SetFontsize (f font.Size) { sc().SetFontsize(f) }

func ActLinewidth() linewd.Linewidth { return sc().ActLinewidth() }
func SetLinewidth (w linewd.Linewidth) { sc().SetLinewidth(w) }
func Point (x, y int) { sc().Point(x,y) }
func PointInv (x, y int) { sc().PointInv(x,y) }
func Points (xs, ys []int) { sc().Points (xs,ys) }
func PointsInv (xs, ys []int) { sc().PointsInv(xs,ys) }
func OnPoint (x, y, a, b int, d uint) bool { return sc().OnPoint(x,y,a,b,d) }
func Line (x, y, x1, y1 int) { sc().Line(x,y,x1,y1) }
func LineInv (x, y, x1, y1 int) { sc().LineInv(x,y,x1,y1) }
func OnLine (x, y, x1, y1, a, b int, t uint) bool { return sc().OnLine(x,y,x1,y1,a,b,t) }
func Lines (xs, ys, xs1, ys1 []int) { sc().Lines(xs,ys,xs1,ys1) }
func LinesInv (xs, ys, xs1, ys1 []int) { sc().LinesInv(xs,ys,xs1,ys1) }
func OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool { return sc().OnLines(xs,ys,xs1,ys1,a,b,t) }
func Segments (xs, ys []int) { sc().Segments(xs,ys) }
func SegmentsInv (xs, ys []int) { sc().SegmentsInv(xs,ys) }
func OnSegments (xs, ys []int, a, b int, t uint) bool { return sc().OnSegments(xs,ys,a,b,t) }
func InfLine (x, y, x1, y1 int) { sc().InfLine (x,y,x1,y1) }
func InfLineInv (x, y, x1, y1 int) { sc().InfLineInv (x,y,x1,y1) }
func OnInfLine (x, y, x1, y1, a, b int, t uint) bool { return sc().OnInfLine(x,y,x1,y1,a,b,t) }
func Triangle (x, y, x1, y1, x2, y2 int) { sc().Triangle(x,y,x1,y1,x2,y2) }
func TriangleInv (x, y, x1, y1, x2, y2 int) { sc().TriangleInv(x,y,x1,y1,x2,y2) }
func TriangleFull (x, y, x1, y1, x2, y2 int) { sc().TriangleFull(x,y,x1,y1,x2,y2) }
func TriangleFullInv (x, y, x1, y1, x2, y2 int) { sc().TriangleFullInv(x,y,x1,y1,x2,y2) }
func Rectangle (x, y, x1, y1 int) { sc().Rectangle(x,y,x1,y1) }
func RectangleInv (x, y, x1, y1 int) { sc().RectangleInv(x,y,x1,y1) }
func RectangleFull (x, y, x1, y1 int) { sc().RectangleFull(x,y,x1,y1) }
func RectangleFullInv (x, y, x1, y1 int) { sc().RectangleFullInv(x,y,x1,y1) }
func OnRectangle (x, y, x1, y1, a, b int, t uint) bool { return sc().OnRectangle(x,y,x1,y1,a,b,t) }
func InRectangle (x, y, x1, y1, a, b int, t uint) bool { return sc().InRectangle(x,y,x1,y1,a,b,t) }
func Polygon (xs, ys []int) { sc().Polygon(xs,ys) }
func PolygonInv (xs, ys []int) { sc().PolygonInv(xs,ys) }
func PolygonFull (xs, ys []int) { sc().PolygonFull(xs,ys) }
func PolygonFullInv (xs, ys []int) { sc().PolygonFullInv(xs,ys) }
func OnPolygon (xs, ys []int, a, b int, t uint) bool { return sc().OnPolygon(xs,ys,a,b,t) }
func Circle (x, y int, r uint) { sc().Circle(x,y,r) }
func CircleInv (x, y int, r uint) { sc().CircleInv(x,y,r) }
func CircleFull (x, y int, r uint) { sc().CircleFull(x,y,r) }
func CircleFullInv (x, y int, r uint) { sc().CircleFullInv(x,y,r) }
func OnCircle (x, y int, r uint, a, b int, t uint) bool { return sc().OnCircle(x,y,r,a,b,t) }
// func InCircle (x, y int, r uint, a, b int) bool { return sc().InCircle(x,y,r,a,b) } // TODO
func Arc (x, y int, r uint, a, b float64) { sc().Arc(x,y,r,a,b) }
func ArcInv (x, y int, r uint, a, b float64) { sc().ArcInv(x,y,r,a,b) }
func ArcFull (x, y int, r uint, a, b float64) { sc().ArcFull(x,y,r,a,b) }
func ArcFullInv (x, y int, r uint, a, b float64) { sc().ArcFullInv(x,y,r,a,b) }
func Ellipse (x, y int, a, b uint) { sc().Ellipse(x,y,a,b) }
func EllipseInv (x, y int, a, b uint) { sc().EllipseInv(x,y,a,b) }
func EllipseFull (x, y int, a, b uint) { sc().EllipseFull(x,y,a,b) }
func EllipseFullInv (x, y int, a, b uint) { sc().EllipseFullInv(x,y,a,b) }
func OnEllipse (x, y int, a, b uint, A, B int, t uint) bool { return sc().OnEllipse(x,y,a,b,A,B,t) }
// func InEllipse (x, y int, a, b uint, A, B int) bool { return sc().InEllipse(x,y,a,b,A,B) } // TODO
func Curve (xs, ys []int) { sc().Curve(xs,ys) }
func CurveInv (xs, ys []int) { sc().CurveInv(xs,ys) }
func OnCurve (xs, ys []int, a, b int, t uint) bool { return sc().OnCurve(xs,ys,a,b,t) }

func MouseEx() bool { return sc().MouseEx() }
func SetPointer (p ptr.Pointer) { sc().SetPointer(p) }
func MousePos() (uint, uint) { return sc().MousePos() }
func MousePosGr() (int, int) { return sc().MousePosGr() }
func WarpMouse (l, c uint) { sc().WarpMouse(l,c) }
func WarpMouseGr (x, y int) { sc().WarpMouseGr(x,y) }
func MousePointer (b bool) { sc().MousePointer(b) }
func MousePointerOn() bool { return sc().MousePointerOn() }
func UnderMouse (l, c, w, h uint) bool { return sc().UnderMouse(l,c,w,h) }
func UnderMouseGr (x, y, x1, y1 int, d uint) bool { return sc().UnderMouseGr(x,y,x1,y1,d) }
func UnderMouse1 (x, y int, d uint) bool { return sc().UnderMouse1(x,y,d) }

func Codelen (w, h uint) uint { return sc().Codelen(w,h) }
func Encode (x, y, w, h uint) Stream { return sc().Encode(x,y,w,h) }
func Decode (bs Stream) { sc().Decode(bs) }

func PPMHeader (w, h uint) string { return sc().PPMHeader(w,h) }
func PPMCodelen (w, h uint) uint { return sc().PPMCodelen(w,h) }
func PPMEncode (x, y, w, h uint) Stream { return sc().PPMEncode(x,y,w,h) }
func PPMDecode (s Stream, x, y uint) { sc().PPMDecode(s,x,y) }
func PPMSize (s Stream) (uint, uint) { return sc().PPMSize(s) }
