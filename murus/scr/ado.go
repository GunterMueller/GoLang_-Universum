package scr

// (c) murus.org  v. 161216 - license see murus.go

import (
  "murus/mode"; . "murus/shape"; . "murus/linewd"; "murus/ptr"; "murus/font"; "murus/col"
)

func Fin() { actual.Fin() }
func Flush() { actual.Flush() }
func Name (s string) { actual.Name(s) }

func Move (x, y int) { actual.Move(x,y) }

func ActMode() mode.Mode { return actual.ActMode() }
func X() uint { return actual.X() }
func Y() uint { return actual.Y() }
func Wd() uint { return actual.Wd() }
func Ht() uint { return actual.Ht() }
func Wd1() uint { return actual.Wd1() }
func Ht1() uint { return actual.Ht1() }
func NLines() uint { return actual.NLines() }
func NColumns() uint { return actual.NColumns() }
func Proportion() float64 { return actual.Proportion() }

func ScrColours (f, b col.Colour) { actual.ScrColours(f,b) }
func ScrColourF (f col.Colour) { actual.ScrColourF(f) }
func ScrColourB (b col.Colour) { actual.ScrColourB(b) }
func ScrCols() (col.Colour, col.Colour) { return actual.ScrCols() }
func ScrColF() col.Colour { return actual.ScrColF() }
func ScrColB() col.Colour { return actual.ScrColB() }
func Colours (f, b col.Colour) { actual.Colours(f,b) }
func ColourF (f col.Colour) { actual.ColourF(f) }
func ColourB (b col.Colour) { actual.ColourB(b) }
func Cols() (col.Colour, col.Colour) { return actual.Cols() }
func ColF() col.Colour { return actual.ColF() }
func ColB() col.Colour { return actual.ColB() }
func Colour (x, y uint) col.Colour { return actual.Colour(x,y) }

func Clr (l, c, w, h uint) { actual.Clr(l,c,w,h) }
func ClrGr (x, y, x1, y1 int) { actual.ClrGr(x,y,x1,y1) }
func Cls() { actual.Cls() }
func Buf (on bool) { actual.Buf(on) }
func Buffered() bool { return actual.Buffered() }
func Save (l, c, w, h uint) { actual.Save(l,c,w,h) }
func SaveGr (x, y, x1, y1 int) { actual.SaveGr(x,y,x1,y1) }
func Save1() { actual.Save1() }
func Restore (l, c, w, h uint) { actual.Restore(l,c,w,h) }
func RestoreGr (x, y, x1, y1 int) { actual.RestoreGr(x,y,x1,y1) }
func Restore1() { actual.Restore1() }

func Warp (l, c uint, s Shape) { actual.Warp(l,c,s) }
func WarpGr (x, y uint, s Shape) { actual.WarpGr(x,y,s) }

func Write1 (b byte, l, c uint) { actual.Write1(b,l,c) }
func Write (s string, l, c uint ) { actual.Write(s,l,c) }
func Write1Gr (b byte, x, y int) { actual.Write1Gr(b,x,y) }
func WriteGr (s string, x, y int) { actual.WriteGr(s,x,y) }
func WriteNat (n, l, c uint) { actual.WriteNat(n,l,c) }
func WriteNatGr (n uint, x, y int) { actual.WriteNatGr(n,x,y) }
func Write1InvGr (b byte, x, y int) { actual.Write1InvGr(b,x,y) }
func WriteInvGr (s string, x, y int) { actual.WriteInvGr(s,x,y) }
func Transparent() bool { return actual.Transparent() }
func Transparence (t bool) { actual.Transparence(t) }

func ActFontsize() font.Size { return actual.ActFontsize() }
func SetFontsize (f font.Size) { actual.SetFontsize(f) }

func ActLinewidth() Linewidth { return actual.ActLinewidth() }
func SetLinewidth (w Linewidth) { actual.SetLinewidth(w) }
func Point (x, y int) { actual.Point(x,y) }
func PointInv (x, y int) { actual.PointInv(x,y) }
func Points (xs, ys []int) { actual.Points (xs,ys) }
func PointsInv (xs, ys []int) { actual.PointsInv(xs,ys) }
func Line (x, y, x1, y1 int) { actual.Line(x,y,x1,y1) }
func LineInv (x, y, x1, y1 int) { actual.LineInv(x,y,x1,y1) }
func OnLine (x, y, x1, y1, a, b int, t uint) bool { return actual.OnLine(x,y,x1,y1,a,b,t) }
func Lines (xs, ys, xs1, ys1 []int) { actual.Lines(xs,ys,xs1,ys1) }
func LinesInv (xs, ys, xs1, ys1 []int) { actual.LinesInv(xs,ys,xs1,ys1) }
func OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool { return actual.OnLines(xs,ys,xs1,ys1,a,b,t) }
func Segments (xs, ys []int) { actual.Segments(xs,ys) }
func SegmentsInv (xs, ys []int) { actual.SegmentsInv(xs,ys) }
func OnSegments (xs, ys []int, a, b int, t uint) bool { return actual.OnSegments(xs,ys,a,b,t) }
func InfLine (x, y, x1, y1 int) { actual.InfLine (x,y,x1,y1) }
func InfLineInv (x, y, x1, y1 int) { actual.InfLineInv (x,y,x1,y1) }
func OnInfLine (x, y, x1, y1, a, b int, t uint) bool { return actual.OnInfLine(x,y,x1,y1,a,b,t) }
func Rectangle (x, y, x1, y1 int) { actual.Rectangle(x,y,x1,y1) }
func RectangleInv (x, y, x1, y1 int) { actual.RectangleInv(x,y,x1,y1) }
func RectangleFull (x, y, x1, y1 int) { actual.RectangleFull(x,y,x1,y1) }
func RectangleFullInv (x, y, x1, y1 int) { actual.RectangleFullInv(x,y,x1,y1) }
func OnRectangle (x, y, x1, y1, a, b int, t uint) bool { return actual.OnRectangle(x,y,x1,y1,a,b,t) }
func InRectangle (x, y, x1, y1, a, b int, t uint) bool { return actual.InRectangle(x,y,x1,y1,a,b,t) }
func Polygon (xs, ys []int) { actual.Polygon(xs,ys) }
func PolygonInv (xs, ys []int) { actual.PolygonInv(xs,ys) }
func PolygonFull (xs, ys []int) { actual.PolygonFull(xs,ys) }
func PolygonFullInv (xs, ys []int) { actual.PolygonFullInv(xs,ys) }
func OnPolygon (xs, ys []int, a, b int, t uint) bool { return actual.OnPolygon(xs,ys,a,b,t) }
func Circle (x, y int, r uint) { actual.Circle(x,y,r) }
func CircleInv (x, y int, r uint) { actual.CircleInv(x,y,r) }
func CircleFull (x, y int, r uint) { actual.CircleFull(x,y,r) }
func CircleFullInv (x, y int, r uint) { actual.CircleFullInv(x,y,r) }
func OnCircle (x, y int, r uint, a, b int, t uint) bool { return actual.OnCircle(x,y,r,a,b,t) }
// func InCircle (x, y int, r uint, a, b int) bool { return actual.InCircle(x,y,r,a,b) } // TODO
func Arc (x, y int, r uint, a, b float64) { actual.Arc(x,y,r,a,b) }
func ArcInv (x, y int, r uint, a, b float64) { actual.ArcInv(x,y,r,a,b) }
func ArcFull (x, y int, r uint, a, b float64) { actual.ArcFull(x,y,r,a,b) }
func ArcFullInv (x, y int, r uint, a, b float64) { actual.ArcFullInv(x,y,r,a,b) }
func Ellipse (x, y int, a, b uint) { actual.Ellipse(x,y,a,b) }
func EllipseInv (x, y int, a, b uint) { actual.EllipseInv(x,y,a,b) }
func EllipseFull (x, y int, a, b uint) { actual.EllipseFull(x,y,a,b) }
func EllipseFullInv (x, y int, a, b uint) { actual.EllipseFullInv(x,y,a,b) }
func OnEllipse (x, y int, a, b uint, A, B int, t uint) bool { return actual.OnEllipse(x,y,a,b,A,B,t) }
// func InEllipse (x, y int, a, b uint, A, B int) bool { return actual.InEllipse(x,y,a,b,A,B) } // TODO
func Curve (xs, ys []int) { actual.Curve(xs,ys) }
func CurveInv (xs, ys []int) { actual.CurveInv(xs,ys) }
func OnCurve (xs, ys []int, a, b int, t uint) bool { return actual.OnCurve(xs,ys,a,b,t) }

func MouseEx() bool { return actual.MouseEx() }
func SetPointer (p ptr.Pointer) { actual.SetPointer(p) }
func MousePos() (uint, uint) { return actual.MousePos() }
func MousePosGr() (int, int) { return actual.MousePosGr() }
func WarpMouse (l, c uint) { actual.WarpMouse(l,c) }
func WarpMouseGr (x, y int) { actual.WarpMouseGr(x,y) }
func MousePointer (b bool) { actual.MousePointer(b) }
func MousePointerOn() bool { return actual.MousePointerOn() }
func UnderMouse (l, c, w, h uint) bool { return actual.UnderMouse(l,c,w,h) }
func UnderMouseGr (x, y, x1, y1 int, d uint) bool { return actual.UnderMouseGr(x,y,x1,y1,d) }
func UnderMouse1 (x, y int, d uint) bool { return actual.UnderMouse1(x,y,d) }

func Codelen (w, h uint) uint { return actual.Codelen(w,h) }
func Encode (x, y, w, h uint) []byte { return actual.Encode(x,y,w,h) }
func Decode (bs []byte) { actual.Decode(bs) }

func P6Codelen (w, h uint) uint { return actual.P6Codelen (w,h) }
func P6Size (ps []byte) (uint, uint) { return actual.P6Size(ps) }
func P6Encode (x, y, w, h uint) []byte { return actual.P6Encode(x,y,w,h) }
func P6Decode (x, y uint, ps []byte) { actual.P6Decode(x,y,ps) }

func WriteGlx() { actual.WriteGlx() }

func Cut (s string) { actual.Cut(s) }
func Paste() string { return actual.Paste() }
