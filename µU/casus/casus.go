package casus

// (c) Christian Maurer   v. 200910 - license see µU.go

import (
//  "µU/ker"
  . "µU/obj"
  "µU/str"
//  "µU/kbd"
  "µU/col"
//  "µU/scr"
  "µU/box"
  "µU/errh"
//  "µU/sel"
  "µU/font"
  "µU/pbox"
)
const (
  Undef = uint8(iota)
  Nom
  Gen
  Dat
  Akk
  Abl
  N
)
type
  casus struct {
             s [NFormats][]string
            wd [NFormats]uint
               uint8
        cF, cB col.Colour
               Format
               }
var (
  bx = box.New()
  pbx = pbox.New()
  wd = [NFormats]uint{4, 9}
  text [NFormats][N]string
)

func init() {
  text[Short] = [N]string{ "    ",
                           "Nom.",
                           "Gen.",
                           "Dat.",
                           "Akk.",
                           "Abl." }
  text[Long] = [N]string{ "         ",
                          "Nominativ",
                          "Genitiv  ",
                          "Dativ    ",
                          "Akkusativ",
                          "Ablativ  " }
}

func new_() Casus {
  x := new (casus)
  x.uint8 = Undef
  x.Format = Short
  return x
}

func (x *casus) imp(Y Any) *casus {
  y, ok := Y.(*casus)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *casus) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.uint8 == y.uint8
}

func (x *casus) Copy (Y Any) {
  y := x.imp(Y)
  x.uint8 = y.uint8
  x.cF, x.cB = y.cF, y.cB
  x.Format = y.Format
}

func (x *casus) Clone() Any {
  y := New().(*casus)
  y.uint8 = x.uint8
  return y
}

func (x *casus) Less (Y Any) bool {
  return x.uint8 < x.imp (Y).uint8
}

func (x *casus) Set (c uint8) bool {
  if c < N {
    x.uint8 = c
    return true
  }
  return false
}

func (x *casus) GetFormat() Format {
  return x.Format
}

func (x *casus) SetFormat (f Format) {
  x.Format = f
}

func (x *casus) Empty() bool {
  return x.uint8 == Undef
}

func (x *casus) Clr() {
  x.uint8 = Undef
}

func (x *casus) Codelen() uint {
  return 1
}

func (x *casus) Encode() Stream {
  s := make ([]byte, 1)
  s[0] = x.uint8
  return s
}

func (x *casus) Decode (s Stream) {
  x.uint8 = Undef
  if s[0] < N {
    x.uint8 = s[0]
  }
}

func (x *casus) String() string {
  return text[x.Format][x.uint8]
}

func (x *casus) Defined (s string) bool {
  for c := Undef; c < N; c++ {
    if p, ok:= str.EquivSub (s, text[x.Format][c]); p == 0 && ok {
      x.uint8 = c
      return true
    }
  }
  return false
}

func (x *casus) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *casus) Write (l, c uint) {
  bx.Colours (x.cF, x.cB)
  bx.Wd (wd[x.Format])
  s := x.String()
  bx.Write (s, l, c)
}

/*/
func (x *casus) selected (l, c uint) bool {
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
/*/

func (x *casus) Edit (l, c uint) {
  x.Write (l, c)
  s := x.String()
  for {
    bx.Edit (&s, l, c)
/*/
    if C, _:= kbd.LastCommand(); C == kbd.Search {
      if x.selected (l, c) {
        break
      }
    }
/*/
    if x.Defined (s) {
      break
    } else {
      errh.Error0("geht nicht")
    }
  }
  x.Write (l, c)
}

func (x *casus) SetFont (f font.Font) {
  pbx.SetFont (f)
}

func (x *casus) Print (l, c uint) {
  pbx.Print (text[x.Format][x.uint8], l, c)
}
