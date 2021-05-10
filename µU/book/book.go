package book

// (c) Christian Maurer   v. 210509 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/text"
  "µU/bn"
  "µU/book/gebiet"
)
const (
  len0 = 30
  len1 = 63
)
type
  ordnung int; const (
  nachGebiet = iota
  nachAutor
  nOrdnungen
)
type
  book struct {
               gebiet.Gebiet
         autor,
       koautor text.Text
               bn.Natural
         titel,
       fundort text.Text
               }
var (
  aktuelleOrdnung ordnung
)

func (x *book) imp (Y Any) *book {
  y, ok := Y.(*book)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Book {
  x := new (book)
  x.Gebiet = gebiet.New()
  x.autor = text.New (len0)
  x.koautor = text.New (len0)
  x.Natural = bn.New (2)
  x.titel = text.New (len1)
  x.fundort = text.New (len0)
  x.Natural.Colours (col.LightWhite(), col.DarkCyan())
  x.autor.Colours (col.Yellow(), col.Red())
  x.koautor.Colours (col.Yellow(), col.Red())
  x.titel.Colours (col.LightWhite(), col.DarkGreen())
  x.fundort.Colours (col.LightWhite(), col.DarkGray())
  return x
}

func (x *book) Empty() bool {
  return x.Gebiet.Empty() &&
         x.autor.Empty() && x.koautor.Empty() &&
         x.Natural.Empty() &&
         x.titel.Empty() &&
         x.fundort.Empty()
}

func (x *book) Clr() {
  x.Gebiet.Clr()
  x.autor.Clr()
  x.koautor.Clr()
  x.Natural.Clr()
  x.titel.Clr()
  x.fundort.Clr()
}

func (x *book) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.Gebiet.Eq (y.Gebiet) &&
         x.autor.Eq (y.autor) && x.koautor.Eq (y.koautor) &&
         x.Natural.Eq (y.Natural) &&
         x.titel.Eq (y.titel) &&
         x.fundort.Eq (y.fundort)
}

func (x *book) Copy (Y Any) {
  y := x.imp(Y)
  x.Gebiet.Copy (y.Gebiet)
  x.autor.Copy (y.autor)
  x.koautor.Copy (y.koautor)
  x.Natural.Copy (y.Natural)
  x.titel.Copy (y.titel)
  x.fundort.Copy (y.fundort)
}

func (x *book) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *book) Less (Y Any) bool {
  y := x.imp(Y)
  switch aktuelleOrdnung {
  case nachGebiet:
    if x.Gebiet.Eq (y.Gebiet) {
      if x.autor.Eq (y.autor) {
//        return x.titel.Less (y.titel)
        return x.Natural.Less (y.Natural)
      }
      return x.autor.Less (y.autor)
    }
    return x.Gebiet.Less (y.Gebiet)
  case nachAutor:
    if x.autor.Eq (y.autor) {
      if x.Gebiet.Eq (y.Gebiet) {
//        return x.titel.Less (y.titel)
        return x.Natural.Less (y.Natural)
      }
      return x.Gebiet.Less (y.Gebiet)
    }
    return x.autor.Less (y.autor)
  }
  return false
}

const (
  lg = 1; cg =  7
  la = 3; ca =  7
  lk = 3; ck = 49
  ln = 5; cn =  7
  lt = 5; ct = 16
  lf = 7; cf = 49
)
/*/
          1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

Gebiet __________________

 Autor ______________________________    Koautor ______________________________

    Nr __ Titel _______________________________________________________________

                                         Fundort ______________________________
/*/
func writeMask() {
  scr.Colours (col.LightGray(), col.Black())
  scr.Write ("Gebiet",     1,  0)
  scr.Write ("Autor",      3,  1)
  scr.Write ("Koautor",    3, 41)
  scr.Write ("Nr",         5,  4)
  scr.Write ("Titel",      5, 10)
  scr.Write ("Fundort",    7, 41)
}

func (x *book) Write (l, c uint) {
  writeMask()
  x.Gebiet.Write (lg, cg)
  x.autor.Write (la, ca)
  x.koautor.Write (lk, ck)
  x.Natural.Write (ln, cn)
  x.titel.Write (lt, ct)
  x.fundort.Write (lf, cf)
}

func (x *book) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  loop:
  for {
    switch i {
    case 0:
      x.Gebiet.Edit (lg, cg)
    case 1:
      x.autor.Edit (la, ca)
      for i := 0; i < len(a); i++ {
        if x.autor.Sub0 (autor[i]) {
          x.autor.Copy (autor[i])
          x.autor.Write (la, ca)
          break
        }
      }
    case 2:
      x.koautor.Edit (lk, ck)
      if ! x.koautor.Empty() {
        for i := 0; i < len(a); i++ {
          if x.koautor.Sub0 (autor[i]) {
            x.koautor.Copy (autor[i])
            x.koautor.Write (lk, ck)
            break
          }
        }
      }
    case 3:
      x.Natural.Edit (ln, cn)
    case 4:
      x.titel.Edit (lt, ct)
    case 5:
      x.fundort.Edit (lf, cf)
    }
    switch k, _ := kbd.LastCommand(); k {
    case kbd.Esc:
      break loop
    case kbd.Enter, kbd.Down:
      if i < 5 {
        i++
      } else {
        break loop
      }
    case kbd.Back, kbd.Up:
			if i > 0 {
        i--
      }
    }
  }
}

func (x *book) String() string {
  s := x.Gebiet.String() + " "
  s += x.autor.String() + " "
  s += x.koautor.String() + " "
  s += x.Natural.String() + " "
  s += x.titel.String() + " "
  s += x.fundort.String()
  return s
}

var letztesGebiet = gebiet.New()

func (x *book) TeX() string {
  s := ""
  if ! x.Gebiet.Eq (letztesGebiet) {
    letztesGebiet.Copy (x.Gebiet)
    s += "\\medskip{\\bf " + x.Gebiet.TeX() + "}\\smallskip\n"
  }
  s += x.autor.TeX()
  if ! x.koautor.Empty() {
    s += "/" + x.koautor.TeX()
  }
  s += ": " + x.titel.TeX()
  if ! x.fundort.Empty() {
    s += " (" + x.fundort.TeX() + ")"
  }
  return s
}

func (x *book) Defined (s string) bool {
  if ! x.Gebiet.(Stringer).Defined (s[0:12]) {
    return false
  }
  if ! x.autor.(Stringer).Defined (s[13:43]) {
    return false
  }
  if ! x.koautor.(Stringer).Defined (s[44:74]) {
    return false
  }
  if ! x.Natural.(Stringer).Defined (s[75:77]) {
    return false
  }
  if ! x.titel.(Stringer).Defined (s[78:141]) {
    return false
  }
  if ! x.fundort.(Stringer).Defined (s[142:172]) {
    return false
  }
  return true
}

func (x *book) Codelen() uint {
  return x.Gebiet.Codelen() +
       2 * len0 +
       x.Natural.Codelen() +
       len1 + len0
}

func (x *book) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.Gebiet.Codelen()
  copy (s[i:i+a], x.Gebiet.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.autor.Encode())
  i += a
  copy (s[i:i+a], x.koautor.Encode())
  i += a
  a = x.Natural.Codelen()
  copy (s[i:i+a], x.Natural.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.titel.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.fundort.Encode())
  return s
}

func (x *book) Decode (s Stream) {
  i, a := uint(0), x.Gebiet.Codelen()
  x.Gebiet.Decode (s[i:i+a])
  i += a
  a = len0
  x.autor.Decode (s[i:i+a])
  i += a
  x.koautor.Decode (s[i:i+a])
  i += a
  a = x.Natural.Codelen()
  x.Natural.Decode (s[i:i+a])
  i += a
  a = len1
  x.titel.Decode (s[i:i+a])
  i += a
  a = len0
  x.fundort.Decode (s[i:i+a])
}

func (x *book) Rotate() {
  aktuelleOrdnung = (aktuelleOrdnung + 1) % nOrdnungen
}

func (x *book) Index() Func {
  return func (a Any) Any {
    return a
  }
}
