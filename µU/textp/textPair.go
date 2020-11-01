package textp

import (
  . "µU/obj"
  "µU/col"
  "µU/font"
  "µU/box"
  "µU/text"
)
type
  textPair struct {
             t, t1 text.Text
             n, n1 uint
                  }
var
  bx = box.New()

func init() {
  bx.Wd (9)
}

func new_(m, n uint) TextPair {
  x := new(textPair)
  x.n, x.n1 = m, n
  x.t = text.New (x.n)
  x.t1 = text.New (x.n1)
  return x
}

func (x *textPair) imp(Y Any) *textPair {
  y, ok := Y.(*textPair)
  if ! ok || x.n != y.n || x.n1 != y.n1 { TypeNotEqPanic (x, Y) }
  return y
}

func (x *textPair) Len() (uint, uint) {
  return x.n, x.n1
}

func (x *textPair) Empty() bool {
  return x.t.Empty() && x.t1.Empty()
}

func (x *textPair) Clr() {
  x.t.Clr()
  x.t1.Clr()
}

func (x *textPair) Copy (Y Any) {
  y := x.imp (Y)
  x.t.Copy (y.t)
  x.t1.Copy (y.t1)
}

func (x *textPair) Clone() Any {
  y := new_(x.n, x.n1)
  y.Copy (x)
  return y
}

func (x *textPair) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.t.Eq (y.t) &&
         x.t1.Eq (y.t1)
}

func (x *textPair) Less (Y Any) bool {
  y := x.imp (Y)
  if x.t.Eq (y.t) {
    return x.t1.Less (y.t1)
  }
  return x.t.Less (y.t)
}

func (x *textPair) Leq (Y Any) bool {
  y := x.imp (Y)
  return x.Eq(y) || x.Less(y)
}

func (x *textPair) Colours (f, b col.Colour) {
  x.t.Colours (f, b)
  x.t1.Colours (f, b)
}

func (x *textPair) SetFont (f font.Font) {
  x.t.SetFont (f)
  x.t1.SetFont (f)
}

// func (x *textPair) SetFontsize (s font.Size) {
//   x.t.SetFontsize (s)
 //  x.t1.SetFontsize (s)
// }

func (x *textPair) Print (l, c uint) {
  x.t.Print (l, c)
  x.t1.Print (l + 1, c)
}

func (x *textPair) Codelen() uint {
  return x.n + x.n1
}

func (x *textPair) Encode() Stream {
  s := make(Stream, x.Codelen())
  m := x.t.Codelen()
  copy (s[:m], x.t.Encode())
  copy (s[m:], x.t1.Encode())
  return s
}

func (x *textPair) Decode (s Stream) {
  m := x.t.Codelen()
  x.t.Decode (s[:m])
  x.t1.Decode (s[m:])
}

func (x *textPair) Write (l, c uint) {
  x.t.Write (l, c)
  x.t1.Write (l + 1, c)
}

func (x *textPair) Edit (l, c uint) {
  x.t.Edit (l, c)
  x.t1.Edit (l + 1, c)
}
