package edg

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/N"
)
type
  edge struct {
     directed bool
              any "value" // uint[8,16,32,64] or Valuator
       x0, y0,
       x1, y1 int
           wd uint
        label bool
 f, b, fm, bm col.Colour
       marked bool
              }
var
  bx = box.New()

func new_(d bool, a any) Edge {
  x := new(edge)
  x.directed = d
  if a == nil {
    a = uint(1)
  }
  CheckUintOrValuator(a)
  x.any = Clone(a)
  x.wd = N.Wd(Val(a))
  x.label = true
  x.f, x.b = col.Black(), col.White()
  x.fm, x.bm = col.StartColsA()
  return x
}

func (x *edge) imp (Y any) *edge {
  y, ok := Y.(*edge)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *edge) Directed() bool {
  return x.directed
}

func (x *edge) Direct (b bool) {
  x.directed = b
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
  return Val(x.any) == 1
}

func (x *edge) Clr() {
  SetVal (&x.any, 1)
}

func (x *edge) Eq (Y any) bool {
  return Eq (x.any, x.imp(Y).any)
}

func (x *edge) Less (Y any) bool {
  return Less (x.any, x.imp(Y).any)
}

func (x *edge) Leq (Y any) bool {
  return Leq (x.any, x.imp(Y).any)
}

func (x *edge) Copy (Y any) {
  y := x.imp (Y)
  x.directed = y.directed
  x.any = Clone(y.any)
  x.wd = y.wd
  x.label = y.label
  x.x0, x.y0, x.x1, x.y1 = y.x0, y.y0, y.x1, y.y1
  x.f, x.b, x.fm, x.bm = y.f, y.b, y.f, y.bm
  x.marked = y.marked
}

func (x *edge) Clone() any {
  y := new_ (x.directed, x.any)
  y.Copy (x)
  return y
}

func (x *edge) Val() uint {
  return Val (x.any)
}

func (x *edge) SetVal (n uint) {
  SetVal (&x.any, n)
}

func (x *edge) Colours (f, b, fm, bm col.Colour) {
  x.f, x.b = f, b
  x.fm, x.bm = fm, bm
}

func (x *edge) Label (b bool) {
  x.label = b
}

func (x *edge) Write() {
  f, b := x.f, x.b
  if x.marked {
    f, b = x.fm, x.bm
  }
  scr.ColourF (f)
  scr.Line (x.x0, x.y0, x.x1, x.y1)
  if x.directed {
//    x0, y0 := (x.x0 + 4 * x.x1) / 5, (x.y0 + 4 * x.y1) / 5
//    scr.CircleFull (x0, y0, 4)
  }
  if x.any == uint(1) { return }
  dx := int(x.wd * scr.Wd1() - scr.Wd1() / 2)
  x0, y0 := (x.x0 + x.x1 - dx) / 2, (x.y0 + x.y1) / 2
  switch x.any.(type) {
  case EditorGr:
    x.any.(EditorGr).WriteGr (x0, y0)
    return
  }
  if x.Empty() { return }
  if ! x.label { return }
  x0 -= int(scr.Wd1()) / 2; y0 -= int(scr.Ht1()) / 2
  bx.Colours (f, b)
  s := N.String (Val(x.any))
  bx.Wd(x.wd)
  bx.WriteGr (s, x0, y0)
}

func (x *edge) Edit() {
  dx := int(x.wd * scr.Wd1() - scr.Wd1() / 2)
  x0, y0 := (x.x0 + x.x1 - dx) / 2, (x.y0 + x.y1) / 2
  x0 -= int(scr.Wd1()) / 2; y0 -= int(scr.Ht1()) / 2
  switch x.any.(type) {
  case EditorGr:
    x.any.(EditorGr).EditGr (x0, y0)
    return
  }
  k := Val (x.any)
  s := N.String(k)
  bx.Wd (x.wd)
  bx.Colours (x.f, x.b)
  for {
    bx.EditGr (&s, x0, y0)
    if i, ok := N.Natural(s); ok {
      SetVal (&x.any, i)
      break
    }
  }
}

func (x *edge) Codelen() uint {
  c := 1 +
       Codelen (x.any) +
       2 * 1 +
       4 * 4 +
       4 * x.f.Codelen()
  return c
}

func (x *edge) Encode() Stream {
  s := make (Stream, x.Codelen())
  s[0] = 0; if x.directed { s[0] = 1 }
  i, a := uint(1), Codelen(x.any)
  copy (s[i:i+a], Encode(x.any))
  i += a
  s[i] = byte(x.wd)
  i++
  s[i] = 0; if x.label { s[i] = 1 }
  i++
  a = 4
  copy (s[i:i+a], Encode(int32(x.x0)))
  i += a
  copy (s[i:i+a], Encode(int32(x.y0)))
  i += a
  copy (s[i:i+a], Encode(int32(x.x1)))
  i += a
  copy (s[i:i+a], Encode(int32(x.y1)))
  i += a
  a = x.f.Codelen()
  copy (s[i:i+a], x.f.Encode())
  i += a
  copy (s[i:i+a], x.b.Encode())
  i += a
  copy (s[i:i+a], x.fm.Encode())
  i += a
  copy (s[i:i+a], x.bm.Encode())
  i += a
  return s
}

func (x *edge) Decode (s Stream) {
  x.directed = s[0] == 1
  i, a := uint(1), Codelen(x.any)
  x.any = Decode (x.any, s[i:i+a])
  i += a
  x.wd = uint(s[i])
  i++
  x.label = s[i] == 1
  i++
  a = 4
  x.x0 = int(Decode(int32(0), s[i:i+a]).(int32))
  i += a
  x.y0 = int(Decode(int32(0), s[i:i+a]).(int32))
  i += a
  x.x1 = int(Decode(int32(0), s[i:i+a]).(int32))
  i += a
  x.y1 = int(Decode(int32(0), s[i:i+a]).(int32))
  i += a
  a = x.f.Codelen()
  x.f.Decode (s[i:i+a])
  i += a
  x.b.Decode (s[i:i+a])
  i += a
  x.fm.Decode (s[i:i+a])
  i += a
  x.bm.Decode (s[i:i+a])
  i += a

}

func (x *edge) Mark (m bool) {
  x.marked = m
}

func (x *edge) Marked () bool {
  return x.marked
}
