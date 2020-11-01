package pt

// (c) Christian Maurer   v. 201027 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/vect"
)
type
  point struct {
             c Class
               uint32 "number"
        colour col.Colour
    vect, norm vect.Vector
               }

func new_() Point {
  x := new (point)
  x.colour, _ = col.StartCols()
  x.vect, x.norm = vect.New(), vect.New()
  return x
}

func (x *point) imp (Y Any) *point {
  y, ok := Y.(*point)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *point) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.c == y.c &&
         x.uint32 == y.uint32 &&
         x.vect.Eq (y.vect) && x.norm.Eq (y.norm)
}

func (x *point) Less (Y Any) bool {
  return false // objects are not comparable
}

func (x *point) Copy (Y Any) {
  y := x.imp (Y)
  x.c = y.c
  x.uint32 = y.uint32
  x.colour = y.colour
  x.vect.Copy (y.vect)
  x.norm.Copy (y.norm)
}

func (x *point) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *point) Empty() bool {
  return x.c == Start
}

func (x *point) Clr() {
  x.c = Start
}

func (x *point) Set (c Class, a uint, f col.Colour, v, n vect.Vector) {
  x.c = c
  x.uint32 = uint32(a)
  x.colour = f
  x.vect.Copy (v)
  x.norm.Copy (n)
}

func (x *point) Class() Class {
  return x.c
}

func (x *point) Number() uint {
  return uint(x.uint32)
}

func (x *point) Colour() col.Colour {
  return x.colour
}

/*/
func (x *point) Write (i uint) {
//  x.vect.Colours (x.colour, scr.ScrColB())
//  x.norm.Colours (x.colour, scr.ScrColB())
  x.vect.Write (i,  0)
  x.norm.Write (i, 20)
}
/*/

func (x *point) Read() vect.Vector {
  return x.vect.Clone().(vect.Vector)
}

func (x *point) Read2() (vect.Vector, vect.Vector) {
  return x.vect.Clone().(vect.Vector), x.norm.Clone().(vect.Vector)
}

func (x *point) Codelen() uint {
  return 1 +
         4 +
         x.colour.Codelen() +
         2 * x.vect.Codelen()
}

func (x *point) Encode() []byte {
  b := make ([]byte, x.Codelen())
  b[0] = byte(x.c)
  i, a := uint(1), uint(4)
  copy (b[i:i+a], Encode (x.uint32))
  i += a
  a = x.colour.Codelen()
  copy (b[i:i+a], Encode (x.colour))
  i += a
  a = x.vect.Codelen()
  copy (b[i:i+a], x.vect.Encode())
  i += a
  copy (b[i:i+a], x.norm.Encode())
  return b
}

func (x *point) Decode (b []byte) {
  x.c = Class(b[0])
  i, a := uint(1), uint(4)
  x.uint32 = Decode (uint32(0), b[i:i+a]).(uint32)
  i += a
  a = x.colour.Codelen()
  x.colour.Decode (b[i:i+a])
  i += a
  a = x.vect.Codelen()
  x.vect.Decode (b[i:i+a])
  i += a
  x.norm.Decode (b[i:i+a])
}
