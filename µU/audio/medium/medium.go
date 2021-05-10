package medium

// (c) Christian Maurer   v. 210509 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/errh"
  "µU/text"
)
const (
  Undef = iota
  SP
  LP
  CD
  DVD
  BR
  nMedien
)
var tx = [nMedien]string {"   ",
                          "SP ",
                          "LP ",
                          "CD ",
                          "DVD",
                          "BR "}
var
  cF, cB = col.LightWhite(), col.Blue()
type
  medium struct {
                int
                text.Text
                }

func new_() Medium {
  x := new(medium)
  x.int = Undef
  x.Text = text.New (3)
  x.Text.Defined (tx[x.int])
  return x
}

func (x *medium) imp (Y Any) *medium {
  y := Y.(*medium)
  return y
}

func (x *medium) Empty() bool {
  return x.int == Undef
}

func (x *medium) Clr() {
  x.int = Undef
}

func (x *medium) Eq (Y Any) bool {
  return x.int == x.imp(Y).int
}

func (x *medium) Less (Y Any) bool {
  return x.int < x.imp(Y).int
}

func (x *medium) Copy (Y Any) {
  x.int = x.imp(Y).int
}

func (x *medium) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *medium) Write (l, c uint) {
  x.Text.Defined (tx[x.int])
  x.Text.Colours (cF, cB)
  x.Text.Write (l, c)
}

func (x *medium) Edit (l, c uint) {
  x.Write (l, c)
  errh.Hint ("S(P L(P C(D D(VD B(R")
  loop:
  for {
    x.Text.Edit (l, c)
    for i := 0; i < nMedien; i++ {
      if str.Sub0 (x.Text.String(), tx[i]) {
        x.int = i
        x.Text.Defined (tx[i])
        x.Text.Write (l, c)
        break loop
      }
    }
    errh.Error0 ("kein Medium")
  }
  errh.DelHint()
}

func (x *medium) TeX() string {
  t := x.Text.String()
  str.OffSpc1 (&t)
  return t
}

func (x *medium) Codelen() uint {
  return 1
}

func (x* medium) Encode() Stream {
  s := make(Stream, x.Codelen())
  s[0] = byte (x.int)
  return s
}

func (x* medium) Decode (s Stream) {
  x.int = int(s[0])
  x.Text.Defined (tx[x.int])
}
