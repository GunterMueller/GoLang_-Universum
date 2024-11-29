package pers

// (c) Christian Maurer   v. 241109 - license see µU.go

import (
  "µU/env"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/str"
  "µU/font"
  "µU/pbox"
  "µU/text"
  "µU/day"
)
const (
  lenn = uint(27)
  lenf = uint(15)
  lent = uint(15)
  sep = ','
  seps = ","
)
const ( // Order
  nameOrder = iota
  ageOrder
)
type (
  person struct {
           name,
      firstName,
          title text.Text
                day.Calendarday "birthdate" // "Geburtsdatum"
          field []any // to [En|De]code
             cl []uint
                Format
                }
)
var (
  actualOrder = nameOrder
  bx = box.New()
  pbx = pbox.New()
  tmp = day.New()
)

func new_() Person {
  x := new(person)
  x.name = text.New (lenn)
  x.firstName = text.New (lenf)
  x.Calendarday = day.New()
  f, b := col.FlashWhite(), col.DarkRed()
  x.name.Colours (f, b)
  x.firstName.Colours (f, b)
  x.Calendarday.Colours (f, b)
  x.title = text.New (lent)
  x.field = []any { x.name, x.firstName, x.Calendarday, x.title }
  x.cl = []uint {lenn, lenf, x.Calendarday.Codelen(), x.title.Codelen()}
//  x.Format = NameB
  x.Format = NameBT
  return x
}

func (x *person) imp (Y any) *person {
  y, ok := Y.(*person)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *person) Empty() bool {
  if x.name.Empty() {
    return x.firstName.Empty()
  }
  return false
}

func (x *person) Clr() {
  x.name.Clr()
  x.firstName.Clr()
  x.Calendarday.Clr()
  x.title.Clr()
}

func (x *person) Identifiable() bool {
  return ! x.name.Empty() &&
         ! x.firstName.Empty() &&
         ! x.Calendarday.Empty()
}

func (x *person) Birthdate() day.Calendarday {
  return x.Calendarday
}

func (x *person) FullAged() bool {
  if x.Calendarday.Empty() { return false }
  tmp.Copy (x)
  for i := uint(0); i < 18; i++ {
    tmp.Inc (day.Yearly)
  }
  return tmp.Elapsed()
}

func (x *person) Age() uint {
  today := day.New()
	today.Update()
	birth := x.Calendarday.Clone().(day.Calendarday)
  var a uint
  for birth.Less (today) {
    a++
    birth.Inc (day.Yearly)
  }
  return a - 1
}

func (x *person) Copy (Y any) {
  y := x.imp (Y)
  x.name.Copy (y.name)
  x.firstName.Copy (y.firstName)
  x.Calendarday.Copy (y.Calendarday)
  x.title.Copy (y.title)
  x.Format = y.Format
}

func (x *person) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *person) Eq (Y any) bool {
  y := x.imp (Y)
  return x.name.Eq (y.name) &&
    x.firstName.Eq (y.firstName) &&
  x.Calendarday.Eq (y.Calendarday) &&
        x.title.Eq (y.title)
}

func (x *person) Equiv (Y Person) bool {
  y := x.imp (Y)
  if actualOrder == nameOrder {
    return x.Eq (y)
  } // actualOrder == AgeOrder
  return x.Calendarday.Eq (y.Calendarday)
}

func (x *person) Less (Y any) bool {
  y := x.imp (Y)
  if actualOrder == nameOrder {
    if x.name.Eq (y.name) {
      if x.firstName.Eq (y.firstName) {
        return x.Calendarday.Less (y.Calendarday)
      }
      return x.firstName.Less (y.firstName)
    }
    return x.name.Less (y.name)
  } // actualOrder == AgeOrder
  if x.Calendarday.Eq (x.imp (Y).Calendarday) {
    if x.name.Eq (y.name) {
      return x.firstName.Less (y.firstName)
    }
    return x.name.Less (y.name)
  }
  return x.Calendarday.Less (x.imp (Y).Calendarday)
}

func (x *person) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *person) String() string {
  s := x.name.String()
  str.OffSpc1 (&s)
  s += seps
  t := x.firstName.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.title.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.Calendarday.String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *person) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != 4 { return false }
  if ! x.name.Defined (ss[0]) { return false }
  if ! x.firstName.Defined (ss[1]) { return false }
  if ! x.title.Defined (ss[2]) { return false }
  if ! x.Calendarday.Defined (ss[3]) { return false }
  return true
}

func (x *person) Sub (Y any) bool {
  y := x.imp (Y)
  if ! x.name.Sub (y.name) {
    return false
  }
  if ! x.firstName.Sub (y.firstName) {
    return false
  }
/*/
  if ! x.Calendarday.Empty() && y.Calendarday.Empty() {
    return false
  }
  if ! x.Calendarday.Eq (y.Calendarday) {
    return false
  }
/*/
  return true
}

func (x *person) GetFormat() Format {
  return x.Format
}

func (x *person) SetFormat (f Format) {
  if f < NFormats {
    x.Format = f
  }
}

func (x *person) Colours (f, b col.Colour) {
  x.name.Colours (f, b)
  x.firstName.Colours (f, b)
  x.Calendarday.Colours (f, b)
  x.title.Colours (f, b)
}

func (x *person) Cols() (col.Colour, col.Colour) {
  return x.name.Cols()
}

const (
  cn = 10; cf = 48; cb = 71; ct = 10
)
/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

Nachname: ___________________________  Vorname: _______________  geb.: ________
  Anrede: _______________
 surname: ___________________________  fstname: _______________  born: ________
   title: _______________
*/

func (x *person) writeMask (l, c uint) {
  var n, f, b, t string
  if env.E() {
    n, f, b, t = " surname", "fstname", "born", " title"
  } else {
    n, f, b, t = "Nachname", "Vorname", "geb.", "Anrede"
  }
  bx.ScrColours()
  bx.Wd (9)
  bx.Write (n + ":", l, c + cn - 10)
  bx.Wd (8)
  bx.Write (f + ":", l, c + cf - 9)
  switch x.Format {
  case NameB, NameBT:
    bx.Write (b + ":", l, c + cb - 6)
  }
  if x.Format == NameBT {
    bx.Wd (7)
    bx.Write (t + ":", l + 1, c + cn - 8)
  }
}

func (x *person) TeX() string {
  s := ""
  if ! x.title.Empty() { s += "{\\bf " + x.title.TeX() + "} " }
  s += "{\\bf " + x.firstName.TeX() + " " + x.name.TeX() + "}"
  if ! x.Calendarday.Empty() { s += " (" + x.Calendarday.String() + ")" }
  return s + "\n"
}

func (x *person) Write (l, c uint) {
  x.writeMask (l, c)
  x.name.Write (l, c + cn)
  x.firstName.Write (l, c + cf)
  switch x.Format {
  case NameB, NameBT:
//    x.Calendarday.Colours (col.FlashWhite(), col.Blue())
    x.Calendarday.Write (l, c + cb)
  }
  if x.Format == NameBT {
    x.title.Write (l + 1, c + cn)
  }
}

func (x *person) Edit (l, c uint) {
  x.Write (l, c)
  i := uint(0)
  if C, _ := kbd.LastCommand(); C == kbd.Up { // see persaddr
    i = N - 1
  }
  loop:
  for {
    switch i {
    case 0:
      x.name.Edit (l, c + cn)
    case 1:
      x.firstName.Edit (l, c + cf)
    case 2:
      if x.Format > Name {
        x.Calendarday.Edit (l, c + cb)
      }
    case 3:
      if x.Format == NameBT {
        x.title.Edit (l + 1, c + cn)
      }
    }
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < N - 1 {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < N - 1 {
        i++
      } else {
        break loop
      }
    case kbd.Up, kbd.Left:
      if i > 0 {
        i--
      } else {
        break loop
      }
    case kbd.Search:
      break loop
    }
  }
}

func (x *person) SetFont (f font.Font) {
  // dummy
}

func (x *person) Print (l, c uint) {
  pbx.Print (x.TeX(), l, c)
}

func (x *person) Codelen() uint {
  return lenn + lenf +
         x.Calendarday.Codelen() +
         lenn
}

func (x *person) Encode() Stream {
  return Encodes (x.field, x.cl)
}

func (x *person) Decode (bs Stream) {
  Decodes (bs, x.field, x.cl)
  x.name = x.field[0].(text.Text)
  x.firstName = x.field[1].(text.Text)
  x.Calendarday = x.field[2].(day.Calendarday)
  x.title = x.field[3].(text.Text)
}

func (x *person) Index() Func {
  if actualOrder == nameOrder {
    return func (a any) any {
      x, ok := a.(*person)
      if ! ok { TypeNotEqPanic(x, a) }
      return x.name
    }
  } // actualOrder == AgeOrder
  return func (a any) any {
    x, ok := a.(*person)
    if ! ok { TypeNotEqPanic(x, a) }
    return x.Calendarday
  }
}

func (x *person) Rotate() {
  actualOrder = 1 - actualOrder
}
