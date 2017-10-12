package text

// (c) Christian Maurer   v. 170810 - license see µu.go

import (
  "µu/rand"
  . "µu/obj"
  "µu/z"
  "µu/str"
  "µu/col"
  "µu/scr"
  "µu/box"
  "µu/font"
  "µu/pbox"
)
type
  text struct {
              uint "uint"
              string "string"
       cF, cB col.Colour
              font.Font
              }
var (
  bx = box.New()
  pbx = pbox.New()
  Vokale, Konsonanten string = "aeiouy", "bcdfghjklmnpqrstvwxz"
  upper, lower string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ", "abcdefghijklmnopqrstuvwxyz"
)

func new_(n uint) Text {
  if n == 0 { return nil }
  x := new (text)
  x.uint = n
  x.string = str.Clr (n)
  x.cF, x.cB = scr.StartCols()
  x.Font = font.Normal
  return x
}

func (x *text) imp(Y Any) *text {
  y, ok := Y.(*text)
  if ! ok || x.uint != y.uint { TypeNotEqPanic (x, Y) }
  return y
}

func randomvokal() byte {
  return Vokale [rand.Natural (6)]
}

func randomkonsonant() byte {
  return Konsonanten [rand.Natural (20)]
}

func (x *text) Generate() {
  b := make ([]byte, x.uint)
  b[0] = upper [rand.Natural (uint(len (upper)))]
  for i := uint(1); i < x.uint; i++ {
    b[i] = lower [rand.Natural (uint(len (lower)))]
  }
  x.string = string (b)
  z.ToHellWithUTF8 (&x.string)
/*
  n := 3 + rand.Natural (x.uint - 2)
  if n >= x.uint {
    n = x.uint - 1
  }
  b := rand.Natural (2) % 2 == 1
  s := x.uint
  for i := 0; i < int(n); i++ {
    if (i % 2 == 1) == b {
      B[i] = randomkonsonant()
      if B[i] == 's' && i + 2 < int(n) {
        s = uint(i)
      }
    } else {
      B[i] = randomvokal()
    }
    if i == 0 {
//      B[i] = CAP (B[i])
    }
  }
  if s < x.uint {
    B[s + 1] = 'c'
    B[s + 2] = 'h'
  }
  for i := n; i <= x.uint; i++ {
//    B[i] = 0C
  }
*/
}

func (x *text) Empty() bool {
  return str.Empty (x.string)
}

func (x *text) Clr() {
  x.string = str.Clr (x.uint)
}

func (x *text) Copy (Y Any) {
  y := x.imp (Y)
  x.string = y.string
  x.cF, x.cB = y.cF, y.cB
}

func (x *text) Clone() Any {
  y := new_(x.uint)
  y.Copy (x)
  return y
}

func (x *text) Eq (Y Any) bool {
  return x.string == x.imp (Y).string
}

func (x *text) Less (Y Any) bool {
  return str.Less (x.string, x.imp (Y).string)
}

func (x *text) Equiv (Y Text) bool {
  return str.Equiv (x.string, x.imp (Y).string)
}

func (x *text) Defined (s string) bool {
  s = str.Lat1 (s)
  str.Norm (&s, x.uint)
  x.string = s
  return true
}

func (x *text) String() string {
  return x.string
}

func (x *text) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *text) Write (l, c uint) {
  bx.Wd (x.uint)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.string, l, c)
}

func (x *text) Edit (l, c uint) {
  bx.Wd (x.uint)
  bx.Colours (x.cF, x.cB)
  bx.Edit (&x.string, l, c)
}

func (x *text) SetFont (f font.Font) {
  x.Font = f
}

func (x *text) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.string, l, c)
}

func (x *text) Codelen() uint {
  return x.uint
}

func (x *text) Encode() []byte {
  return ([]byte)(x.string)
}

func (x *text) Decode (b []byte) {
  if uint(len (b)) == x.uint {
    x.string = string(b)
//    str.Lat1 (&x.string)
  } else {
    x.string = str.Clr (x.uint)
  }
}

func (x *text) Split() []Text {
  s := x.string
  n, ws, ps := str.Split (s)
  ts := make ([]Text, n)
  for i := uint(0); i < n; i++ {
    ts[i] = New (x.uint)
    ts[i].Defined (str.Part (s, ps[i], uint(len(ws[i]))))
  }
  return ts
}

func (x *text) Len() uint {
  return x.uint
}

/////////////////////////////////////////////////////

func (x *text) Sub (Y Text) bool {
  _, ok := str.Sub (x.string, x.imp(Y).string)
  return ok
}

func (x *text) Sub0 (Y Text) bool {
  return str.Sub0 (x.string, x.imp(Y).string)
}

func (x *text) EquivSub (Y Text) (uint, bool) {
  return str.EquivSub (x.string, x.imp(Y).string)
}

func (x *text) Pos (b byte) (uint, bool) {
  return str.Pos (x.string, b)
}

func (x *text) Rem (p, n uint) {
  str.Rem (&x.string, p, n)
  str.Norm (&x.string, x.uint)
}

func (x *text) Cut (Y Text, p, n uint) {
  y := x.imp (Y)
  x.string = str.Part (y.String(), p, n)
  str.Norm (&x.string, x.uint)
}

func (x *text) ProperLen() uint {
  return str.ProperLen (x.string)
}

func (x *text) IsCap0() bool {
  s := x.String()
  return z.IsCap (s[0])
}

func (x *text) ToUpper() {
  s := x.String()
  str.ToUpper (&s)
  x.Defined (s)
}

func (x *text) ToLower() {
  s := x.String()
  str.ToLower (&s)
  x.Defined (s)
}

func (x *text) ToUpper0() {
  s := x.String()
  str.ToUpper0 (&s)
  x.Defined (s)
}

func (x *text) ToLower0() {
  s := x.String()
  str.ToLower0 (&s)
  x.Defined (s)
}

func (X *text) WriteGr (x, y int) {
  bx.Wd (X.uint)
  bx.Colours (X.cF, X.cB)
  bx.WriteGr (X.string, x, y)
}

func (X *text) EditGr (x, y int) {
  bx.Wd (X.uint)
  bx.Colours (X.cF, X.cB)
  bx.EditGr (&X.string, x, y)
}
