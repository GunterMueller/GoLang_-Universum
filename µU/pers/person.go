package pers

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/text"
  "µU/sex"
  "µU/day"
)
const (
  lenName = uint(27)
  lenFirstName = uint(15)
  lenShort = lenName + lenFirstName + 2 // ", "
)
const ( // Order
  nameOrder = iota
  ageOrder
)
type (
  person struct {
        surname,
      firstName text.Text
                sex.Sex
                day.Calendarday "birthday"
          title text.Text
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
  x.surname = text.New (lenName)
  x.firstName = text.New (lenFirstName)
  x.Sex = sex.New()
  x.Calendarday = day.New()
  x.title = text.New (lenName)
  x.field = []any { x.surname, x.firstName, x.Calendarday, x.Sex, x.title }
  x.cl = []uint {lenName, lenFirstName, x.Calendarday.Codelen(),
                 x.Sex.Codelen(), x.title.Codelen()}
  x.Format = LongB
  return x
}

func (x *person) imp(Y any) *person {
  y, ok := Y.(*person)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *person) Empty() bool {
  if x.surname.Empty() {
    return x.firstName.Empty()
  }
  return false
}

func (x *person) Clr() {
  x.surname.Clr()
  x.firstName.Clr()
  x.Sex.Clr()
  x.Calendarday.Clr()
  x.title.Clr()
}

func (x *person) Identifiable() bool {
  return ! x.surname.Empty() &&
         ! x.firstName.Empty() &&
         ! x.Calendarday.Empty()
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
  x.surname.Copy (y.surname)
  x.firstName.Copy (y.firstName)
  x.Sex.Copy (y.Sex)
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
  g := x.surname.Eq (y.surname) &&
       x.firstName.Eq (y.firstName)
  switch x.Format {
  case ShortB, LongB, LongTB:
    g = g && x.Calendarday.Eq (y.Calendarday)
  }
  switch x.Format {
  case LongTB:
    g = g && x.title.Eq (y.title) &&
             x.Sex.Eq (y.Sex)
  }
  return g
}

func (x *person) Equiv (Y Person) bool {
  y := x.imp (Y)
  if actualOrder == nameOrder {
    return x.surname.Eq (y.surname) &&
           x.firstName.Eq (y.firstName) &&
           x.Calendarday.Eq (y.Calendarday)
  } // actualOrder == AgeOrder
  return x.Calendarday.Eq (y.Calendarday)
}

func (x *person) Less (Y any) bool {
  y := x.imp (Y)
  if actualOrder == nameOrder {
    if x.surname.Eq (y.surname) {
      if x.firstName.Eq (y.firstName) {
        return x.Calendarday.Less (y.Calendarday)
      }
      return x.firstName.Less (y.firstName)
    }
    return x.surname.Less (y.surname)
  } // actualOrder == AgeOrder
  if x.Calendarday.Eq (x.imp (Y).Calendarday) {
    if x.surname.Eq (y.surname) {
      return x.firstName.Less (y.firstName)
    }
    return x.surname.Less (y.surname)
  }
  return x.Calendarday.Less (x.imp (Y).Calendarday)
}

func (x *person) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *person) Sub (Y any) bool {
  y := x.imp (Y)
  if ! x.surname.Sub (y.surname) {
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
  x.surname.Colours (f, b)
  x.firstName.Colours (f, b)
  x.Sex.Colours (f, b)
  x.Calendarday.Colours (f, b)
  x.title.Colours (f, b)
}

func (x *person) Cols() (col.Colour, col.Colour) {
  return x.surname.Cols()
}

var
  ln, lv, lg, lb, la, cn, cv, cg, cb, ca uint
/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

without mask:

Short: Name, Vorname
__________________________, _______________

ShortB: Name, Vorname (GebDat) 
__________________________, _______________ (________)

with mask:

Long:
 Name: ___________________________   Vorname: _______________

LongB:
 Name: ___________________________   Vorname: _______________    geb.: ________

LongTB:
 Name: ___________________________   Vorname: _______________    geb.: ________
 Anr.: ___________________________       m/w: _
*******************************************************************************/

func (x *person) writeMask (l, c uint) {
  cn = 7; cv = 45; cg = cv; cb = 70; ca = cn
  switch x.Format {
  case Short, ShortB:
    cn = 0; cv = 28; cb = 45
  }
  bx.Wd (1)
  bx.ScrColours()
  switch x.Format {
  case Long, LongB, LongTB:
    bx.Wd (5)
    bx.Write ("Name:", l, c + cn - 6)
    bx.Wd (8)
    bx.Write ("Vorname:", l, c + cv - 9)
  }
  switch x.Format {
  case Short:
    bx.Write (",", l, c + cv - 2)
    return
  case ShortB:
    bx.Write (",", l, c + cv - 2)
    bx.Write ("(", l, c + cb - 1)
    bx.Write (")", l, c + cb + 8)
  case LongB:
    bx.Wd (4)
    bx.Write ("geb.:", l, c + cb - 6)
  case LongTB:
    bx.Wd (5)
    bx.Write ("geb.:", l, c + cb - 6)
    bx.Wd (5)
    bx.Write ("Anr.:", l + 1, c + ca - 6)
    bx.Wd (4)
    bx.Write ("m/w:", l + 1, c + cg - 5)
  }
}

func (x *person) TeX() string {
  s := ""
  if ! x.title.Empty() { s += "{\\bf " + x.title.TeX() + "} " }
  s += "{\\bf " + x.firstName.TeX() + " " + x.surname.TeX() + "}"
//  if ! x.Sex.Empty() { s += " (" + x.Sex.String() + ")" }
  if ! x.Calendarday.Empty() { s += " (" + x.Calendarday.String() + ")" }
  s += "\\newline\n"
  return s
}

func (x *person) Write (l, c uint) {
  x.writeMask (l, c)
  x.surname.Write (l, c + cn)
  x.firstName.Write (l, c + cv)
  switch x.Format {
  case ShortB, LongB:
    x.Calendarday.Write (l, c + cb)
  case LongTB:
    x.Calendarday.Write (l, c + cb)
    x.title.Write (l + 1, c + ca)
    x.Sex.Write (l + 1, c + cg)
  }
}

func (x *person) Edit (l, c uint) {
  x.Write (l, c)
  i := uint(0)
  if C, _ := kbd.LastCommand(); C == kbd.Up { // see persaddr
    i = 4
  }
  loop:
  for {
    switch i {
    case 0:
      x.surname.Edit (l + ln, c + cn)
    case 1:
      x.firstName.Edit (l + lv, c + cv)
    case 2:
      x.Calendarday.Edit (l + lg, c + cb)
    case 3:
      if x.Format == LongTB {
        x.title.Edit (l + 1, c + ca)
      }
    case 4:
      if x.Format == LongTB {
        x.Sex.Edit (l + 1, c + cg)
      }
    }
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < 4 {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < 4 {
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
  return lenName + lenFirstName +
         x.Calendarday.Codelen() +
         x.Sex.Codelen() +
         lenName
}

func (x *person) Encode() Stream {
  return Encodes (x.field, x.cl)
}

func (x *person) Decode (bs Stream) {
  Decodes (bs, x.field, x.cl)
  x.surname = x.field[0].(text.Text)
  x.firstName = x.field[1].(text.Text)
  x.Calendarday = x.field[2].(day.Calendarday)
  x.Sex = x.field[3].(sex.Sex)
  x.title = x.field[4].(text.Text)
}

func (x *person) Index() Func {
  if actualOrder == nameOrder {
    return func (a any) any {
      x, ok := a.(*person)
      if ! ok { TypeNotEqPanic(x, a) }
      return x.surname
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
