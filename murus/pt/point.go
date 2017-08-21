package pt

// (c) murus.org  v. 170820 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/scr"
  "murus/vect"
)
type
  point struct {
         class Class
               uint32 "number"
        colour col.Colour
        vector,
        normal vect.Vector
               }

func new_() Point {
  x := new(point)
  x.colour, _ = scr.StartCols()
  x.vector, x.normal = vect.New(), vect.New()
  return x
}

func (x *point) imp (Y Any) *point {
  y, ok := Y.(*point)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *point) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.class == y.class &&
         x.uint32 == y.uint32 &&
         x.vector.Eq (y.vector) && x.normal.Eq (y.normal)
}

func (x *point) Less (Y Any) bool {
  return false // objects are not comparable
}

func (x *point) Copy (Y Any) {
  y := x.imp(Y)
  x.class = y.class
  x.uint32 = y.uint32
  x.colour = y.colour
  x.vector.Copy (y.vector)
  x.normal.Copy (y.normal)
}

func (x *point) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *point) Empty() bool {
  return x.class == Undef
}

func (x *point) Clr() {
  x.class = Undef
}

func (x *point) Set (c Class, a uint, f col.Colour, v, n vect.Vector) {
  x.class = c
  x.uint32 = uint32(a)
  x.colour = f
  x.vector.Copy (v)
  x.normal.Copy (n)
}

func (x *point) Class() Class {
  return x.class
}

func (x *point) Number() uint {
  return uint(x.uint32)
}

func (x *point) Colour() col.Colour {
  return x.colour
}

func (x *point) Write (i uint) {
  x.vector.Colours (x.colour, scr.ScrColB())
  x.normal.Colours (x.colour, scr.ScrColB())
  x.vector.Write (i,  0)
  x.normal.Write (i, 20)
}

func (x *point) Read() vect.Vector {
  return x.vector.Clone().(vect.Vector)
}

func (x *point) Read2() (vect.Vector, vect.Vector) {
  return x.vector.Clone().(vect.Vector), x.normal.Clone().(vect.Vector)
}

func (x *point) Codelen() uint {
  return 1 +
         4 +
         col.Codelen() +
         2 * x.vector.Codelen()
}

func (x *point) Encode() []byte {
  b := make ([]byte, x.Codelen())
  b[0] = byte(x.class)
  i, a := uint(1), uint(4)
  copy (b[i:i+a], Encode (x.uint32))
  i += a
  a = col.Codelen()
  copy (b[i:i+a], col.Encode (x.colour))
  i += a
  a = x.vector.Codelen()
  copy (b[i:i+a], x.vector.Encode())
  i += a
  copy (b[i:i+a], x.normal.Encode())
  return b
}

func (x *point) Decode (b []byte) {
  x.class = Class(b[0])
  i, a := uint(1), uint(4)
  x.uint32 = Decode (uint32(0), b[i:i+a]).(uint32)
  i += a
  a = col.Codelen()
  col.Decode (&x.colour, b[i:i+a])
  i += a
  a = x.vector.Codelen()
  x.vector.Decode (b[i:i+a])
  i += a
  x.normal.Decode (b[i:i+a])
}
