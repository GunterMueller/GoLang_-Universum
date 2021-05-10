package gebiet

// (c) Christian Maurer   v. 210510 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/errh"
  "µU/text"
)
const (
  Undef = iota
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
  SciencFiction
  Märchen
  nGebiete
)
var tx = [nGebiete]string {"                  ",
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
var kx = [nGebiete]string {"  ",
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
  gebiet struct {
                int
                text.Text
                }
var
  cF, cB = col.LightWhite(), col.Blue()

func new_() Gebiet {
  x := new(gebiet)
  x.int = Undef
  x.Text = text.New (18)
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

func (x *gebiet) Write (l, c uint) {
  x.Text.Colours (cF, cB)
  x.Text.Defined (tx[x.int])
  x.Text.Write (l, c)
}

func (x *gebiet) Edit (l, c uint) {
  x.Write (l, c)
  errh.Hint ("p kl r h rr ir k rk ik k ä e s t sf ki m")
  loop:
  for {
    x.Text.Colours (cF, cB)
    x.Text.Edit (l, c)
    for i := 0; i < nGebiete; i++ {
      s := x.Text.String()
      if s[0:2] == kx[i] || str.Sub0 (s, tx[i]) {
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

func (x *gebiet) String() string {
  return x.Text.String()
}

func (x *gebiet) TeX() string {
  x.Text.Defined (tx[x.int])
  return x.Text.TeX()
}

func (x *gebiet) Defined (s string) bool {
  for i := 0; i < nGebiete; i++ {
	  if s == tx[i] {
      x.int = i
      x.Text.Defined (tx[i])
      return true
    }
  }
  return false
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
