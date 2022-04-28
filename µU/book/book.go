package book

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/text"
  "µU/str"
  "µU/bn"
  "µU/book/field"
)
const (
  lena = 30
  lent = 63
  sep = ';'
  seps = ";"
  nosep = " darf kein " + seps + " enthalten"
)
const (
  fieldOrder = iota
  authorOrder
)
type
  book struct {
               field.Field
        author,
      coauthor text.Text
               bn.Natural
         title,
      cupboard,
         floor text.Text
              }
var
  actOrder = fieldOrder

func (x *book) imp (Y any) *book {
  y, ok := Y.(*book)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Book {
  x := new (book)
  x.Field = field.New()
  x.author = text.New (lena)
  x.coauthor = text.New (lena)
  x.Natural = bn.New (2)
  x.title = text.New (lent)
  x.cupboard = text.New (3)
  x.floor = text.New (3)
  x.Natural.Colours (col.LightWhite(), col.DarkCyan())
  x.author.Colours (col.Yellow(), col.Red())
  x.coauthor.Colours (col.Yellow(), col.Red())
  x.title.Colours (col.LightWhite(), col.DarkGreen())
  x.cupboard.Colours (col.LightWhite(), col.DarkGray())
  x.floor.Colours (col.LightWhite(), col.DarkGray())
  return x
}

func (x *book) Empty() bool {
  return x.Field.Empty() &&
         x.author.Empty() && x.coauthor.Empty() &&
         x.Natural.Empty() &&
         x.title.Empty() &&
         x.cupboard.Empty() &&
         x.floor.Empty()
}

func (x *book) Clr() {
  x.Field.Clr()
  x.author.Clr()
  x.coauthor.Clr()
  x.Natural.Clr()
  x.title.Clr()
  x.cupboard.Clr()
  x.floor.Clr()
}

func (x *book) Eq (Y any) bool {
  y := x.imp(Y)
  return x.Field.Eq (y.Field) &&
         x.author.Eq (y.author) && x.coauthor.Eq (y.coauthor) &&
         x.Natural.Eq (y.Natural) &&
         x.title.Eq (y.title) &&
         x.cupboard.Eq (y.cupboard) &&
         x.floor.Eq (y.floor)
}

func (x *book) Copy (Y any) {
  y := x.imp(Y)
  x.Field.Copy (y.Field)
  x.author.Copy (y.author)
  x.coauthor.Copy (y.coauthor)
  x.Natural.Copy (y.Natural)
  x.title.Copy (y.title)
  x.cupboard.Copy (y.cupboard)
  x.floor.Copy (y.floor)
}

func (x *book) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *book) Less (Y any) bool {
  y := x.imp(Y)
  switch actOrder {
  case fieldOrder:
    if x.Field.Eq (y.Field) {
      if x.author.Eq (y.author) {
        return x.Natural.Less (y.Natural)
      }
      return x.author.Less (y.author)
    }
    return x.Field.Less (y.Field)
  case authorOrder:
    if x.author.Eq (y.author) {
      if x.Field.Eq (y.Field) {
        return x.Natural.Less (y.Natural)
      }
      return x.Field.Less (y.Field)
    }
    return x.author.Less (y.author)
  }
  panic ("")
}

func (x *book) String() string {
  s := x.Field.String()
  str.OffSpc1 (&s)
  s += seps
  t := x.author.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.coauthor.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.Natural.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.title.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.cupboard.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.floor.String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *book) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != 7 { return false }
  if ! x.Field.Defined (ss[0]) { return false }
  if ! x.author.Defined (ss[1]) { return false }
  if ! x.coauthor.Defined (ss[2]) { return false }
  if ! x.Natural.Defined (ss[3]) { return false }
  if ! x.title.Defined (ss[4]) { return false }
  if ! x.cupboard.Defined (ss[5]) { return false }
  if ! x.floor.Defined (ss[6]) { return false }
  return true
}

func (x *book) Sub (Y any) bool {
  y := x.imp(Y)
  s := false
  if ! x.Field.Empty() {
    s = s || x.Field.Eq (y.Field)
  }
  if ! x.author.Empty() {
    s = s || x.author.Sub (y.author)
  }
  if ! x.title.Empty() {
    s = s || x.title.Sub (y.title)
  }
  if ! x.cupboard.Empty() {
    s = s || x.cupboard.Sub (y.cupboard)
  }
  return s
}

const (
  lg = 1; cg =  7
  la = 3; ca =  7
  lk = 3; ck = 49
  ln = 5; cn =  7
  lt = 5; ct = 16
  lc = 7; cc = 49
  lf = 7; cf = 71
)

/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

Gebiet __________________

 Autor ______________________________    Koautor ______________________________

    Nr __ Titel _______________________________________________________________

                                 Schrank / Regal ___________________ / ________ */
//                                               Papas Zimmer rechts   6 hinten

func writeMask (l, c uint) {
  scr.Colours (col.LightGray(), col.Black())
  scr.Write ("Gebiet",  l + 1, c +  0)
  scr.Write ("Autor",   l + 3, c +  1)
  scr.Write ("Koautor", l + 3, c + 41)
  scr.Write ("Nr",      l + 5, c +  4)
  scr.Write ("Titel",   l + 5, c + 10)
  scr.Write ("Schrank / Regal", l + 7, c + 33)
  scr.Write ("/", l + 7, c + 69)
}

var maskWritten = false

func (x *book) longC() string {
  n := x.cupboard.ProperLen()
  s := ""
  if n == 0 { return s }
  switch x.cupboard.Byte(0) {
  case 'w':
    s = "Wohnzimmer"
  case 'p':
    s = "Papas Zimmer"
  case 'g':
    s = str.Lat1("Gästezimmer")
  }
  if n == 3 {
    s += " "
    switch x.cupboard.Byte(2) {
    case 'l':
      s += "links"
    case 'r':
      s += "rechts"
    }
  }
  return s
}

func (x *book) longF() string {
  n := x.floor.ProperLen()
  s := ""
  if n == 0 { return s }
  s += string(x.floor.Byte(0))
  if n == 3 {
    s += " "
    switch x.floor.Byte(2) {
    case 'h':
      s += "hinten"
    case 'v':
      s += "vorne"
    }
  }
  return s
}

func (x *book) Write (l, c uint) {
  if ! maskWritten {
    writeMask (l, c)
    maskWritten = true
  }
  writeMask (l, c)
  x.Field.Write (l + lg, c + cg)
  x.author.Write (l + la, c + ca)
  x.coauthor.Write (l + lk, c + ck)
  if x.Natural.Val() != 0 {
    x.Natural.Write (l + ln, c + cn)
  }
  x.title.Write (l + lt, c + ct)
/*/
  x.cupboard.Write (l + lc, c + cc)
  x.floor.Write (l + lf, c + cf)
/*/
  scr.Write (str.New(19), l + lc, c + cc)
  scr.Write (x.longC(), l + lc, c + cc)
  scr.Write (str.New(8), l + lf, c + cf)
  scr.Write (x.longF(), l + lf, c + cf)
}

func containsSep (t text.Text) bool {
  _, c := str.Pos (t.String(), sep)
  return c
}

func edit (t text.Text, s string, l, c uint) {
  for {
    t.Edit (l, c)
    if containsSep (t) {
      errh.Error0 (s + nosep)
    } else {
      break
    }
  }
}

func (x *book) Edit (l, c uint) {
  i := 0
  loop:
  for {
    x.Write (l, c)
    switch i {
    case 0:
      x.Write (l, c)
      x.Field.Edit (l + lg, c + cg)
    case 1:
      edit (x.author, "Autor", l + la, c + ca)
/*/
      if k, _ := kbd.LastCommand(); k == kbd.Tab {
        for i := 0; i < len(a); i++ {
          if x.author.Sub0 (author[i]) {
            x.author.Copy (author[i])
            x.author.Write (l + la, c + ca)
            break
          }
        }
      }
/*/
    case 2:
      edit (x.coauthor, "Koautor", l + lk, c + ck)
/*/
      if k, _ := kbd.LastCommand(); k == kbd.Tab {
        if ! x.coauthor.Empty() {
          for i := 0; i < len(a); i++ {
            if x.coauthor.Sub0 (author[i]) {
              x.coauthor.Copy (author[i])
              x.coauthor.Write (l + lk, c + ck)
              break
            }
          }
        }
      }
/*/
    case 3:
      x.Natural.Edit (l + ln, c + cn)
    case 4:
      edit (x.title, "Titel", l + lt, c + ct)
    case 5:
      edit (x.cupboard, "Schrank", l + lc, c + cc)
    case 6:
      edit (x.floor, "Regal", l + lf, c + cf)
    }
    switch k, _ := kbd.LastCommand(); k {
    case kbd.Esc:
      break loop
    case kbd.Enter, kbd.Down:
      if i < 6 {
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

var lastField = field.New()

func (x *book) TeX() string {
  s := ""
  if ! x.Field.Eq (lastField) {
    lastField.Copy (x.Field)
    s += "\\bigskip\n"
    s += "\\line{\\bfbig\\hfil " + x.Field.TeX() + "\\hfil}\\medskip\\nopagebreak\n"
  }
  s += "\\vskip-6pt\n{\\bi " + x.author.TeX()
  if ! x.coauthor.Empty() {
    s += "/" + x.coauthor.TeX()
  }
  s += "}\\newline\\nopagebreak\n"
  sn := x.Natural.String()
  if sn == "0" { sn = "" }
  s += "\\hbox to 16pt{\\hfil"
  s += sn
  s += "}\\quad " + x.title.TeX()
  if ! x.cupboard.Empty() {
    s += " (" + x.cupboard.TeX()
  }
  if x.floor.Empty() {
    s += ")"
  } else {
    s += " " + x.floor.TeX() + ")"
  }
  s += "\n\\par\\smallpagebreak"
  return s
}

func (x *book) Codelen() uint {
  return x.Field.Codelen() +
       2 * lena +
       x.Natural.Codelen() +
       2 * 3
}

func (x *book) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.Field.Codelen()
  copy (s[i:i+a], x.Field.Encode())
  i += a
  a = lena
  copy (s[i:i+a], x.author.Encode())
  i += a
  copy (s[i:i+a], x.coauthor.Encode())
  i += a
  a = x.Natural.Codelen()
  copy (s[i:i+a], x.Natural.Encode())
  i += a
  a = lent
  copy (s[i:i+a], x.title.Encode())
  i += a
  a = 3
  copy (s[i:i+a], x.cupboard.Encode())
  i += a
  copy (s[i:i+a], x.floor.Encode())
  return s
}

func (x *book) Decode (s Stream) {
  i, a := uint(0), x.Field.Codelen()
  x.Field.Decode (s[i:i+a])
  i += a
  a = lena
  x.author.Decode (s[i:i+a])
  i += a
  x.coauthor.Decode (s[i:i+a])
  i += a
  a = x.Natural.Codelen()
  x.Natural.Decode (s[i:i+a])
  i += a
  a = lent
  x.title.Decode (s[i:i+a])
  i += a
  a = 3
  x.cupboard.Decode (s[i:i+a])
  i += a
  x.floor.Decode (s[i:i+a])
}

func (x *book) Rotate() {
  actOrder = 1 - actOrder
}

func (x *book) Index() Func {
  return Id
}
