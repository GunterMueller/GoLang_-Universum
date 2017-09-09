package xwin

// (c) Christian Maurer   v. 170814 - license see murus.go

// #include <stdlib.h>
// #include <X11/X.h>
// #include <X11/Xlib.h>
// #include <GL/glx.h>
import
  "C"
import (
  "math"
  "murus/ker"
  . "murus/linewd"
)

func (X *xwindow) ActLinewidth() Linewidth {
  return X.lineWd
}

func (X *xwindow) SetLinewidth (w Linewidth) {
  X.lineWd = w
  cw:= C.uint(0)
  switch w { case Thick:
    cw = C.uint(2)
  case Thicker:
    cw = C.uint(3)
  }
  C.XSetLineAttributes (dpy, X.gc, cw, C.LineSolid, C.CapRound, C.JoinRound)
}

func intord (x, y, x1, y1 *int) {
  if *x > *x1 { *x, *x1 = *x1, *x }
  if *y > *y1 { *y, *y1 = *y1, *y }
}

func (X *xwindow) point (x, y int, n bool) {
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawPoint (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y)) }
  C.XDrawPoint (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Point (x, y int) {
  X.point (x, y, true)
}

func near (x, y, a, b int, d uint) bool {
  dx, dy := x - a, y - b
  return dx * dx + dy * dy <= int(d * d)
}

func (X *xwindow) PointInv (x, y int) {
  X.point (x, y, false)
}

func ok2 (xs, ys []int) bool {
  return len (xs) == len (ys)
}

func ok4 (xs, ys, xs1, ys1 []int) bool {
  return len (xs) == len (ys) &&
         len (xs1) == len (ys1) &&
         len (xs) == len (xs1)
}

func (X *xwindow) points (xs, ys []int, b bool) {
  n:= len (xs)
  if n == 0 { return }
  if ! ok2 (xs, ys) { return }
  if n == 1 { X.point (xs[0], ys[0], b) }
  p:= make ([]C.XPoint, n)
  for i:= 0; i < n; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! b { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawPoints (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(n), C.CoordModeOrigin) }
  C.XDrawPoints (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(n), C.CoordModeOrigin)
  if ! b { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Points (xs, ys []int) {
  X.points (xs, ys, true)
}

func (X *xwindow) PointsInv (xs, ys []int) {
  X.points (xs, ys, false)
}

func (X *xwindow) OnPoint (x, y, a, b int, d uint) bool {
  return near (x, y, a, b, d)
}

func (X *xwindow) line (x, y, x1, y1 int, n bool) {
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawLine (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.int(x1), C.int(y1)) }
  C.XDrawLine (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.int(x1), C.int(y1))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Line (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, true)
}

func (X *xwindow) LineInv (x, y, x1, y1 int) {
  X.line (x, y, x1, y1, false)
}

// Returns true, if m is - up to tolerance t - between i and k.
func between (i, k, m, t int) bool {
  return i <= m + t && m <= k + t || k <= m + t && m <= i + t
}

func (X *xwindow) OnLine (x, y, x1, y1, a, b int, t uint) bool {
  if x > x1 { x, x1 = x1, x; y, y1 = y1, y }
  if ! (between (x, x1, a, int(t)) && between (y, y1, b, int(t))) {
    return false
  }
  if x == x1 {
    return between (x, x, a, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t))
  }
  if near (x, y, a, b, t) || near (x1, y1, a, b, t) { return true }
  m:= float64(y1 - y) / float64(x1 - x)
  return near (a, b, a, y + int(m * float64(a - x) + 0.5), t)
}

func (X *xwindow) lines (xs, ys, xs1, ys1 []int, n bool) {
  l:= len (xs); if len (ys) != l { return }
  s:= make ([]C.XSegment, l)
  for i:= 0; i < l; i++ {
    s[i].x1, s[i].y1, s[i].x2, s[i].y2 = C.short(xs[i]), C.short(ys[i]), C.short(xs1[i]), C.short(ys1[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawSegments (dpy, C.Drawable(X.win), X.gc, &s[0], C.int(l)) }
  C.XDrawSegments (dpy, C.Drawable(X.buffer), X.gc, &s[0], C.int(l))
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Lines (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, true)
}

func (X *xwindow) LinesInv (xs, ys, xs1, ys1 []int) {
  X.lines (xs, ys, xs1, ys1, false)
}

func (X *xwindow) OnLines (xs, ys, xs1, ys1 []int, a, b int, t uint) bool {
  if len (xs) == 0 { return false }
  if ! ok4 (xs, ys, xs1, ys1) { return false }
  for i:= 0; i < len (xs); i++ {
    if X.OnLine (xs[i], ys[i], xs1[i], ys1[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *xwindow) segments (xs, ys []int, n bool) {
  l:= len (xs); if len (ys) != l { return }
  p:= make ([]C.XPoint, l)
  for i:= 0; i < l; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
  if ! X.buff { C.XDrawLines (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(l), C.CoordModeOrigin) }
  C.XDrawLines (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(l), C.CoordModeOrigin)
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Segments (xs, ys []int) {
  X.segments (xs, ys, true)
}

func (X *xwindow) SegmentsInv (xs, ys []int) {
  X.segments (xs, ys, false)
}

func (X *xwindow) OnSegments (xs, ys []int, a, b int, t uint) bool {
  if ! ok2 (xs, ys) { return false }
  if len (xs) == 1 { return xs[0] == a && ys[0] == b }
  for i:= 1; i < len (xs); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return false
}

func (X *xwindow) border (x, y, x1, y1 *int) {
  if *x > *x1 { *x, *x1 = *x1, *x; *y, *y1 = *y1, *y }
  for *x > 0 {
    *x -= *x1 - *x
    *y -= *y1 - *y
  }
  for *x1 < int(X.wd) {
    *x1 += *x1 - *x
    *y1 += *y1 - *y
  }
}

func (X *xwindow) InfLine (x, y, x1, y1 int) {
  if x == x1 {
    if y == y1 { return }
    X.Line (x, 0, x, int(X.ht) - 1)
    return
  }
  if y == y1 {
    X.Line (0, y, int(X.wd) - 1, y)
    return
  }
  X.border (&x, &y, &x1, &y1)
  X.Line (x, y, x1, y1)
}

func (X *xwindow) InfLineInv (x, y, x1, y1 int) {
  if x == x1 {
    if y == y1 { return }
    X.LineInv (x, 0, x, int(X.ht) - 1)
    return
  }
  if y == y1 {
    X.LineInv (0, y, int(X.wd) - 1, y)
    return
  }
  X.border (&x, &y, &x1, &y1)
  X.LineInv (x, y, x1, y1)
}

func (X *xwindow) OnInfLine (x, y, x1, y1, a, b int, t uint) bool {
  if x > x1 { x, x1 = x1, x; y, y1 = y1, y }
  if x == x1 {
    return between (x, x, a, int(t))
  }
  if y == y1 {
    return between (y, y, b, int(t))
  }
  if near (x, y, a, b, t) || near (x1, y1, a, b, t) { return true }
  X.border (&x, &y, &x1, &y1)
  m:= float64(y1 - y) / float64(x1 - x)
  return near (a, b, a, y + int(m * float64(a - x) + 0.5), t)
}

func (X *xwindow) rectangle (x, y, w, h int, n, f bool) {
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ? 
    if ! X.buff { C.XFillRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XFillRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawRectangle (dpy, C.Drawable(X.win), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h)) }
    C.XDrawRectangle (dpy, C.Drawable(X.buffer), X.gc, C.int(x), C.int(y), C.uint(w), C.uint(h))
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Rectangle (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1 - x + 1, y1 - y + 1, true, false)
}

func (X *xwindow) RectangleInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1 - x + 1, y1 - y + 1, false, false)
}

func (X *xwindow) RectangleFull (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1 - x + 1, y1 - y + 1, true, true)
}

func (X *xwindow) RectangleFullInv (x, y, x1, y1 int) {
  intord (&x, &y, &x1, &y1)
  X.rectangle (x, y, x1 - x + 1, y1 - y + 1, false, true)
}

func (X *xwindow) OnRectangle (x, y, x1, y1, a, b int, t uint) bool {
  if ! X.InRectangle (x, y, x1, y1, a, b, t) { return false }
  return between (x, x, a, int(t)) || between (x1, x1, a, int(t)) ||
         between (y, y, b, int(t)) || between (y1, y1, b, int(t))
}

func (X *xwindow) InRectangle (x, y, x1, y1, a, b int, t uint) bool {
  return between (x, x1, a, int(t)) && between (y, y1, b, int(t))
}

func (X *xwindow) Polygon (xs, ys []int) {
  X.segments (xs, ys, true)
}

func (X *xwindow) PolygonInv (xs, ys []int) {
  X.segments (xs, ys, false)
}

func (X *xwindow) polygonFull (xs, ys []int, n bool) {
  l:= len (xs); if len (ys) != l { return }
  p:= make ([]C.XPoint, l)
  for i:= 0; i < l; i++ {
    p[i].x, p[i].y = C.short(xs[i]), C.short(ys[i])
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopyInverted) }
  if ! X.buff { C.XFillPolygon (dpy, C.Drawable(X.win), X.gc, &p[0], C.int(l), C.Convex, C.CoordModeOrigin) }
  C.XFillPolygon (dpy, C.Drawable(X.buffer), X.gc, &p[0], C.int(l), C.Convex, C.CoordModeOrigin)
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) PolygonFull (xs, ys []int) {
  X.polygonFull (xs, ys, true)
}

func (X *xwindow) PolygonFullInv (xs, ys []int) {
  X.polygonFull (xs, ys, false)
}

func (X *xwindow) OnPolygon (xs, ys []int, a, b int, t uint) bool {
  n:= len (xs)
  if n == 0 { return false }
  if ! ok2 (xs, ys) { return false }
  if n == 1 { return xs[0] == a && ys[0] == b }
  for i:= 1; i < int(n); i++ {
    if X.OnLine (xs[i-1], ys[i-1], xs[i], ys[i], a, b, t) {
      return true
    }
  }
  return X.OnLine (xs[n-1], ys[n-1], xs[0], ys[0], a, b, t)
}

func (X *xwindow) ellipse (x, y int, a, b uint, n, f bool) {
  x0, y0:= C.int(x) - C.int(a), C.int(y) - C.int(b)
  aa, bb:= C.uint(2 * a), C.uint(2 * b)
  const a0 = C.int(0)
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ?
    if ! X.buff { C.XFillArc (dpy, C.Drawable(X.win), X.gc, x0, y0, aa, bb, 0, 64 * 360) }
    C.XFillArc (dpy, C.Drawable(X.buffer), X.gc, C.int(x0), y0, aa, bb, 0, 64 * 360)
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawArc (dpy, C.Drawable(X.win), X.gc, x0, y0, aa, bb, 0, 64 * 360) }
    C.XDrawArc (dpy, C.Drawable(X.buffer), X.gc, C.int(x0), y0, aa, bb, 0, 64 * 360)
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Circle (x, y int, r uint) {
  X.ellipse (x, y, r, r, true, false)
}

func (X *xwindow) CircleInv (x, y int, r uint) {
  X.ellipse (x, y, r, r, false, false)
}

func (X *xwindow) CircleFull (x, y int, r uint) {
  X.ellipse (x, y, r, r, true, true)
}

func (X *xwindow) CircleFullInv (x, y int, r uint) {
  X.ellipse (x, y, r, r, false, true)
}

func (X *xwindow) OnCircle (x, y int, r uint, a, b int, t uint) bool {
  return X.OnEllipse (x, y, r, r, a, b, t)
}

func (X *xwindow) arc (x, y int, r uint, a, b float64, n, f bool) {
  for a >= 360 { a -= 360 }
  for a <= -360 { a += 360 }
  x0, y0:= C.int(x) - C.int(r), C.int(y) - C.int(r)
  rr, aa, bb:= C.uint(2 * r), C.int(64 * a + 0.5), C.int(64 * b + 0.5)
  if f {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) } // C.GXcopyInverted ?
    if ! X.buff { C.XFillArc (dpy, C.Drawable(X.win), X.gc, x0, y0, rr, rr, aa, bb) }
    C.XFillArc (dpy, C.Drawable(X.buffer), X.gc, x0, y0, rr, rr, aa, bb)
  } else {
    if ! n { C.XSetFunction (dpy, X.gc, C.GXinvert) }
    if ! X.buff { C.XDrawArc (dpy, C.Drawable(X.win), X.gc, x0, y0, rr, rr, aa, bb) }
    C.XDrawArc (dpy, C.Drawable(X.buffer), X.gc, x0, y0, rr, rr, aa, bb)
  }
  if ! n { C.XSetFunction (dpy, X.gc, C.GXcopy) }
  C.XFlush (dpy)
}

func (X *xwindow) Arc (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, true, false)
}

func (X *xwindow) ArcInv (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, false, false)
}

func (X *xwindow) ArcFull (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, true, true)
}

func (X *xwindow) ArcFullInv (x, y int, r uint, a, b float64) {
  X.arc (x, y, r, a, b, false, true)
}

func (X *xwindow) Ellipse (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, true, false)
}

func (X *xwindow) EllipseInv (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, false, false)
}

func (X *xwindow) EllipseFull (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, true, true)
}

func (X *xwindow) EllipseFullInv (x, y int, a, b uint) {
  X.ellipse (x, y, a, b, false, true)
}

func dist (x, y, x1, y1 int) int {
  return int((math.Sqrt(float64((x1 - x) * (x1 - x) + (y1 - y) * (y1 - y))) + 0.5))
}

// work around Bresenham ellipse
func (X *xwindow) OnEllipse (x, y int, a, b uint, A, B int, t uint) bool {
  e:= int(math.Sqrt(math.Abs(float64(a * a) - float64(b * b))) + 0.5)
  r:= 2 * int(a); z:= 2 * dist (x, y, A, B) // if a == b
  if a > b {
    z = dist (x - e, y, A, B) + dist (x + e, y, A, B)
  }
  if a < b {
    z = dist (x, y - e, A, B) + dist (x, y + e, A, B)
    r = 2 * int(b)
  }
  return between (r, r, z, int(t))
}

func (X *xwindow) curve (xs, ys []int, xs1, ys1 *[]int) {
  m:= len (xs)
  if m == 0 || m != len (ys) { return }
  n:= ker.ArcLen (xs, ys)
  *xs1, *ys1 = make ([]int, n), make ([]int, n)
  for i:= uint(0); i < n; i++ {
    (*xs1)[i], (*ys1)[i] = ker.Bezier (xs, ys, uint(m), n, i)
  }
  C.XFlush (dpy)
}

func (X *xwindow) Curve (xs, ys []int) {
  var xs1, ys1 []int
  X.curve (xs, ys, &xs1, &ys1)
  X.Point (xs[0], ys[0])
  X.Points (xs1, ys1)
}

func (X *xwindow) CurveInv (xs, ys []int) {
  var xs1, ys1 []int
  X.curve (xs, ys, &xs1, &ys1)
  X.PointInv (xs[0], ys[0])
  X.PointsInv (xs1, ys1)
}

func (X *xwindow) OnCurve (xs, ys []int, a, b int, t uint) bool {
  var xs1, ys1 []int
  X.curve (xs, ys, &xs1, &ys1)
  if near (xs[0], ys[0], a, b, t) { return true }
  for i:= 0; i < len (xs1); i++ {
    if near (xs1[i], ys1[i], a, b, t) { return true }
  }
  return false
}
