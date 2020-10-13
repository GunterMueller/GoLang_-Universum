package breal

// (c) Christian Maurer   v. 200913 - license see µU.go

// TODO: more than 2 digits after the decimal point

import (
//  "math"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/pbox"
  "µU/nat"
  "µU/real"
)
const
  m = 9 // 
type
  breal struct {
               float64
       pre, wd uint
       invalid float64
        cF, cB col.Colour
               font.Font
               }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_(d uint) Real {
  x := new(breal)
  if d == 0 { d = 1 }
  if d > m { d = m }
  x.pre, x.wd = d, 1 + d + 1 + 2
  x.invalid = exp(d)
  x.cF, x.cB = scr.StartCols()
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

func (x *breal) Encode() []byte {
  b := make ([]byte, x.Codelen())
  copy (b[:8], Encode (x.float64))
  return b
}

func (x *breal) Decode (b []byte) {
  x.float64 = float64(Decode (float64(0), b[:8]).(float64))
}

func (x *breal) Defined (s string) bool {
  if uint(len (s)) > x.wd { return false }
  str.OffBytes (&s, ' ')
//  n := x.wd / 2
//  P, L := make ([]uint, n), make ([]uint, n)
//  n = nat.NDigitSequences (s, &P, &L)
  n, t, p, l := nat.DigitSequences (s)
  if n == 0 || n > 2 || l[0] > x.pre {
    return false
  }
  if n == 2 {
    c := s[p[1] - 1]
    if l[1] > 2 || ! (c == '.' || c == ',') {
      return false
    }
  }
  n1, _ := nat.Natural (t[0])
  x.float64 = float64(n1)
  if n == 2 {
    n, _ = nat.Natural (t[1])
    if n < 10 { n *= 10 }
    x.float64 = x.float64 + float64(n) / 100
  }
  if s[0] == '-' { x.float64 = - x.float64 }
  return true
}

func (x *breal) String() string {
  if x.float64 == x.invalid {
    return str.New (x.wd)
  }
  s := real.String (x.float64)
  str.OffSpc (&s)
  str.Norm (&s, x.wd)
  str.Move (&s, false)
  return s
}

func (x *breal) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *breal) Write (l, c uint) {
  bx.Wd (x.wd)
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *breal) Edit (l, c uint) {
  x.Write (l, c)
  s := x.String()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error0Pos ("Eingabe falsch", l + 1, c)
    }
  }
  x.Write (l, c)
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

func (x *breal) SetReal (r float64) bool {
  if r < x.invalid {
    x.float64 = r
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
