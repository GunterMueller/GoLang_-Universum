package stru

// (c) Christian Maurer   v. 230326 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
)
type
  structure struct {
                   int // typ - see µU/atom
           l, c, w uint
              f, b col.Colour 
                   bool // isIndex
                   }

func new_() Structure {
  x := new(structure)
  x.int = 0 // typ atom.String
  x.f, x.b = col.FlashWhite(), col.Black()
  return x
}

func (x *structure) imp (Y any) *structure {
  y := Y.(*structure)
  return y
}

func (x *structure) Typ() int {
  return x.int
}

func (x *structure) Empty() bool {
  return x.int == 0 && x.bool == false
}

func (x *structure) Clr() {
  x.int = 0 // typ atom.Char
  x.f, x.b = scr.ColF(), scr.ColB()
  x.bool = false
}

func (x *structure) Eq (Y any) bool {
  y := x.imp (Y)
  return x.int == y.int &&
         x.l == y.l && x.c == y.c &&
         x.f == y.f && x.b == y.b &&
         x.bool == y.bool
}

func (x *structure) Copy (Y any) {
  y := x.imp (Y)
  x.int = y.int
  x.l, x.c = y.l, y.c
  x.f.Copy (y.f)
  x.b.Copy (y.b)
  x.bool = y.bool
}

func (x *structure) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *structure) Less (Y any) bool {
  return false
}

func (x *structure) Leq (Y any) bool {
  return false
}

func (x *structure) Colours (f, b col.Colour) {
  x.f.Copy (f)
  x.b.Copy (b)
}

func (x *structure) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *structure) Codelen() uint {
  return 4 + 2 * x.f.Codelen() + 1
}

func (x *structure) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), uint(1)
  copy (s[i:i+a], Encode(uint8(x.int)))
  i++
  copy (s[i:i+a], Encode(uint8(x.l)))
  i++
  copy (s[i:i+a], Encode(uint8(x.c)))
  i++
  copy (s[i:i+a], Encode(uint8(x.w)))
  i++
  a = x.f.Codelen()
  copy (s[i:i+a], x.f.Encode())
  i += a
  copy (s[i:i+a], x.b.Encode())
  i += a
  s[i] = 0
  if x.bool { s[i] = 1 }
  return s
}

func (x *structure) Decode (s Stream) {
  i, a := uint(0), uint(1)
  x.int = int(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  x.l = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  x.c = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  x.w = uint(Decode (uint8(0), s[i:i+a]).(uint8))
  i++
  a = x.f.Codelen()
  x.f.Decode (s[i:i+a])
  i += a
  x.b.Decode (s[i:i+a])
  i += a
  x.bool = false
  if s[i] == 1 { x.bool = true }
}

func (x *structure) Define (t int, n uint) {
  x.int = t
  x.w = n
}

func (x *structure) Index (b bool) {
  x.bool = b
}

func (x *structure) IsIndex() bool {
  return x.bool
}

func (x *structure) Place (l, c uint) {
  x.l, x.c = l, c
}

func (x *structure) Pos() (uint, uint) {
  return x.l, x.c
}

func (x *structure) Width() uint {
  return x.w
}
