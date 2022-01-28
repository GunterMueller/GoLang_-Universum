package fig2

// (c) Christian Maurer   v. 220124 - license see µU.go

import (
  "math"
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/font"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/sel"
  "µU/ppm"
  "µU/psp"
)
const (
  lenName = 10
  lenText = 40
)
type
  figure struct {
                typ
         colour col.Colour
           x, y []int
         marked,
         filled bool
                string
                ppm.Image
                }
var (
  wd, ht int
  bx = box.New()
  name = []string {"Punktfolge",
                   "Strecke(n)",
                   "Polygon   ",
                   "Kurve     ",
                   "Gerade    ",
                   "Rechteck  ",
                   "Kreis     ",
                   "Ellipse   ",
                   "Text      ",
                   "Bild      "}
  fontsize = font.Normal
)

func init() {
  bx.Transparence (true)
  bx.Wd (lenText)
}

func new_() Figure {
  wd, ht = int(scr.Wd()), int(scr.Ht())
  f := new(figure)
  f.x, f.y = nil, nil
  f.marked, f.filled = false, false
  f.string = ""
  f.typ = Segments
  c, _ := col.StartCols()
  f.colour = c
  f.Image = ppm.New()
  return f
}

func newPoints (xs, ys []int, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Points
  f.colour = c
  n := len(xs)
  if n == 0 || len(ys) != n { return nil }
  f.x, f.y = make([]int, n), make([]int, n)
  for i := 0; i < n; i++ {
    f.x[i], f.y[i] = xs[i], ys[i]
  }
  return f
}

func newSegments (xs, ys []int, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Segments
  f.colour = c
  n := len(xs)
  if n == 0 || len(ys) != n { return nil }
  f.x, f.y = make([]int, n), make([]int, n)
  for i := 0; i < n; i++ {
    f.x[i], f.y[i] = xs[i], ys[i]
  }
  return f
}

func newPolygon (xs, ys []int, b bool, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Polygon
  f.colour = c
  n := len(xs)
  if n == 0 || len(ys) != n { return nil }
  f.x, f.y = make([]int, n), make([]int, n)
  for i := 0; i < n; i++ {
    f.x[i], f.y[i] = xs[i], ys[i]
  }
  f.filled = b
  return f
}

func newCurve (xs, ys []int, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Curve
  f.colour = c
  n := len(xs)
  if n == 0 || len(ys) != n { return nil }
  f.x, f.y = make([]int, n), make([]int, n)
  for i := 0; i < n; i++ {
    f.x[i], f.y[i] = xs[i], ys[i]
  }
  return f
}

func newInfLine (x, y, x1, y1 int, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = InfLine
  f.colour = c
  f.x, f.y = make([]int, 2), make([]int, 2)
  f.x[0], f.y[0] = x, y
  f.x[1], f.y[1] = x1, y1
  return f
}

func newRectangle (x, y, x1, y1 int, b bool, c col.Colour) Figure {
  if x > x1 { x, x1 = x1, x }
  if y > y1 { y, y1 = y1, y }
  f := new_().(*figure)
  f.typ = Rectangle
  f.colour = c
  f.x, f.y = make([]int, 2), make([]int, 2)
  f.x[0], f.y[0] = x, y
  f.x[1], f.y[1] = x1, y1
  f.filled = b
  return f
}

func newCircle (x, y, r int, b bool, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Circle
  f.colour = c
  f.x, f.y = make([]int, 2), make([]int, 2)
  f.x[0], f.y[0] = x, y
  f.x[1], f.y[1] = x + r, y - r
  f.filled = b
  return f
}

func newEllipse (x, y, a, b int, f bool, c col.Colour) Figure {
  e := new_().(*figure)
  e.typ = Ellipse
  e.colour = c
  e.x, e.y = make([]int, 3), make([]int, 3)
  e.x[0], e.y[0] = x, y
  e.x[1], e.y[1] = x + a, y
  e.x[2], e.y[2] = x, y - b
  e.filled = f
  return e
}

func newText (x, y int, s string, c col.Colour) Figure {
  f := new_().(*figure)
  f.typ = Text
  f.colour = c
  f.x, f.y = make([]int, 1), make([]int, 1)
  f.x[0], f.y[0] = x, y
  f.string = s
  return f
}

func newImage (x, y int, n string) Figure {
  f := new_().(*figure)
  f.typ = Image
  f.x, f.y = make([]int, 1), make([]int, 1)
  f.x[0], f.y[0] = x, y
  f.string = n
  f.Image.Load (f.string)
  W, H := f.Image.Size()
  w, h := int(W) - 1, int(H) - 1
  if f.x[0] + w <= wd && f.y[0] + h <= ht {
    f.x, f.y = append (f.x, f.x[0] + w), append (f.y, f.y[0] + h)
  } else {
    f.x, f.y = nil, nil
  }
  return f
}

func (f *figure) imp (Y Any) *figure {
  y, ok := Y.(*figure)
  if ! ok { TypeNotEqPanic (f, Y) }
  return y
}

func (f *figure) Empty() bool {
  return len(f.x) == 0
}

func (f *figure) Clr() {
  f.x, f.y = nil, nil
  f.marked, f.filled = false, false
  f.string = ""
}

func (f *figure) Typ() typ {
  return f.typ
}

func (f *figure) SetTyp (t typ) {
  f.Clr()
  f.typ = t
}

func (f *figure) Select() {
  f.Clr()
  scr.SetFontsize (font.Normal)
  n := uint(Rectangle)
  y, x := scr.MousePos()
  sel.Select1 (name, Ntypes, lenName, &n, y, x, col.LightWhite(), col.Blue())
  if n < Ntypes {
    f.typ = typ (n)
  }
}

func (f *figure) Eq (Y Any) bool {
  f1 := f.imp (Y)
  n, n1 := uint(len (f.x)), uint(len (f1.x))
  if f.typ != f1.typ || n != n1 || f.filled != f1.filled {
    return false
  }
  if n == 0 { return true } // ?
  if f.x[0] != f1.x[0] || f.y[0] != f1.y[0] {
    return false
  }
  switch f.typ {
  case Text, Image:
    if f.x[1] != f1.x[1] || f.y[1] != f1.y[1] ||
       f.string != f1.string {
      return false
    }
  default:
    for i := uint(1); i < n; i++ {
      if f.x[i] != f1.x[i] || f.y[i] != f1.y[i] {
        return false
      }
    }
  }
  return true
}

func (x *figure) Less (Y Any) bool {
  return false
}

func (f *figure) Copy (Y Any) {
  f1 := f.imp (Y)
  f.typ = f1.typ
  f.colour.Copy (f1.colour)
  n1 := uint(len(f1.x))
  f.x, f.y = make ([]int, n1), make ([]int, n1)
  for i := uint(0); i < n1; i++ {
    f.x[i] = f1.x[i]
    f.y[i] = f1.y[i]
  }
  f.filled = f1.filled
  f.marked = f1.marked
  f.string = f1.string
  if f.typ == Image {
    // TODO copy the image
  }
}

func (x *figure) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (f *figure) Pos() (int, int) {
  return f.x[0], f.y[0]
}

func (f *figure) SetPos (x, y int) {
  f.WriteInv()
  dx, dy := x - f.x[0], y - f.y[0]
  n := len(f.x)
  for i := 0; i < n; i++ {
    f.x[i] += dx
    f.y[i] += dy
  }
  f.Write()
}

var a0, b0, a1, b1 int

func (f *figure) ShowPoints (v bool) {
  switch f.typ {
  case Polygon, Rectangle, Circle:
    return
  }
  n := len(f.x)
  if f.typ == InfLine { n = 1 }
  if n == 0 { return }
  scr.ColourF (f.colour)
  x0, y0 := f.x[0], f.y[0]
  x1, y1 := f.x[1], f.y[1]
  for i := 1; i < n; i++ {
    if f.x[i] <= x0 { x0 = f.x[i] }
    if f.y[i] <= y0 { y0 = f.y[i] }
    if f.x[i] >= x1 { x1 = f.x[i] }
    if f.y[i] >= y1 { y1 = f.y[i] }
  }
  const d = 4
  if x0 >= d { x0 -= d }
  if y0 >= d { y0 -= d }
  if x1 + d <= int(scr.Wd()) { x1 += d }
  if y1 + d <= int(scr.Ht()) { y1 += d }
  if v {
    scr.SaveGr (x0, y0, uint(x1 - x0), uint(y1 - y0))
    a0, b0, a1, b1 = x0, y0, x1, y1
    for i := 0; i < n; i++ {
      scr.CircleFull (f.x[i], f.y[i], d)
    }
  } else {
    scr.RestoreGr (a0, b0, uint(a1 - a0), uint(b1 - b0))
  }
}

func (f *figure) On (a, b int, t uint) bool {
  switch f.typ {
  case Points:
    return scr.OnPoints (f.x, f.y, a, b, t)
  case Segments:
    return scr.OnSegments (f.x, f.y, a, b, t)
  case Polygon:
    return scr.OnPolygon (f.x, f.y, a, b, t)
  case Curve:
    return scr.OnCurve (f.x, f.y, a, b, t)
  case InfLine:
    return scr.OnInfLine (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
  case Rectangle:
    if f.filled {
      return scr.InRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    }
    return scr.OnRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
  case Circle:
    if f.filled {
      return scr.InCircle (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), a, b, t)
    }
    return scr.OnCircle (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), a, b, t)
  case Ellipse:
    if f.filled {
      return scr.InEllipse (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]), a, b, t)
    }
    return scr.OnEllipse (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]), a, b, t)
  case Text:
    return scr.InRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, 0)
  case Image:
    return scr.InRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, 0)
  }
  return false
}

func angle (x, y, x1, y1 int) float64 {
  a, b, c, d := float64(x), float64(y), float64(x1), float64(y1)
  return math.Acos ((a * c + b * d) / math.Sqrt ((a * a + b * b) * (c * c + d * d)))
}

func (f *figure) convex() bool {
  switch f.typ {
  case Rectangle, Circle, Ellipse, Image:
    return true
  case Polygon:
    return scr.Convex (f.x, f.y)
  }
  return false
}

func (f *figure) UnderMouse (t uint) bool {
  a, b := scr.MousePosGr()
  return f.On (a, b, t)
}

func (f *figure) focus() (int, int) {
  n := len(f.x)
  x, y := 0, 0
  for i := 0; i < n; i++ {
    x += f.x[i]
    y += f.y[i]
  }
  x /= n
  y /= n
  return x, y
}

func (f *figure) Move (dx, dy int) {
  f.WriteInv()
  n := len(f.x)
  for i := 0; i < n; i++ {
    f.x[i] += dx
    f.y[i] += dy
  }
  if f.typ == Image {
    scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
  } else {
    f.Write()
  }
}

func (f *figure) Marked() bool {
  return f.marked
}

func (f *figure) Mark (m bool) {
  f.marked = m
}

func (f *figure) SetColour (c col.Colour) {
  if f.typ != Image {
    f.colour.Copy (c)
  }
}

func (f *figure) Colour() col.Colour {
  return f.colour
}

func (f *figure) String() string {
  return name[f.typ]
}

func (f *figure) Defined (s string) bool {
  str.OffSpc1 (&s)
  str.Norm (&s, lenName)
  for t := typ(0); t < Ntypes; t++ {
    if s == name[t] {
      f.typ = t
      return true
    }
  }
  return false
}

func (f *figure) NumPoints() int {
  return len(f.x)
}

func (f *figure) writePoints() {
  scr.Points (f.x, f.y)
}

func (f *figure) writeSegments() {
  scr.Segments (f.x, f.y)
}

func (f *figure) writePolygon() {
  scr.Polygon (f.x, f.y)
  if f.filled {
    scr.PolygonFull (f.x, f.y)
  }
}

func (f *figure) writeCurve() {
  scr.Curve (f.x, f.y)
}

func (f *figure) writeInfLine() {
  scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
}

func (f *figure) writeRectangle() {
  if f.filled {
    scr.RectangleFull (f.x[0], f.y[0], f.x[1], f.y[1])
  } else {
    scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
  }
}

func (f *figure) writeCircle() {
  if f.filled {
    scr.CircleFull (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
  } else {
    scr.Circle (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
  }
}

func (f *figure) writeEllipse() {
  if f.filled {
    scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
  } else {
    scr.Ellipse (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
  }
}

func (f *figure) writeText() {
  bx.Wd (str.ProperLen (f.string))
  bx.ColourF (f.colour)
  bx.WriteGr (f.string, f.x[0], f.y[0])
}

func (f *figure) writeImage() {
  f.Image.Load (f.string)
  scr.WriteImage (f.Image.Colours(), f.x[0], f.y[0])
}

func (f *figure) Write() {
  if f.Empty() { return }
  scr.ColourF (f.colour)
  switch f.typ {
  case Points:
    f.writePoints()
  case Segments:
    f.writeSegments()
  case Polygon:
    f.writePolygon()
  case Curve:
    f.writeCurve()
  case InfLine:
    f.writeInfLine()
  case Rectangle:
    f.writeRectangle()
  case Circle:
    f.writeCircle()
  case Ellipse:
    f.writeEllipse()
  case Text:
    f.writeText()
  case Image:
    f.writeImage()
  }
}

func (f *figure) writePointsInv() {
  scr.PointsInv (f.x, f.y)
}

func (f *figure) writeSegmentsInv() {
  scr.SegmentsInv (f.x, f.y)
}

func (f *figure) writePolygonInv() {
  scr.PolygonInv (f.x, f.y)
  if f.filled {
    scr.PolygonFullInv (f.x, f.y)
  }
}

func (f *figure) writeCurveInv() {
  scr.CurveInv (f.x, f.y)
}

func (f *figure) writeInfLineInv() {
  scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
}

func (f *figure) writeRectangleInv() {
  if f.filled {
    scr.RectangleFullInv (f.x[0], f.y[0], f.x[1], f.y[1])
  } else {
    scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
  }
}

func (f *figure) writeCircleInv() {
  if f.filled {
    scr.CircleFullInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
  } else {
    scr.CircleInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
  }
}

func (f *figure) writeEllipseInv() {
  if f.filled {
    scr.EllipseFullInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
  } else {
    scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
  }
}

func (f *figure) writeTextInv() {
  scr.Transparence (true)
  scr.WriteInvGr (f.string, f.x[0], f.y[0])
}

func (f *figure) writeImageInv() {
//  ppm.Get (f.string, uint(f.x[0]), uint(f.y[0]))
}

func (f *figure) WriteInv() {
  if f.Empty() { return }
  scr.ColourF (f.colour)
  switch f.typ {
  case Points:
    f.writePointsInv()
  case Segments:
    f.writeSegmentsInv()
  case Polygon:
    f.writePolygonInv()
  case Curve:
    f.writeCurveInv()
  case InfLine:
    f.writeInfLineInv()
  case Rectangle:
    f.writeRectangleInv()
  case Circle:
    f.writeCircleInv()
  case Ellipse:
    f.writeEllipseInv()
  case Text:
    f.writeTextInv()
  case Image:
//  images cannot be written inversely yet.
  }
}

func (f *figure) editPoints() {
  xm, ym := 0, 0
  scr.ColourF (f.colour)
  loop:
  for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.To:
      break loop
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      scr.Point (xm, ym)
    }
  }
}

func (f *figure) editSegments() {
  scr.ColourF (f.colour)
  xm, ym, n := f.x[0], f.y[0], 0
  loop:
    for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.This:
      xm, ym = scr.MousePosGr()
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      break loop
    case kbd.Go:
      n = len(f.x)
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      scr.Line (f.x[n-1], f.y[n-1], xm, ym)
    case kbd.Here:
      if n == 0 { continue }
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      n = len(f.x)
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      scr.Line (f.x[n-1], f.y[n-1], f.x[n], f.y[n])
    }
  }
}

func (f *figure) editCurve() {
  scr.ColourF (f.colour)
  xm, ym, n := f.x[0], f.y[0], 0
  loop:
    for {
    c, _ := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.This:
      xm, ym = scr.MousePosGr()
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      break loop
    case kbd.Go:
      n = len(f.x)
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      scr.Line (f.x[n-1], f.y[n-1], xm, ym)
    case kbd.Here:
      if n == 0 { continue }
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      n = len(f.x)
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      scr.Line (f.x[n-1], f.y[n-1], f.x[n], f.y[n])
    }
  }
}

func (f *figure) changeCurve() {
  xm, ym := scr.MousePosGr()
  if f.On (xm, ym, 7) {
    f.ShowPoints (true)
  }
  scr.MousePointer (true)
  loop:
  for {
    i, ok := f.pointUnderMouse()
//    scr.CircleInv (f.x[i], f.y[i], 4)
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      if ok {
        xm, ym = scr.MousePosGr()
        scr.CircleInv (f.x[i], f.y[i], 4)
        f.x[i], f.y[i] = xm, ym
        f.Write()
      }
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) editPolygon() {
  scr.ColourF (f.colour)
  xm, ym, n := f.x[0], f.y[0], 0
  loop:
    for {
    c, d := kbd.Command()
    scr.MousePointer (true)
    switch c {
    case kbd.Here:
      if n == 0 { continue }
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      n = len(f.x)
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      scr.Line (f.x[n-1], f.y[n-1], f.x[n], f.y[n])
    case kbd.Go:
      n = len(f.x)
      scr.LineInv (f.x[n-1], f.y[n-1], xm, ym)
      xm, ym = scr.MousePosGr()
      scr.Line (f.x[n-1], f.y[n-1], xm, ym)
    case kbd.This:
      xm, ym = scr.MousePosGr()
      f.x, f.y = append (f.x, xm), append (f.y, ym)
      f.filled = d > 0
      if ker.UnderC() { // console cannot fill polygons with shape C.Complex
        f.filled = d > 0 && f.convex()
      }
      break loop
    }
  }
}

func (f *figure) changeSegPol() {
  f.ShowPoints (true)
  i, ok := f.pointUnderMouse()
  if ! ok { return }
  loop:
  for {
    scr.MousePointer (true)
    xm, ym := scr.MousePosGr()
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      f.x[i], f.y[i] = xm, ym
      f.Write()
    case kbd.To:
      break loop
    }
  }
  f.ShowPoints (false)
}

func (f *figure) editInfLine() {
  scr.ColourF (f.colour)
  f.x, f.y = append (f.x, f.x[0]), append (f.y, f.x[0])
  f.x[1], f.y[1] = f.x[0] + 8, f.y[0]
  loop:
  for {
    c, _ := kbd.Command()
    scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
    switch c {
    case kbd.Drag:
      if f.x[1] != f.x[0] || f.y[1] != f.y[0] {
        scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
      }
      f.x[1], f.y[1] = scr.MousePosGr()
      if f.x[1] != f.x[0] || f.y[1] != f.y[0] {
        scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
      }
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) changeInfLine() {
  xm, ym := scr.MousePosGr()
  if scr.OnInfLine (f.x[0], f.y[0], f.x[1], f.y[0], xm, ym, 7) {
    scr.ColourF (col.Yellow())
    scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[0])
  }
  loop:
  for {
    scr.MousePointer (true)
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      if f.x[1] != f.x[0] || f.y[1] != f.y[0] {
        f.x[1], f.y[1] = xm, ym
        f.Write()
      }
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) editRectangle() {
  scr.ColourF (f.colour)
  f.x = append (f.x, f.x[0])
  f.y = append (f.y, f.x[0])
  loop:
  for {
    c, d := kbd.Command()
    switch c {
    case kbd.Drag:
      scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
      f.x[1], f.y[1] = scr.MousePosGr()
      scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
    case kbd.To:
      f.filled = d > 0
      break loop
    }
  }
}

func (f *figure) changeRectangle() {
  xm, ym := scr.MousePosGr()
  oben   := scr.OnLine (f.x[0], f.y[0], f.x[1], f.y[0], xm, ym, 7)
  unten  := scr.OnLine (f.x[0], f.y[1], f.x[1], f.y[1], xm, ym, 7)
  rechts := scr.OnLine (f.x[1], f.y[1], f.x[1], f.y[0], xm, ym, 7)
  links  := scr.OnLine (f.x[0], f.y[0], f.x[0], f.y[1], xm, ym, 7)
  loop:
  for {
    scr.MousePointer (true)
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      if oben {
        f.y[0] = ym
      }
      if unten {
        f.y[1] = ym
      }
      if rechts {
        f.x[1] = xm
      }
      if links {
        f.x[0] = xm
      }
      f.WriteInv()
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) editCircle() {
  scr.ColourF (f.colour)
  f.x = append (f.x, 0); f.y = append (f.y, 0)
  loop:
  for {
    c, d := kbd.Command()
    switch c {
    case kbd.Drag:
      scr.CircleInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
      xm, ym := scr.MousePosGr()
      dx := float64((xm - f.x[0]) * (xm - f.x[0]))
      dy := float64((ym - f.y[0]) * (ym - f.y[0]))
      f.x[1] = f.x[0] + int(math.Sqrt (dx + dy) + 0.5)
      f.y[1] = f.y[0]
      scr.Circle (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
    case kbd.To:
      f.filled = d > 0
      break loop
    }
  }
}

func (f *figure) changeCircle() {
  loop:
  for {
    scr.MousePointer (true)
    xm, ym := scr.MousePosGr()
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      r := dist (f.x[0], f.y[0], xm, ym)
      f.x[1] = f.x[0] + r
      f.Write()
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) editEllipse() {
  scr.ColourF (f.colour)
  f.x, f.y = append (f.x, 0), append (f.y, 0)
  f.x, f.y = append (f.x, 0), append (f.y, 0)
  loop:
  for {
    c, d := kbd.Command()
    switch c {
    case kbd.Drag:
      scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
      xm, ym := scr.MousePosGr()
      f.x[2] = f.x[0]
      f.y[1] = f.y[0]
      if xm >= f.x[0] {
        f.x[1] = xm
      } else {
        f.x[1] = 2 * f.x[0] - xm
      }
      if ym >= f.y[0] {
        f.y[2] = 2 * f.y[0] - ym
      } else {
        f.y[2] = ym
      }
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
    case kbd.To:
      f.filled = d > 0
      break loop
    }
  }
}

func (f *figure) changeEllipse() {
  i, ok := f.pointUnderMouse()
  if ! ok { return }
  if i == 0 { return }
  loop:
  for {
    scr.MousePointer (true)
    xm, ym := scr.MousePosGr()
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
      xm, ym = scr.MousePosGr()
      if i == 1 {
        if xm >= f.x[0] {
          f.x[1] = xm
        } else {
          f.x[1] = 2 * f.x[0] - xm
        }
      } else { // i == 2
        if ym >= f.y[0] {
          f.y[2] = 2 * f.y[0] - ym
        } else {
          f.y[2] = ym
        }
      }
      f.Write()
    case kbd.To:
      break loop
    }
  }
}

func (f *figure) editText() {
  ht1, wd1 := int(scr.Ht1()), int(scr.Wd1())
  scr.MousePointer (false)
  bx.Wd (lenText)
  bx.ColourF (f.colour)
  x1 := f.x[0] + int(lenText * scr.Wd1())
  if x1 >= wd { x1 = wd - 1 }
  y1 := f.y[0] + int(scr.Ht1()) - 1
  if y1 >= ht { y1 = ht - 1 }
  scr.SaveGr (f.x[0], f.y[0], uint(x1 - f.x[0]), uint(y1 - f.y[0]))
  bx.Transparence (false)
  bx.EditGr (&f.string, f.x[0], f.y[0])
  str.OffSpc1 (&f.string)
  k := len(f.string)
  bx.Transparence (true)
  scr.RestoreGr (f.x[0], f.y[0], uint(x1 - f.x[0]), uint(y1 - f.y[0]))
  switch c, _ := kbd.LastCommand(); c {
  case kbd.Enter:
    bx.Transparence (true)
    scr.RestoreGr (f.x[0], f.y[0], uint(x1 - f.x[0]), uint(y1 - f.y[0]))
    bx.WriteGr (f.string, f.x[0], f.y[0])
    f.x, f.y = append (f.x, f.x[0] + k * wd1), append (f.y, f.y[0] + ht1)
    f.Write()
  default:
    f.Clr()
  }
  scr.MousePointer (true)
}

func (f *figure) changeText() {
  f.editText()
}

func (f *figure) EditImage (n string) {
  f.x, f.y = make ([]int, 1), make ([]int, 1)
  f.x[0], f.y[0] = scr.MousePosGr()
  f.string = n
  f.Image.Load (f.string)
  w, h := f.Image.Size()
//  f.x, f.y = append (f.x, f.x[0] + int(w) - 1), append (f.y, f.y[0] + int(h) - 1)
  f.x, f.y = append (f.x, f.x[0] + int(w)), append (f.y, f.y[0] + int(h))
}

func (f *figure) pointUnderMouse() (uint, bool) {
  xm, ym := scr.MousePosGr()
  n := len(f.x)
  for i := 0; i < n; i++ {
    if scr.InCircle (f.x[i], f.y[i], 4, xm, ym, 4) {
      return uint(i), true
    }
  }
  return uint(n), false
}

func (f *figure) Edit() {
  if f.Empty() {
    scr.ColourF (f.colour)
    f.x, f.y = make ([]int, 1), make ([]int, 1)
    f.x[0], f.y[0] = scr.MousePosGr()
    switch f.typ {
    case Points:
      f.editPoints()
    case Segments:
      f.editSegments()
    case Polygon:
      f.editPolygon()
    case Curve:
      f.editCurve()
    case InfLine:
      f.editInfLine()
    case Rectangle:
      f.editRectangle()
    case Circle:
      f.editCircle()
    case Ellipse:
      f.editEllipse()
    case Text:
      f.editText()
    }
    if f.x == nil {
      f.Clr()
    }
    return
  }
  f.Change()
}

func (f *figure) Change() {
  switch f.typ {
  case Points:
    return // Points cannot be changed
  case Segments, Polygon:
    f.changeSegPol()
  case Curve:
    f.changeCurve()
  case InfLine:
    f.changeInfLine()
  case Rectangle:
    f.changeRectangle()
  case Circle:
    f.changeCircle()
  case Ellipse:
    f.changeEllipse()
  case Text:
    f.changeText()
  case Image:
    return // images cannot be changed
  }
}

func (f *figure) SetFont (s font.Size) {
  if f.typ == Text {
    fontsize = s
  }
}

func (f *figure) printPoints (p psp.PostscriptPage) {
  n := uint(len (f.x))
  x, y := make ([]float64, n), make ([]float64, n)
  for i := uint(0); i < n; i++ {
    x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
  }
  p.Points (x, y)
}

func (f *figure) printSegments (p psp.PostscriptPage) {
  n := uint(len (f.x))
  x, y := make ([]float64, n), make ([]float64, n)
  for i := uint(0); i < n; i++ {
    x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
  }
  p.Segments (x, y)
}

func (f *figure) printPolygon (p psp.PostscriptPage) {
  n := uint(len (f.x))
  x, y := make ([]float64, n), make ([]float64, n)
  for i := uint(0); i < n; i++ {
    x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
  }
  p.Polygon (x, y, f.filled)
}

func (f *figure) printCurve (p psp.PostscriptPage) {
  n := uint(len (f.x))
  x, y := make ([]float64, n), make ([]float64, n)
  for i := uint(0); i < n; i++ {
    x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
  }
  p.Curve (x, y)
}

func (f *figure) printInfLine (p psp.PostscriptPage) {
  x, y, x1, y1 := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1]), p.Y (f.y[1])
  p.Line (x, y, x1, y1)
}

func (f *figure) printRectangle (p psp.PostscriptPage) {
  x, y, x1, y1 := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1]), p.Y (f.y[1])
  p.Rectangle (x, y, x1 - x, y1 - y, f.filled)
}

func (f *figure) printCircle (p psp.PostscriptPage) {
  x, y, r := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1] - f.x[0])
  p.Circle (x, y, r, f.filled)
}

func (f *figure) printEllipse (p psp.PostscriptPage) {
  x, y, a, b := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1] - f.x[0]), p.X (f.y[0] - f.y[2])
  p.Ellipse (x, y, a, b, f.filled)
}

func (f *figure) printText (p psp.PostscriptPage) {
  x, y := p.X (f.x[0]), p.Y (f.y[0])
  p.SetFontsize (fontsize)
  p.Write (f.string, x, y)
}

func (f *figure) printImage (p psp.PostscriptPage) {
  f.Image.Print (f.x[0], f.y[0])
}

func (f *figure) Print (p psp.PostscriptPage) {
  if f.Empty() { return }
  p.SetColour (f.colour)
  switch f.typ {
  case Points:
    f.printPoints (p)
  case Segments:
    f.printSegments (p)
  case Polygon:
    f.printPolygon (p)
  case Curve:
    f.printCurve (p)
  case InfLine:
    f.printInfLine (p)
  case Rectangle:
    f.printRectangle (p)
  case Circle:
    f.printCircle (p)
  case Ellipse:
    f.printEllipse (p)
  case Text:
    f.printText (p)
  case Image:
    f.printImage (p)
  }
}

func (f *figure) Codelen() uint {
  n := uint(1) // f.typ
  n += f.colour.Codelen()
  n += C0 // len(f.x)
  switch f.typ {
  case Text:
    n += 2 * C0 + 1 + uint(len (f.string))
  case Image:
    n += 4 * C0 + 1 + uint(len (f.string))
  default:
    n += 2 * uint(len (f.x)) * C0
  }
  n += 2 * C0 // Reserve
  return uint(n)
}

func (f *figure) Encode() Stream {
  bs := make (Stream, f.Codelen())
  a := uint(0)
  bs[a] = byte(f.typ)
  a++
  copy (bs[a:a+f.colour.Codelen()], f.colour.Encode())
  a += f.colour.Codelen()
  n := uint(len(f.x))
  if f.typ == Text || f.typ == Image {
    n = uint(len(f.string))
  }
  copy (bs[a:a+C0], Encode (n))
  a += C0
  switch f.typ {
  case Text, Image:
    copy (bs[a:a+C0], Encode (f.x[0]))
    a += C0
    copy (bs[a:a+C0], Encode (f.y[0]))
    a += C0
    if f.typ == Image {
      copy (bs[a:a+C0], Encode (f.x[1]))
      a += C0
      copy (bs[a:a+C0], Encode (f.y[1]))
      a += C0
    }
    copy (bs[a:a+n], Stream(f.string))
    a += n
  default:
    for i := uint(0); i < n; i++ {
      copy (bs[a:a+C0], Encode (f.x[i]))
      a += C0
      copy (bs[a:a+C0], Encode (f.y[i]))
      a += C0
    }
  }
  bs[a] = 0
  if f.filled { bs[a] ++ }
  if f.marked { bs[a] += 2 }
  return bs
}

func (f *figure) Decode (bs Stream) {
  a := uint(0)
  f.typ = typ(bs[a])
  a++
  f.colour.Decode (bs[a:a+f.colour.Codelen()])
  a += f.colour.Codelen()
  n := Decode (uint(0), bs[a:a+C0]).(uint)
  a += C0
  if f.typ < Text {
    f.x, f.y = make ([]int, n), make ([]int, n)
    for i := uint(0); i < n; i++ {
      f.x[i] = Decode (f.x[i], bs[a:a+C0]).(int)
      a += C0
      f.y[i] = Decode (f.y[i], bs[a:a+C0]).(int)
      a += C0
    }
  } else { // Text, Image
    f.x, f.y = make ([]int, 2), make ([]int, 2)
    f.x[0] = Decode (f.x[0], bs[a:a+C0]).(int)
    a += C0
    f.y[0] = Decode (f.y[0], bs[a:a+C0]).(int)
    a += C0
    if f.typ == Image {
      f.x[1] = Decode (f.x[1], bs[a:a+C0]).(int)
      a += C0
      f.y[1] = Decode (f.y[1], bs[a:a+C0]).(int)
      a += C0
    }
    f.string = string(bs[a:a+n])
    a += n
    if f.typ == Text {
      f.x[1] = f.x[0] + int(scr.Wd1()) * int(n) - 1
      f.y[1] = f.y[0] + int(scr.Ht1()) - 1
    }
  }
  f.filled = bs[a] % 2 == 1
  f.marked = (bs[a] / 2) % 2 == 1
}

func dist (x, y, x1, y1 int) int {
  return int(math.Sqrt(float64((x-x1)*(x-x1)+(y-y1)*(y-y1)))+0.5)
}
