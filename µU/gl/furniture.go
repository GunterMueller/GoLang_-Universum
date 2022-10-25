package gl

// (c) Christian Maurer   v. 221023 - license see µU.go

import (
  "math"
  "µU/col"
)
const pi180 = math.Pi / 180

func fsin (x float64) float64 { return math.Sin (x * pi180) }
func fcos (x float64) float64 { return math.Cos (x * pi180) }

func fcuboid (x, y, z, w, d, h, a float64, f col.Colour) {
  Colour (f)
  Cuboid1 (x, y, z, w, d, h, a)
}

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

func roundTable (x, y, z, r, rf, rb, h, hf, hp float64, c col.Colour) {
  Colour (c)
  if hf > 0 { Cylinder (x, y, z, rf, hf) } // foot
  Cylinder (x, y, z + hf, rb, h - hf - hp) // leg
  Cylinder (x, y, z + h - hp, r, hp) // olate 
}

func ovalTable (x, y, z, w, d, h, a float64, f col.Colour) {
  if d > w { panic ("oops") }
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

func chair (x, y, z, w, h, p, a float64, c col.Colour) {
  bench (x, y, z, w, w, h/2, p, h/20, h/20, h/2, a, c)
}

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
