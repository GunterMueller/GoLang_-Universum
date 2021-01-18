package gl

// (c) Christian Maurer   v. 191027 - license see µU.go

import (
  "math"
  "µU/col"
)
const
  p = math.Pi / 180
var (
  cWall col.Colour = col.LightWhite()
  xx, yy, height0, alpha float64
  height float64 = 2.75
)

func sin_(a float64) float64 { return math.Sin(p * a) }
func cos_(a float64) float64 { return math.Cos(p * a) }

func wal (x, y *float64, x1, y1, h float64) {
  Colour (cWall)
  VertRectangle (*x, *y, height0, x1, y1, h)
  *x, *y = x1, y1
}

func wall (w float64) {
  wal (&xx, &yy, xx + cos_(alpha) * w, yy + sin_(alpha) * w, height)
}

func dor (x, y *float64, w, r, d, j, h, a float64, f col.Colour) {
  s, c := sin_(a), cos_(a)
  w1, d1 := w + 2 * r, d + j
  z1 := height0 + h
  z2 := z1 + r
  x1, y1 := *x + w1 * c, *y + w1 * s
  Colour (f)
  VertRectangle (*x, *y, z2, x1, y1, height)
  Colour (f)
  x2, y2 := *x + j * s, *y - j * c
  x3, y3 := x1 + j * s, y1 - j * c
  x4, y4 := x2 + r * c, y2 + r * s
  x5, y5 := x3 - r * c, y3 - r * s
  VertRectangle (x2, y2, height0, x4, y4, z1)
  VertRectangle (x5, y5, height0, x3, y3, z1)
  VertRectangle (x2, y2, z1, x3, y3, z2)
  if j > 0. {
    VertRectangle (*x, *y,  height0, x2, y2, z2)
    VertRectangle (x3, y3, height0, x1, y1, z2)
    Quad (*x, *y, z2, x2, y2, z2, x3, y3, z2, x1, y1, z2)
  }
  if d1 > 0. {
    x6, y6 := x4 - d1 * s, y4 + d1 * c
    x7, y7 := x5 - d1 * s, y5 + d1 * c
    VertRectangle (x4, y4, height0, x6, y6, z1)
    VertRectangle (x7, y7, height0, x5, y5, z1)
    Quad (x4, y4, z1, x6, y6, z1, x7, y7, z1, x5, y5, z1)
  }
  *x, *y = x1, y1
}

func door (w, m, d, j, h float64, c col.Colour) {
  dor (&xx, &yy, w, m, d, j, h, alpha, c)
}

func win (x, y *float64, w, r, d, fb, h, rb, rt, a float64, b bool, f col.Colour) {
  s, c := sin_(a), cos_(a)
  x1, y1 := *x + w * c, *y + w * s
  x2, y2 := *x - d * s, *y + d * c
  x3, y3 := x1 - d * s, y1 + d * c
  x4, y4 := x2 + r * c, y2 + r * s
  x5, y5 := x2 + (w - r) * c, y2 + (w - r) * s
  zf := height0 + fb
  zu, zo, zh := zf + rb, zf + h - rt, zf + h
  Colour (cWall)
  if b {
    VertRectangle (x2, y2, height0, x3, y3, zf - 0.03)
  } else {
    VertRectangle (*x, *y, height0, x1, y1, zf - 0.03)
  } // Wand unter Fenster
  Colour (f)
  Quad (*x, *y, zf, x1, y1, zf - 0.03, x3, y3, zf, x2, y2, zf) // Fensterbrett
  VertRectangle (*x, *y, height0, x2, y2, zh) // linke Seitenwand
  VertRectangle (x3, y3, height0, x1, y1, zh) // rechte Seitenwand
  Quad (*x, *y, zh, x2, y2, zh, x3, y3, zh, x1, y1, zh) // Seitendecke
  if rb > 0.0 {
    VertRectangle (x2, y2, zf, x3, y3, zu)
  }
  if rt > 0.0 {
    VertRectangle (x2, y2, zo, x3, y3, zh)
  }
  if r > 0.0 {
    VertRectangle (x2, y2, zu, x4, y4, zo)
    VertRectangle (x5, y5, zu, x3, y3, zo)
  }
  Colour (cWall)
  VertRectangle (*x, *y, zh, x1, y1, height) // Wand über Fenster
  *x, *y = x1, y1
}

func window (w, d, h, fb, r, rb, rt float64, b bool, c col.Colour) {
  win (&xx, &yy, w, r, d, fb, h, rb, rt, alpha, b, c)
}

func win1 (x, y *float64, w, d, h, fb, a float64, f col.Colour) {
  s, c := sin_(a), cos_(a)
  x1, y1 := *x + w * c, *y + w * s
  x2, y2 := *x - d * s, *y + d * c
  x3, y3 := x1 - d * s, y1 + d * c
  zf := height0 + fb
  zh := zf + h
  Colour (cWall)
  VertRectangle (*x, *y, height0, x1, y1, zf) // Wand unter Fenster
  Colour (f)
  Quad (*x, *y, zf, x1, y1, zf, x3, y3, zf, x2, y2, zf) // Fensterbrett
  VertRectangle (*x, *y, zf, x2, y2, zh) // linke Seitenwand
  VertRectangle (x3, y3, zf, x1, y1, zh) // rechte Seitenwand
  Quad (*x, *y, zh, x2, y2, zh, x3, y3, zh, x1, y1, zh) // Seitendecke
  Colour (cWall)
  VertRectangle (*x, *y, zh, x1, y1, height) // Wand über Fenster
  *x, *y = x1, y1
}

func window1 (w, d, h, fb float64, c col.Colour) {
  win1 (&xx, &yy, w, d, h, fb, alpha, c)
}
