package book

// (c) Christian Maurer   v. 240318 - license see µU.go

import (
  "µU/env"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/scr"
  "µU/errh"
  "µU/str"
  "µU/text"
  "µU/bn"
  "µU/enum"
)
const (
  lenf = 20
  lena = 30
  lenn =  2
  lent = 63
  lenl = 22
  lenlk = 3
  sep = ';'
  seps = ";"
)
const (
  fieldIndex = iota
  authorIndex
  numberIndex
  titleIndex
  locationIndex
  nIndices
)
const
  field = "Gebiet"
type
  book struct {
        field enum.Enum
       author,
     coauthor text.Text
              bn.Natural
        title,
     location text.Text
              }
var
  actIndex = fieldIndex

func new_() Book {
  x := new (book)
  x.field = enum.Newk (lenf, lenn)
  if env.E() {
// If you want to use the program "books",
// you should adapt the following to your personal requirements.
    x.field.Set ("Biography", "Classical Literature", "Crime Novel", "Roman")
    x.field.Setk ("bi", "cl", "cn", "ro")
  } else {
// Wenn Sie das Programm "bücher" benutzen wollen,
// passen Sie das Folgende an Ihre eigenen Anforderungen an.
    x.field.Set ("Ägypten", "Antike Biographie", "Etrurien", "Griechenland",
                 "Griechisch", "Griechischer Text", "Historischer Roman", "Horror",
                 "Italien", "Italien-Krimi", "Italien-Roman", "Jugendbuch",
                 "Klassische Literatur", "Krimi", "Kunst", "Lateinisch", "Lateinischer Text",
                 "Neuere Literatur", "Orient", "Pflanzen", "Rom-Krimi",
                 "Rom-Roman", "Sachbuch", "Science Fiction", "Theaterstück(e)", "XX")
    x.field.Setk ("äg", "ab", "et", "gl",
                  "gr", "gt", "hr", "ho",
                  "il", "ik", "ir", "jb",
                  "kl", "kr", "ku", "la", "lt",
                  "nl", "or", "pf", "rk",
                  "rr", "sb", "sf", "th", "xx")
  }
  x.field.Colours (col.FlashWhite(), col.Blue())
  x.author = text.New (lena)
  x.author.Colours (col.Yellow(), col.Red())
  x.coauthor = text.New (lena)
  x.coauthor.Colours (col.Yellow(), col.Red())
  x.Natural = bn.New (lenn)
  x.Natural.Colours (col.FlashWhite(), col.DarkGray())
  x.title = text.New (lent)
  x.title.Colours (col.FlashWhite(), col.DarkGreen())
  x.location = text.New (lenl)
  x.location.Colours (col.FlashWhite(), col.Brown())
  return x
}

func (x *book) imp (Y any) *book {
  y, ok := Y.(*book)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *book) Empty() bool {
  return x.field.Empty() &&
         x.author.Empty() &&
         x.coauthor.Empty() &&
         x.Natural.Empty() &&
         x.title.Empty() &&
         x.location.Empty()
}

func (x *book) Clr() {
  x.field.Clr()
  x.author.Clr()
  x.coauthor.Clr()
  x.Natural.Clr()
  x.title.Clr()
  x.location.Clr()
}

func (x *book) Eq (Y any) bool {
  y := x.imp(Y)
  return x.field.Eq (y.field) &&
         x.author.Eq (y.author) &&
         x.coauthor.Eq (y.coauthor) &&
         x.title.Eq (y.title) &&
         x.Natural.Eq (y.Natural) &&
         x.location.Eq (y.location)
}

func (x *book) Copy (Y any) {
  y := x.imp(Y)
  x.field.Copy (y.field)
  x.author.Copy (y.author)
  x.coauthor.Copy (y.coauthor)
  x.Natural.Copy (y.Natural)
  x.title.Copy (y.title)
  x.location.Copy (y.location)
}

func (x *book) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *book) Less (Y any) bool {
  y := x.imp(Y)
  switch actIndex {
  case fieldIndex:
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
    if ! x.author.Eq (y.author) {
      return x.author.Less (y.author)
    }
    if ! x.Natural.Eq (y.Natural) {
      return x.Natural.Less (y. Natural)
    }
    if ! x.title.Eq (y.title) {
      return x.title.Less (y.title)
    }
  case authorIndex, numberIndex:
    if ! x.author.Eq (y.author) {
      return x.author.Less (y.author)
    }
    if ! x.Natural.Eq (y.Natural) {
      return x.Natural.Less (y. Natural)
    }
    if ! x.title.Eq (y.title) {
      return x.title.Less (y.title)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
  case titleIndex:
    if ! x.title.Eq (y.title) {
      return x.title.Less (y.title)
    }
    if ! x.author.Eq (y.author) {
      return x.author.Less (y.author)
    }
    if ! x.Natural.Eq (y.Natural) {
      return x.Natural.Less (y.Natural)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
  case locationIndex:
    if ! x.location.Eq (y.title) {
      return x.location.Less (y.title)
    }
    if ! x.author.Eq (y.author) {
      return x.author.Less (y.author)
    }
    if ! x.Natural.Eq (y.Natural) {
      return x.Natural.Less (y.Natural)
    }
    if ! x.title.Eq (y.title) {
      return x.title.Less (y.title)
    }
    if ! x.field.Eq (y.field) {
      return x.field.Less (y.field)
    }
  }
  return false
}

func (x *book) longL() string {
  ds := x.location.String()
  n := str.ProperLen(ds)
  if n == 0 { return "" }
//  if n != lenlk { errh.Error (ds, n) }
  s := ""
  if env.E() {
    switch ds[0] {
// If you want to use the program "books",
// you should adapt the following to your personal requirements.
    case 'l':
      s = "living room"
    case 's':
      s = "study"
    case 'g':
      s = str.Lat1 ("guest room")
    default:
      s = ""
    }
    s += " "
    switch ds[1] {
    case 'l':
      s += "left"
    case 'r':
      s += "right"
    }
    s += " "
    str.Append (&s, ds[2])
  } else {
    switch ds[0] {
// Wenn Sie das Programm "bücher" benutzen wollen,
// passen Sie das Folgende an Ihre eigenen Anforderungen an.
    case 'w':
      s = "Wohnzimmer"
    case 'a':
      s = "Arbeitszimmer"
    case 'g':
      s = str.Lat1 ("Gästezimmer")
    default:
      s = ""
    }
    s += " "
    switch ds[1] {
    case 'l':
      s += "links"
    case 'r':
      s += "rechts"
    }
    s += " "
    str.Append (&s, ds[2])
  }
  return s
}

func (x *book) string0() string {
  s := x.field.String()
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
  return s
}

func (x *book) String() string {
  s := x.string0()
  t := x.longL()
  s += t + seps
  return s
}

func (x *book) String1() string {
  s := x.string0()
  t := x.location.String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *book) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != 6 { return false }
  if ! x.field.Defined (ss[0]) { return false }
  if ! x.author.Defined (ss[1]) { return false }
  if ! x.coauthor.Defined (ss[2]) { return false }
  if ! x.Natural.Defined (ss[3]) { return false }
  if ! x.title.Defined (ss[4]) { return false }
  if ! x.location.Defined (ss[5]) { return false }
  return true
}

func (x *book) Sub (Y any) bool {
  y := x.imp(Y)
  s := false
  if ! x.field.Empty() {
    s = s || x.field.Eq (y.field)
  }
  if ! x.author.Empty() {
    s = s || x.author.Sub (y.author)
  }
  if ! x.title.Empty() {
    s = s || x.title.Sub (y.title)
  }
  if ! x.location.Empty() {
    s = s || x.location.Sub (y.location)
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
)

/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

 field __________________

author ______________________________   coauthor ______________________________

    nr __ title _______________________________________________________________

                                        location _____________________
Gebiet __________________

 Autor ______________________________    Koautor ______________________________

    Nr __ Titel _______________________________________________________________

                                       Ablageort _____________________ */

func writeMask (l, c uint) {
  scr.Colours (col.LightGray(), col.Black())
  if env.E() {
    scr.Write ("field",    l + 1, c +  1)
    scr.Write ("author",   l + 3, c +  9)
    scr.Write ("coauthor", l + 3, c + 40)
    scr.Write ("nr",       l + 5, c +  4)
    scr.Write ("title",    l + 5, c + 10)
    scr.Write ("location", l + 7, c + 40)
  } else {
    scr.Write ("Gebiet",  l + 1, c +  0)
    scr.Write ("Autor",   l + 3, c +  1)
    scr.Write ("Koautor", l + 3, c + 41)
    scr.Write ("Nr",      l + 5, c +  4)
    scr.Write ("Titel",   l + 5, c + 10)
    scr.Write ("Ablageort", l + 7, c + 39)
  }
}

var
  maskWritten = false

func (x *book) Write (l, c uint) {
  if ! maskWritten {
    writeMask (l, c)
    maskWritten = true
  }
  writeMask (l, c)
  x.field.Write (l + lg, c + cg)
  x.author.Write (l + la, c + ca)
  x.coauthor.Write (l + lk, c + ck)
  if x.Natural.Val() != 0 {
    x.Natural.Write (l + ln, c + cn)
  }
  x.title.Write (l + lt, c + ct)
  scr.Colours (col.FlashWhite(), col.Brown())
  scr.Write (str.New (lenl), l + lc, c + cc)
  scr.Colours (col.FlashWhite(), col.Brown())
  scr.Write (x.longL(), l + lc, c + cc)
}

func containsSep (t text.Text) bool {
  _, c := str.Pos (t.String(), sep)
  return c
}

func edit (t text.Text, s string, l, c uint) {
  for {
    t.Edit (l, c)
    if containsSep (t) {
      if env.E() {
        errh.Error0 (s + " must not contain " + seps)
      } else {
        errh.Error0 (s + " darf kein " + seps + " enthalten")
      }
    } else {
      break
    }
  }
}

func (x *book) Edit (l, c uint) {
  i, s := 0, ""
  loop:
  for {
    x.Write (l, c)
    switch i {
    case 0:
      x.Write (l, c)
      x.field.Edit (l + lg, c + cg)
    case 1:
      if env.E() { s = "author" } else { s = "Autor" }
      edit (x.author, s, l + la, c + ca)
    case 2:
      if env.E() { s = "coauthor" } else { s = "Koautor" }
      edit (x.coauthor, s, l + lk, c + ck)
    case 3:
      x.Natural.Edit (l + ln, c + cn)
    case 4:
      if env.E() { s = "title" } else { s = "Titel" }
      edit (x.title, s, l + lt, c + ct)
    case 5:
      if env.E() {
        errh.Hint ("edit a string of 3 characters")
      } else {
        errh.Hint ("Geben Sie eine Zeichenkette aus 3 Zeichen ein")
      }
      var ok0, ok1, ok2 bool
      for {
        edit (x.location, s, l + lc, c + cc)
        n := x.location.ProperLen()
        if n == 3 {
          t := x.location.String()
          if env.E() {
// If you want to use the program "books",
// you should adapt the following to your personal requirements.
            switch t[0] {
            case 'l', 's', 'g':
              ok0 = true
            default:
              ok0 = false
              errh.Error0 ("as 1st character only \"l\", \"s\" or \"g\" permissible")
            }
          } else {
// Wenn Sie das Programm "bücher" benutzen wollen,
// passen Sie das Folgende an Ihre eigenen Anforderungen an.
            switch t[0] {
            case 'w', 'a', 'g':
              ok0 = true
            default:
              ok0 = false
              errh.Error0 ("als erste Zeichen nur \"w\", \"a\" oder \"g\" zulässig")
            }
          }
          switch t[1] {
          case 'l', 'r':
            ok1 = true
          default:
            ok1 = false
            if env.E() {
              errh.Error0 ("as 2nd character only \"l\" or \"r\" permissible")
            } else {
              errh.Error0 ("als 2. Zeichen nur \"l\" oder \"r\" zulässig")
            }
          }
          switch t[2] {
          case '1', '2', '3', '4', '5', '6', '7', '8', '9':
            ok2 = true
          default:
            ok2 = false
            if env.E() {
              errh.Error0 ("as 3rd character only a digit permissible")
            } else {
              errh.Error0 ("als 3. Zeichen ist nur eine Ziffer zulässig")
            }
          }
          if ok0 && ok1 && ok2 {
            scr.Colours (col.FlashWhite(), col.Brown())
            scr.Write (x.longL(), l + lc, c + cc)
            errh.DelHint()
            break
          }
        } else {
          if env.E() {
            errh.Error0 (x.location.String() + " does not contain exactly 3 characters")
          } else {
            errh.Error0 (x.location.String() + " enthält nicht genau 3 Zeichen")
          }
        }
      }
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

var
  lastField = enum.New (lenf)

func (x *book) TeX() string {
  s := ""
  if ! x.field.Eq (lastField) {
    lastField.Copy (x.field)
    s += "\\bigskip\n" + "\\line{\\bfbig\\hfil "
    s += x.field.(TeXer).TeX()
    s += "\\hfil}\\medskip\\nopagebreak\n"
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
  if ! x.location.Empty() {
    s += " (" + x.location.TeX() + ")"
  }
  s += "\n\\par\\smallpagebreak"
  return s
}

func (x *book) Codelen() uint {
  return x.field.Codelen() +
       2 * lena +
       x.Natural.Codelen() +
       lent +
       lenl
}

func (x *book) Encode() Stream {
  s := make(Stream, x.Codelen())
  i, a := uint(0), x.field.Codelen()
  copy (s[i:i+a], x.field.Encode())
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
  a = lenl
  copy (s[i:i+a], x.location.Encode())
  return s
}

func (x *book) Decode (s Stream) {
  i, a := uint(0), x.field.Codelen()
  x.field.Decode (s[i:i+a])
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
  a = lenl
  x.location.Decode (s[i:i+a])
}

func (x *book) Rotate() {
  actIndex = (actIndex + 1) % nIndices
}

func (x *book) Index() Func {
  return Id
}
