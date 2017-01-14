package lint

// (c) murus.org  v. 170107 - license see murus.go
//
// >>> lots of things TODO, particularly new packages lnat and lreal (and lrat (?)

import (
//  "math"
  . "math/big"
//  "strconv"
  . "murus/obj"
//  "murus/str"
  "murus/col"; "murus/scr"; "murus/box"; "murus/errh"
  "murus/font"; "murus/prt" // ; "murus/pbox";
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
//  px = pbox.New()
)

func newInt(n int) LongInteger {
  x:= new(longInteger)
  x.n = NewInt(int64(n))
  x.cF, x.cB = col.StartCols()
  return x
}

func new64 (n int64) LongInteger {
  x:= new(longInteger)
  x.n = NewInt(n)
  x.cF, x.cB = col.StartCols()
  return x
}

/* deprecated, has to be moved into a package lnat !

func newNat (n uint) LongInteger {
  x:= new(longInteger)
  x.n = NewInt(int64(n))
  x.cF, x.cB = col.StartCols()
  return x
}
*/

func (x *longInteger) imp (Y Any) *Int {
  x, ok:= Y.(*longInteger)
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

func (x *longInteger) Eq (Y Any) bool {
  return x.n.Cmp (x.imp (Y)) == 0
}

func (x *longInteger) Copy (Y Any) {
  x.n.Set (x.imp (Y))
}

func (x *longInteger) Clone() Any {
  y:= New (0)
  y.Copy (x)
  return y
}

func (x *longInteger) Less (Y Any) bool {
  return x.n.Cmp (x.imp (Y)) == -1
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

func (x *longInteger) Encode() []byte {
  return x.n.Bytes()
}

func (x *longInteger) Decode (b []byte) {
  x.n.SetBytes (b)
}

func (x *longInteger) SetVal (n uint) bool {
  x.n.SetInt64(int64(n))
  return true
}

func (x *longInteger) SetInt (n int) {
  x.n.SetInt64 (int64(n))
}

func (x *longInteger) SetInt64 (n int64) {
  x.n.SetInt64 (int64(n))
}

func (x *longInteger) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *longInteger) Write (l, c uint) {
  s:= x.String()
  c0:= c
  scr.Colours (x.cF, x.cB)
  for n:= 0; n < len (s); n++ {
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
  s:= x.String()
  w:= uint(len (s))
  N:= scr.NColumns()
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
  s:= x.String()
  c0:= c
  for i:= 0; i < len (s); i++ {
    prt.Print1 (s[i], l, c, x.Font)
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

func (x *longInteger) Len() uint {
  return uint(len (x.String()))
}

func (x *longInteger) Odd() bool {
  return x.n.Bit (0) == 1
}

func (x *longInteger) Val() int {
  n:= x.n.Int64()
  if Codelen(int(0)) == 4 {
    if n < 1<<32 - 1 && n > -1<<32 {
      return int(n)
    }
  } else { // 8
    if n < int64(1<<63 - 1) && n > int64(-1<<63) { // XXX
      return int(n)
    }
  }
  return 0
}

func (x *longInteger) ValInt() int {
  n:= x.n.Int64()
  if n < 1<<32 - 1 {
    return int(n)
  }
  return 0
}

func (x *longInteger) Val64() int64 {
  if x.Less(max) {
    return x.n.Int64()
  }
  return 0 // ? TODO
}

/* deprecated, has to be moved into a package lreal !

func newReal (r float64) LongInteger {
  x:= New(0)
  x.SetReal(r)
  return x
}
 
func (x *longInteger) ValReal() float64 {
  r, err:= strconv.ParseFloat (x.n.String(), 64)
  if err != nil { return math.NaN() }
  if x.n.Sign() < 0 { r = -r }
  return r
}

func (x *longInteger) SetReal (r float64) {
  i, _:= math.Modf (r + 0.5)
  s:= strconv.FormatFloat (i, 'f', -1, 64)
  if p, ok:= str.Pos (s, '.'); ok {
    s = str.Part (s, 0, p)
  }
  if _, ok:= x.n.SetString (s, 10); ! ok {
    x.n.SetInt64 (0)
  }
}
*/

func (x *longInteger) Defined (s string) bool {
  _, ok:= x.n.SetString (s, 10)
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
  a:= uint(0)
  for _, c:= range (tmp.n.String()) {
    a += uint(c)
  }
  return a
}

func (x *longInteger) Null() bool {
  return x.n.Sign() == 0
}

func (x *longInteger) Add (Y ...Adder) Adder {
  for _, y:= range Y {
    x.n.Add (x.n, y.(*longInteger).n)
  }
  return x.Clone().(Adder)
}

func (x *longInteger) Inc() {
  x.n.Add (x.n, one.n)
}

func (x *longInteger) Sub (Y ...Adder) Adder {
  for _, y:= range Y {
    x.n.Sub (x.n, y.(*longInteger).n)
  }
  return x.Clone().(Adder)
}

func (x *longInteger) Dec() {
  x.n.Sub (x.n, one.n)
}

func (x *longInteger) One() bool {
  return x.Eq (one)
}

func (x *longInteger) Mul (Y ...Multiplier) Multiplier {
  for _, y:= range Y {
    x.n.Mul (x.n, y.(*longInteger).n)
  }
  return x.Clone().(Multiplier)
}

func (x *longInteger) Sqr() {
  x.n.Mul (x.n, x.n)
}

func (x *longInteger) Div (Y, Z Multiplier) {
  zn:= x.imp (Z)
  if zn.Cmp (zero.n) == 0 { DivBy0Panic() }
  x.n.Quo (x.imp (Y), zn)
}

func (x *longInteger) DivBy (Y Multiplier) {
  x.Div (x, Y)
}

func (x *longInteger) Mod (Y LongInteger) {
  x.n.Mod (x.n, x.imp (Y))
}

func (x *longInteger) MulMod (Y, M LongInteger) {
  x.n.Mul (x.n, x.imp (Y)) // not efficient
  x.n.Mod (x.n, x.imp (M))
}

func (x *longInteger) Div2 (Y, R LongInteger) {
  yn:= x.imp (Y)
  r, ok:= R.(*longInteger)
  if ! ok { TypeNotEqPanic (r, R) }
  if yn.Cmp (zero.n) == 0 { DivBy0Panic() }
  _, r.n = x.n.QuoRem (x.n, yn, one.n)
}

func (x *longInteger) Gcd (Y LongInteger) {
  yn:= x.imp (Y)
  if x.n.Sign() <= 0 || yn.Sign() <= 0 {
    return
  }
  x.n.GCD (tmp.n, tmp1.n, x.n, yn)
}

func (x *longInteger) Lcm (Y LongInteger) {
  yn:= x.imp (Y)
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
  e:= k % 2 != 0
  nn, ii:= new64(int64(n)).(*longInteger), new64(1).(*longInteger)
  for i:= uint(1); i <= k; i++ {
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

func (x *longInteger) Bitlen() uint {
  return uint(x.n.BitLen())
}

func (x *longInteger) Bit (i int) uint {
  return x.n.Bit (i)
}

func (x *longInteger) SetBit (i int, b bool) {
  u:= uint(0)
  if b { u++ }
  x.n.SetBit (x.n, i, u)
}

func init() {
  bx.Wd (64)
}
