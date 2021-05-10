package pers

// (c) Christian Maurer   v. 210510 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/text"
  "µU/tval"
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
                tval.TruthValue "female/indetermined/male"
                day.Calendarday "birthday"
          title text.Text
          field []Any // to [En|De]code
             cl []uint
                Format
                }
)
var (
  actualOrder = nameOrder
  bx, shbx = box.New(), box.New()
  pbx = pbox.New()
  tmp = day.New()
)

func init() {
  shbx.Wd (lenShort)
}

func new_() Person {
  x := new(person)
  x.surname = text.New (lenName)
  x.firstName = text.New (lenFirstName)
  x.TruthValue = tval.New()
  x.TruthValue.SetFormat (" ", "m", "w")
  x.Calendarday = day.New()
  x.title = text.New (lenName)
  x.field = []Any { x.surname, x.firstName, x.Calendarday, x.TruthValue, x.title }
  x.cl = []uint {lenName, lenFirstName, x.Calendarday.Codelen(),
                 x.TruthValue.Codelen(), x.title.Codelen()}
  x.Format = LongB
  return x
}

func (x *person) imp(Y Any) *person {
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
  x.TruthValue.Clr()
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

func (x *person) Copy (Y Any) {
  y := x.imp (Y)
  x.surname.Copy (y.surname)
  x.firstName.Copy (y.firstName)
  x.TruthValue.Copy (y.TruthValue)
  x.Calendarday.Copy (y.Calendarday)
  x.title.Copy (y.title)
  x.Format = y.Format
}

func (x *person) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *person) Eq (Y Any) bool {
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
             x.TruthValue.Eq (y.TruthValue)
  }
  return g
}

func (x *person) Equiv (Y Person) bool {
  y := x.imp (Y)
  if actualOrder == nameOrder {
    return x.surname.Eq (y.surname) &&
           x.firstName.Eq (y.firstName) &&
           x.Calendarday.Eq (y.Calendarday)
  }
  return x.Calendarday.Eq (y.Calendarday)
}

func (x *person) Less (Y Any) bool {
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

func (x *person) Sub (Y Person) bool {
  y := x.imp (Y)
  if ! x.surname.Sub (y.surname) {
    return false
  }
  if ! x.firstName.Sub (y.firstName) {
    return false
  }
  if ! x.Calendarday.Empty() && y.Calendarday.Empty() {
    return false
  }
  if ! x.Calendarday.Eq (y.Calendarday) {
    return false
  }
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
  shbx.Colours (f, b)
  x.TruthValue.Colours (f, b)
  x.Calendarday.Colours (f, b)
  x.title.Colours (f, b)
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
 Anr.: ___________________________      m/w: _

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

func (x *person) String() string {
// TODO
  n, f := x.surname.String(), x.firstName.String()
  str.OffSpc (&n); str.OffSpc (&f)
  nf, b := n + ", " + f, x.Calendarday.String()
  switch x.Format {
  case Short:
    return nf
  case ShortB:
    return nf + " (" + b + ")"
/* mit Maske:
  case Long:    // Name, Vorname, m/w      1 line, 64 columns
    return ""
  case LongB:   // lang, GebDat            1 line, 80 columns
    return ""
  case LongT:   // lang, Anrede            2 line, 64 columns
    return ""
  case LongTB:  // lang, GebDat, Anrede   2 lines, 80 columns
    return ""
*/
  }
  return nf
}

func (x *person) Defined (s string) bool {
// TODO
  if ! x.surname.Defined (s[:26]) { return false }
  if ! x.firstName.Defined (s[26:41]) { return false }
  if ! x.TruthValue.Defined (s[41:42]) { return false }
  if ! x.Calendarday.Defined (s[42:50]) { return false }
//  if ! x.title.Defined (s[49:]) { return false }
  return true
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
    x.TruthValue.Write (l + 1, c + cg)
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
        x.TruthValue.Edit (l + 1, c + cg)
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
  x.surname.SetFont (f)
  x.firstName.SetFont (f)
  x.TruthValue.SetFont (f)
  x.Calendarday.SetFont (f)
  x.title.SetFont (f)
}

func (x *person) printMask (l, c uint) {
  switch x.Format {
  case Short, ShortB:
    cn = 0; cv = 28; cb = 44
  default:
    cn = 6; cv = 42; cg = 63; cb = 71; ca = cn
  }
  switch x.Format {
  case Short:
    pbx.Print (",", l, c + cv - 2)
    return
  case ShortB:
    pbx.Print (",", l, c + cv - 2)
    pbx.Print ("(", l, c + cb - 1)
    pbx.Print (")", l, c + cb + 8)
    return
  default:
    pbx.Print ("Name:", l, c + cn - 6)
    pbx.Print ("Vorname:", l, c + cv - 9)
    pbx.Print ("u/m/w:", l, c + cg - 5)
  }
  switch x.Format {
  case LongB, LongTB:
    pbx.Print ("geb.:", l, c + cb - 6)
  }
  switch x.Format {
  case LongTB:
    pbx.Print ("Anr.:", l + 1, c + ca - 6)
  }
}

func (x *person) Print (l, c uint) {
  x.printMask (l, c)
  x.surname.SetFont (font.Bold)
  x.surname.Print (l, c + cn)
  x.firstName.SetFont (font.Bold)
  x.firstName.Print (l, c + cv)
  switch x.Format {
  case Short, ShortB:
  default:
    x.TruthValue.Print (l, c + cg)
  }
  switch x.Format {
  case ShortB, LongB, LongTB:
    x.Calendarday.Print (l, c + cb)
  }
  switch x.Format {
  case LongTB:
    x.title.Print (l + 1, c + ca)
  }
}

func (x *person) Codelen() uint {
  return lenName + lenFirstName +
         x.Calendarday.Codelen() +
         x.TruthValue.Codelen() +
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
  x.TruthValue = x.field[3].(tval.TruthValue)
  x.title = x.field[4].(text.Text)
}

func (x *person) Rotate() {
  actualOrder = 1 - actualOrder
}
