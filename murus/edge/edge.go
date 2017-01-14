package edge

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/ker"
  "murus/col"; "murus/scr"; "murus/box"
  "murus/nat"
)
const
  w = 2
type
  edge struct {
              Any "value" // uint or Valuator
         f, b,
       fa, ba col.Colour
              }
var
  bx = box.New()

func init() {
//  bx.SetNumerical()
  bx.Wd(w)
}

func newEdge(v Any) Edge {
  x:= new (edge)
  if v == nil {
    x.Any = nil // nothing to write/edit - edge has value 1
  } else {
    CheckUintOrValuator(v)
    x.Any = Clone(v)
  }
  x.f, x.b = col.StartCols()
  x.fa, x.ba = col.Red, x.b
  return x
}

func (x *edge) imp (Y Any) *edge {
  y, ok:= Y.(*edge)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *edge) Empty() bool {
  return Val(x.Any) == 1 // XXX
}

func (x *edge) Clr() {
  if x.Any == nil {
    // XXX ?
  } else {
    SetVal (&x.Any, 1) // XXX
  }
}

func (x *edge) Eq (Y Any) bool {
  return Eq (x.Any, x.imp(Y).Any)
}

func (x *edge) Less (Y Any) bool {
  if x.Any == nil {
    return false
  }
  return Less (x.Any, x.imp(Y).Any)
}

func (x *edge) Copy (Y Any) {
  y:= x.imp (Y)
  x.Any = Clone(y.Any)
  x.f, x.b, x.fa, x.ba = y.f, y.b, y.fa, y.ba
}

func (x *edge) Clone() Any {
  y:= newEdge (x.Any)
  y.Copy (x)
  return y
}

func (x *edge) Val() uint {
  if x.Any == nil {
    return 1
  }
  return Val(x.Any)
}

func (x *edge) SetVal (n uint) bool {
  if x.Any == nil {
    return false
  }
  SetVal (&x.Any, n)
  return true
}

func (x *edge) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *edge) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
}

func (e *edge) Write (x, y, x1, y1 int, a bool) {
  f, b := e.f, e.b; if a { f, b = e.fa, e.ba }
  scr.ColourF (f)
  scr.Line (x, y, x1, y1)
  if e.Any == nil { return }
  x0, y0 := (x + x1) / 2, (y + y1) / 2
  switch e.Any.(type) {
  case EditorGr:
    e.Any.(EditorGr).WriteGr (x0, y0)
    return
  }
  w1, h1 := scr.Wd1(), scr.Ht1()
  x0 -= int(w1) / 2; y0 -= int(h1) / 2
  bx.Colours (f, b)
  bx.WriteGr (nat.String(Val (e.Any)), x0, y0)
}

func (e *edge) Edit (x, y, x1, y1 int) {
  if e.Any == nil { return }
  a, b := (x + x1) / 2, (y + y1) / 2
  w1, h1 := scr.Wd1(), scr.Ht1()
  a -= int(w1) / 2; b -= int(h1) / 2
//  e.Write (x, y, x1, y1, false) // XXX false ?
  switch e.Any.(type) {
  case EditorGr:
    e.Any.(EditorGr).EditGr (a, b)
    return
  }
  t:= nat.StringFmt (Val(e.Any), w, false)
  bx.Wd (w)
  bx.Colours (e.f, e.b)
  for {
    bx.EditGr (&t, a, b)
    if n, ok := nat.Natural(t); ok {
      SetVal (&e.Any, n)
      break
    }
  }
}

func (x *edge) Codelen() uint {
  return Codelen (x.Any) // XXX
}

func (x *edge) Encode() []byte {
  bs:= make ([]byte, x.Codelen())
  if x.Any != nil {
    copy (bs[:], Encode (x.Any))
  }
  return bs
}

func (x *edge) Decode (bs []byte) {
  if len (bs) == 0 {
    if x.Any != nil { ker.Panic ("huch") }
  } else {
    x.Any = Decode (x.Any, bs[:])
  }
}
