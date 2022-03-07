package field

// (c) Christian Maurer   v. 220228 - license see µU.go

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
  lx = []string {"                      ", // W l   W r  P    G r   G l
                 "Ägypten             äg", // 6 h
                 "Antike Biographie   ab", // 6 v
                 "Griechischer Text   gt", // 5 h
                 "Lateinischer Text   lt", // 5 h
                 "Rom-Roman           rr", // 5 v        5
                 "Rom-Krimi           rk", //            5 6
                 "Historischer Roman  hr", // 4 hv 
//                                         // 3 h
                 "Italien-Roman       ir", // 3 v
                 "Klassische Literaturkl", //       8 h
                 "Neuere Literatur    nl", //       8 v
                 "Theaterstück(e)     th", //       7 hv
                 "Science Fiction     sf", //       7 v
                 "Italien-Krimi       ik", //       6 hv  4
//                                         //       5 hv  3
                 "Horror              ho", //       5 v
                 "Krimi               KR", //       4 hv
                 "Jugendbuch          ju", //       3 hv
                 "Englisch            en", // 2 
                 "Französisch         fr", // 2 
                 "Italienisch         it", // 2 
                 "Griechisch          gr", // 2 
                 "Lateinisch          la", // 2 
                 "Sachbuch            sb", // 
                 "Pflanzen            pf", // 
                }
/*/
  lx = []string {"                      ", // W l   W r  P    G r   G l
                 "Ägypten             äg", // 6 h
                 "Antike Biographie   ab", // 6 v
                 "Griechischer Text   gt", // 5 h
                 "Lateinischer Text   lt", // 5 h
                 "Rom-Roman           rr", // 5 v        5
                 "Rom-Krimi           rk", //            5 6
                 "Historischer Roman  hr", // 4 hv 
//                                         // 3 h
                 "Italien-Roman       ir", // 3 v
                 "Klassische Literaturkl", //                 4 hv
                 "Neuere Literatur    nl", //       8 hv
                 "Theaterstück(e)     th", //       7 hv
                 "Science Fiction     sf", //       7 v
                 "Italien-Krimi       ik", //       6 hv  4
//                                         //       5 hv  3
                 "Horror              ho", //       5 v
                 "Krimi               KR", //       4 hv
                 "Jugendbuch          ju", //       3 hv
                 "Englisch            en", // 2 
                 "Französisch         fr", // 2 
                 "Italienisch         it", // 2 
                 "Griechisch          gr", // 2 
                 "Lateinisch          la", // 2 
                 "Sachbuch            sb", // 
                 "Pflanzen            pf", // 
                }
/*/
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

func (x *field) imp (Y Any) *field {
  y := Y.(*field)
  return y
}

func (x *field) Empty() bool {
  return x.Text.Empty()
}

func (x *field) Clr() {
  x.Text.Clr()
}

func (x *field) Eq (Y Any) bool {
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

func (x *field) Less (Y Any) bool {
  return x.pos() < x.imp(Y).pos()
}

func (x *field) Copy (Y Any) {
  x.Text.Copy (x.imp(Y).Text)
}

func (x *field) Clone() Any {
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
