package pos

// (c) Christian Maurer   v. 161216 - license see murus.go

import (
  "math"
  . "murus/obj"
  "murus/col"
  "murus/scr"
//  "murus/errh"
//  "murus/str"
//  "murus/integ"
)
const
  inf = math.MaxInt32
type
  position struct {
             x, y int
             w, h uint8
             f, b col.Colour
                  }

func new_(w, h uint) Position {
  x := new(position)
  x.x, x.y = inf, inf
  x.w, x.h = uint8 (w % (1<<8)), uint8(h % (1<<8))
  x.f, x.b = col.StartCols()
  return x
}

func (x *position) imp (Y Any) *position {
  y, ok := Y.(*position)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (p *position) Set (x, y int) {
  if x > inf { x = inf }
  if y > inf { y = inf }
  p.x, p.y = x, y
}

func (x *position) Pos() (int, int) {
  return x.x, x.y
}

func (x *position) Contour() (uint, uint) {
  return uint(x.w), uint(x.h)
}

func (x *position) Empty() bool {
  return x.x == inf
}

func (x *position) Clr() {
  x.x, x.y = inf, inf
}

func (x *position) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.x == y.x &&
         x.y == y.y
}

func (x *position) Copy (Y Any) {
  y := x.imp (Y)
  x.x, x.y = y.x, y.y
  x.w, x.h = y.w, y.h
  x.f, x.b = y.f, y.b
}

func (x *position) Less (Y Any) bool {
  y := x.imp (Y)
  if x.y < y.y {
    return true
  }
	if x.y == y.y {
    return x.x < y.x
  }
  return false
}

func (x *position) Clone() Any {
  y := new_(uint(x.w), uint(x.h))
  y.Copy (x)
  return y
}

func (x *position) Codelen() uint {
  return 3 * 2
}

func (x *position) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  i, a := uint(0), uint(2)
  copy (bs[i:i+a], Encode (int16(x.x)))
  i += a
  copy (bs[i:i+a], Encode (int16(x.y)))
  i += a
  copy (bs[i:i+a], Encode (uint16(x.w) + 1<<8 * uint16(x.h)))
  return bs
}

func (x *position) Decode (bs []byte) {
  i, a := uint(0), uint(2)
  x.x = int(Decode (int16(0), bs[i:i+a]).(int16))
  i += a
  x.y = int(Decode (int16(0), bs[i:i+a]).(int16))
  i += a
  wh  := Decode (uint16(0), bs[i:i+a]).(uint16)
  x.w, x.h = uint8(wh % (1<<8)), uint8(wh / (1<<8))
}

func (x *position) Mouse() {
  xm, ym := scr.MousePosGr()
  x.x, x.y = xm, ym
  wd, ht := int(scr.Wd()), int(scr.Ht())
  w1, h1 := int(scr.Wd1()), int(scr.Ht1())
  w, h := int(x.w) * w1, int(x.h) * h1
  if x.x < w { x.x = w }
  if x.x + w >= wd { x.x = wd - 1 - w }
  if x.y < h { x.y = h }
  if x.y + h >= ht { x.y = ht - 1 - h }
}

func (x *position) UnderMouse () bool {
  return scr.UnderMouseGr (x.x, x.y, x.x, x.y, uint(x.h) * scr.Ht1())
}

func (x *position) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *position) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

/* func (x *position) Defined (s string) bool {
  n, ok  := str.Pos (s, byte(','))
  if ! ok { return false }
  x.x, ok = integ.Integer(s[:n])
  if ! ok { return false }
  x.y, ok = integ.Integer(s[n+1:])
  if ! ok { return false }
  return true
}

func (x *position) String() string {
  return integ.String (x.x) + ", " + integ.String (x.y)
} */

func (x *position) Write() {
  if x.w * x.h == 0 {
    return
  }
  const R0 = 3
  w1, h1  := scr.Wd1(), scr.Ht1()
  scr.Colours (x.f, x.b)
  if uint(x.w) * w1 <= R0 {
    scr.EllipseFull (x.x, x.y, R0, uint(x.h) * h1 / 2)
  } else {
    scr.ColourF (x.b)
    scr.EllipseFull (x.x, x.y, uint(x.w) * w1 * 7 / 10, uint(x.h) * h1 * 8 / 10)
    scr.ColourF (x.f)
    scr.Ellipse (x.x, x.y, uint(x.w) * w1 * 7 / 10, uint(x.h) * h1 * 8 / 10)
  }
}

func (x *position) Edit() {
  x.Write()
// TODO
}

var
  xx, yy int

func (x *position) Save() {
  xx, yy = x.x, x.y
}

func (x *position) Restore() {
  x.x, x.y = xx, yy
}
