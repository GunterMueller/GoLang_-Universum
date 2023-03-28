package flat

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  "µU/col"
)

// walls and windows ///////////////////////////////////////////////////

func SetHt (f, c float64) { setHt(f,c) }

func SetAng (a float64) { alpha = a }

func SetXY (x, y float64) { xx, yy = x, y }

func XY() (float64, float64) { return xx, yy }

func Move (x, y float64) { xx += x; yy += y }

func SetColW (c col.Colour) { cWall = c }

func Wall (w float64) { wall (w) }

// w, d, h = width, depth, height, f = door frame, p = door protrusion
func Door (w, f, d, p, h float64, c col.Colour) {
  door (w, f, d, p, h, c)
}

// f, fb, ft = window frame left/right, bottom, top; wc = height of window sill
func Window (w, d, h, fb, r, rb, rt float64, b bool, c col.Colour) {
  window (w, d, h, fb, r, rb, rt, b, c)
}

func Window1 (w, d, h, f float64, c col.Colour) {
  window1 (w, d, h, f, c)
}

// furniture ///////////////////////////////////////////////////////////


// (x, y, z) = left front corner, (w, d, h) = width, depth and height; a = angle

// p, l = thickness of table plate and legs
func Table (x, y, z, w, d, h, p, l, a float64, f col.Colour) {
  table (x, y, z, w, d, h, p, l, a, f)
}

// r, rf, rl = radius of table plate, foot and leg, hf, hp = height (thickness) of foot and plate
func RoundTable (x, y, z, r, rf, rb, h, hf, hp float64, c col.Colour) {
  roundTable (x, y, z, r, rf, rb, h, hf, hp, c)
}

// Pre: d <= w.
// d = length of straight part
func OvalTable (x, y, z, w, d, h, a float64, f col.Colour) {
  ovalTable (x, y, z, w, d, h, a, f)
}

// w, h = width and height of back and arm rests, p = thickness of seat plate
func Chair (x, y, z, w, h, p, a float64, c col.Colour) {
  chair (x, y, z, w, h, p, a, c)
}

// w, d, h = width, depth and height of back and arm rests,
// wb = TODO, wd = TODO,  hs = seat height, hb = TODO
func ArmChair (x, y, z, w, d, h, wb, wd, hs, hb, a float64, f col.Colour) {
  armChair (x, y, z, w, d, h, wb, wd, hs, hb, a, f)
}

// h = seat height, p, l = thickness of seat plate, legs
// db, hb = depth, height of back rest
func Bench (x, y, z, w, d, h, p, l, db, hb, a float64, f col.Colour) {
  bench (x, y, z, w, d, h, p, l, db, hb, a, f)
}
