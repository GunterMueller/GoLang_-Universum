package masks

// (c) Christian Maurer   v. 210413 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/font"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/pbox"
)
type
  maskSet struct {
            mask []string
                 uint "number of masks"
            l, c []uint
            f, b col.Colour
                 font.Font
                 }
var
  pbx = pbox.New()

func new_() MaskSet {
  x := new (maskSet)
  x.uint = 0
  x.mask = make([]string, 0)
  x.f, x.b = col.LightWhite(), col.Black()
  x.l, x.c = make([]uint, 0), make([]uint, 0)
  x.Font = font.Roman
  return x
}

func (x *maskSet) imp (Y Any) *maskSet {
  y, ok := Y.(*maskSet)
  if ! ok { TypeNotEqPanic (x, Y) }
  if y == nil { ker.Oops() }
//  if len (x.mask) != len (y.mask) { TypeNotEqPanic (x, Y) }
  return y
}

func (x *maskSet) Num() uint {
  return x.uint
}

func (x *maskSet) Empty() bool {
  return x.uint == 0
}

func (x *maskSet) Clr() {
  x.mask = make([]string, 0)
  x.l, x.c = make([]uint, 0), make([]uint, 0)
  x.uint = 0
}

func (x *maskSet) Eq (Y Any) bool {
  y := x.imp (Y)
  for i := uint(0); i < x.uint; i++ {
    if x.mask[i] != y.mask[i] || x.l[i] != y.l[i] || x.c[i] != y.c[i] {
      return false
    }
  }
  return false
}

func (x *maskSet) Less (Y Any) bool {
  return false
}

func (x *maskSet) Copy (Y Any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.mask = make ([]string, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.mask[i] = y.mask[i]
    x.l[i], x.c[i] = y.l[i], y.c[i]
  }
  x.f, x.b = y.f, y.b
}

func (x *maskSet) Clone() Any {
  y := new (maskSet)
  y.Copy (x)
  return y
}

func (x *maskSet) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *maskSet) Write (l, c uint) {
  if x.uint == 0 { return }
  scr.Colours (x.f, x.b)
  for i := uint(0); i < x.uint; i++ {
    scr.Write (x.mask[i], l + x.l[i], c + x.c[i])
  }
}

func (x *maskSet) Edit (l, c uint) {
// TODO
}

func (x *maskSet) SetFont (f font.Font) {
  x.Font = f
}

func (x *maskSet) Print (l, c uint) {
  if x.uint == 0 { return }
  pbx.SetFont (x.Font)
  for i := uint(0); i < x.uint; i++ {
    pbx.Print (x.mask[i], x.l[i], x.c[i])
  }
}

func (x *maskSet) Codelen() uint {
  c := uint(2)
  for k := uint(0); k < x.uint; k++ {
    c += 1 + Codelen(x.mask[k]) + 2
  }
  c += 2 * x.f.Codelen()
  return c
}

func (x *maskSet) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint(0), uint(2)
  copy (s[i:i+a], Encode (uint16(x.uint)))
  i += a
  for k := uint(0); k < x.uint; k++ {
    a = Codelen(x.mask[k])
    copy (s[i:i+1], Encode (uint8(a)))
    i++
    copy (s[i:i+a], Stream(x.mask[k]))
    i += a
    copy (s[i:i+1], Encode (uint8(x.l[k])))
    i++
    copy (s[i:i+1], Encode (uint8(x.c[k])))
    i++
  }
  a = x.f.Codelen()
  copy (s[i:i+a], x.f.Encode())
  i += a
  copy (s[i:i+a], x.b.Encode())
  return s
}

func (x *maskSet) Decode (s Stream) {
  i, a := uint(0), uint(2)
  x.uint = uint(Decode (uint16(0), s[i:i+a]).(uint16))
  i += a
  for k := uint(0); k < x.uint; k++ {
    a = uint(Decode (uint8(0), s[i:i+1]).(uint8))
    i++
    x.mask[k] = Decode (str.New (a), s[i:i+a]).(string)
    i += a
    x.l[k] = uint(Decode(uint8(0), s[i:i+1]).(uint8))
    i++
    x.c[k] = uint(Decode(uint8(0), s[i:i+1]).(uint8))
    i++
  }
  a = x.f.Codelen()
  x.f.Decode (s[i:i+a])
  i += a
  x.b.Decode (s[i:i+a])
}

func (x *maskSet) Ins (s string, l, c uint) {
  x.mask = append (x.mask, s)
  x.l, x.c = append (x.l, l), append (x.c, c)
  x.uint++
}

func (x *maskSet) Del (n uint) {
  if n >= x.uint { return }
  for i := uint(0); i + 1 < x.uint; i++ {
    x.mask[i] = x.mask[i + 1]
    x.l[i], x.c[i] = x.l[i + 1], x.c[i + 1]
  }
  x.uint--
}

func (x *maskSet) Size (n uint) (uint, uint) {
  if n >= x.uint { return 0, 0 }
  return uint(len(x.mask[n])), 1
}

func (x *maskSet) Ex (l, c uint) (uint, bool) {
  for i := uint(0); i < x.uint; i++ {
    if x.l[i] == l && x.c[i] == c {
      return i, true
    }
  }
  return x.uint, false
}
