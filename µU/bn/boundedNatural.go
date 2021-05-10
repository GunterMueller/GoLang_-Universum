package bn

// (c) Christian Maurer   v. 210509 - license see µU.go

import (
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
  if n == 0 || n > M { ker.PrePanic() }
  x := new(natural)
  x.uint = invalid
  x.wd = n
  x.f, x.b = col.StartCols()
  return x
}

func (x *natural) imp (Y Any) *natural {
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

func (x *natural) Copy (Y Any) {
  y := x.imp (Y)
  x.uint = y.uint
  x.wd = y.wd
  x.f, x.b = y.f, y.b
}

func (x *natural) Clone() Any {
  y := new_(x.wd)
  y.Copy (x)
  return y
}

func (x *natural) Eq (Y Any) bool {
  return x.uint == x.imp(Y).uint
}

// non-empty Less than Empty
func (x *natural) Less (Y Any) bool {
  return x.uint < x.imp(Y).uint
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

func (x *natural) String() string {
  if x.uint == invalid {
    return str.New (M)
  }
  return st (x.uint)
}

func (n *natural) Colours (f, b col.Colour) {
  n.f, n.b = f, b
}

func (n *natural) Write (l, c uint) {
  bx.Wd (n.wd)
  bx.Colours (n.f, n.b)
  bx.Write (str.New(n.wd), l, c)
  bx.Write (n.String(), l, c)
}

func (x *natural) Edit (l, c uint) {
  x.EditGr (8 * int(c), 16 * int(l))
}

func (n *natural) WriteGr (x, y int) {
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

func (x *natural) SetFont (f font.Font) {
  x.Font = f
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
