package vtx

// (c) Christian Maurer   v. 221021 - license see µU.go

import (. "nU/obj"; "nU/col"; "nU/scr")

type vertex struct {
  uint "content"
  x, y uint // position
}

func new_(n uint) Vertex {
  x := new (vertex)
  x.uint = n
  return x
}

func (v *vertex) Set (x, y uint) {
  v.x, v.y = x, y
}

func (x *vertex) Pos() (uint, uint) {
  return x.x, x.y
}

func (x *vertex) Empty() bool {
  return x.uint == 0
}

func (x *vertex) Clr() {
  x.uint = 0
}

func (x *vertex) Eq (Y any) bool {
  y := Y.(*vertex)
  return x.uint == y.uint &&
         x.x == y.x && x.y == y.y
}

func (x *vertex) Less (Y any) bool {
  return x.uint < Y.(*vertex).uint
}

func (x *vertex) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *vertex) Copy (Y any) {
  y := Y.(*vertex)
  x.uint = y.uint
  x.x, x.y = y.x, y.y
}

func (x *vertex) Clone() any {
  y  := new_(x.uint)
  y.Copy (x)
  return y
}

var c0 = C0()

func (x *vertex) Codelen() uint {
  return 3 * c0
}

func (x *vertex) Encode() Stream {
  bs := make (Stream, x.Codelen())
  i, a  := uint(0), c0
  copy (bs[i:i+a], Encode(x.uint))
  i += a
  copy (bs[i:i+a], Encode(x.x))
  i +=a
  copy (bs[i:i+a], Encode(x.y))
  return bs
}

func (x *vertex) Decode (bs Stream) {
  i, a  := uint(0), c0
  x.uint = Decode (uint(0), bs[i:i+a]).(uint)
  i += a
  x.x = Decode(uint(0), bs[i:i+a]).(uint)
  i += a
  x.y = Decode(uint(0), bs[i:i+a]).(uint)
}

func (x *vertex) Val() uint {
  return x.uint
}

func (x *vertex) SetVal (n uint) bool {
  x.uint = n
  return true
}

func (x *vertex) Write() {
  x.Write1 (false)
}

func (x *vertex) Write1 (a bool) {
  f := col.White()
  if a { f = col.Red() }
  scr.ColourF (f)
  scr.WriteNat (x.uint, x.x, x.y)
}

func w (v any, a bool) {
  f := col.White()
  if a { f = col.Red() }
  scr.ColourF (f)
  v.(*vertex).Write1(a)
}

func w2 (v0, v1 any, a bool) {
  x0, y0 := v0.(*vertex).Pos()
  x1, y1 := v1.(*vertex).Pos()
  f := col.White()
  if a { f = col.Red() }
  scr.ColourF (f)
  scr.Line (x0, y0, x1, y1)
}
