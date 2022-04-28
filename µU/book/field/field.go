package field

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/str"
  "µU/errh"
  "µU/text"
)
const
  l = 20
var (
  lx = []string {"                      ",
                 "Ägypten             äg", // w l 6 h
                 "Antike Biographie   ab", // w l 6 v
                 "Griechischer Text   gt", // w l 5 h
                 "Lateinischer Text   lt", // w l 5 h
                 "Rom-Roman           rr", // w l 5 v    p 5
                 "Rom-Krimi           rk", //            p 5 6
                 "Neuere Literatur    nl", // w r 8 h
                 "Italien-Roman       ir", // w r 8 v
//                                         // w r 7 h
                 "Theaterstück(e)     th", // w r 7 hv
                 "Italien-Krimi       ik", // w r 6 hv   p 4
//                                         // w r 5 hv   p 3
                 "Horror              ho", // w r 5 v
                 "Krimi               kr", // w r 4 hv
                 "Jugendbuch          ju", // w r 3 hv
                 "Englisch            en", // w r 2 h
                 "Französisch         fr", // w r 2 h
                 "Italienisch         it", // w r 2 h
                 "Griechisch          gr", // w r 2 h
                 "Lateinisch          la", // w r 2 h
                 "Sachbuch            sb", // w r 2 h
                 "Pflanzen            pf", // w r 2 v
                 "Science Fiction     sf", // w r 2 v
                 "Klassische Literaturkl", // g r 4 hv
                 "Historischer Roman  hr", // g 5 3 hv 
//                                         // g r 2 h
                }
  n = len(lx)
  tx = make([]string, n)
  cx = make([]Stream, n)
)
type
  field struct {
                text.Text
               }
var
  cF, cB = col.LightWhite(), col.Blue()

func init() {
  for i := 0; i < n; i++ {
    lx[i] = str.Lat1 (lx[i])
  }
  tx = make([]string, n)
  for i := 0; i < n; i++ {
    tx[i] = lx[i][0:l]
    cx[i] = Stream(lx[i][l:l+2])
  }
}

func new_() Field {
  x := new(field)
  x.Text = text.New (l)
  x.Text.Defined (tx[0])
  return x
}

func (x *field) imp (Y any) *field {
  y := Y.(*field)
  return y
}

func (x *field) Empty() bool {
  return x.Text.Empty()
}

func (x *field) Clr() {
  x.Text.Clr()
}

func (x *field) Eq (Y any) bool {
  return x.Text.Eq (x.imp(Y).Text)
}

func (x* field) pos() int {
  if x.Text.String()[0] == 'Ä' {
    return 1
  }
  for i := 0; i < n; i++ {
    if x.Text.String() == tx[i] {
      return i
    }
  }
  return l
}

func (x *field) Less (Y any) bool {
  return x.pos() < x.imp(Y).pos()
}

func (x *field) Copy (Y any) {
  x.Text.Copy (x.imp(Y).Text)
}

func (x *field) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *field) Write (l, c uint) {
  x.Text.Colours (cF, cB)
  x.Text.Write (l, c)
}

func (x *field) Edit (l, c uint) {
  x.Write (l, c)
  h := ""
  for i := 1; i < n; i++ {
    h += string(cx[i]) + " "
  }
  errh.Hint (h)
  loop:
  for {
    x.Text.Colours (cF, cB)
    x.Text.Edit (l, c)
    s := x.Text.String()
    for i := 0; i < n; i++ {
      if s[0:2] == string(cx[i]) || str.Sub0 (s, tx[i]) {
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
  return x.Text.String()
}

func (x *field) TeX() string {
  return x.Text.TeX()
}

func (x *field) Defined (s string) bool {
  s = str.Lat1 (s)
  for i := 0; i < n; i++ {
    t := tx[i]
    str.OffSpc1 (&t)
	  if s == t || s == string(cx[i]) {
      x.Text.Defined (tx[i])
      return true
    }
  }
  return false
}

func (x *field) Codelen() uint {
  return 2
}

func (x* field) Encode() Stream {
  s := x.Text.String()
  for i := 0; i < n; i++ {
    if s == tx[i] {
      return cx[i]
    }
  }
  return Stream("xx")
}

func (x* field) Decode (s Stream) {
  x.Text.Clr()
  for i := 0; i < n; i++ {
    if string(s) == string(cx[i]) {
      x.Text.Defined (tx[i])
    }
  }
}
