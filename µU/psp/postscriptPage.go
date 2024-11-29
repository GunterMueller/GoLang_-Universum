package psp

// (c) Christian Maurer   v. 230112 - license see µU.go

import (
  "os"
  "strconv"
  "math"
  "µU/ker"
  . "µU/obj"
  "µU/char"
  "µU/col"
  "µU/scr"
  "µU/fontsize"
  "µU/font"
  "µU/N"
)
const (
  dx = -72
  dy = 342
)
type
  postscriptPage struct {
                   file *os.File
                        float64 "linewidth"
                        }

func new_() PostscriptPage {
  x := new (postscriptPage)
  x.float64 = 0.4
  const ppi = ker.PointsPerInch
  return x
}

func (p *postscriptPage) X (x int) float64 {
  return float64(x) / float64(scr.Wd()) * ker.A4wdPt
}

func (p *postscriptPage) Y (y int) float64 {
  y = ker.A4htPt - y
  return float64(y) / float64(scr.Wd()) * ker.A4wdPt
}

func (x *postscriptPage) Name (n string) {
  var err error
  x.file, err = os.Create (n + ".ps")
  if err != nil { panic ("os.Create error") }
  x.write ("%!PS-Adobe-2.0\n")
  x.write ("%%Creator Christian Maurer with µU/psp\n")
  x.write ("%%BoundingBox: 0 0 596 842\n") // A4
  x.write ("%%DocumentPaperSize: a4\n")
  x.write ("%%EndComments\n")
  x.write (x.f(x.float64) + " setlinewidth\n")
  x.SetFont (font.Roman)
  x.SetFontsize (fontsize.Normal)
  x.write ("72 72 translate\n")
}

func (x *postscriptPage) Fin() {
  x.write ("showpage\n")
  x.file.Close()
}

func (x *postscriptPage) SetUnit (pt float64) {
  x.write (x.f(pt) + " dup scale\n")
}

func (x *postscriptPage) Translate (l, c float64) {
  x.write (x.f(l) + " " + x.f(c) + " translate\n")
}

func (x *postscriptPage) write (s string) {
  x.file.Write (Stream(s))
}

func (x *postscriptPage) newpath() {
  x.write ("newpath\n")
}

func (x *postscriptPage) closepath() {
  x.write ("closepath\n")
}

func (x *postscriptPage) fill() {
  x.write ("fill\n")
}

func (x *postscriptPage) stroke() {
  x.write ("stroke\n")
}

func (x *postscriptPage) f (r float64) string {
  return strconv.FormatFloat (r, 'f', 4, 64)
}

func (x *postscriptPage) moveto (a, b float64) {
  x.write (x.f(a) + " " + x.f(b) + " moveto\n")
}

func (x *postscriptPage) lineto (a, b float64) {
  x.write (x.f(a) + " " + x.f(b) + " lineto\n")
}

func (x *postscriptPage) rmoveto (a, b float64) {
  x.write (x.f(a) + " " + x.f(b) + " rmoveto\n")
}

func (x *postscriptPage) rlineto (a, b float64) {
  x.write (x.f(a) + " " + x.f(b) + " rlineto\n")
}

func (x *postscriptPage) arc (x0, x1, r, a, b float64) {
  x.write (x.f(x0) + " " + x.f(x1) + " " + x.f(r) + " " + x.f(a) + " " + x.f(b) + " arc\n")
}

func (x *postscriptPage) scale (s float64) {
  x.write ("1 " + strconv.FormatFloat (s, 'f', 4, 64) + " scale\n")
}

func g (n uint8) string {
  return strconv.FormatFloat (float64(n) / 255, 'f', 4, 64)
}

func (x *postscriptPage) SetColour (c col.Colour) {
  x.write (g (c.R()) + " " + g (c.G()) + " " + g (c.B()) + " setrgbcolor\n")
}

func (x *postscriptPage) SetFont (f font.Font) {
  var s string
  switch f {
  case font.Roman:
    s = "terminus-normal 16"
  case font.Bold:
    s = "terminus-bold"
  case font.Italic:
    s = "Times-Roman-Italic" // nonsense
  }
  x.write ("/" + s + " findfont\n")
}

func (x *postscriptPage) SetFontsize (s fontsize.Size) {
  h := 0
  switch s {
  case fontsize.Tiny:
    h = 4
  case fontsize.Small:
    h = 5
  case fontsize.Normal:
    h = 8
  case fontsize.Big:
    h = 12
  case fontsize.Large:
    h = 14
  case fontsize.Huge:
    h = 16
  }
  x.write ("/terminus-normal 16 findfont\n")
  x.write (strconv.Itoa(h) + " scalefont setfont\n")
}

func (x *postscriptPage) Write (s string, x0, y0 float64) {
  x.newpath()
  x0 += dx; y0 += dy
  x.moveto (x0, y0)
  for i := 0; i < len (s); i++ {
    if char.IsLatin1 (s[i]) {
      x.write ("/" + char.Postscript (s[i]) + " glyphshow\n")
    } else {
      x.write ("(" + string(s[i]) + ") show\n")
    }
  }
  x.stroke()
}

func (x *postscriptPage) SetLinewidth (w float64) {
  x.float64 = w
  x.write (x.f(x.float64) + " setlinewidth\n")
}

func (x *postscriptPage) Point (x0, y0 float64) {
  x.newpath()
  x0 += dx; y0 += dy
  x.arc (x0, y0, x.float64, 0, 360)
  x.fill()
  x.stroke()
}

func (x *postscriptPage) Points (xs, ys []float64) {
  n := len(xs)
  if n == 0 || len(ys) != n { return }
  x.newpath()
  for i := 0; i < n; i++ {
    xs[i] += dx; ys[i] += dy
    x.arc (xs[i], ys[i], 2 * x.float64, 0, 360)
    x.fill()
  }
  x.stroke()
}

func (x *postscriptPage) Line (x1, y1, x2, y2 float64) {
println ("psp.Line")
  x.newpath()
  x.moveto (x1, y1)
  x.lineto (x2, y2)
  x.stroke()
}

func (x *postscriptPage) Lines (x0, y0, x1, y1 []float64) {
  n := len(x0)
  if n < 1 || len(y0) != n || len(x1) != n || len(y1) != n { return }
  x.newpath()
  for i := 0; i < n; i++ {
    x0[i] += dx; y0[i] += dy
    x1[i] += dx; y1[i] += dy
    x.moveto (x0[i], y0[i])
    x.lineto (x1[i], y1[i])
  }
  x.closepath()
  x.stroke()
}

func (x *postscriptPage) Segments (xs, ys []float64) {
  n := len (xs)
  if n < 1 || len (ys) != n { return }
  xs[0] += dx; ys[0] += dy
  if n == 1 {
    x.Point (xs[0], ys[0])
    return
  }
  x.newpath()
  x.moveto (xs[0], ys[0])
  for i := 1; i < n; i++ {
    xs[i] += dx; ys[i] += dy
    x.lineto (xs[i], ys[i])
  }
  x.stroke()
}

func (x *postscriptPage) Rectangle (x0, y0, w, h float64, f bool) {
  x.newpath()
  x0 += dx; y0 += dy
  x.moveto (x0, y0)
  x.rlineto (w, 0)
  x.rlineto (0, h)
  x.rlineto (-w, 0)
  x.closepath()
  if f { x.fill() }
  x.stroke()
}

func (x *postscriptPage) Polygon (xs, ys []float64, f bool) {
  n := len (xs)
  if n < 1 || len (ys) != n { return }
  xs[0] += dx; ys[0] += dy
  if n == 1 {
    x.Point (xs[0], ys[0])
    return
  }
  x.newpath()
  x.moveto (xs[0], ys[0])
  for i := 1; i < n; i++ {
    xs[i] += dx; ys[i] += dy
    x.lineto (xs[i], ys[i])
  }
  x.closepath()
  if f { x.fill() }
  x.stroke()
}

func (x *postscriptPage) Arc (x0, y0, r, a, b float64) {
  x.newpath()
  x0 += dx; y0 += dy
  x.arc (x0, y0, r, a, b)
  x.stroke()
}

func (x *postscriptPage) Circle (x0, y0, r float64, f bool) {
println ("psp.Circle")
  x0 += dx; y0 += dy
  x.newpath()
  x.arc (x0, y0, r, 0, 360)
  if f { x.fill() }
  x.stroke()
}

func (x *postscriptPage) Ellipse (x0, y0, a, b float64, f bool) {
  x0 += dx; y0 += dy
  x.write ("/ellipse { 7 dict begin\n")
  x.write ("/" + x.f(b) + " exch def\n")
  x.write ("/" + x.f(a) + " exch def\n")
  x.write ("/" + x.f(y0) + " exch def\n")
  x.write ("/" + x.f(x0) + " exch def\n")
  x.write ("/mat matrix currentmatrix def\n")
  x.write (x.f(x0) + " " + x.f(y0) + " translate\n")
  x.write (x.f(a) + " " + x.f(b) + " scale\n")
  x.write ("0 0 1 0 360 arc\n")
  x.write ("mat setmatrix\n")
  x.write ("end\n")
  x.write ("} def\n")
  x.newpath()
  x.write (x.f(x0) + " " + x.f(y0) + " " + x.f(a) + " " + x.f(b) + " ellipse\n")
  if f { x.fill() }
  x.stroke()
}

func p (t, a float64, k uint) float64 {
  if k == 0 { return a }
  if k % 2 == 0 {
    return p (t * t, a, k / 2)
  }
  return p (t * t, t * a, k / 2)
}

func (x *postscriptPage) nodes (xs, ys []float64) int {
  n := len (xs)
  if n == 0 || n != len (ys) { return 0 }
  a, b := 0., 0.
  ns := 0
  for i := 1; i < n; i++ {
    a, b = math.Abs (xs[i] - xs[i-1]), math.Abs (ys[i] - ys[i-1])
    ns += int(math.Sqrt (a * a + b * b + 0.5))
  }
  return ns
}

func bezier (t float64, k uint, xs, ys []float64) (float64, float64) {
  x, y := 0.0, 0.0
  for i := uint(0); i <= k; i++ {
    a := float64(N.Binom (k, i)) * p (1 - t, 1, k - i) * p (t, 1, i)
    x += a * xs[i]
    y += a * ys[i]
  }
  x += dx
  y += dy
  return x, y
}

func (x *postscriptPage) Curve (xs, ys []float64) {
  n := len (xs)
  if len (ys) != n { return }
  x.newpath()
  m := x.nodes (xs, ys)
  if m == 0 { return }
  xb, yb := bezier (float64(1) / float64(m), uint(n - 1), xs, ys)
  x.moveto (xb, yb)
  for i := 2; i < m; i++ {
    xb, yb = bezier (float64(i) / float64(m), uint(n - 1), xs, ys)
    x.lineto (xb, yb)
  }
  x.stroke()
}
