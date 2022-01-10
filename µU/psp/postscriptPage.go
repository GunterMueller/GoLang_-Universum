package psp

// (c) Christian Maurer   v. 220109 - license see µU.go

import (
  "os"
  "strconv"
  "math"
  "µU/ker"
  . "µU/obj"
  "µU/char"
  "µU/col"
  "µU/scr"
  "µU/font"
  "µU/n"
)
type
  postscriptPage struct {
                   file *os.File
                        float64 "linewidth"
                        }

func new_() PostscriptPage {
  p := new (postscriptPage)
  p.float64 = 0.4
  const ppi = ker.PointsPerInch
  return p
}

func (p *postscriptPage) X (x int) float64 {
  x -= 120 // XXX suboptimal
  return float64(x) / float64(scr.Wd()) * ker.A4wdPt
}

func (p *postscriptPage) Y (y int) float64 {
  y = ker.A4htPt - y
  y += 600 // XXX supoptimal
  return float64(y) / float64(scr.Wd()) * ker.A4wdPt
}

func (p *postscriptPage) Name (n string) {
  var err error
  p.file, err = os.Create (n + ".ps")
  if err != nil { panic ("os.Create error") }
  p.write ("%!PS-Adobe-2.0\n")
  p.write ("%%Creator µU/psp.go (c) Christian Maurer \n")
  p.write ("%%BoundingBox: 0 0 596 842 \n") // A4
  p.write ("%%DocumentPaperSize: a4\n")
  p.write ("%%EndComments\n")
  p.write (p.f(p.float64) + " setlinewidth\n")
  p.SetFont (font.Normal)
  p.write ("72 72 translate\n")
}

func (p *postscriptPage) Fin() {
  p.write ("showpage\n")
  p.file.Close()
}

func (p *postscriptPage) SetUnit (pt float64) {
  p.write (p.f(pt) + " dup scale\n")
}

func (p *postscriptPage) Translate (l, b float64) {
  p.write (p.f(l) + " " + p.f(b) + " translate\n")
}

func (p *postscriptPage) write (s string) {
  p.file.Write (Stream(s))
}

func (p *postscriptPage) newpath() {
  p.write ("newpath\n")
}

func (p *postscriptPage) closepath() {
  p.write ("closepath\n")
}

func (p *postscriptPage) fill() {
  p.write ("fill\n")
}

func (p *postscriptPage) stroke() {
  p.write ("stroke\n")
}

func (p *postscriptPage) f (r float64) string {
  return strconv.FormatFloat (r, 'f', 4, 64)
}

func (p *postscriptPage) moveto (a, b float64) {
  p.write (p.f(a) + " " + p.f(b) + " moveto\n")
}

func (p *postscriptPage) lineto (a, b float64) {
  p.write (p.f(a) + " " + p.f(b) + " lineto\n")
}

func (p *postscriptPage) rmoveto (a, b float64) {
  p.write (p.f(a) + " " + p.f(b) + " rmoveto\n")
}

func (p *postscriptPage) rlineto (a, b float64) {
  p.write (p.f(a) + " " + p.f(b) + " rlineto\n")
}

func (p *postscriptPage) arc (x0, x1, r, a, b float64) {
  p.write (p.f(x0) + " " + p.f(x1) + " " + p.f(r) + " " + p.f(a) + " " + p.f(b) + " arc\n")
}

func (p *postscriptPage) scale (s float64) {
  p.write ("1 " + strconv.FormatFloat (s, 'f', 4, 64) + " scale\n")
}

func g (n uint8) string {
  return strconv.FormatFloat (float64(n) / 255, 'f', 4, 64)
}

func (p *postscriptPage) SetColour (c col.Colour) {
  p.write (g (c.R()) + " " + g (c.G()) + " " + g (c.B()) + " setrgbcolor\n")
}

func (p *postscriptPage) SetFont (s font.Size) {
  h := 0
  switch s {
  case font.Tiny:
    h = 4
  case font.Small:
    h = 5
  case font.Normal:
    h = 8
  case font.Big:
    h = 12
  case font.Large:
    h = 14
  case font.Huge:
    h = 16
  }
  p.write ("/terminus-normal 16 findfont\n")
  p.write (strconv.Itoa(h) + " scalefont setfont\n")
}

func (p *postscriptPage) Write (s string, x0, y0 float64) {
  p.newpath()
  p.moveto (x0, y0)
  for i := 0; i < len (s); i++ {
    if char.IsLatin1 (s[i]) {
      p.write ("/" + char.Postscript (s[i]) + " glyphshow\n")
    } else {
      p.write ("(" + string(s[i]) + ") show\n")
    }
  }
  p.stroke()
}

func (p *postscriptPage) SetLinewidth (w float64) {
  p.float64 = w
  p.write (p.f(p.float64) + " setlinewidth\n")
}

func (p *postscriptPage) Point (x1, y1 float64) {
  p.newpath()
  p.arc (x1, y1, p.float64, 0, 360)
panic ("Point")
  p.fill()
  p.stroke()
}

func (p *postscriptPage) Points (xs, ys []float64) {
  n := len(xs)
  if n == 0 || len(ys) != n { return }
  p.newpath()
  for i := 0; i < n; i++ {
    p.arc (xs[i], ys[i], 2 * p.float64, 0, 360)
panic ("Points")
    p.fill()
  }
  p.stroke()
}

func (p *postscriptPage) Line (x1, y1, x2, y2 float64) {
  p.newpath()
  p.moveto (x1, y1)
  p.lineto (x2, y2)
  p.stroke()
}

func (p *postscriptPage) Lines (x0, y0, x1, y1 []float64) {
  n := len(x0)
  if n < 1 || len(y0) != n || len(x1) != n || len(y1) != n { return }
  p.newpath()
  for i := 0; i < n; i++ {
    p.moveto (x0[i], y0[i])
    p.lineto (x1[i], y1[i])
  }
  p.closepath()
  p.stroke()
}

func (p *postscriptPage) Segments (xs, ys []float64) {
  n := len (xs)
  if n < 1 || len (ys) != n { return }
  if n == 1 {
    p.Point (xs[0], ys[0])
    return
  }
  p.newpath()
  p.moveto (xs[0], ys[0])
  for i := 1; i < n; i++ {
    p.lineto (xs[i], ys[i])
  }
  p.stroke()
}

func (p *postscriptPage) Rectangle (x0, y0, w, h float64, f bool) {
  p.newpath()
  p.moveto (x0, y0)
  p.rlineto (w, 0)
  p.rlineto (0, h)
  p.rlineto (-w, 0)
  p.closepath()
  if f { p.fill() }
  p.stroke()
}

func (p *postscriptPage) Polygon (xs, ys []float64, f bool) {
  n := len (xs)
  if n < 1 || len (ys) != n { return }
  if n == 1 {
    p.Point (xs[0], ys[0])
    return
  }
  p.newpath()
  p.moveto (xs[0], ys[0])
  for i := 1; i < n; i++ {
    p.lineto (xs[i], ys[i])
  }
  p.closepath()
  if f { p.fill() }
  p.stroke()
}

func (p *postscriptPage) Arc (x0, y0, r, a, b float64) {
  p.newpath()
  p.arc (x0, y0, r, a, b)
panic ("Arc")
  p.stroke()
}

func (p *postscriptPage) Circle (x0, y0, r float64, f bool) {
  p.newpath()
  p.arc (x0, y0, r, 0, 360)
  if f { p.fill() }
  p.stroke()
}

func (p *postscriptPage) Ellipse (x0, y0, a, b float64, f bool) {
  p.write ("/ellipse { 7 dict begin\n")
  p.write ("/" + p.f(b) + " exch def\n")
  p.write ("/" + p.f(a) + " exch def\n")
  p.write ("/" + p.f(y0) + " exch def\n")
  p.write ("/" + p.f(x0) + " exch def\n")
  p.write ("/mat matrix currentmatrix def\n")
  p.write (p.f(x0) + " " + p.f(y0) + " translate\n")
  p.write (p.f(a) + " " + p.f(b) + " scale\n")
  p.write ("0 0 1 0 360 arc\n")
  p.write ("mat setmatrix\n")
  p.write ("end\n")
  p.write ("} def\n")
  p.newpath()
  p.write (p.f(x0) + " " + p.f(y0) + " " + p.f(a) + " " + p.f(b) + " ellipse\n")
  if f { p.fill() }
  p.stroke()
}

func p (t, a float64, k uint) float64 {
  if k == 0 { return a }
  if k % 2 == 0 {
    return p (t * t, a, k / 2)
  }
  return p (t * t, t * a, k / 2)
}

func (p *postscriptPage) nodes (xs, ys []float64) int {
  l := len (xs)
  if l == 0 || l != len (ys) { return 0 }
  n := 0
  for i := 1; i < l; i++ {
    dx, dy := math.Abs (xs[i] - xs[i-1]), math.Abs (ys[i] - ys[i-1])
    n += int(math.Sqrt (dx * dx + dy * dy + 0.5))
  }
  return n
}

func bezier (t float64, k uint, xs, ys []float64) (float64, float64) {
  var x, y float64
  for i := uint(0); i <= k; i++ {
    a := float64(n.Binom (k, i)) * p (1 - t, 1, k - i) * p (t, 1, i)
    x += a * xs[i]
    y += a * ys[i]
  }
  return x, y
}

func (p *postscriptPage) Curve (xs, ys []float64) {
  n := len (xs)
  if len (ys) != n { return }
  p.newpath()
  m := p.nodes (xs, ys)
  if m == 0 { return }
  p.moveto (xs[0], ys[0])
  for i := 1; i < m; i++ {
    xb, yb := bezier (float64(i) / float64(m), uint(n - 1), xs, ys)
    p.lineto (xb, yb)
  }
  p.stroke()
}
