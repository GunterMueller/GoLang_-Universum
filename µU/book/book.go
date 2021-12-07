package book

// (c) Christian Maurer   v. 211126 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/text"
  "µU/bn"
  "µU/book/field"
)
const (
  len0 = 30
  len1 = 63
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
       location text.Text
               }
var
  actOrder = fieldOrder

func (x *book) imp (Y Any) *book {
  y, ok := Y.(*book)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func new_() Book {
  x := new (book)
  x.Field = field.New()
  x.author = text.New (len0)
  x.coauthor = text.New (len0)
  x.Natural = bn.New (2)
  x.title = text.New (len1)
  x.location = text.New (len0)
  x.Natural.Colours (col.LightWhite(), col.DarkCyan())
  x.author.Colours (col.Yellow(), col.Red())
  x.coauthor.Colours (col.Yellow(), col.Red())
  x.title.Colours (col.LightWhite(), col.DarkGreen())
  x.location.Colours (col.LightWhite(), col.DarkGray())
  return x
}

func (x *book) Empty() bool {
  return x.Field.Empty() &&
         x.author.Empty() && x.coauthor.Empty() &&
         x.Natural.Empty() &&
         x.title.Empty() &&
         x.location.Empty()
}

func (x *book) Clr() {
  x.Field.Clr()
  x.author.Clr(); x.coauthor.Clr()
  x.Natural.Clr()
  x.title.Clr()
  x.location.Clr()
}

func (x *book) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.Field.Eq (y.Field) &&
         x.author.Eq (y.author) && x.coauthor.Eq (y.coauthor) &&
         x.Natural.Eq (y.Natural) &&
         x.title.Eq (y.title) &&
         x.location.Eq (y.location)
}

func (x *book) Copy (Y Any) {
  y := x.imp(Y)
  x.Field.Copy (y.Field)
  x.author.Copy (y.author); x.coauthor.Copy (y.coauthor)
  x.Natural.Copy (y.Natural)
  x.title.Copy (y.title)
  x.location.Copy (y.location)
}

func (x *book) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *book) Less (Y Any) bool {
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

func (x *book) Sub (Y Any) bool {
  y := x.imp(Y)
  s := false
  if ! x.author.Empty() {
    s = s || x.author.Sub (y.author)
  }
  if ! x.title.Empty() {
    s = s || x.title.Sub (y.title)
  }
  return s
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

func writeMask (l, c uint) {
  scr.Colours (col.LightGray(), col.Black())
  scr.Write ("Gebiet",  l + 1, c + 0)
  scr.Write ("Autor",   l + 3, c + 1)
  scr.Write ("Koautor", l + 3, c +41)
  scr.Write ("Nr",      l + 5, c + 4)
  scr.Write ("Titel",   l + 5, c +10)
  scr.Write ("Fundort", l + 7, c +41)
}

var maskWritten = false

func (x *book) Write (l, c uint) {
  if ! maskWritten {
    writeMask (l, c)
    maskWritten = true
  }
  writeMask (l, c)
  x.Field.Write (l + lg, c + cg)
  x.author.Write (l + la, c + ca)
  x.coauthor.Write (l + lk, c + ck)
  x.Natural.Write (l + ln, c + cn)
  x.title.Write (l + lt, c + ct)
  x.location.Write (l + lf, c + cf)
}

func (x *book) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  loop:
  for {
    switch i {
    case 0:
      x.Field.Edit (l + lg, c + cg)
    case 1:
      x.author.Edit (l + la, c + ca)
      if k, _ := kbd.LastCommand(); k == kbd.Tab {
        for i := 0; i < len(a); i++ {
          if x.author.Sub0 (author[i]) {
            x.author.Copy (author[i])
            x.author.Write (l + la, c + ca)
            break
          }
        }
      }
    case 2:
      x.coauthor.Edit (l + lk, c + ck)
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
    case 3:
      x.Natural.Edit (l + ln, c + cn)
    case 4:
      x.title.Edit (l + lt, c + ct)
    case 5:
      x.location.Edit (l + lf, c + cf)
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

var lastField = field.New()

func (x *book) TeX() string {
  s := ""
  if ! x.Field.Eq (lastField) {
    lastField.Copy (x.Field)
    s += "\\bigskip\\line{\\bfbig\\hfil " + x.Field.TeX() + "\\hfil}\\medskip\\nopagebreak\n"
  }
  s += "{\\bi " + x.author.TeX()
  if ! x.coauthor.Empty() {
    s += "/" + x.coauthor.TeX()
  }
  s += "}\n\\newline"
  sn := x.Natural.String()
  if sn == "0" { sn = "" }
  s += "\\hbox to 16pt{\\hfil"
  s += sn
  s += "}\\quad " + x.title.TeX()
  if ! x.location.Empty() {
    s += " (" + x.location.TeX() + ")"
  }
  s += "\n\\par\\smallpagebreak\n"
  return s
}

func (x *book) Codelen() uint {
  return x.Field.Codelen() +
       2 * len0 +
       x.Natural.Codelen() +
       len1 + len0
}

func (x *book) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.Field.Codelen()
  copy (s[i:i+a], x.Field.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.author.Encode())
  i += a
  copy (s[i:i+a], x.coauthor.Encode())
  i += a
  a = x.Natural.Codelen()
  copy (s[i:i+a], x.Natural.Encode())
  i += a
  a = len1
  copy (s[i:i+a], x.title.Encode())
  i += a
  a = len0
  copy (s[i:i+a], x.location.Encode())
  return s
}

func (x *book) Decode (s Stream) {
  i, a := uint(0), x.Field.Codelen()
  x.Field.Decode (s[i:i+a])
  i += a
  a = len0
  x.author.Decode (s[i:i+a])
  i += a
  x.coauthor.Decode (s[i:i+a])
  i += a
  a = x.Natural.Codelen()
  x.Natural.Decode (s[i:i+a])
  i += a
  a = len1
  x.title.Decode (s[i:i+a])
  i += a
  a = len0
  x.location.Decode (s[i:i+a])
}

func (x *book) Rotate() {
  actOrder = 1 - actOrder
}

func (x *book) Index() Func {
  return Id
}
