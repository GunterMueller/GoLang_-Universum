package li

// (c) Christian Maurer   v. 230217 - license see µU.go

import (
  . "math/big"
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/prt"
)
type (
  longInteger struct {
                   n *Int
              cF, cB col.Colour
                     font.Font
//               nan bool
                     }
)
var (
  zero, one = New(0).(*longInteger), New(1).(*longInteger)
  max32, max = New(1 << 32 - 1).(*longInteger), New(1 << 63 - 1).(*longInteger)
  tmp, tmp1 = New(0).(*longInteger), New(0).(*longInteger)
  bx = box.New()
)

func init() {
  bx.Wd (64)
}

func new_(n int) LongInteger {
  x := new(longInteger)
  x.n = NewInt (int64(n))
  x.cF, x.cB = col.StartCols()
  return x
}

func (x *longInteger) imp (Y any) *Int {
  x, ok := Y.(*longInteger)
  if ! ok { TypeNotEqPanic (x, Y) }
  return x.n
}

func (x *longInteger) Empty() bool {
  return x.n.Cmp (zero.n) == 0
//  return x.nan
}

func (x *longInteger) Clr() {
  x.n.SetInt64 (0)
//  x.nan = true
}

func (x *longInteger) Eq (Y any) bool {
  return x.n.Cmp (x.imp (Y)) == 0
}

func (x *longInteger) Copy (Y any) {
  x.n.Set (x.imp (Y))
}

func (x *longInteger) Clone() any {
  y := New (0)
  y.Copy (x)
  return y
}

func (x *longInteger) Less (Y any) bool {
  return x.n.Cmp (x.imp (Y)) == -1
}

func (x *longInteger) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *longInteger) Geq0() bool {
  return x.n.Sign() > 0
}

func (x *longInteger) ChSign() {
  x.n.Neg (x.n)
}

func (x *longInteger) Codelen() uint {
  return uint(len (x.n.Bytes()))
}

func (x *longInteger) Encode() Stream {
  return x.n.Bytes()
}

func (x *longInteger) Decode (s Stream) {
  x.n.SetBytes (s)
}

func (x *longInteger) SetVal (n uint) {
  x.n.SetInt64(int64(n))
}

func (x *longInteger) Cols() (col.Colour, col.Colour) {
  return x.cF, x.cB
}

func (x *longInteger) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *longInteger) Write (l, c uint) {
  s := x.String()
  c0 := c
  scr.Colours (x.cF, x.cB)
  for n := 0; n < len (s); n++ {
    scr.Write1 (s[n], l, c)
    if c + 1 < scr.NColumns() {
      c ++
    } else if l + 2 < scr.NLines() {
      l ++
      c = c0
    } else {
      break
    }
  }
}

func (x *longInteger) Edit (l, c uint) {
  s := x.String()
  w := uint(len (s))
  N := scr.NColumns()
  if c >= N - w {
    x.Write (l, c)
    errh.Error0("zu wenig Platz auf dem Bildschirm") // TODO
    return
  }
  bx.Wd (N - 1 - c)
  bx.Edit (&s, l, c)
  for {
    if x.Defined (s) {
      break
    } else {
      errh.Error0("keine Zahl")
    }
  }
  x.Write (l, c)
}

func (x *longInteger) SetFont (f font.Font) {
  x.Font = f
}

func (x *longInteger) Print (l, c uint) {
  s := x.String()
  c0 := c
  for i := 0; i < len (s); i++ {
    prt.SetFont (x.Font)
    prt.Print1 (s[i], l, c)
    if c + 1 < prt.NColumns() {
      c ++
    } else if l + 2 < prt.NLines() {
      l ++
      c = c0
    } else {
      break
    }
  }
  prt.GoPrint()
}

func (x *longInteger) Odd() bool {
  return x.n.Bit (0) == 1
}

/*/
// Abs sets z to |x| (the absolute value of x) and returns z.
func (z *Int) Abs(x *Int) *Int {
	z.Set(x)
	z.neg = false
	return z
}
/*/

func (x *longInteger) Abs0() {
  x.n.Abs (x.n)
}

func (x *longInteger) Abs() LongInteger {
  x.n.Abs (x.n)
  return x
}

func (x *longInteger) Val() int {
  n := x.n.Int64()
  if n <= int64(1<<63 - 1) && n >= int64(-1<<63) {
    return int(n)
  }
  return 0
}

func (x *longInteger) Defined (s string) bool {
  _, ok := x.n.SetString (s, 10)
  if ok {
    return true
  }
  x.n.SetInt64 (0) // nan
  return false
}

func (x *longInteger) String() string {
  return x.n.String()
}

func (x *longInteger) SumDigits() uint {
  tmp.n.Abs (x.n)
  a := uint(0)
  for _, c := range (tmp.n.String()) {
    a += uint(c)
  }
  return a
}
// func sumDigits (x LongInteger) uint {
//   return uint(len(x.(*longInteger).n.String()))
// }

func (x *longInteger) Zero() bool {
  return x.n.Sign() == 0
}

func (x *longInteger) Add (Y ...Adder) {
  for _, y := range Y {
    x.n.Add (x.n, y.(*longInteger).n)
  }
}

func (x *longInteger) Sum (Y, Z Adder) {
  y, z := Y.(Adder), Z.(Adder)
  x.Copy (y)
  x.Add (z)
}

func (x *longInteger) Inc() {
  x.n.Add (x.n, one.n)
}

func (x *longInteger) Sub (Y ...Adder) {
  for _, y := range Y {
    x.n.Sub (x.n, y.(*longInteger).n)
  }
}

func (x *longInteger) Diff (Y, Z Adder) {
  y, z := Y.(Adder), Z.(Adder)
  x.Copy (y)
  x.Sub (z)
}

func (x *longInteger) Dec() {
  x.n.Sub (x.n, one.n)
}

func (x *longInteger) One() bool {
  return x.Eq (one)
}

func (x *longInteger) Mul (Y ...Multiplier) {
  for _, y := range Y {
    x.n.Mul (x.n, y.(*longInteger).n)
  }
}

func (x *longInteger) Prod (Y, Z Multiplier) {
  y, z := Y.(Multiplier), Z.(Multiplier)
  x.Copy (y)
  x.Mul (z)
}

func (x *longInteger) Quot (Y, Z Multiplier) {
  y, z := Y.(Multiplier), Z.(Multiplier)
  x.Copy (y)
  x.DivBy (z)
}

func (x *longInteger) Sqr() {
  x.n.Mul (x.n, x.n)
}

func (x *longInteger) Power (n uint) {
  switch n {
  case 0:
    x.n.SetInt64 (1)
  case 1:
    return
  default:
    q := x.n
    for i := uint(1); i < n; i++ {
      q.Mul (q, x.n)
    }
    x.n = q
  }
}

func (x *longInteger) Invertible() bool {
  return x.n.Cmp (zero.n) != 0
}

func (x *longInteger) Invert() {
  e := new_(1)
  e.DivBy (x)
  x.Copy (e)
}

func (x *longInteger) div (Y, Z Multiplier) {
  z := Z.(*longInteger)
  if ! z.Invertible() { DivBy0Panic() }
  x.n.Quo (x.imp (Y), z.imp (Z))
}

func (x *longInteger) DivBy (Y Multiplier) {
  x.div (x, Y)
}

func (x *longInteger) Mod (Y LongInteger) {
  x.n.Mod (x.n, x.imp (Y))
}

func (x *longInteger) MulMod (Y, M LongInteger) {
  x.n.Mul (x.n, x.imp (Y)) // not efficient
  x.n.Mod (x.n, x.imp (M))
}

func (x *longInteger) Div2 (Y, R LongInteger) {
  yn := x.imp (Y)
  r, ok := R.(*longInteger)
  if ! ok { TypeNotEqPanic (r, R) }
  if yn.Cmp (zero.n) == 0 { DivBy0Panic() }
  _, r.n = x.n.QuoRem (x.n, yn, one.n)
}

func (x *longInteger) Gcd (Y LongInteger) {
  yn := x.imp (Y)
  if x.n.Sign() <= 0 || yn.Sign() <= 0 {
    return
  }
  x.n.GCD (tmp.n, tmp1.n, x.n, yn)
}

func (x *longInteger) Lcm (Y LongInteger) {
  yn := x.imp (Y)
  if x.n.Sign() <= 0 || yn.Sign() <= 0 {
    return
  }
  x.n.Mul (x.n, yn)
  tmp.n.Set (yn)
  tmp.Gcd (x)
  x.n.Quo (x.n, tmp.n)
}

func (x *longInteger) Pow (Y LongInteger) {
  x.n.Exp (x.n, x.imp (Y), nil)
}

func (x *longInteger) PowMod (Y, M LongInteger) {
  x.n.Exp (x.n, x.imp (Y), x.imp (M))
}

func (x *longInteger) Fak (n uint) {
  x.n.MulRange (1, int64(n))
}

func (x *longInteger) ProbabylPrime (n int) bool {
  return x.n.ProbablyPrime (n)
}

func (x *longInteger) Binom (n, k uint) {
  x.n.Binomial (int64(n), int64(k))
}

func (x *longInteger) LowFak (n, k uint) {
  if n < k {
    x.n.SetInt64 (0)
    return
  }
  if k == 0 {
    x.n.SetInt64 (1)
    return
  }
  x.n.MulRange (int64(n - k + 1), int64(n))
}

func (x *longInteger) Stirl2 (n, k uint) {
  x.n.SetInt64 (0)
  if n < k {
    return
  }
  if k == 0 {
    if n == 0 {
      x.n.SetInt64 (1)
    }
    return
  }
  tmp.n.SetInt64 (1)
  e := k % 2 != 0
  nn, ii := new_(int(n)).(*longInteger), new_(1).(*longInteger)
  for i := uint(1); i <= k; i++ {
    tmp.n.Mul (tmp.n, tmp1.n.SetInt64 (int64(k - i + 1)))
    tmp.n.Div (tmp.n, ii.n)
    tmp1.n.Mul (tmp1.n.Exp (ii.n, nn.n, nil), tmp.n)
    if e {
      x.n.Add (x.n, tmp1.n)
    } else {
      x.n.Sub (x.n, tmp1.n)
    }
    e = ! e
    ii.Inc()
  }
  x.n.Div (x.n, tmp.n.MulRange (1, int64(k)))
}
