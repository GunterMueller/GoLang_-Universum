package br

// (c) Christian Maurer   v. 210311 - license see µU.go

import (
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
       pre, wd uint
       invalid float64
        cF, cB col.Colour
               font.Font
               }
var
  pbx = pbox.New()

func new_(d uint) Real {
  x := new(breal)
  x.float64 = r.NaN()
  if d == 0 { d = 1 }
  if d > m { d = m }
  x.pre, x.wd = d, 1 + d + 1 + 2
  x.invalid = exp(d)
  x.cF, x.cB = col.StartCols()
  return x
}

func (x *breal) imp (Y Any) *breal {
  y, ok := Y.(*breal)
  if ! ok || x.wd != y.wd { TypeNotEqPanic (x, Y) }
  return y
}

func exp (n uint) float64 {
  if n == 0 { return 1 }
  return 10 * exp (n - 1)
}

func (x *breal) Empty() bool {
  return x.float64 == x.invalid
}

func (x *breal) Clr() {
  x.float64 = x.invalid
}

func (x *breal) Copy (Y Any) {
  x.float64 = x.imp (Y).float64
}

func (x *breal) Clone() Any {
  y := new_(x.pre)
  y.Copy(x)
  return y
}

func (x *breal) Eq (Y Any) bool {
  return x.float64 == x.imp (Y).float64
}

func (x *breal) Less (Y Any) bool {
  y := x.imp (Y)
  if x.float64 == x.invalid || y.float64 == y.invalid {
    return false
  }
  return x.float64 < y.float64
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
  x.float64 = r.NaN()
  return false
/*/
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
  x.cF, x.cB = f, b
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

func (x *breal) Float64() float64 {
  return x.float64
}

func (x *breal) SetFloat64 (f float64) bool {
  if f < x.invalid {
    x.float64 = f
    return true
  }
  return false
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

func (x *breal) Sub (Y ...Adder) {
  n := len(Y)
  y := make([]*breal, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.float64 -= y[i].float64
  }
}

func (x *breal) One() bool {
  return x.float64 == 1. // XXX epsilon
}

func (x *breal) Mul (Y ...Any) {
  n := len(Y)
  y := make([]*breal, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.float64*= y[i].float64
  }
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

func (x *breal) DivBy (Y Any) {
  y := x.imp(Y)
  if Zero(y) { DivBy0Panic() }
  x.float64 /= x.imp(Y).float64
}
