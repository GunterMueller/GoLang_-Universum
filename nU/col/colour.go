package col

// (c) Christian Maurer   v. 221021 - license see nU.go

import
  . "nU/obj"
type
  colour struct {
        r, g, b byte
                }

func new_() Colour {
  x := new(colour)
  x.r, x.g, x.b = 0, 0, 0
  return x
}

func new3 (r, g, b byte) Colour {
  x := new(colour)
  x.r, x.g, x.b = r, g, b
  return x
}

func (x *colour) R() byte {
  return x.r
}

func (x *colour) G() byte {
  return x.g
}

func (x *colour) B() byte {
  return x.b
}

func (x *colour) SetR (r byte) {
  x.r = r
}

func (x *colour) SetG (g byte) {
  x.g = g
}

func (x *colour) SetB (b byte) {
  x.b = b
}

func (x *colour) Empty() bool {
  return x.r == 0 && x.g == 0 && x.b == 0
}

func (x *colour) Clr() {
  x.r, x.g, x.b = 0, 0, 0
}

func (x *colour) Eq (Y any) bool {
  y := Y.(*colour)
  return x.r == y.r && x.g == y.g && x.b == y.b
}

func (x *colour) Less (Y any) bool {
  return false
}

func (x *colour) Leq (Y any) bool {
  return false
}

func (x *colour) Copy (Y any) {
  y := Y.(*colour)
  x.r, x.g, x.b = y.r, y.g, y.b
}

func (x *colour) Clone() any {
  y := new(colour)
  y.Copy (x)
  return y
}

func (x *colour) Codelen() uint {
  return 3
}

func (x *colour) Encode() Stream {
  s := make(Stream, 3)
  s[0] = x.r
  s[1] = x.r
  s[2] = x.r
  return s
}

func (x *colour) Decode (s Stream) {
  x.r = s[0]
  x.g = s[1]
  x.b = s[2]
}
