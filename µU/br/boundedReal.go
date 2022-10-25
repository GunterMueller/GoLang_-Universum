package br

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/font"
  "µU/pbox"
//  "µU/n"
  "µU/r"
)
const
  m = 20 // XXX ?
type
  breal struct {
               float64
             d uint // number of digits before the dot
            wd uint // 
       invalid float64
        f, b col.Colour
               font.Font
               }
var
  pbx = pbox.New()

func new_(d uint) Real {
  x := new(breal)
  x.float64 = 0 // r.NaN()
  x.d = d
  x.wd = 1 + d + 1 + 2 // 1 for sign, 1 for dot, 2 digits after the dot
  x.invalid = exp (d + 1)
  x.f, x.b = col.StartCols()
  return x
}

func (x *breal) imp (Y any) *breal {
  y, ok := Y.(*breal)
  if ! ok || x.wd != y.wd { TypeNotEqPanic (x, Y) }
  return y
}

func exp (n uint) float64 {
  if n == 0 { return 1 }
  return 10 * exp (n - 1)
}

func (x *breal) Width() uint {
  return x.wd
}

func (x *breal) Empty() bool {
  return x.float64 == x.invalid
}

func (x *breal) Clr() {
  x.float64 = x.invalid
}

func (x *breal) Copy (Y any) {
  y := x.imp (Y)
  x.float64 = y.float64
  x.d, x.wd = y.d, y.wd
  x.invalid = y.invalid
  x.f, x.b = y.f, y.b
  x.Font = y.Font
}

func (x *breal) Clone() any {
  y := new_(x.d)
  y.Copy(x)
  return y
}

func (x *breal) Eq (Y any) bool {
  return x.float64 == x.imp (Y).float64
}

func (x *breal) Less (Y any) bool {
  y := x.imp (Y)
  if x.float64 == x.invalid || y.float64 == y.invalid {
    return false
  }
  return x.float64 < y.float64
}

func (x *breal) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *breal) Codelen() uint {
  return 8
}

func (x *breal) Encode() Stream {
  b := make (Stream, x.Codelen())
  copy (b[:8], Encode (x.float64))
  return b
}

func (x *breal) Decode (b Stream) {
  x.float64 = float64(Decode (float64(0), b[:8]).(float64))
}

func (x *breal) Defined (s string) bool {
  ok := true
  if x.float64, _, ok = r.Real (s); ok {
    return true
  }
  x.float64 = 0 // r.NaN()
  return false
/*/ TODO
  if uint(len (s)) > x.wd { return false }
  str.OffSpc (&s)
//  n := x.wd / 2
//  P, L := make ([]uint, n), make ([]uint, n)
//  n = n.NDigitSequences (s, &P, &L)
  k, t, p, l := n.DigitSequences (s)
  if k == 0 || k > 2 || l[0] > x.pre {
    return false
  }
  if k == 2 {
    c := s[p[1] - 1]
    if l[1] > 2 || ! (c == '.' || c == ',') {
      return false
    }
  }
  n1, _ := n.Natural (t[0])
  x.float64 = float64(n1)
  if k == 2 {
    k, _ = n.Natural (t[1])
    if k < 10 { k *= 10 }
    x.float64 = x.float64 + float64(k) / 100
  }
  if s[0] == '-' { x.float64 = - x.float64 }
  return true
/*/
}

func (x *breal) String() string {
  if x.float64 == x.invalid {
    return str.New (x.wd)
  }
  s := r.String (x.float64)
  str.OffSpc (&s)
  str.Norm (&s, x.wd)
  str.Move (&s, false)
  return s
}

func (x *breal) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *breal) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *breal) Write (l, c uint) {
  r.Write (x.float64, l, c)
}

func (x *breal) Edit (l, c uint) {
  r.Wd (x.wd)
//  r.SetFormat (2)
  r.Edit (&x.float64, l, c)
}

func (x *breal) SetFont (f font.Font) {
  x.Font = f
}

func (x *breal) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *breal) RealVal() float64 {
  return x.float64
}

func (x *breal) SetRealVal (r float64) {
  if r >= x.invalid { ker.PrePanic() }
  x.float64 = r
}

func (x *breal) Zero() bool {
  return x.float64 == 0. // XXX epsilon
}

func (x *breal) Add (Y ...Adder) {
  n := len(Y)
  y := make([]*breal, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.float64 += y[i].float64
  }
}

func (x *breal) Sum (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Add (z)
}

func (x *breal) Sub (Y ...Adder) {
  n := len(Y)
  y := make([]*breal, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.float64 -= y[i].float64
  }
}

func (x *breal) Diff (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Sub (z)
}

func (x *breal) One() bool {
  return x.float64 == 1. // XXX epsilon
}

func (x *breal) Mul (Y ...Multiplier) {
  n := len(Y)
  y := make([]*breal, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.float64*= y[i].float64
  }
}

func (x *breal) Prod (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Mul (z)
}

func (x *breal) Quot (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.DivBy (z)
}

func (x *breal) Sqr() {
  q := x.float64 * x.float64
  x.float64 = q
}

func (x *breal) Power (n uint) {
  switch n {
  case 0:
    x.float64 = 1
  case 1:
    return
  default:
    q := x.float64
    for i := uint(1); i < n; i++ {
      q *= x.float64
    }
    x.float64 = q
  }
}

func (x *breal) Invertible() bool {
  return ! x.Zero()
}

func (x *breal) Invert() {
  e := new_(1)
  e.DivBy (x)
  x.Copy (e)
}

func (x *breal) DivBy (Y Multiplier) {
  y := x.imp(Y)
  if ! y.Invertible() { DivBy0Panic() }
  x.float64 /= y.float64
}
