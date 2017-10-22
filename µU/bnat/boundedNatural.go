package bnat

// (c) Christian Maurer   v. 170810 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/scr"
  "µU/str"
  "µU/box"
  "µU/errh"
  "µU/font"
  "µU/pbox"
  "µU/nat"
)
type
  natural struct {
               n,
         invalid uint64
              wd uint
            f, b col.Colour
                 font.Font
                 }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_(w uint) Natural {
  x := new(natural)
  if w == 0 { w = 1 } else if w > 19 { w = 19 }
  x.wd = w
  x.invalid = uint64(10)
  for i := x.wd; i > 1; i-- {
    x.invalid *= 10
  }
  x.n = x.invalid
  x.f, x.b = scr.StartCols()
  return x
}

func (x *natural) imp (Y Any) *natural {
  y, ok := Y.(*natural)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *natural) Empty() bool {
  return x.n == x.invalid
}

func (x *natural) Clr() {
  x.n = x.invalid
}

func (x *natural) Copy (Y Any) {
  y := x.imp (Y)
  x.wd = y.wd
  x.n = y.n
}

func (x *natural) Clone() Any {
  y := new_(x.wd)
  y.Copy (x)
  return y
}

func (x *natural) Eq (Y Any) bool {
  return x.n == x.imp(Y).n
}

// non-empty Leess than Empty
func (x *natural) Less (Y Any) bool {
  return x.n < x.imp(Y).n
}

func (x *natural) Codelen() uint {
  return 8
}

func (x *natural) Encode() []byte {
  if x.n == x.invalid { println("bnat.Enc", x.n) }
  return Encode (uint64(x.n))
}

func (x *natural) Decode (bs []byte) {
  x.n = Decode (uint64(0), bs).(uint64)
  if x.n == x.invalid { println("bnat.Dec", x.n) }
}

func (x *natural) Defined (s string) bool {
  if n, ok:= nat.Natural (s); ok {
    x.n = uint64(n)
    return true
  }
  return false
}

func (x *natural) String() string {
  if x.n == x.invalid {
    return str.Clr (x.wd)
  }
  return nat.StringFmt (uint(x.n), x.wd, false)
}

func (x *natural) Colours (f, b col.Colour) {
  x.f, x.b = f, b
}

func (x *natural) Write (l, c uint) {
  x.WriteGr (8 * int(c), 16 * int(l))
}

func (x *natural) Edit (l, c uint) {
  x.EditGr (8 * int(c), 16 * int(l))
}

func (n *natural) WriteGr (x, y int) {
  bx.Wd (n.wd)
  bx.Colours (n.f, n.b)
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
  return uint(x.n)
}

func (x *natural) SetVal (n uint) bool {
  n1 := uint64(n)
  if n1 < x.invalid {
    x.n = n1
    return true
  }
  x.n = n1 % x.invalid
  return false
}

func (x *natural) Width() uint {
  return x.wd
}
