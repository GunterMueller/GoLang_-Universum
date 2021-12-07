package field

// (c) Christian Maurer   v. 210515 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/errh"
  "µU/text"
)
const (
  undef = iota
  Prosa
  Klassik
  Roman
  HistoRoman
  RomRoman
  ItalienRoman
  Krimi
  RomKrimi
  ItalienKrimi
  Kunst
  Ägypten
  Etrurien
  Sachbuch
  Theater
  Kinderbuch
  ScienceFiction
  Märchen
  nFields
)
var tx = [nFields]string {"                  ",
                          "Prosa             ",
                          "Klassik           ",
                          "Roman             ",
                          "Historischer Roman",
                          "Rom-Roman         ",
                          "Italien-Roman     ",
                          "Krimi             ",
                          "Rom-Krimi         ",
                          "Italien-Krimi     ",
                          "Kunst             ",
                          "Ägypten           ",
                          "Etrurien          ",
                          "Sachbuch          ",
                          "Theaterstück(e)   ",
                          "Science Fiction   ",
                          "Kinderbuch        ",
                          "Märchenbuch       "}
var kx = [nFields]string {"  ",
                          "p ",
                          "kl",
                          "r ",
                          "h ",
                          "rr",
                          "ir",
                          "k ",
                          "rk",
                          "ik",
                          "ku",
                          "ä ",
                          "e ",
                          "s ",
                          "t ",
                          "sf",
                          "ki",
                          "m "}
type
  field struct {
                int "ordinal of the constant"
                text.Text
                }
var
  cF, cB = col.LightWhite(), col.Blue()

func new_() Field {
  x := new(field)
  x.int = undef
  x.Text = text.New (18)
  x.Text.Defined (tx[undef])
  return x
}

func (x *field) imp (Y Any) *field {
  y := Y.(*field)
  return y
}

func (x *field) Empty() bool {
  return x.int == undef
}

func (x *field) Clr() {
  x.int = undef
}

func (x *field) Eq (Y Any) bool {
  return x.int == x.imp(Y).int
}

func (x *field) Less (Y Any) bool {
  return x.int < x.imp(Y).int
}

func (x *field) Copy (Y Any) {
  x.int = x.imp(Y).int
}

func (x *field) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *field) Write (l, c uint) {
  x.Text.Colours (cF, cB)
  x.Text.Defined (tx[x.int])
  x.Text.Write (l, c)
}

func (x *field) Edit (l, c uint) {
  x.Write (l, c)
  errh.Hint ("p kl r h rr ir k rk ik k ä e s t sf ki m")
  loop:
  for {
    x.Text.Colours (cF, cB)
    x.Text.Edit (l, c)
    for i := 0; i < nFields; i++ {
      s := x.Text.String()
      if s[0:2] == kx[i] || str.Sub0 (s, tx[i]) {
        x.int = i
        x.Text.Defined (tx[i])
        x.Write (l, c)
        break loop
      }
    }
    errh.Error0 ("falsche Eingabe")
  }
  errh.DelHint()
}

func (x *field) String() string {
  x.Text.Defined (tx[x.int])
  t := x.Text.String()
  str.OffSpc1 (&t)
  return t
}

func (x *field) TeX() string {
  x.Text.Defined (tx[x.int])
  return x.Text.TeX()
}

func (x *field) Defined (s string) bool {
  for i := 0; i < nFields; i++ {
	  if s == tx[i] {
      x.int = i
      x.Text.Defined (tx[i])
      return true
    }
  }
  return false
}

func (x *field) Codelen() uint {
  return 1
}

func (x* field) Encode() Stream {
  s := make(Stream, 1)
  s[0] = byte(x.int)
  return s
}

func (x* field) Decode (s Stream) {
  x.int = int(s[0])
  x.Text.Defined (tx[x.int])
}
