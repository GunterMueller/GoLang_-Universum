package fig2

// (c) Christian Maurer   v. 220111 - license see µU.go

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
  f.string = str.New (lenText)
  f.string = ""
  f.typ = Segments
  c, _ := col.StartCols()
  f.colour = c
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
    scr.SaveGr (x0, y0, x1, y1)
    a0, b0, a1, b1 = x0, y0, x1, y1
    for i := 0; i < n; i++ {
      scr.CircleFull (f.x[i], f.y[i], d)
    }
  } else {
    scr.RestoreGr (a0, b0, a1, b1)
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

func (f *figure) remainsOnEboard (dx, dy int) bool {
  n := len(f.x)
  W, H := scr.Wd(), scr.Ht()
  w, h := int(W), int(H)
  switch f.typ {
  case Points, Segments, Polygon, Curve, InfLine, Rectangle:
    x, y := f.focus()
    return x + dx >= 0 && x + dx <= w && y + dy >= 0 && y + dy <= h
  case Circle, Ellipse:
    n = 1 // only the center has to remain on the eboard
  }
  for i := 0; i < n; i++ {
    if f.x[i] + dx < 0 || f.x[i] + dx >= w || f.y[i] + dy < 0 || f.y[i] + dy >= h {
      return false
    }
  }
  return true
}

func (f *figure) Move (dx, dy int) {
  if f.remainsOnEboard (dx, dy) {
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

func (f *figure) Write() {
  if f.Empty() { return }
  scr.ColourF (f.colour)
  switch f.typ {
  case Points:
    scr.Points (f.x, f.y)
  case Segments:
    scr.Segments (f.x, f.y)
  case Polygon:
    scr.Polygon (f.x, f.y)
    if f.filled {
      scr.PolygonFull (f.x, f.y)
    }
  case Curve:
    scr.Curve (f.x, f.y)
  case InfLine:
    scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFull (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle:
    if f.filled {
      scr.CircleFull (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
    } else {
      scr.Circle (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
    } else {
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
    }
  case Text:
    bx.Wd (str.ProperLen (f.string))
    bx.ColourF (f.colour)
    bx.WriteGr (f.string, f.x[0], f.y[0])
  case Image:
    ppm.Get (f.string, uint(f.x[0]), uint(f.y[0]))
  }
}

func (f *figure) WriteInv() {
  if f.Empty() { return }
  switch f.typ {
  case Points:
    scr.PointsInv (f.x, f.y)
  case Segments:
    scr.SegmentsInv (f.x, f.y)
  case Polygon:
    if f.filled {
      scr.PolygonFullInv (f.x, f.y)
    } else {
      scr.PolygonInv (f.x, f.y)
    }
  case Curve:
    scr.CurveInv (f.x, f.y)
  case InfLine:
    scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFullInv (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle:
    if f.filled {
      scr.CircleFullInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
    } else {
      scr.CircleInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFullInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
    } else {
      scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1] - f.x[0]), uint(f.y[0] - f.y[2]))
    }
  case Text:
    scr.Transparence (true)
    scr.WriteInvGr (f.string, f.x[0], f.y[0])
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

func (f *figure) changeCurve() { // TODO
  xm, ym := scr.MousePosGr()
  if f.On (xm, ym, 7) {
    f.ShowPoints (true)
  }
  loop:
  for {
//  i, ok := f.pointUnderMouse()
//  if ! ok { return }
    scr.MousePointer (true)
    xm, ym = scr.MousePosGr()
    c, _ := kbd.Command()
    switch c {
    case kbd.Drag:
/*/
      xm, ym = scr.MousePosGr()
      f.x[i], f.y[i] = xm, ym
      f.Write()
/*/
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
  scr.SaveGr (f.x[0], f.y[0], x1, y1)
  bx.Transparence (false)
  bx.EditGr (&f.string, f.x[0], f.y[0])
  str.OffSpc1 (&f.string)
  k := len(f.string)
  bx.Transparence (true)
  scr.RestoreGr (f.x[0], f.y[0], x1, y1)
  switch c, _ := kbd.LastCommand(); c {
  case kbd.Enter:
    bx.Transparence (true)
    scr.RestoreGr (f.x[0], f.y[0], x1, y1)
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
  W, H := ppm.Size (f.string)
  w, h := int(W) - 1, int(H) - 1
  if f.x[0] + w <= wd && f.y[0] + h <= ht {
    f.x, f.y = append (f.x, f.x[0] + w), append (f.y, f.y[0] + h)
  } else {
    f.x, f.y = nil, nil
  }
}

func (f *figure) pointUnderMouse() (uint, bool) {
  xm, ym := scr.MousePosGr()
  n := len(f.x)
  for i := 0; i < n; i++ {
    if scr.InCircle (f.x[i], f.y[i], 4, xm, ym, 4) {
      return uint(i), true
    }
  }
  return 0, false
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

func delta (s string, x int) {
  println (s, float64 (x) / float64(scr.Wd()) * 596) // ker.A4wdPt
}

func (f *figure) Print (p psp.PostscriptPage) {
  if f.Empty() { return }
  n := uint(len (f.x))
  p.SetColour (f.colour)
  switch f.typ {
  case Points:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
    }
    p.Points (x, y)
  case Segments:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
    }
    p.Segments (x, y)
  case Polygon:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
    }
    p.Polygon (x, y, f.filled)
  case Curve:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.X (f.x[i]), p.Y (f.y[i])
    }
    p.Curve (x, y)
  case InfLine:
    x, y, x1, y1 := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1]), p.Y (f.y[1])
    p.Line (x, y, x1, y1)
  case Rectangle:
    x, y, x1, y1 := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1]), p.Y (f.y[1])
    p.Rectangle (x, y, x1 - x, y1 - y, f.filled)
  case Circle:
    x, y, r := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1] - f.x[0])
    p.Circle (x, y, r, f.filled)
  case Ellipse:
    x, y, a, b := p.X (f.x[0]), p.Y (f.y[0]), p.X (f.x[1] - f.x[0]), p.X (f.y[0] - f.y[2])
    p.Ellipse (x, y, a, b, f.filled)
  case Text:
    x, y := p.X (f.x[0]), p.Y (f.y[0])
    p.SetFontsize (fontsize)
    p.Write (f.string, x, y)
  case Image:
// TODO print the image
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
