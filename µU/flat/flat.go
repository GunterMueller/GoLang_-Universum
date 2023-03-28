package flat

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  "math"
  "µU/ker"
  "µU/col"
  . "µU/gl"
)
const
  p = math.Pi / 180
var (
  cWall col.Colour = col.FlashWhite()
  xx, yy, height0, alpha float64
  height float64 = 2.75
)

func sin_(a float64) float64 { return math.Sin(p * a) }
func cos_(a float64) float64 { return math.Cos(p * a) }

// walls and windows ///////////////////////////////////////////////////

func setHt (f, c float64) { height0, height = f, f + c }

func setAng (a float64) { alpha = a }

func setPos (x, y float64) { xx, yy = x, y }

func pos() (float64, float64) { return xx, yy }

func move (x, y float64) { xx += x; yy += y }

func setColW (c col.Colour) { cWall = c }

func wal (x, y *float64, x1, y1, h float64) {
  Colour (cWall)
// if *y == y1 { y1 += 0.1 }
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
// if x2 == x4 { x4 = x4 + 0.1}
  VertRectangle (x2, y2, height0, x4, y4, z1)
// if x5 == x3 { x5 = x5 + 0.1 }
  VertRectangle (x5, y5, height0, x3, y3, z1)
  VertRectangle (x2, y2, z1, x3, y3, z2)
  if j > 0. {
// if *y == y2 { y2 += 0.1 }
    VertRectangle (*x, *y,  height0, x2, y2, z2)
// if y3 == y1 { y3 += 0.1 }
    VertRectangle (x3, y3, height0, x1, y1, z2)
    Quad (*x, *y, z2, x2, y2, z2, x3, y3, z2, x1, y1, z2)
  }
  if d1 > 0. {
    x6, y6 := x4 - d1 * s, y4 + d1 * c
    x7, y7 := x5 - d1 * s, y5 + d1 * c
// if y4 == y6 { y4 += 0.1 }
    VertRectangle (x4, y4, height0, x6, y6, z1)
// if y7 == y5 { y7 += 0.1 }
    VertRectangle (x7, y7, height0, x5, y5, z1)
    Quad (x4, y4, z1, x6, y6, z1, x7, y7, z1, x5, y5, z1)
  }
  *x, *y = x1, y1
}

// w, d, h = width, depth, height, f = door frame, p = door protrusion
func door (w, f, d, p, h float64, c col.Colour) {
  dor (&xx, &yy, w, f, d, p, h, alpha, c)
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
  } // wall below window
  Colour (f)
// window recess:
  Quad (*x, *y, zf, x1, y1, zf - 0.03, x3, y3, zf, x2, y2, zf) // sill
  VertRectangle (*x, *y, height0, x2, y2, zh) // left wall
  VertRectangle (x3, y3, height0, x1, y1, zh) // rigth wall
  Quad (*x, *y, zh, x2, y2, zh, x3, y3, zh, x1, y1, zh) // ceiling
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
  VertRectangle (*x, *y, zh, x1, y1, height) // wall above window
  *x, *y = x1, y1
}

// f, fb, ft = window frame left/right, bottom, top; wc = height of window sill
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

func window1 (w, d, h, f float64, c col.Colour) {
  win1 (&xx, &yy, w, d, h, f, alpha, c)
}

// furniture ///////////////////////////////////////////////////////////

const pi180 = math.Pi / 180

func fsin (x float64) float64 { return math.Sin (x * pi180) }
func fcos (x float64) float64 { return math.Cos (x * pi180) }

func fcuboid (x, y, z, w, d, h, a float64, f col.Colour) {
  Colour (f)
  Cuboid1 (x, y, z, w, d, h, a)
}

// (x, y, z) = left front corner, (w, d, h) = width, depth and height; a = angle

// p, l = thickness of table plate and legs
func table (x, y, z, w, d, h, p, l, a float64, f col.Colour) {
  Colour (f)
  s, c := fsin(a), fcos(a)
  hl := h - p // height of the legs
  Cuboid1 (x, y, z + hl, w, d, p, a) // table plate
  Cuboid1 (x + l * c - l * s,
           y + l * c + l * s,                     z, l, l, hl, a) // left front leg
  Cuboid1 (x + (w - 2 * l) * c - l * s,
           y + l * c + (w - 2 * l) * s,           z, l, l, hl, a) // right front
  Cuboid1 (x + l * c - (d - 2 * l) * s,
           y + l * s + (d - 2 * l) * c,           z, l, l, hl, a) // left back
  Cuboid1 (x + (w - 2 * l) * c - (d - 2 * l) * s,
           y + (d - 2 * l) * c + (w - 2 * l) * s, z, l, l, hl, a) // right back
}

// r, rf, rl = radius of table plate, foot and leg, hf, hp = height (thickness) of foot and plate
func roundTable (x, y, z, r, rf, rb, h, hf, hp float64, c col.Colour) {
  Colour (c)
  if hf > 0 { Cylinder (x, y, z, rf, hf) } // foot
  Cylinder (x, y, z + hf, rb, h - hf - hp) // leg
  Cylinder (x, y, z + h - hp, r, hp) // olate 
}

// Pre: d <= w.
// d = length of straight part
func ovalTable (x, y, z, w, d, h, a float64, f col.Colour) {
  if d > w { ker.Oops() }
  s, c := fsin(a), fcos(a)
  r := d / 2
  x1, y1 := x + r * c, y + r * s
  x2, y2 := x + (w - r) * c, y + (w - r) * s
  x3, y3 := x1 - d * s, y1 + d * c
  x4, y4 := x2 - d * s, y2 + d * c
  Colour (f)
  Cylinder ((x1 + x3) / 2, (y1 + y3) / 2, z, r, h)
  Cylinder ((x2 + x4) / 2, (y2 + y4) / 2, z, r, h)
  Quad (x1, y1, z + h, x2, y2, z + h, x4, y4, z + h, x3, y3, z + h)
  VertRectangle (x1, y1, z, x2, y2, z + h)
  VertRectangle (x4, y4, z, x3, y3, z + h)
}

// w, h = width and height of back and arm rests, p = thickness of seat plate
func chair (x, y, z, w, h, p, a float64, c col.Colour) {
  Bench (x, y, z, w, w, h/2, p, h/20, h/20, h/2, a, c)
}

// w, d, h = width, depth and height of back and arm rests,
// wb = TODO, wd = TODO,  hs = seat height, hb = TODO
func armChair (x, y, z, w, d, h, wb, wd, hs, hb, a float64, f col.Colour) {
  Colour (f)
  s, c := fsin(a), fcos(a)
  Cuboid1 (x, y, z, w, d, hs, a) // seat
  Cuboid1 (x - (d - wd) * s, y + (d - wd) * c,
           z + hs, w, wd, h - hs, a) // back rest
  Cuboid1 (x, y, z + hs, wb, d - wd, hb - hs, a) // left arm rest
  Cuboid1 (x + (w - wb) * c, y + (w - wb) * s,
           z + hs, wb, d - wd, hb - hs, a) // right arm rest
}

// h = seat height, p, l = thickness of seat plate, legs
// db, hb = depth, height of back rest
func bench (x, y, z, w, d, h, p, l, db, hb, a float64, f col.Colour) {
  Colour (f)
  s, c := fsin(a), fcos(a)
  hl := h - p // height of legs
  Cuboid1 (x, y, z + hl, w, d, p, a) // seat pad
  Cuboid1 (x,               y,               z, l, l, hl, a) // left front leg
  Cuboid1 (x + (w - l) * c, y + (w - l) * s, z, l, l, hl, a) // right front
  Cuboid1 (x - (d - l) * s, y + (d - l) * c, z, l, l, hl, a) // left back
  Cuboid1 (x + (w - l) * c - (d - l) * s,
                            y + (d - l) * c + (w - l) * s,
                                             z, l, l, hl, a) // right back
  if hb > 0 { // with back rest
    Cuboid1 (x - (d - p) * s, y + (d - p) * c, z + h, w, db, hb, a)
  }
}
