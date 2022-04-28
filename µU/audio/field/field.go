package field

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/errh"
  "µU/text"
)
const (
  undef = iota
  alteMusik
  Barock
  Klassik
  Romantik
  neueMusik
  Beat_Rock
  Folklore
  Jazz
  Italien
  Frankreich
  Weihnachten
  Kinder
  nFields
)
var tx = [nFields]string {"           ",
                          "alte Musik ",
                          "Barock     ",
                          "Klassik    ",
                          "Romantik   ",
                          "neue Musik ",
                          "Beat / Rock",
                          "Folklore   ",
                          "Jazz       ",
                          "Italien    ",
                          "Frankreich ",
                          "Weihnachten",
                          "Kinder     "}
type
  field struct {
                byte
                text.Text
                }
var
  cF, cB = col.LightWhite(), col.Blue()

func new_() Field {
  x := new(field)
  x.byte = undef
  x.Text = text.New (11)
  x.Text.Defined (tx[undef])
  return x
}

func (x *field) imp (Y any) *field {
  y, ok := Y.(*field)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *field) Empty() bool {
  return x.byte == undef
}

func (x *field) Clr() {
  x.byte = undef
}

func (x *field) Eq (Y any) bool {
  return x.byte == x.imp(Y).byte
}

func (x *field) Less (Y any) bool {
  return x.byte < x.imp(Y).byte
}

func (x *field) Copy (Y any) {
  x.byte = x.imp(Y).byte
}

func (x *field) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *field) TeX() string {
  x.Text.Defined (tx[x.byte])
  return x.Text.TeX()
}

func (x *field) Write (l, c uint) {
  x.Text.Colours (cF, cB)
  x.Text.Defined (tx[x.byte])
  x.Text.Write (l, c)
}

func (x *field) Edit (l, c uint) {
  x.Write (l, c)
  errh.Hint ("A(lteMusik B(arock K(lassik R(omantik N(eueMusik Be(at J(azz F(olklore I(talien Ki(nder")
  loop:
  for {
    x.Text.Colours (cF, cB)
    x.Text.Edit (l, c)
    for i := byte(0); i < nFields; i++ {
      if str.Sub0 (x.Text.String(), tx[i]) {
        x.byte = i
        x.Text.Defined (tx[i])
        x.Write (l, c)
        break loop
      }
    }
    errh.Error0 ("falsche Eingabe")
  }
  errh.DelHint()
}

func (x *field) Codelen() uint {
  return 1
}

func (x* field) Encode() Stream {
  s := make(Stream, 1)
  s[0] = x.byte
  return s
}

func (x* field) Decode (s Stream) {
  x.byte = s[0]
  x.Text.Defined (tx[x.byte])
}
