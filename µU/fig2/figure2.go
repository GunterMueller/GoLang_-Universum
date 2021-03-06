package fig2

// (c) Christian Maurer   v. 210408 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/font"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/sel"
  "µU/ppm"
  "µU/psp"
)
const
  lenText = 40 // maximal len of text
type (
  figure struct {
                Type
         colour col.Colour
           x, y []int
         marked,
//        moved,
         filled bool
                string
                }
)
var (
  xx, yy int
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
)

func init() {
  bx.Transparence (true)
  bx.Wd (lenText)
}

func new_() Figure {
  xx, yy = int(scr.Wd()), int(scr.Ht())
  f := new (figure)
  f.Clr()
//  f.Type = Points
  f.Type = Segments
  f.colour, _ = col.StartCols()
  return f
}

func (f *figure) imp (Y Any) *figure {
  y, ok := Y.(*figure)
  if ! ok { TypeNotEqPanic (f, Y) }
  return y
}

func (f *figure) Empty() bool {
  return len (f.x) == 0
}

func (f *figure) Clr() {
  f.x, f.y = nil, nil
  f.marked, f.filled = false, false
  f.string = ""
}

func (f *figure) SetType (k Type) {
  f.Clr()
  f.Type = k
}

func (f *figure) Select() {
  f.Clr()
  Acolour := f.colour
  Hcolour := Acolour
  Hcolour.Contrast()
  scr.SetFontsize (font.Normal)
  n := uint(Rectangle)
  y, x := scr.MousePos()
  sel.Select1 (name, NTypes, uint(len(name[0])), &n, y, x, Acolour, Hcolour)
  if n < NTypes {
    f.Type = Type(n)
  }
}

func (f *figure) Eq (Y Any) bool {
  f1 := f.imp (Y)
  n, n1 := uint(len (f.x)), uint(len (f1.x))
  if f.Type != f1.Type || n != n1 || f.filled != f1.filled {
    return false
  }
  if n == 0 { return true } // ?
  if f.x[0] != f1.x[0] || f.y[0] != f1.y[0] {
    return false
  }
  switch f.Type {
  case Text:
    if f.string != f1.string {
      return false
    }
  case Image:
    if f.x[1] != f1.x[1] || f.y[1] != f1.y[1] {
      return false
    } else {
      // Vergleich der Images fehlt
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
  f.Type = f1.Type
  f.colour = f1.colour
  n1 := uint(len (f1.x))
  f.x, f.y = make ([]int, n1), make ([]int, n1)
  for i := uint(0); i < n1; i++ {
    f.x[i] = f1.x[i]
    f.y[i] = f1.y[i]
  }
  f.filled = f1.filled
  f.string = f1.string
  if f.Type == Image {
    // Kopieren des Image fehlt
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

func (f *figure) On (a, b int, t uint) bool {
  if ! f.Empty() {
    switch f.Type {
    case Points, Segments:
      return scr.OnSegments (f.x, f.y, a, b, t)
    case Polygon:
      return scr.OnPolygon (f.x, f.y, a, b, t)
    case Curve:
      return scr.OnCurve (f.x, f.y, a, b, t)
    case InfLine:
      return scr.OnInfLine (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    case Rectangle:
      return scr.OnRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    case Circle:
      return scr.OnCircle (f.x[0], f.y[0], uint(f.x[1]), a, b, t)
    case Ellipse:
      return scr.OnEllipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]), a, b, t)
    case Text:
      if len (f.x) != 2 { errh.Error ("Incident case Text: len (f.x) ==", uint(len(f.x))) }
      return scr.OnRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t) // crash: TODO
    case Image:
      return scr.InRectangle (f.x[0], f.y[0], f.x[1], f.y[1], a, b, t)
    }
  }
  return false
}

func (f *figure) convex() bool {
  n := uint(len (f.x))
  switch f.Type {
  case Rectangle, Circle, Ellipse, Image:
    return true
  case Polygon:
    switch n {
    case 0, 1:
      return false
    case 2:
      return true
    }
  default:
    return false
  }
 // polygon with 3 or more nodes
/*
 // TODO
  dxi := f.x[0] - f.x[n - 1]
  dxk := f.x[1] - f.x[0]
  dyi := f.y[0] - f.y[n - 1]
  dyk := f.y[1] - f.y[0]
  z := uint(0)
  if dxi * dxk + dyi * dyk < 0 { z = 1 }
  a := dxi * dyk
  b := dxk * dyi
  if a == b { // polygon reduced by a node
    return true
    // for n > 3 we are going to roasted in devils oven ...
  }
  gr := a > b
  var k uint
  for i := uint(1); i < n; i++ {
    if i < n { k = i + 1 } else { k = 0 }
    dxi = f.x[i] - f.x[i - 1]
    dyi = f.y[i] - f.y[i - 1]
    dxk = f.x[k] - f.x[i]
    dyk = f.y[k] - f.y[i]
    if dxi * dxk + dyi * dyk < 0 { // Winkel < 90 Grad
      z++
      if z > 3 {  // if more than 3 angles are < 90°, then
        return false // the angle sum is < (n - 1) * 180° !
      }
    }
    a = dxi * dyk
    b = dxk * dyi
    if a != b {
      if (a > b) != gr { return false }
    }
  }
*/
  return true
}

func (f *figure) rectangular() bool {
  switch f.Type {
  case Rectangle, Image:
    return true
  }
  if f.Type != Polygon { return false }
  if len (f.x) != 4 { return false }
  return f.x[1] + f.x[3] == f.x[0] + f.x[2] && f.y[1] + f.y[3] == f.y[0] + f.y[2] &&
         f.x[1] * f.x[1] + f.x[0] * f.x[2] + f.y[1] * f.y[1] + f.y[0] * f.y[2] ==
           f.x[1] * (f.x[0] + f.x[2]) + f.y[1] * (f.y[0] * f.y[2])
}

func (f *figure) UnderMouse (t uint) bool {
  a, b := scr.MousePosGr()
  return f.On (a, b, t)
}

// Locate (a, b) = Relocate (a - x[0], b - y[0])
func (f *figure) Move (a, b int) {
  var n uint
  switch f.Type {
  case Points, Segments, Polygon, Curve, InfLine, Rectangle:
    n = uint(len (f.x))
  case Circle, Ellipse:
    n = 1
  case Text, Image:
    n = 2
  }
  for i := uint(0); i < n; i++ {
    f.x[i] += a
    f.y[i] += b
  }
}

func (f *figure) Marked() bool {
  return f.marked
}

func (f *figure) Mark (m bool) {
  f.marked = m
}

func (f *figure) SetColour (c col.Colour) {
  f.colour = c
  bx.ColourF (f.colour)
  if f.Type == Image {
    // what ?
  }
}

func (f *figure) Colour() col.Colour {
  return f.colour
}

func (f *figure) Erase() {
  switch f.Type {
  case Image:
    scr.ClrGr (f.x[0], f.y[0], f.x[1], f.y[1])
  default:
    c := f.colour
    f.SetColour (scr.ScrColB())
    f.Write()
    f.SetColour (c)
  }
}

func (f *figure) Write() {
  if f.Empty() { return }
  scr.ColourF (f.colour)
  switch f.Type {
  case Points:
    scr.Points (f.x, f.y)
  case Segments:
    scr.Segments (f.x, f.y)
  case Polygon:
    scr.Polygon (f.x, f.y)
    if f.filled {
//      scr.PolygonFull (f.x, f.y) // not yet implemented
    }
  case Curve:
    scr.Curve (f.x, f.y)
    if f.filled {
      n := len (f.x) - 1
      scr.CircleFull (f.x[n], f.y[n], 4) // ?
    }
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
      scr.CircleFull (f.x[0], f.y[0], uint(f.x[1]))
    } else {
      scr.Circle (f.x[0], f.y[0], uint(f.x[1]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  case Text:
    bx.Wd (str.ProperLen (f.string))
    bx.ColourF (f.colour)
    bx.WriteGr (f.string, f.x[0], f.y[0])
  case Image:
//    if moved {
//      scr.RectangleFullInv (...)
//    } else {
//      copy from Imageptr in Framebuffer
//    }
    ppm.Get (f.string, uint(f.x[0]), uint(f.y[0]))
  }
}

func (f *figure) Print (p psp.PostscriptPage) {
  if f.Empty() { return }
  n := uint(len (f.x))
  p.SetColour (f.colour)
  switch f.Type {
  case Points:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.S (f.x[i]), p.Sy (f.y[i])
    }
    p.Points (x, y)
  case Segments:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.S (f.x[i]), p.Sy (f.y[i])
    }
    p.Segments (x, y)
  case Polygon:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.S (f.x[i]), p.Sy (f.y[i])
    }
    p.Polygon (x, y, f.filled)
  case Curve:
    x, y := make ([]float64, n), make ([]float64, n)
    for i := uint(0); i < n; i++ {
      x[i], y[i] = p.S (f.x[i]), p.Sy (f.y[i])
    }
    p.Curve (x, y)
  case InfLine:
    x, y, x1, y1 := p.S (f.x[0]), p.Sy (f.y[0]), p.S (f.x[1]), p.Sy (f.y[1])
    p.Line (x, y, x1, y1)
  case Rectangle:
    x, y, x1, y1 := p.S (f.x[0]), p.Sy (f.y[0]), p.S(f.x[1]), p.Sy (f.y[1])
    p.Rectangle (x, y, x1 - x, y1 - y, f.filled)
  case Circle:
    x, y, r := p.S (f.x[0]), p.Sy (f.y[0]), p.S (f.x[1])
    p.Circle (x, y, r, f.filled)
  case Ellipse:
    x, y, a, b := p.S (f.x[0]), p.Sy (f.y[0]), p.S (f.x[1]), p.S (f.y[1])
    p.Ellipse (x, y, a, b, f.filled)
  case Text:
    x, y := p.S (f.x[0]), p.Sy (f.y[0])
    p.Write (f.string, x, y)
  case Image:
// TODO
  }
}

func (f *figure) Invert() {
  if f.Empty() { return }
  switch f.Type {
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
    if f.filled {
      n := len (f.x) - 1
      scr.CircleInv (f.x[n], f.y[n], 4) // TODO ?
    }
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
      scr.CircleFullInv (f.x[0], f.y[0], uint(f.x[1]))
    } else {
      scr.CircleInv (f.x[0], f.y[0], uint(f.x[1]))
    }
  case Ellipse:
    if f.filled {
      scr.EllipseFullInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  case Text:
// >>>  sollte in bx integriert werden:
//  bx.WriteInvGr (string, x[0], y[0])
    scr.Transparence (true)
    scr.WriteInvGr (f.string, f.x[0], f.y[0])
  case Image:
    scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
  }
}

func (f *figure) invertN() {
  switch f.Type {
  case Points:
    scr.PointsInv (f.x, f.y)
  case Segments:
    scr.SegmentsInv (f.x, f.y)
  case Polygon:
    scr.PolygonInv (f.x, f.y)
  case Curve:
    scr.CurveInv (f.x, f.y)
    if f.filled {
      n := len (f.x) - 1
      scr.CircleInv (f.x[n], f.y[n], 4) // TODO ?
    }
  }
}

func (f *figure) editN() {
  switch f.Type {
  case Points, Segments, Polygon, Curve: default: return }
  x0 := make ([]int, 2); x0[0] = f.x[0]; f.x = x0
  y0 := make ([]int, 2); y0[0] = f.y[0]; f.y = y0
  f.x[1], f.y[1] = scr.MousePosGr()
  f.invertN()
  var ( K kbd.Comm; T uint )
  loop: for {
    K, T = kbd.Command()
    scr.MousePointer (true)
    n := uint(len (f.x))
    switch K {
    case kbd.Esc:
      break loop
    case kbd.Go,
         kbd.Here, kbd.This, kbd.That,
         kbd.There, kbd.Drop, kbd.Move,
         kbd.To: // kbd.ToThis:
      f.invertN()
//      if f.Type == Curve {
//        if n == scr.MaxBezierdegree { break loop }
//      }
      if f.Type == Points {
        if K != kbd.Go {
          n++
        }
      } else {
        if K == kbd.Here { // TODO Curve: missing
          n++
        }
      }
      if K == kbd.To {
        n := len (f.x)
        if n == 0 {
          break loop
        } else { // TODO
          n--
          if n == 0 {
            break loop
//          } else {
//            x0 = make ([]int, n); copy (x0, f.x[:n]); f.x = x0
//            y0 = make ([]int, n); copy (y0, f.y[:n]); f.y = y0
            }
        }
      }
      if n > uint(len (f.x)) {
        x0 = make ([]int, n); copy (x0, f.x); f.x = x0
        y0 = make ([]int, n); copy (y0, f.y); f.y = y0
      }
      f.x[n-1], f.y[n-1] = scr.MousePosGr()
      f.invertN()
      if f.Type == Points {
        if K == kbd.That { break loop }
      } else {
        if K == kbd.Move { break loop }
      }
    }
  }
  if f.x == nil {
    f.Clr()
    return
  }
  scr.ColourF (f.colour)
  switch f.Type {
  case Points:
    scr.Points (f.x, f.y)
  case Segments:
    scr.Segments (f.x, f.y)
  case Polygon:
    scr.Polygon (f.x, f.y)
    f.filled = T > 0 && f.convex()
    if f.filled {
      scr.PolygonFull (f.x, f.y) // not yet implemented
    }
  case Curve:
    scr.Curve (f.x, f.y)
    f.filled = T > 0
    if f.filled {
      n := len (f.x) - 1
      scr.CircleFull (f.x[n], f.y[n], 4)
    }
  }
}

func (f *figure) invert1() {
  switch f.Type {
  case InfLine:
    scr.InfLineInv (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    scr.RectangleInv (f.x[0], f.y[0], f.x[1], f.y[1])
  default:
    scr.EllipseInv (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
  }
}

func (f *figure) edit1() {
  x0 := make ([]int, 2); x0[0] = f.x[0]; f.x = x0
  y0 := make ([]int, 2); y0[0] = f.y[0]; f.y = y0
  switch f.Type {
  case InfLine:
    if f.x[0] == 0 {
      f.x[1] = 1
    } else {
      f.x[1] = f.x[0] - 1
    }
    f.y[1] = f.y[0]
  case Rectangle:
    f.x[1] = f.x[0]
    f.y[1] = f.y[0]
  case Circle, Ellipse:
    f.x[1] = 0
    f.y[1] = 0
  default:
    return
  }
//    scr.PointInv (f.x[0], f.y[0])
  f.invert1()
  loop: for {
    K, T := kbd.Command()
    switch K {
    case kbd.This, kbd.That:
      f.invert1()
      f.x[1], f.y[1] = scr.MousePosGr()
      switch f.Type {
      case InfLine:
        if f.x[1] == f.x[0] && f.y[1] == f.y[0] {
          if f.x[0] == 0 {
            f.x[1] = 1
          } else {
            f.x[1] = f.x[0] - 1
          }
        }
      case Rectangle:
        ;
      case Circle, Ellipse:
        if f.x[1] > f.x[0] {
          f.x[1] -= f.x[0]
        } else {
          f.x[1] = f.x[0] - f.x[1]
        }
        if f.y[1] > f.y[0] {
          f.y[1] -= f.y[0]
        } else {
          f.y[1] = f.y[0] - f.y[1]
        }
        if f.Type == Circle {
          if f.x[1] > f.y[1] {
            f.y[1] = f.x[1]
          } else {
            f.x[1] = f.y[1]
          }
        }
      default:
        // stop (Modul, 1)
      }
      f.invert1()
      if K == kbd.That {
        f.filled = T > 0
        break loop
      }
    }
  }
  switch f.Type {
  case InfLine:
    scr.InfLine (f.x[0], f.y[0], f.x[1], f.y[1])
  case Rectangle:
    if f.filled {
      scr.RectangleFull (f.x[0], f.y[0], f.x[1], f.y[1])
    } else {
      scr.Rectangle (f.x[0], f.y[0], f.x[1], f.y[1])
    }
  case Circle, Ellipse:
    if f.filled {
      scr.EllipseFull (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    } else {
      scr.Ellipse (f.x[0], f.y[0], uint(f.x[1]), uint(f.y[1]))
    }
  }
}

func (f *figure) editText() {
  if f.Type != Text { return }
  scr.MousePointer (false)
  bx.Wd (lenText)
  bx.ColourF (f.colour)
  x1 := f.x[0] + int(lenText * scr.Wd1()) - 1
  if x1 >= xx { x1 = xx - 1 }
  y1 := f.y[0] + int(scr.Ht1()) - 1
  if y1 >= yy { y1 = yy - 1 }
  scr.SaveGr (f.x[0], f.y[0], x1, y1)
  bx.Transparence (false)
  f.string = str.New (lenText) // wörkeraunt
  bx.EditGr (&f.string, f.x[0], f.y[0])
  bx.Transparence (true)
  scr.RestoreGr (f.x[0], f.y[0], x1, y1)
  if C, _ := kbd.LastCommand(); C == kbd.Enter {
    bx.Transparence (true)
//    scr.RestoreGr (f.x[0], f.y[0], x1, y1)
    bx.WriteGr (f.string, f.x[0], f.y[0])
    k := str.ProperLen (f.string)
    x0 := make ([]int, 2); x0[0] = f.x[0]; f.x = x0
    y0 := make ([]int, 2); y0[0] = f.y[0]; f.y = y0
    f.x[1] = f.x[0] + int(scr.Wd1() * k) - 1
    f.y[1] = f.y[0] + int(scr.Ht1()) - 1
    scr.WarpMouseGr (f.x[0], f.y[1])
  } else {
//    f.string = str.New (lenText)
//    bx.WriteGr (f.string, f.x[0], f.y[0])
//    f.string = ""
//    f.x, f.y = nil, nil
  }
  scr.MousePointer (true)
}

func (f *figure) editImage() {
  if f.Type != Image { return }
  scr.MousePointer (false)
  errh.Hint ("Name des Bildes eingeben")
  bx.Wd (32) // reine Willkür
  bx.Colours (f.colour, scr.ScrColB())
  f.string = str.New (uint(len(name[0])))
  bx.EditGr (&f.string, f.x[0], f.y[0])
  str.OffSpc (&f.string)
  W, H := ppm.Size (f.string)
  w, h := int(W), int(H)
  if w <= xx && h <= yy {
    x0 := make ([]int, 2); x0[0] = f.x[0]; f.x = x0
    y0 := make ([]int, 2); y0[0] = f.y[0]; f.y = y0
    f.x[1] = f.x[0] + w - 1
    f.y[1] = f.y[0] + h - 1
    if f.x[1] >= xx {
      f.x[0] = xx - w
      f.x[1] = xx - 1
    }
    if f.y[1] >= yy {
      f.y[0] = yy - h
      f.y[1] = yy - 1
    }
    errh.DelHint()
//  besser:
//    ppm.Get ...
//    NEW (Imagespeicher)
//    ppm.Get ( ... dort rein ...)
//    ppm.Get (string, x[0], y[0])
  } else {
    errh.DelHint()
  }
  scr.MousePointer (true)
}

func (f *figure) uM() uint {
  const ( r = 4; t = 4 )
  a, b := scr.MousePosGr()
  n := uint(len (f.x))
  for i := uint(0); i < n; i++ {
    if scr.OnCircle (f.x[i], f.y[i], r, a, b, t) {
      return uint(i)
    }
  }
  return n + 1 // ?
}

func (f *figure) mark (i uint) {
//  if f.Type != Curve { return }
  for r := uint(3); r <= 4; r++ {
    scr.CircleInv (f.x[i], f.y[i], r)
  }
}

func (f *figure) Edit() {
  if f.Empty() {
    scr.ColourF (f.colour)
    f.x, f.y = make ([]int, 1), make ([]int, 1)
    f.x[0], f.y[0] = scr.MousePosGr()
    switch f.Type {
    case Points, Segments, Polygon, Curve:
      f.editN()
    case InfLine, Rectangle, Circle, Ellipse:
      f.edit1()
    case Text:
      f.editText()
    case Image:
//      ALLOCATE (Imageptr, Groesse())
//      ppm.Get (string [...], Imageptr)
      f.editImage()
    }
    if f.x == nil {
      f.Clr()
    }
  } else {
    n := uint(len (f.x))
errh.Error ("Figur hat Länge", n)
    switch f.Type {
    case Text:
      f.editText()
    case Image:
      f.editImage()
    default:
      f.Erase()
      f.Invert()
      if true { // f.Type == Curve {
        for i := uint(0); i < n; i++ { f.mark (i) }
      }
      i := f.uM()
      f.x[i], f.y[i] = scr.MousePosGr()
      loop: for {
        scr.MousePointer (true)
        c, _ := kbd.Command()
        switch c {
        case kbd.Esc:
          break loop
        case kbd.Enter, kbd.Tab, kbd.Search:
          f.colour = sel.Colour()
        case kbd.Here:
          break loop
        case kbd.There:
          i = f.uM()
        case kbd.Drop, kbd.Move:
          if i < n {
            f.Invert()
            f.mark (i)
            f.x[i], f.y[i] = scr.MousePosGr()
            f.mark (i)
            f.Invert()
            if c == kbd.Move { i = n } // ? ? ?
          }
        case kbd.To:
          switch f.Type {
          case Points, Segments, Polygon, Curve:
            if f.x == nil {
              f.Clr()
            } else {
              for i := uint(0); i < n; i++ { f.mark (i) }
              f.Erase()
              n-- // ? ? ?
              f.Invert()
              for i := uint(0); i < n; i++ { f.mark (i) }
            }
          }
        }
        errh.Hint (c.String())
      }
      f.Invert()
      if true { // kind != Text {
        for i := uint(0); i < n; i++ { f.mark (i) }
      }
      f.Write()
    }
  }
}

func (f *figure) Codelen() uint {
  n := 1 + uint32(f.colour.Codelen()) + 4
  switch f.Type {
  case Text:
    n += 2 * 4 + 1 + uint32(len (f.string)) // 4 = Codelen (uint32(0))
  case Image:
    n += 4 * 4 + 1 + uint32(len (f.string))
  default:
    n += 2 * uint32(len (f.x)) * 4
  }
  n += 2 * 4 // Reserve
  return uint(n)
}

func (f *figure) Encode() Stream {
  bs := make (Stream, f.Codelen())
  a := uint32(0)
  bs[a] = byte(f.Type)
  a++
  copy (bs[a:a+3], f.colour.Encode())
  a += 3
  var n uint32
  if f.Type < Text {
    n = uint32(len (f.x))
  } else {
    n = uint32(len (f.string))
  }
  copy (bs[a:a+4], Encode (n))
  a += 4
  if f.Type < Text {
    for i := uint32(0); i < n; i++ {
      copy (bs[a:a+4], Encode (int32(f.x[i])))
      a += 4
      copy (bs[a:a+4], Encode (int32(f.y[i])))
      a += 4
    }
  } else { // Text, Image
    copy (bs[a:a+4], Encode (int32(f.x[0])))
    a += 4
    copy (bs[a:a+4], Encode (int32(f.y[0])))
    a += 4
    if f.Type == Image {
      copy (bs[a:a+4], Encode (int32(f.x[1])))
      a += 4
      copy (bs[a:a+4], Encode (int32(f.y[1])))
      a += 4
    }
    copy (bs[a:a+n], Stream(f.string))
    a += n
  }
  bs[a] = 0
  if f.filled { bs[a] ++ }
  if f.marked { bs[a] += 2 }
  return bs
}

func (f *figure) Decode (bs Stream) {
  a := uint32(0)
  f.Type = Type(bs[a])
  a ++
  f.colour.Decode (bs[a:a+3])
  a += 3
  n := uint32(0)
  n = uint32(Decode (uint32(0), bs[a:a+4]).(uint32))
  a += 4
  if f.Type < Text {
    f.x, f.y = make ([]int, n), make ([]int, n)
    for i := uint32(0); i < n; i++ {
      f.x[i] = int(Decode (int32(f.x[i]), bs[a:a+4]).(int32))
      a += 4
      f.y[i] = int(Decode (int32(f.y[i]), bs[a:a+4]).(int32))
      a += 4
    }
  } else { // kind == Text, Image
    f.x, f.y = make ([]int, 2), make ([]int, 2)
    f.x[0] = int(Decode (int32(f.x[0]), bs[a:a+4]).(int32))
    a += 4
    f.y[0] = int(Decode (int32(f.y[0]), bs[a:a+4]).(int32))
    a += 4
    if f.Type == Image {
      f.x[1] = int(Decode (int32(f.x[1]), bs[a:a+4]).(int32))
      a += 4
      f.y[1] = int(Decode (int32(f.y[1]), bs[a:a+4]).(int32))
      a += 4
    }
    f.string = string(bs[a:a+n])
    a += n
    if f.Type == Text {
      f.x[1] = f.x[0] + int(scr.Wd1()) * int(n) - 1
      f.y[1] = f.y[0] + int(scr.Ht1()) - 1
    }
  }
  f.filled = bs[a] % 2 == 1
  f.marked = (bs[a] / 2) % 2 == 1
}
