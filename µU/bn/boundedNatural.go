package bn

// (c) Christian Maurer   v. 220924 - license see µU.go

import (
  "math"
  "µU/ker"
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/pbox"
)
const
  invalid = uint(1<<64 - 1)
type
  natural struct {
                 uint
              wd uint
            f, b col.Colour
                 font.Font
                 }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_(n uint) Natural {
  x := new(natural)
  x.uint = invalid
  x.wd = n
  x.f, x.b = col.StartCols()
  x.Font = font.Roman
  return x
}

func (x *natural) imp (Y any) *natural {
  y, ok := Y.(*natural)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *natural) Width() uint {
  return x.wd
}

func (x *natural) Empty() bool {
  return x.uint == invalid
}

func (x *natural) Clr() {
  x.uint = invalid
}

func (x *natural) Copy (Y any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.wd = y.wd
  x.f, x.b = y.f, y.b
}

func (x *natural) Clone() any {
  y := new_(x.wd)
  y.Copy (x)
  return y
}

func (x *natural) Eq (Y any) bool {
  return x.uint == x.imp(Y).uint
}

// non-empty Less than Empty
func (x *natural) Less (Y any) bool {
  return x.uint < x.imp(Y).uint
}

func (x *natural) Leq (Y any) bool {
  return x.uint <= x.imp(Y).uint
}

func (x *natural) Codelen() uint {
  return C0
}

func (x *natural) Encode() Stream {
  return Encode (x.uint)
}

func (x *natural) Decode (s Stream) {
  x.uint = Decode (uint(0), s).(uint)
}

func (x *natural) Defined (s string) bool {
  if s == "" {
    return false
  }
  str.Move (&s, true)
  n := str.ProperLen (s)
  x.uint = 0
  for i := 0; i < int(n); i++ {
    if '0' <= s[i] && s[i] <= '9' {
      b := s[i] - '0'
      if x.uint < (invalid - uint(b)) / 10 {
        x.uint = 10 * x.uint + uint(b)
      } else {
        return false
      }
    } else {
      return false
    }
  }
  return true
}

func st (n uint) string {
  if n < 10 {
    return string(n + '0')
  }
  return st (n / 10) + st (n % 10)
}

func (n *natural) String() string {
  if n.uint == invalid {
    return str.New (M)
  }
  return st (n.uint)
}

func (n *natural) Colours (f, b col.Colour) {
  n.f.Copy (f)
  n.b.Copy (b)
}

func (n *natural) Cols() (col.Colour, col.Colour) {
  return n.f, n.b
}

func (n *natural) ActFont() font.Font {
  return n.Font
}

func (n *natural) SetFont (f font.Font) {
  n.Font = f
}

func (n *natural) Write (l, c uint) {
//  bx.SetFont (n.Font)
  bx.Wd (n.wd)
  bx.Colours (n.f, n.b)
  bx.Write (str.New (n.wd), l, c)
  bx.Write (n.String(), l, c)
}

func (n *natural) Edit (l, c uint) {
  n.EditGr (8 * int(c), 16 * int(l))
}

func (n *natural) WriteGr (x, y int) {
//  bx.SetFont (n.Font)
  bx.Wd (n.wd)
  bx.Colours (n.f, n.b)
  bx.WriteGr (str.New(n.wd), x, y)
  bx.WriteGr (n.String(), x, y)
}

func (n *natural) EditGr (x, y int) {
  n.WriteGr (x, y)
  s := n.String()
  for {
    bx.EditGr (&s, x, y)
    if str.Empty (s) {
      n.Clr()
      return
    }
    if n.Defined (s) {
      break
    } else {
      errh.Error0 ("keine Zahl")
    }
  }
  n.WriteGr (x, y)
}

func (x *natural) Print (l, c uint) {
  pbx.SetFont (x.Font)
  pbx.Print (x.String(), l, c)
}

func (x *natural) Val() uint {
  return uint(x.uint)
}

func (x *natural) SetVal (n uint) {
  x.uint = n
}

func (x *natural) Dual() string {
  const M = 64
  s := ""
  n := x.uint
  for i := M - 1; i > 0; i-- {
    c := "0"
    if n % 2 == 1 { c = "1" }
    s = c + s
    n /= 2
  }
  return s
}

func (x *natural) Decimal (s string) {
  n := uint(0)
  for i := 0; i < len(s); i++ {
    k := uint(0)
    if s[i] == '1' { k = 1 }
    n = 2 * n + k
  }
  x.uint = n
}

func (x *natural) Zero() bool {
  return x.uint == 0
}

func (x *natural) Add (Y ...Adder) {
  n := len(Y)
  y := make([]*natural, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.uint += y[i].uint
  }
}

func (x *natural) Sum (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Add (z)
}

func (x *natural) Sub (Y ...Adder) {
  n := len(Y)
  y := make([]*natural, n)
  for i:= 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.uint -= y[i].uint
  }
}

func (x *natural) Diff (Y, Z Adder) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Sub (z)
}

func (x *natural) One() bool {
  return x.uint == 1
}

func (x *natural) Mul (Y ...Multiplier) {
  n := len(Y)
  y := make([]*natural, n)
  for i := 0; i < n; i++ {
    y[i] = x.imp(Y[i])
    x.uint*= y[i].uint
  }
}

func (x *natural) Prod (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.Mul (z)
}

func (x *natural) Div (Y, Z Multiplier) {
  y, z := x.imp(Y), x.imp(Z)
  x.Copy (y)
  x.DivBy (z)
}

func (x *natural) Mod (Y, Z Multiplier) {
  q, z := x.imp(Y).uint, x.imp(Z).uint
  for q > z {
    q -= z
  }
  x.uint = q
}

func (x *natural) Quot (Y, Z Multiplier) {
  if Z.One() {
    x.Copy (Y)
  } else {
    ker.PrePanic()
  }
}

func (x *natural) Sqr() {
  q := x.uint * x.uint
  x.uint = q
}

func (x *natural) Power (n uint) {
  switch n {
  case 0:
    x.uint = 1
  case 1:
    return
  default:
    q := x.uint
    for i := uint(1); i < n; i++ {
      q *= x.uint
    }
    x.uint = q
  }
}

func (x *natural) Invertible() bool {
  return x.One()
}

func (x *natural) Invert() {
  if ! x.Invertible() {
    ker.PrePanic()
  }
}

func (x *natural) DivBy (Y Multiplier) {
  y := x.imp(Y)
  q := float64(x.uint) / float64(y.uint)
  if q == math.Trunc (q) {
    x.uint = uint(q)
  } else {
    x.uint = invalid
  }
/*/
  y := x.imp(Y)
  if ! y.Invertible() { DivBy0Panic() }
  x.uint /= y.uint
/*/
}
