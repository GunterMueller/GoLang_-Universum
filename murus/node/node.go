package node

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/scr"
  "murus/pos"
)
type
  node struct {
              EditorGr
        width,
       height uint
              pos.Position
         f, b,
       fa, ba col.Colour
              }

func new_(e EditorGr, w, h uint) Node {
  x:= new (node)
  x.EditorGr = e.Clone().(EditorGr)
  x.width, x.height = w, h
  x.Position = pos.New (x.width, x.height)
  x.f, x.b = col.StartCols()
  x.fa, x.ba = col.StartColsA()
  return x
}

func (x *node) imp (Y Any) *node {
  y, ok:= Y.(*node)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *node) Content() EditorGr {
  return x.EditorGr.Clone().(EditorGr)
}

func (x *node) Size() (uint, uint) {
  return x.width, x.height
}

func (n *node) Set (x, y int) {
  n.Position.Set (x, y)
}

func (x *node) Pos() (int, int) {
  return x.Position.Pos()
}

func (x *node) Contour() (uint, uint) {
  return x.Position.Contour()
}

func (x *node) Empty() bool {
  return x.EditorGr.Empty()
}

func (x *node) Clr() {
  x.EditorGr.Clr()
}

func (x *node) Eq (Y Any) bool {
  y:= x.imp(Y)
  return x.EditorGr.Eq (y.EditorGr) &&
         x.Position.Eq (y.Position)
}

func (x *node) Copy (Y Any) {
  y:= x.imp(Y)
  yy := y.EditorGr
  x.EditorGr.Copy (yy)
  x.width, x.height = y.width, y.height
  x.Position.Copy (y.Position)
  x.f, x.b = y.f, y.b
  x.fa, x.ba = y.fa, y.ba
}

func (x *node) Less (Y Any) bool {
  return x.EditorGr.Less (x.imp(Y).EditorGr)
}

func (x *node) Clone() Any {
  w, h := x.Position.Contour()
  y := new_(x.EditorGr, w, h).(*node)
  y.Copy (x)
  return y
}

func (x *node) Mouse() {
  x.Position.Mouse()
}

func (x *node) UnderMouse () bool {
  return x.Position.UnderMouse()
}

func (x *node) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *node) ColoursA (f, b col.Colour) {
  x.fa, x.ba = f, b
}

/* func (x *node) Defined (s string) bool {
  switch x.EditorGr.(type) {
  case Stringer:
    return x.EditorGr.(Stringer).Defined (s)
  }
  return false
}

func (x *node) String() string {
  switch x.EditorGr.(type) {
  case Stringer:
    return x.EditorGr.(Stringer).String()
  }
  return ""
} */

func (n *node) Write() {
  n.Write1 (false)
}

func (n *node) Write1 (a bool) {
  f, b := n.f, n.b; if a { f, b = n.fa, n.ba }
  x, y := n.Position.Pos()
  w, h := n.Position.Contour()
  w1, h1 := int(scr.Wd1()), int(scr.Ht1())
  n.Position.Colours (f, b)
  n.Position.Write()
  n.EditorGr.Colours (f, b)
  n.EditorGr.WriteGr (x - int(w) * w1 / 2 + 1, y - int(h) * h1 / 2)
}

func (n *node) Write3 (e Any, N Node, a bool) {
  n.Write3dir (e, false, N, a)
}

func (n *node) Write3dir (e Any, d bool, N Node, a bool) {
  f := n.f; if a { f = n.fa }
  x, y := n.Position.Pos()
  x1, y1 := N.(*node).Position.Pos()
  scr.ColourF (f)
  scr.Line (x, y, x1, y1)
  if d {
    x0, y0 := (x + 4 * x1) / 5, (y + 4 * y1) / 5
    scr.CircleFull (x0, y0, 4)
  }
}

func (n *node) Edit() {
  n.Write()
  x, y := n.Position.Pos()
  w, h := n.Position.Contour()
  w1, h1 := int(scr.Wd1()), int(scr.Ht1())
  n.EditorGr.EditGr (x - int(w) * w1 / 2 + 1, y - int(h) * h1 / 2)
}

func (x *node) Codelen() uint {
  return x.EditorGr.Codelen() +
         x.Position.Codelen() +
         4 * col.Codelen()
}

func (x *node) Encode() []byte {
  bs:= make ([]byte, x.Codelen())
  i, a := uint(0), x.EditorGr.Codelen()
  copy (bs[i:i+a], x.EditorGr.Encode())
  i += a
  a = x.Position.Codelen()
  copy (bs[i:i+a], x.Position.Encode())
  i += a
  a = col.Codelen()
  copy (bs[i:i+a], col.Encode (x.f))
  i += a
  copy (bs[i:i+a], col.Encode (x.b))
  i += a
  copy (bs[i:i+a], col.Encode (x.fa))
  i += a
  copy (bs[i:i+a], col.Encode (x.ba))
  return bs
}

func (x *node) Decode (bs []byte) {
  i, a := uint(0), x.EditorGr.Codelen()
  x.EditorGr.Decode (bs[i:i+a])
  i += a
  a = x.Position.Codelen()
  x.Position.Decode (bs[i:i+a])
  i += a
  a = col.Codelen()
  col.Decode (&x.f, bs[i:i+a])
  i += a
  col.Decode (&x.b, bs[i:i+a])
  i += a
  col.Decode (&x.fa, bs[i:i+a])
  i += a
  col.Decode (&x.ba, bs[i:i+a])
}
