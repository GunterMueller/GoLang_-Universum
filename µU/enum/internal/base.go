package internal

// (c) Christian Maurer   v. 210311 - license see µU.go

import (
  "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/errh"
  "µU/sel"
  "µU/font"
  "µU/pbox"
)
type
  base struct {
          typ,
            b uint8
            s [NFormats][]string
          num uint8
           wd [NFormats]uint
       cF, cB col.Colour
              Format
              }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_(t uint8, s [NFormats][]string) Base {
  x := new (base)
  x.typ, x.s = t, s
  x.num = uint8(len (s[Short]))
  m := [NFormats]uint { uint(0), uint(0) }
  for f := Short; f < NFormats; f++ {
    for i, t := range (s[f]) {
      s[f][i] = str.Lat1 (t)
      w := uint(len (s[f][i]))
      if m[f] < w { m[f] = w }
    }
    for i, _ := range (s[f]) { str.Norm (&s[f][i], m[f]) } // TODO gefährlich
    x.wd[f] = m[f]
  }
//  x.wd = m
  x.cF, x.cB = col.StartCols()
  x.Format = Short
  return x
}

func (x *base) imp (Y Any) *base {
  y, ok := Y.(*base)
  if ! ok {
    if x == Y { ker.Panic ("enum/internal/imp: x == Y") }
    TypeNotEqPanic (x, Y)
  }
  return y
}

func (x *base) GetFormat() Format {
  return x.Format
}

func (x *base) SetFormat (f Format) {
  x.Format = f
}

func (x *base) Typ() uint8 {
  return x.typ
}

func (x *base) Empty() bool {
  return x.b == 0
}

func (x *base) Clr() {
  x.b = 0
}

func (x *base) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.typ == y.typ && x.b == y.b
}

func (x *base) Copy (Y Any) {
  y := x.imp (Y)
  x.typ, x.b = y.typ, y.b
//  x.s = y.s // [NFormats][]string
  x.num = y.num
//  x.wd = y.wd // [NFormats]uint
  x.cF, x.cB = y.cF, y.cB
  x.Format = y.Format
}

func (x *base) Clone() Any {
  y := new_(x.typ, x.s)
  y.Copy (x)
  return x
}

func (x *base) Codelen() uint {
  return 1
}

func (x *base) Encode() Stream {
  s := make (Stream, 1)
  s[0] = byte(x.b)
  return s
}

func (x *base) Decode (s Stream) {
  if s[0] < x.num {
    x.b = uint8(s[0])
  } else {
    x.b = 0
  }
}

func (x *base) Less (Y Any) bool {
  return x.b < x.imp (Y).b
}

func (x *base) String() string {
  f := x.s[x.Format]
  s := f[x.b]
  str.OffSpc (&s)
  return s
}

func (x *base) Defined (s string) bool {
  for b := uint8(0); b < x.num; b++ {
    if p, ok := str.EquivSub (s, x.s[x.Format][b]); p == 0 && ok {
      x.b = b
      return true
    }
  }
  return false
}

func (x *base) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *base) Write (l, c uint) {
  bx.Colours (x.cF, x.cB)
  bx.Wd (x.wd[x.Format])
  s := x.String()
  str.Norm (&s, x.wd[x.Format])
  bx.Write (s, l, c)
}

func (x *base) selected (l, c uint) bool {
  if x.num == 0 { return false }
  if x.num == 1 { return true }
  x.Write (l, c)
  i := uint(0)
  h := x.num / 2
  if h < 5 { h = 5 }
  if h > x.num { h = x.num }
  errh.Hint (errh.ToSelect)
//  f, b:= x.cF, x.cB
//  f, b = col.Pink, col.Darkmagenta
  sel.Select (func (p, l, c uint, f, b col.Colour) {
                bx.Colours (f, b); bx.Write (x.s[x.Format][p], l, c)
              }, uint(x.num), uint(h), x.wd[x.Format], &i, l, c, x.cB, x.cF)
  errh.DelHint()
  if uint8(i) < x.num { x.b = uint8(i) }
  return uint8(i) < x.num
}

func (x *base) Edit (l, c uint) {
  x.Write (l, c)
  s := x.String()
  for {
    bx.Edit (&s, l, c)
    if C, _ := kbd.LastCommand(); C == kbd.Search {
      if x.selected (l, c) {
        break
      }
    }
    if x.Defined (s) {
      break
    } else {
      errh.Error0("geht nicht")
    }
  }
  x.Write (l, c)
}

func (x *base) SetFont (f font.Font) {
  pbx.SetFont (f)
}

func (x *base) Print (l, c uint) {
  pbx.Print (x.s[x.Format][x.b], l, c)
}

func (x *base) Ord() uint8 {
  return x.b
}

func (x *base) Num() uint8 {
  return x.num
}

func (x *base) Wd() uint {
  return x.wd[x.Format]
}

func (x *base) Set (n uint8) bool {
  if n < x.num {
    x.b = n
    return true
  }
  x.b = 0
  return false
}
