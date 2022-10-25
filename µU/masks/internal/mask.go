package mask

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
)
type
  mask struct {
              string
      l, c, w uint
              }
var (
  bx = box.New()
  colF, colB = col.LightWhite(), col.Black()
)

func new_() Mask {
  x := new(mask)
  x.string = str.New (M)
  x.w = 0
  return x
}

func (x *mask) imp (Y any) *mask {
  y, ok := Y.(*mask)
  if ! ok { TypeNotEqPanic (y, Y) }
  return y
}

func (x *mask) Empty() bool {
  return str.Empty (x.string)
}

func (x *mask) Clr() {
  x.string = str.New (M)
  x.l, x.c, x.w = 0, 0, 0
}

func (x *mask) Eq (Y any) bool {
  y := x.imp (Y)
  return x.string == y.string &&
         x.l == y.l &&
         x.c == y.c &&
         x.w == y.w
}

func (x *mask) Copy (Y any) {
  y := x.imp (Y)
  x.string = y.string
  x.l, x.c = y.l, y.c
  x.w = y.w
}

func (x *mask) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *mask) Less (Y any) bool {
  y := x.imp (Y)
  if x.l == x.l {
    return x.c + x.w <= y.c
  }
  return x.l < y.l
}

func (x *mask) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *mask) Codelen() uint {
  return M + 3
}

func (x *mask) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), M
  t := x.string
  str.Norm (&t, M)
  copy (s[i:i+a], Stream(t))
  i += a
  a = 1
  copy (s[i:i+a], Encode(uint8(x.l)))
  i++
  copy (s[i:i+a], Encode(uint8(x.c)))
  i++
  copy (s[i:i+a], Encode(uint8(x.w)))
  return s
}

func (x *mask) Decode (s Stream) {
  i, a := uint(0), M
  x.string = string(s[i:i+a])
  i += a
  a = 1
  x.l = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  x.c = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  x.w = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  str.Norm (&x.string, x.w)
}

func (x *mask) Place (l, c uint) {
  x.l, x.c = l, c
}

func (x *mask) Pos() (uint, uint) {
  return x.l, x.c
}

func (x *mask) Wd() uint {
  return x.w
}

func (x *mask) Del() {
  n := str.ProperLen (x.string)
  bx.Wd (n)
  bx.Colours (col.Black(), col.Black())
  bx.Write (x.string, x.l, x.c)
  bx.Colours (colF, colB)
}

func (x *mask) Write() {
  n := str.ProperLen (x.string)
  if n == 0 { return }
  bx.Wd (n)
  bx.Write (x.string, x.l, x.c)
}

func (x *mask) Edit() {
  n := M
  if M + x.w >= scr.NColumns() {
    n = scr.NColumns() - M - 1
  }
  bx.Wd (n)
  bx.Edit (&x.string, x.l, x.c)
  str.OffSpc1 (&x.string)
  x.w = uint(len(x.string))
  bx.Colours (colF, colB)
  bx.Write (x.string, x.l, x.c)
}

func (x *mask) Print() {
  n := str.ProperLen (x.string)
  if n == 0 { return }
//  prt.Print (x.string, xl, x.c, font.Normal)
}

func init() {
  bx.Colours (colF, colB)
}
