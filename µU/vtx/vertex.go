package vtx

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
)
type
  vertex struct {
                EditorGr
       valuable bool
          width,
         height uint
           x, y int
   f, b, fa, ba col.Colour
                }

func new_(e EditorGr, w, h uint) Vertex {
  x := new (vertex)
  x.EditorGr = e.(Object).Clone().(EditorGr) // is Valuator, if of type *bnat.Natural
  switch x.EditorGr.(type) {
  case Valuator:
    x.valuable = true
  default:
    x.valuable = false
  }
  x.width, x.height = w, h
  x.f, x.b = col.StartCols()
  x.fa, x.ba = col.StartColsA()
  return x
}

func (x *vertex) imp (Y any) *vertex {
  y, ok := Y.(*vertex)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *vertex) Content() EditorGr {
  return x.EditorGr.(Object).Clone().(EditorGr)
}

func (x *vertex) Wd() uint {
  return x.width
}

func (x *vertex) Ht() uint {
  return x.height
}

func (x *vertex) Size() (uint, uint) {
  return x.width, x.height
}

func (v *vertex) Set (x, y int) {
  v.x, v.y = x, y
}

func (x *vertex) Pos() (int, int) {
  return x.x, x.y
}

func (x *vertex) Empty() bool {
  return x.EditorGr.(Object).Empty()
}

func (x *vertex) Clr() {
  x.EditorGr.(Object).Clr()
}

func (x *vertex) Eq (Y any) bool {
  y := x.imp(Y)
  return x.EditorGr.(Object).Eq (y.EditorGr) &&
         x.x == y.x && x.y == y.y
}

func (x *vertex) Less (Y any) bool {
  return x.EditorGr.(Object).Less (x.imp(Y).EditorGr)
}

func (x *vertex) Copy (Y any) {
  y := x.imp(Y)
  x.EditorGr.(Object).Copy (y.EditorGr.(Object))
  x.width, x.height = y.width, y.height
  x.x, x.y = y.x, y.y
  x.f, x.b = y.f, y.b
  x.fa, x.ba = y.fa, y.ba
}

func (x *vertex) Clone() any {
  y := new_(x.EditorGr, x.width, x.height).(*vertex)
  y.Copy (x)
  return y
}

func (x *vertex) Mouse() {
  xm, ym  := scr.MousePosGr()
  x.x, x.y = xm, ym
  wd, ht  := int(scr.Wd()), int(scr.Ht())
  w1, h1  := int(scr.Wd1()), int(scr.Ht1())
  w, h  := int(x.width) * w1, int(x.height) * h1
  if x.x < w { x.x = w }
  if x.x + w >= wd { x.x = wd - 1 - w }
  if x.y < h { x.y = h }
  if x.y + h >= ht { x.y = ht - 1 - h }
}

func (x *vertex) UnderMouse() bool {
  t := int(x.height * scr.Ht1())
  xm, ym  := scr.MousePosGr()
  dx := xm - x.x; if dx < 0 { dx = -dx }
  dy := ym - x.y; if dy < 0 { dy = -dy }
  return dx < t && dy < t
}

func (x *vertex) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *vertex) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *vertex) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
}

func (x *vertex) Defined (s string) bool {
  switch x.EditorGr.(type) {
  case Stringer:
    return x.EditorGr.(Stringer).Defined (s)
  }
  return false
}

func (x *vertex) String() string {
  switch x.EditorGr.(type) {
  case Stringer:
    return x.EditorGr.(Stringer).String()
  }
  return ""
}

func (x *vertex) Write() {
  x.Write1 (false)
}

func (x *vertex) Write1 (b bool) {
  cf, cb := x.f, x.b
  if b {
    cf, cb = x.fa, x.ba
  }
  w1, h1 := scr.Wd1() * x.width, scr.Ht1() * x.height
  r := w1
  if h1 > r {
    r = h1
  }
  if x.width * x.height > 0 {
    const r0 = 3
    scr.Colours (cf, cb)
    if r <= r0 {
      panic ("vertex.go: r <= r0")
      scr.CircleFull (x.x, x.y, r0)
    } else {
      scr.ColourF (cf)
      scr.CircleFull (x.x, x.y, r / 2 + 1)
    }
  }
  x.EditorGr.(col.Colourer).Colours (cb, cf)
  x.EditorGr.WriteGr (x.x - int(w1) / 2, x.y - int(h1) / 2 + 1)
}

func (x *vertex) Edit() {
  x.Write()
  w1, h1  := scr.Wd1() * x.width, scr.Ht1() * x.height
  x.EditorGr.EditGr (x.x - int(w1) / 2 - int(scr.Wd1()), x.y - int(h1) / 2)
}

func (x *vertex) Codelen() uint {
  return x.EditorGr.(Object).Codelen() +
         2 +               // width, height
         2 * C0 +          // x, y
         4 * x.f.Codelen()
}

func (x *vertex) Encode() Stream {
  bs := make (Stream, x.Codelen())
  i, a  := uint(0), x.EditorGr.(Object).Codelen()
  copy (bs[i:i+a], x.EditorGr.(Object).Encode())
  i += a
  bs[i] = uint8(x.width)
  i++
  bs[i] = uint8(x.height)
  i++
  a = C0
  copy (bs[i:i+a], Encode(x.x))
  i +=a
  copy (bs[i:i+a], Encode(x.y))
  i +=a
  a = x.f.Codelen()
  copy (bs[i:i+a], x.f.Encode())
  i += a
  copy (bs[i:i+a], x.b.Encode())
  i += a
  copy (bs[i:i+a], x.fa.Encode())
  i += a
  copy (bs[i:i+a], x.ba.Encode())
  return bs
}

func (x *vertex) Decode (bs Stream) {
  i, a  := uint(0), x.EditorGr.(Object).Codelen()
//  if a + 2 + 2 * C0 + 4 * col.Codelen() >= uint(len(bs)) { return false }
  x.EditorGr.(Object).Decode (bs[i:i+a])
  i += a
  x.width = uint(bs[i])
  i++
  x.height = uint(bs[i])
  i++
  a = C0
  x.x = Decode(0, bs[i:i+a]).(int)
  i += a
  x.y = Decode(0, bs[i:i+a]).(int)
  i += a
  a = x.f.Codelen()
  x.f.Decode (bs[i:i+a])
  i += a
  x.b.Decode (bs[i:i+a])
  i += a
  x.fa.Decode (bs[i:i+a])
  i += a
  x.ba.Decode (bs[i:i+a])
}

func (x *vertex) Val() uint {
  if x.valuable {
    return x.EditorGr.(Valuator).Val()
  }
  return 0
}

func (x *vertex) SetVal (n uint) {
  if x.valuable {
    x.EditorGr.(Valuator).SetVal (n)
  }
}
