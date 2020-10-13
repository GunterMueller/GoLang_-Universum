package pers

<<<<<<< HEAD
// (c) Christian Maurer   v. 201004 - license see µU.go
=======
// (c) Christian Maurer   v. 200908 - license see µU.go
>>>>>>> a13d69ba2d9c50112f2390abda13b4352cfd3a84

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
  "µU/enum"
  "µU/day"
)
const (
  lenName = uint(26)
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
                enum.Enumerator "title"
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
  x.Enumerator = enum.New (enum.Title)
  x.field = []Any { x.surname, x.firstName, x.Calendarday, x.TruthValue, x.Enumerator }
  x.cl = []uint {lenName, lenFirstName, x.Calendarday.Codelen(),
                 x.TruthValue.Codelen(), x.Enumerator.Codelen()}
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
  x.Enumerator.Clr()
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
  x.Enumerator.Copy (y.Enumerator)
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
       x.firstName.Eq (y.firstName) &&
       x.Calendarday.Eq (y.Calendarday)
  switch x.Format {
  case LongB, LongTB:
    g = g && x.TruthValue.Eq (y.TruthValue)
  }
  switch x.Format {
  case LongT, LongTB:
    g = g && x.Enumerator.Eq (y.Enumerator)
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
  x.Enumerator.Colours (f, b)
}

var
  csn, cfn, cmf, cbd, can uint
/* without mask:
          1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789
VeryShort    Name, Vorname                                   1 line, 44 columns
__________________________, _______________

Name: __________________________ Vorname: _______________ m/w: _ geb.: ________
Anr.: _____________ 

Short, ShortB: Name, Vorname (GebDat)                   1 line, 43 (54) columns
__________________________, _______________ (________)

ShortT, ShortTB: Name, Vorname, Anrede (GebDat)
__________________________, _______________, _____________ (________)

with mask:
Long        Name, Vorname, m/w     1 line,  64 columns
LongB       Lang, GebDat           1 line,  80 columns
LongT       Lang, Anrede           2 lines, 64 columns

LongT, LongTB: Name, Vorname, m/w, (geb)               2 lines, 64 (79) columns
Name: __________________________ Vorname: _______________ m/w: _ geb.: ________
Anr.: _____________
******************************************************************************/

func (x *person) writeMask (l, c uint) {
  switch x.Format {
  case Short, ShortB:
    csn = 0; cfn = 28; cbd = 44
  default:
    csn = 6; cfn = 42; cmf = 63; cbd = 71; can = csn
  }
  bx.Wd (1)
  bx.ScrColours()
  switch x.Format {
  case Short:
    bx.Write (",", l, c + cfn - 2)
    return
  case ShortB:
    bx.Write (",", l, c + cfn - 2)
    bx.Write ("(", l, c + cbd - 1)
    bx.Write (")", l, c + cbd + 8)
    return
  default:
    bx.Wd (5)
    bx.Write ("Name:", l, c + csn - 6)
    bx.Wd (8)
    bx.Write ("Vorname:", l, c + cfn - 9)
    bx.Wd (4)
    bx.Write ("m/w:", l, c + cmf - 5)
    bx.Wd (5)
  }
  switch x.Format {
  case LongB, LongTB:
    bx.Wd (5)
    bx.Write ("geb.:", l, c + cbd - 6)
  }
  switch x.Format {
  case LongT, LongTB:
    bx.Wd (5)
    bx.Write ("Anr.:", l + 1, c + can - 6)
  }
}

func (x *person) String() string {
  n, f := x.surname.String(), x.firstName.String()
  str.OffSpc (&n); str.OffSpc (&f)
  nf, fn, b := n + ", " + f, f + " " + n, x.Calendarday.String()
  switch x.Format {
  case VeryShort, Short:
    return nf
  case ShortB:
    return nf + " (" + b + ")"
  case ShortT:
    if ! x.Enumerator.Empty() { nf = x.Enumerator.String() + " " + nf }
    return nf
  case ShortTB:
    if ! x.Enumerator.Empty() { fn = x.Enumerator.String() + " " + fn }
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

func (x *person) Defined (s string) bool { // trivial version, better one TODO
  if ! x.surname.Defined (s[:26]) { return false }
  if ! x.firstName.Defined (s[26:41]) { return false }
  if ! x.TruthValue.Defined (s[41:42]) { return false }
  if ! x.Calendarday.Defined (s[42:50]) { return false }
//  if ! x.Enumerator.Defined (s[49:]) { return false }
  return true
}

func (x *person) Write (l, c uint) {
  if x.Format == VeryShort {
    shbx.Write (x.String(), l, c)
    return
  }
  x.writeMask (l, c)
  x.surname.Write (l, c + csn)
  x.firstName.Write (l, c + cfn)
  switch x.Format {
  case Short, ShortB:
  default:
    x.TruthValue.Write (l, c + cmf)
  }
  switch x.Format {
  case ShortB, LongB, LongTB:
    x.Calendarday.Write (l, c + cbd)
  }
  switch x.Format {
  case LongT, LongTB:
    x.Enumerator.Write (l + 1, c + can)
  }
}

func (x *person) Edit (l, c uint) {
  x.Write (l, c)
  if x.Format == VeryShort { return }
  i := uint(0)
  if C, _ := kbd.LastCommand(); C == kbd.Up {
    i = 4
  }
<<<<<<< HEAD
  loop:
  for {
=======
  loop: for {
>>>>>>> a13d69ba2d9c50112f2390abda13b4352cfd3a84
    switch i {
    case 0:
      x.surname.Edit (l, c + csn)
    case 1:
      x.firstName.Edit (l, c + cfn)
    case 2:
      switch x.Format {
      case Short, ShortB:
        ;
      default:
        x.TruthValue.Edit (l, c + cmf)
      }
    case 3:
      switch x.Format {
      case ShortB, LongB, LongTB:
        x.Calendarday.Edit (l, c + cbd)
      }
    case 4:
      switch x.Format {
      case LongT, LongTB:
        x.Enumerator.Edit (l + 1, c + can)
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
    case kbd.Tab:
      if d == 0 {
        if i < 4 {
          i++
        }
      } else {
        if i > 0 {
          i--
        }
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
  x.Enumerator.SetFont (f)
}

func (x *person) printMask (l, c uint) {
  switch x.Format {
  case Short, ShortB:
    csn = 0; cfn = 28; cbd = 44
  default:
    csn = 6; cfn = 42; cmf = 63; cbd = 71; can = csn
  }
  switch x.Format {
  case Short:
    pbx.Print (",", l, c + cfn - 2)
    return
  case ShortB:
    pbx.Print (",", l, c + cfn - 2)
    pbx.Print ("(", l, c + cbd - 1)
    pbx.Print (")", l, c + cbd + 8)
    return
  default:
    pbx.Print ("Name:", l, c + csn - 6)
    pbx.Print ("Vorname:", l, c + cfn - 9)
    pbx.Print ("u/m/w:", l, c + cmf - 5)
  }
  switch x.Format {
  case LongB, LongTB:
    pbx.Print ("geb.:", l, c + cbd - 6)
  }
  switch x.Format {
  case LongT, LongTB:
    pbx.Print ("Anr.:", l + 1, c + can - 6)
  }
}

func (x *person) Print (l, c uint) {
  x.printMask (l, c)
  if x.Format == VeryShort {
    pbx.Print (x.String(), l, c)
    return
  }
  x.surname.SetFont (font.Bold)
  x.surname.Print (l, c + csn)
  x.firstName.SetFont (font.Bold)
  x.firstName.Print (l, c + cfn)
  switch x.Format {
  case Short, ShortB:
  default:
    x.TruthValue.Print (l, c + cmf)
  }
  switch x.Format {
  case ShortB, LongB, LongTB:
    x.Calendarday.Print (l, c + cbd)
  }
  switch x.Format {
  case LongT, LongTB:
    x.Enumerator.Print (l + 1, c + can)
  }
}

func (x *person) Codelen() uint {
/*
  return lenName +                 // 26
         lenFirstName +            // 15
         x.Calendarday.Codelen() + //  2
         x.TruthValue.Codelen() +  //  1
         x.Enumerator.Codelen()    //  1
*/
  return                              45
}

func (x *person) Encode() []byte {
  return Encodes (x.field, x.cl)
}

func (x *person) Decode (bs []byte) {
  Decodes (bs, x.field, x.cl)
  x.surname = x.field[0].(text.Text)
  x.firstName = x.field[1].(text.Text)
  x.Calendarday = x.field[2].(day.Calendarday)
  x.TruthValue = x.field[3].(tval.TruthValue)
  x.Enumerator = x.field[4].(enum.Enumerator)
}

func (x *person) RotOrder() {
  actualOrder = 1 - actualOrder
}
