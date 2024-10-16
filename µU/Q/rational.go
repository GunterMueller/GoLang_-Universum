package Q

// (c) Christian Maurer   v. 241011 - license see µU.go

import (
  "math"
  . "µU/obj"
  "µU/ker"
  "µU/str"
  "µU/col"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/N"
  "µU/errh"
)
const (
  d = 9
  w = 1 + d + 1 + d // numerator and denominator with maximal 9 digits
)
type
  rational struct {
       num, denom uint
             geq0 bool
             f, b col.Colour
                  font.Font
                  }
var (
  bx = box.New()
  pbx = pbox.New()
)

func init() {
  bx.Wd (w) // sign, numerator, fraction bar, denominator
}

func new_() Rational {
  x := new(rational)
  x.f, x.b = col.StartCols()
  x.geq0 = true
  return x
}

func (x *rational) imp (Y any) *rational {
  y, ok := Y.(*rational)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *rational) Empty() bool {
  return x.denom == 0
}

func (x *rational) Numbers() (uint, uint) {
  return x.num, x.denom
}

func (x *rational) Clr() {
  x.num, x.denom = 0, 0
  x.geq0 = true
}

func (x *rational) Copy (Y any) {
  y := x.imp (Y)
  x.num, x.denom = y.num, y.denom
  x.geq0 = y.geq0
}

func (x *rational) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *rational) Eq (Y any) bool {
  y := x.imp (Y)
  if y.denom == 0 {
    return x.denom == 0
  } else if x.denom == 0 {
    return false
  }
  if x.num == 0 { return y.num == 0 }
  if y.num == 0 { return x.num == 0 }
  if x.geq0 != y.geq0 { return false }
  return uint (x.num) * uint (y.denom) == uint (y.num) * uint (x.denom)
}

func (x *rational) Less (Y any) bool {
  y := x.imp (Y)
  if y.denom == 0 {
    return x.denom == 0
  } else if x.denom == 0 {
    return false
  }
  if x.num == 0 && y.num == 0 { return false }
  if x.num == 0 { return y.geq0 }
  if y.num == 0 { return ! x.geq0 }
  if x.geq0 && ! y.geq0 { return false }
  if ! x.geq0 && y.geq0 { return true }
  p, q := uint (x.num) * uint (y.denom), uint (y.num) * uint (x.denom)
  if x.geq0 { return p < q } // y.geq0
  return p > q // ! x.geq0, ! y.geq0
}

func (x *rational) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *rational) Vals() (bool, uint, uint) {
  return x.geq0, x.num, x.denom
}

func (x *rational) Set1 (n int) bool {
  return x.Set (n, 1)
}

func (x *rational) Set (n, d int) bool {
  if d == 0 {
    x.num, x.denom = 1, 0
    return false
  }
  if n == 0 {
    x.num, x.denom = 0, 1
    return true
  }
  x.geq0 = n > 0
  if ! x.geq0 {
    n = -n
  }
  if d < 0 {
    d = -d
    x.geq0 = ! x.geq0
  }
  x.num, x.denom = uint(n), uint(d)
  x.reduce()
  return true
}

func (x *rational) SetNat (n, d uint, geq0 bool) {
  x.num, x.denom = n, d
  x.reduce()
  x.geq0 = geq0
  if x.num == 0 { x.geq0 = true }
}

func (x *rational) RealVal() float64 {
  if x.denom == 0 {
    return math.Inf(1)
  }
  r := float64 (x.num) / float64 (x.denom)
  if ! x.geq0 {
    r = - r
  }
  return r
}

func (x *rational) Defined (s string) bool {
  x.Clr()
  if str.Empty (s) {
    return true
  }
  str.Move (&s, true)
  x.geq0 = s[0] != '-'
  switch s[0] {
  case '+', '-':
    str.Rem (&s, 0, 1)
  }
  k := str.ProperLen (s)
  p, ok := str.Pos (s, '/')
  if ok {
    s1 := str.Part (s, p + 1, k - p - 1)
    if x.denom, ok = N.Natural (s1); ! ok {
      return false
    }
  } else {
    p = k
    x.denom = 1
  }
  s1 := str.Part (s, 0, p)
  if x.num, ok = N.Natural (s1); ! ok {
    return false
  }
  return true
}

func (x *rational) String() string {
  if x.denom == 0 {
    return str.Const ('-', w)
  }
  s := N.StringFmt (x.num, 2, false)
  str.OffSpc (&s)
  if ! x.geq0 {
    s = "-" + s
  }
  if x.denom == 1 {
    return s + "   "
  }
  s += "/"
  t := N.StringFmt (x.denom, 2, false)
  str.OffSpc (&t)
  if len(t) == 1 { t += " " }
  return s + t
}

func (x *rational) Wd() uint {
  return uint(len(x.String()))
}

func (x *rational) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *rational) Cols() (col.Colour, col.Colour) {
  return x.f, x.b
}

func (x *rational) Write (l, c uint) {
  bx.Colours (x.f, x.b)
  bx.Write (str.New(w), l, c)
  bx.Write (x.String(), l, c)
}

func (x *rational) Edit (l, c uint) {
  s := x.String()
  bx.Colours (x.f, x.b)
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
      break
    } else {
      errh.Error0("Format: Z = \"0\"|\"1\"|...|\"9\", [\"+\"|\"-\"] Z{Z}[/Z{Z}]; jeweils maximal 9 Ziffern")
    }
  }
  x.reduce()
  x.Write (l, c)
}

func (x *rational) SetFont (f font.Font) {
  x.Font = f
}

func (x *rational) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *rational) Codelen() uint {
  return 2 * C0 + 1
}

func (x *rational) Encode() Stream {
  b := make (Stream, x.Codelen())
  i, a := uint(0), C0
  copy (b[i:i+a], Encode (x.num))
  i += a
  copy (b[i:i+a], Encode (x.denom))
  i += a
  b[i] = 0
  if x.geq0 { b[i] = 1 }
  return b
}

func (x *rational) Decode (b Stream) {
  i, a := uint(0), C0
  x.num = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  x.denom = Decode (uint(0), b[i:i+a]).(uint)
  i += a
  x.geq0 = b[i] == 1
}

func (x *rational) Integer() bool {
  return x.denom == 1
}

func (x *rational) Zero() bool {
  if x.denom == 0 { return false }
  return x.num == 0
}

func (x *rational) One() bool {
  if x.denom == 0 { return false }
  return x.num == x.denom
}

func (x *rational) GeqNull() bool {
  if x.denom == 0 { return false }
  return x.geq0
}

func (x *rational) reduce() {
  if x.denom == 0 {
    x.num = 0
    x.geq0 = true
    return
  }
  if x.num == 0 {
    x.denom = 1
    x.geq0 = true
    return
  }
  g := N.Gcd (x.num, x.denom)
  x.num, x.denom = x.num / g, x.denom / g
  if x.num == 0 {
    x.geq0 = true
  }
}

func gcd (a, b uint) uint {
  if a < b { a, b = b, a }
  if b == 0 { return a }
  return gcd (a % b, b)
}

const
  max64 = uint(1<<64 - 1)

func (x *rational) add (Y Adder) {
  y := x.imp(Y)
  if y.num == 0 {
    return
  }
  if y.denom == 0 {
    x.Clr()
    return
  }
  a, b := uint (x.num) * uint (y.denom), uint (y.num) * uint (x.denom)
  var n uint
  if x.geq0 {
    if y.geq0 {
      if a <= max64 - b {
        n = a + b
        x.geq0 = true
      } else {
        x.Clr()
        return
      }
    } else {
      if a >= b {
        n = a - b
        x.geq0 = true
      } else {
        n = b - a
        x.geq0 = false
      }
    }
  } else { // ! x.geq0
    if y.geq0 {
      if a < b {
        n = b - a
        x.geq0 = true
      } else {
        n = a - b
        x.geq0 = false
      }
    } else {
      if a < max64 - b {
        n = a + b
        x.geq0 = false
      } else {
        x.Clr()
        return
      }
    }
  }
  d := uint (x.denom) * uint (y.denom)
  g := gcd (n, d)
  if g != 0 {
    n, d = n / g, d / g
  }
  x.num, x.denom = uint(n), uint(d)
  x.reduce()
}

func (x *rational) Add (Y ...Adder) {
  for i, _ := range Y {
    x.add (Y[i])
  }
}

func (x *rational) Sum (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Add (z)
}

func (x *rational) changeSign() {
  x.geq0 = ! x.geq0
  if x.num == 0 { x.geq0 = true }
}

func (x *rational) Sub (Y ...Adder) {
  for i, _ := range Y {
    y := x.imp(Y[i])
    yg := y.geq0
    y.geq0 = ! y.geq0
    x.Add (y)
    y.geq0 = yg
  }
}

func (x *rational) Diff (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Sub (z)
}

func (x *rational) mul (Y Multiplier) {
  y := x.imp(Y)
  if y.denom == 0 {
    x.Clr()
    return
  }
  if y.num == 0 {
    x.num, x.denom, x.geq0 = 0, 1, true
    return
  }
  n, d := uint(x.num) * uint(y.num), uint(x.denom) * uint(y.denom)
  g := gcd (n, d)
  if g != 0 {
    n, d = n / g, d / g
  }
  x.num, x.denom = uint(n), uint(d)
  x.geq0 = x.geq0 == y.geq0
  x.reduce()
}

func (x *rational) Mul (Y ...Multiplier) {
  for i, _ := range Y {
    x.mul (Y[i])
  }
}

func (x *rational) Prod (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Mul (z)
}

func (x *rational) Sqr() {
  x.Mul (x)
}

func (x *rational) Power (n uint) {
  c, d := x.num, x.denom
  switch n {
  case 0:
    x.num, x.denom = 1, 1
  case 1:
    return
  default: // n > 1
    for i := uint(1); i < n; i++ {
      c *= x.num
      d *= x.denom
    }
    x.num, x.denom = c, d
    x.reduce()
  }
}

func (x *rational) Invertible() bool {
  return x.num != 0 && x.denom != 0
}

func (x *rational) Invert() {
  if x.denom == 0 {
    x.num = 0
  } else {
    x.num, x.denom = x.denom, x.num
  }
}

func (x *rational) DivBy (Y Multiplier) {
  y := x.imp(Y)
  if ! y.Invertible() {
    x.Clr()
    return
  }
  inv := y.Clone().(*rational)
  inv.Invert()
  x.Mul (inv)
}

func (x *rational) Div (Y, Z Multiplier) {
  ker.PrePanic()
}

func (x *rational) Mod (Y, Z Multiplier) {
  ker.PrePanic()
}

func (x *rational) Quot (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.DivBy (z)
}
