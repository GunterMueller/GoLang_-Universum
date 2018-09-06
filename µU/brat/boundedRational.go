package brat

// (c) Christian Maurer   v. 180902 - license see µU.go

import (
  "math"
  . "µU/obj"
  . "µU/add"
  . "µU/mul"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/nat"
  "µU/errh"
)
const
  max = 1e9 // numerator and denominator with at most 9 digits
type
  rational struct {
       num, denom uint
             geq0 bool
           cF, cB col.Colour
                  font.Font
                  }
var (
  reciprocal = new_().(*rational)
  bx = box.New()
  pbx = pbox.New()
)

func init() {
  bx.Wd (1 + 9 + 1 + 9) // sign, numerator, fraction bar, denominator
}

func new_() Rational {
  x := new(rational)
  x.cF, x.cB = scr.StartCols()
  if x.cF.IsWhite() && x.cB.IsBlack() { x.cF = col.LightWhite() } // Firlefanz
  x.geq0 = true
  return x
}

func (x *rational) imp (Y Any) *rational {
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

func (x *rational) Copy (Y Any) {
  y := x.imp (Y)
  x.num, x.denom = y.num, y.denom
  x.geq0 = y.geq0
}

func (x *rational) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *rational) Eq (Y Any) bool {
  y := x.imp (Y)
  if y.denom == 0 {
    return x.denom == 0
  } else if x.denom == 0 {
    return false
  }
  if x.num == 0 { return y.num == 0 }
  if y.num == 0 { return x.num == 0 }
  if x.geq0 != y.geq0 { return false }
  return uint64 (x.num) * uint64 (y.denom) == uint64 (y.num) * uint64 (x.denom)
}

func (x *rational) Less (Y Any) bool {
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
  p, q := uint64 (x.num) * uint64 (y.denom), uint64 (y.num) * uint64 (x.denom)
  if x.geq0 { return p < q } // y.geq0
  return p > q // ! x.geq0, ! y.geq0
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
  switch s[0] { case '+', '-':
    str.Rem (&s, 0, 1)
  }
  n := str.ProperLen (s)
  p, ok := str.Pos (s, '/')
  if ok {
    s1 := str.Part (s, p + 1, n - p - 1)
    if x.denom, ok = nat.Natural (s1); ok {
      if x.denom >= max {
        return false
      }
    } else {
      return false
    }
  } else {
    p = n
    x.denom = 1
  }
  s1 := str.Part (s, 0, p)
  if x.num, ok = nat.Natural (s1); ok {
    if x.num >= max {
      return false
    }
  } else {
    return false
  }
  return true
}

func (x *rational) String() string {
  if x.denom == 0 {
    return "Zähler/Nenner"
  }
  s := nat.StringFmt (x.num, 9, false)
  str.OffSpc (&s)
  if ! x.geq0 {
    s = "-" + s
  }
  if x.denom == 1 {
    return s
  }
  s += "/"
  t := nat.StringFmt (x.denom, 9, false)
  str.OffSpc (&t)
  return s + t
}

func (x *rational) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *rational) Write (l, c uint) {
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *rational) Edit (l, c uint) {
  s := x.String()
  bx.Colours (x.cF, x.cB)
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
  return 2 * 4 + 1
}

func (x *rational) Encode() []byte {
  b := make ([]byte, x.Codelen())
  copy (b[:4], Encode (x.num))
  copy (b[4:8], Encode (x.denom))
  b[9] = 0
  if x.geq0 { b[9] = 1 }
  return b
}

func (x *rational) Decode (b []byte) {
  x.num = Decode (uint(0), b[:4]).(uint)
  x.denom = Decode (uint(0), b[4:8]).(uint)
  x.geq0 = b[9] == 1
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
    return
  }
  if x.num == 0 {
    x.denom = 1
    return
  }
  g := nat.Gcd (x.num, x.denom)
  x.num, x.denom = x.num / g, x.denom / g
}

func gcd (a, b uint64) uint64 {
  if a < b { a, b = b, a }
  if b == 0 { return a }
  return gcd (a % b, b)
}

const
  max64 = uint64(1<<64 - 1)

func (x *rational) add (Y Adder) {
  y := x.imp(Y)
  if y.num == 0 {
    return
  }
  if y.denom == 0 {
    x.Clr()
    return
  }
  a, b := uint64 (x.num) * uint64 (y.denom), uint64 (y.num) * uint64 (x.denom)
  var n uint64
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
  d := uint64 (x.denom) * uint64 (y.denom)
  g := gcd (n, d)
  n, d = n / g, d / g
  if n > uint64(max) || d > uint64(max) {
    x.Clr()
    return
  }
  x.num, x.denom = uint(n), uint(d)
  x.reduce()
}

func (x *rational) Add (Y ...Adder) {
  for i, _ := range Y {
    x.add (Y[i])
  }
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
  if y.num == y.denom {
    return
  }
  n, d := uint64(x.num) * uint64(y.num), uint64(x.denom) * uint64(y.denom)
  g := gcd (n, d)
  n, d = n / g, d / g
  if n > uint64(max) || d > uint64(max) {
    x.Clr()
    return
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

func (x *rational) Sqr() {
  x.Mul (x)
}

func (x *rational) Invert() {
  if x.denom == 0 {
    x.num = 0
  } else {
    x.num, x.denom = x.denom, x.num
  }
}

func (x *rational) Div (Y, Z Multiplier) {
  y, z := x.imp (Y), x.imp (Z)
  if y.denom == 0 || z.num == 0 || z.denom == 0 {
    x.Clr()
    return
  }
  inv := z.Clone().(*rational)
  inv.Invert()
  x.Mul (y, inv)
}

func (x *rational) DivBy (Y Multiplier) {
  x.Div (x, Y)
}

/*
type Operation byte; const (ADD = iota; SUB; MUL; DIV ) // TODO interface in µU/obj

func (x *rational) Operate (Y ...Rational, op Operation) { // TODO Rational -> ...
  switch op {
  case ADD:
    x.Sum (Y...)
  case SUB:
    x.Difference (Y)
  case MUL:
    x.Product (Y)
  case DIV:
    x.Quotient (Y)
  }
}
*/
