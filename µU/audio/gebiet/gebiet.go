package gebiet

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
  Konzert
  Oper
  Beat
  Jazz
  Folklore
  Italien
  Kinder
  nGebiete
)
var tx = [nGebiete]string {"        ",
                           "Konzert ",
                           "Oper    ",
                           "Beat    ",
                           "Jazz    ",
                           "Folklore",
                           "Italien ",
                           "Kinder  "}
type
  gebiet struct {
                int
                text.Text
                }
var
  cF, cB = col.LightWhite(), col.Blue()

func new_() Gebiet {
  x := new(gebiet)
  x.int = Undef
  x.Text = text.New (8)
  x.Text.Defined (tx[Undef])
  return x
}

func (x *gebiet) imp (Y Any) *gebiet {
  y := Y.(*gebiet)
  return y
}

func (x *gebiet) Empty() bool {
  return x.int == Undef
}

func (x *gebiet) Clr() {
  x.int = Undef
}

func (x *gebiet) Eq (Y Any) bool {
  return x.int == x.imp(Y).int
}

func (x *gebiet) Less (Y Any) bool {
  return x.int < x.imp(Y).int
}

func (x *gebiet) Copy (Y Any) {
  x.int = x.imp(Y).int
}

func (x *gebiet) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *gebiet) TeX() string {
  x.Text.Defined (tx[x.int])
  return x.Text.TeX()
}

func (x *gebiet) Write (l, c uint) {
  x.Text.Colours (cF, cB)
  x.Text.Defined (tx[x.int])
  x.Text.Write (l, c)
}

func (x *gebiet) Edit (l, c uint) {
  x.Write (l, c)
  errh.Hint ("K(onzert O(per B(eat J(azz F(olklore I(talien Ki(nder")
  loop:
  for {
    x.Text.Colours (cF, cB)
    x.Text.Edit (l, c)
    for i := 0; i < nGebiete; i++ {
      if str.Sub0 (x.Text.String(), tx[i]) {
        x.Text.Defined (tx[i])
        x.int = i
        x.Write (l, c)
        break loop
      }
    }
    errh.Error0 ("kein Gebiet")
  }
  errh.DelHint()
}

func (x *gebiet) Codelen() uint {
  return 1
}

func (x* gebiet) Encode() Stream {
  s := make(Stream, 1)
  s[0] = byte (x.int)
  return s
}

func (x* gebiet) Decode (s Stream) {
  x.int = int(s[0])
  x.Text.Defined (tx[x.int])
}
