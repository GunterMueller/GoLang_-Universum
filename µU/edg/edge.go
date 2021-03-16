package edg

// (c) Christian Maurer   v. 210311 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/box"
  "µU/n"
)
type
  edge struct {
              bool "directed"
              Any "value" // uint[8,16,32,64] or Valuator
       x0, y0,
       x1, y1 int
           wd uint
        label bool
       cf, cb,
       fa, ba col.Colour
              }
var
  bx = box.New()

func new_(d bool, a Any) Edge {
  x := new(edge)
  x.bool = d
  if a == nil {
    x.Any = uint(1)
  }
  CheckUintOrValuator(a)
  x.Any = Clone(a)
  x.wd = n.Wd(Val(a))
  x.label = true
  x.cf, x.cb = col.StartCols()
  x.fa, x.ba = col.StartColsA()
  return x
}

func (x *edge) imp (Y Any) *edge {
  y, ok := Y.(*edge)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *edge) Directed() bool {
  return x.bool
}

func (x *edge) Direct (b bool) {
  x.bool = b
}

func (e *edge) SetPos0 (x, y int) {
  e.x0, e.y0 = x, y
}

func (e *edge) SetPos1 (x, y int) {
  e.x1, e.y1 = x, y
}

func (x *edge) Pos0() (int, int) {
  return x.x0, x.y0
}

func (x *edge) Pos1() (int, int) {
  return x.x1, x.y1
}

func (x *edge) Empty() bool {
  return Val(x.Any) == 1
}

func (x *edge) Clr() {
  SetVal (&x.Any, 1)
}

func (x *edge) Eq (Y Any) bool {
  return Eq (x.Any, x.imp(Y).Any)
}

func (x *edge) Less (Y Any) bool {
  return Less (x.Any, x.imp(Y).Any)
}

func (x *edge) Copy (Y Any) {
  y := x.imp (Y)
  x.bool = y.bool
  x.Any = Clone(y.Any)
  x.wd, x.label = y.wd, y.label
  x.x0, x.y0, x.x1, x.y1 = y.x0, y.y0, y.x1, y.y1
  x.cf, x.cb, x.fa, x.ba = y.cf, y.cb, y.fa, y.ba
}

func (x *edge) Clone() Any {
  y := new_ (x.bool, x.Any)
  y.Copy (x)
  return y
}

func (x *edge) Val() uint {
  return Val(x.Any)
}

func (x *edge) SetVal (n uint) {
  SetVal (&x.Any, n)
}

func (x *edge) Colours (f, b col.Colour) {
  x.cf, x.cb = f, b
}

func (x *edge) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
}

func (x *edge) Label (b bool) {
  x.label = b
}

func (x *edge) Write () {
  x.Write1 (false)
}

func (x *edge) Write1 (a bool) {
  f, b := x.cf, x.cb
  if a {
    f, b = x.fa, x.ba
  }
  scr.ColourF (f)
  scr.Line (x.x0, x.y0, x.x1, x.y1)
  if x.bool {
    x0, y0 := (x.x0 + 4 * x.x1) / 5, (x.y0 + 4 * x.y1) / 5
    scr.CircleFull (x0, y0, 4)
  }
  dx := int(x.wd * scr.Wd1() - scr.Wd1() / 2)
  x0, y0 := (x.x0 + x.x1 - dx) / 2, (x.y0 + x.y1) / 2
  switch x.Any.(type) {
  case EditorGr:
    x.Any.(EditorGr).WriteGr (x0, y0)
    return
  }
  if x.Empty() { return }
  if ! x.label { return }
  x0 -= int(scr.Wd1()) / 2; y0 -= int(scr.Ht1()) / 2
  bx.Colours (f, b)
  s := n.String(Val(x.Any))
  bx.Wd(x.wd)
  bx.WriteGr (s, x0, y0)
}

func (x *edge) Edit() {
  dx := int(x.wd * scr.Wd1() - scr.Wd1() / 2)
  x0, y0 := (x.x0 + x.x1 - dx) / 2, (x.y0 + x.y1) / 2
  x0 -= int(scr.Wd1()) / 2; y0 -= int(scr.Ht1()) / 2
//  x.Write1 (false) // XXX false ?
  switch x.Any.(type) {
  case EditorGr:
    x.Any.(EditorGr).EditGr (x0, y0)
    return
  }
  k := Val(x.Any)
  s := n.String(k)
  bx.Wd (x.wd)
  bx.Colours (x.cf, x.cb)
  for {
    bx.EditGr (&s, x0, y0)
    if i, ok := n.Natural(s); ok {
      SetVal (&x.Any, i)
      break
    }
  }
}

func (x *edge) Codelen() uint {
  c := 1 +
       Codelen (x.Any) +
       2 * 1 +
       4 * 4 +
       4 * x.cf.Codelen()
  if c != 35 { errh.Error("Kacke", c) }
  return 35
}

func (x *edge) Encode() Stream {
  bs := make (Stream, x.Codelen())
  bs[0] = 0; if x.bool { bs[0] = 1 }
  i, a := uint(1), Codelen(x.Any)
  copy (bs[i:i+a], Encode(x.Any))
  i += a
  bs[i] = byte(x.wd)
  i++
  bs[i] = 0; if x.label { bs[i] = 1 }
  i++
  a = 4
  copy (bs[i:i+a], Encode(int32(x.x0)))
  i += a
  copy (bs[i:i+a], Encode(int32(x.y0)))
  i += a
  copy (bs[i:i+a], Encode(int32(x.x1)))
  i += a
  copy (bs[i:i+a], Encode(int32(x.y1)))
  i += a
  a = x.cf.Codelen()
  copy (bs[i:i+a], x.cf.Encode())
  i += a
  copy (bs[i:i+a], x.cb.Encode())
  i += a
  copy (bs[i:i+a], x.fa.Encode())
  i += a
  copy (bs[i:i+a], x.ba.Encode())
  return bs
}

func (x *edge) Decode (bs Stream) {
  x.bool = bs[0] == 1
  i, a := uint(1), Codelen(x.Any)
  x.Any = Decode (x.Any, bs[i:i+a])
  i += a
  x.wd = uint(bs[i])
  i++
  x.label = bs[i] == 1
  i++
  a = 4
  x.x0 = int(Decode(int32(0), bs[i:i+a]).(int32))
  i += a
  x.y0 = int(Decode(int32(0), bs[i:i+a]).(int32))
  i += a
  x.x1 = int(Decode(int32(0), bs[i:i+a]).(int32))
  i += a
  x.y1 = int(Decode(int32(0), bs[i:i+a]).(int32))
  i += a
  a = x.cf.Codelen()
  x.cf.Decode (bs[i:i+a])
  i += a
  x.cb.Decode (bs[i:i+a])
  i += a
  x.fa.Decode (bs[i:i+a])
  i += a
  x.ba.Decode (bs[i:i+a])
}
