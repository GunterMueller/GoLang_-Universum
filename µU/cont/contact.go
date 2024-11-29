package cont

// (c) Christian Maurer   v. 240427 - license see µU.go

import (
  "µU/env"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/str"
  "µU/text"
  "µU/phone"
)
const (
  lenEmail  = uint(40)
  sep = ','
  seps = ","
)
type
  contact struct {
     phonenumber,
      cellnumber phone.PhoneNumber
           email text.Text
                 }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_() Contact {
  x := new(contact)
  f, b := col.FlashWhite(), col.DarkBlue()
  x.phonenumber = phone.New()
  x.phonenumber.Colours (f, b)
  x.cellnumber = phone.New()
  x.cellnumber.Colours (f, b)
  x.email = text.New (lenEmail)
  x.email.Colours (f, b)
  return x
}

func (x *contact) imp(Y any) *contact {
  y, ok := Y.(*contact)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *contact) Empty() bool {
  return x.phonenumber.Empty() &&
         x.cellnumber.Empty() &&
         x.email.Empty()
}

func (x *contact) Clr() {
  x.phonenumber.Clr()
  x.cellnumber.Clr()
  x.email.Clr()
}

func (x *contact) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *contact) Copy (Y any) {
  y := x.imp (Y)
  x.phonenumber.Copy (y.phonenumber)
  x.cellnumber.Copy (y.cellnumber)
  x.email.Copy (y.email)
}

func (x *contact) Eq (Y any) bool {
  y := x.imp (Y)
  return x.phonenumber.Eq (y.phonenumber) &&
         x.cellnumber.Eq (y.cellnumber) &&
         x.email.Eq (y.email)
}

func (x *contact) Less (Y any) bool {
  return false
}

func (x *contact) Leq (Y any) bool {
  return false
}

func (x *contact) String() string {
  s := x.phonenumber.String()
  str.OffSpc1 (&s)
  s += seps
  t := x.cellnumber.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.email.String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *contact) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != N { println("n"); return false }
  if ! x.phonenumber.Defined (ss[0]) { return false }
  if ! x.cellnumber.Defined (ss[1]) { return false }
  email := ss[2]
  if ! str.Empty (email) {
    if _, ok := str.Sub ("@", email); ! ok { return false }
  }
  if ! x.email.Defined (email) { return false }
  return true
}

func (x *contact) Colours (f, b col.Colour) {
  x.phonenumber.Colours (f, b)
  x.cellnumber.Colours (f, b)
  x.email.Colours (f, b)
}

func (x *contact) Cols() (col.Colour, col.Colour) {
  return x.phonenumber.Cols()
}

const (
  ct = 8; cf = 32
)

func writeMask (l, c uint) {
/*        1         2         3         4
01234567890123456789012345678901234567890123456789
 phone: ________________  cell: ________________ if env.E()
  Tel.: ________________  Funk: ________________
E-Mail: ________________________________________
*/
  bx.Wd (6)
  if env.E() {
    bx.Write ("phone:",  l, c + ct - 7)
  } else {
    bx.Write (" Tel.:",  l, c + ct - 7)
  }
  bx.Wd (5)
  if env.E() {
    bx.Write ("cell:",   l, c + cf - 6)
  } else {
    bx.Write ("Funk:",   l, c + cf - 6)
  }
  bx.Wd (7)
  bx.Write ("E-Mail:", l + 1, c + ct - 8)
}

func (x *contact) Write (l, c uint) {
  writeMask (l, c)
  x.phonenumber.Write (l, c + ct)
  x.cellnumber.Write (l, c + cf)
  x.email.Write (l + 1, c + ct)
}

func (x *contact) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  if C, _ := kbd.LastCommand(); C == kbd.Up {
    i = N - 1
  }
  loop:
  for {
    switch i {
    case 0:
      x.phonenumber.Edit (l, c + ct)
    case 1:
      x.cellnumber.Edit (l, c + cf)
    case 2:
      x.email.Edit (l + 1, c + ct)
    }
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < N - 1 { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < N - 1 { i++ } else { break loop }
    case kbd.Up, kbd.Left:
      if i > 0 { i-- } else { break loop }
    }
  }
}

func (x *contact) SetFont (f font.Font) {
  x.phonenumber.SetFont (f)
  x.cellnumber.SetFont (f)
  x.email.SetFont (f)
}

func (x *contact) Print (l, c uint) {
  if env.E() {
    pbx.Print ("phone:", l, c + ct - 6)
    pbx.Print ("cell:", l + 1, c + ct - 6)
  } else {
    pbx.Print ("Tel.:", l, c + ct - 6)
    pbx.Print ("Funk:", l + 1, c + ct - 6)
  }
  pbx.Print ("E-Mail:", l + 2, c + ct - 8)
  x.phonenumber.Print (l, c + ct)
  x.cellnumber.Print (l + 1, c + ct)
  x.email.Print (l + 2, c + ct)
}

func (x *contact) Codelen() uint {
  return 2 * x.phonenumber.Codelen() +
         lenEmail
}

func (x *contact) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint(0), x.phonenumber.Codelen()
  copy (s[i:i+a], x.phonenumber.Encode())
  i += a
  copy (s[i:i+a], x.cellnumber.Encode())
  i += a
  a = lenEmail
  copy (s[i:i+a], x.email.Encode())
  return s
}

func (x *contact) Decode (s Stream) {
  i, a := uint(0), x.phonenumber.Codelen()
  x.phonenumber = Decode (x.phonenumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  x.cellnumber = Decode (x.cellnumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  a = lenEmail
  x.email = Decode (x.email, s[i:i+a]).(text.Text)
}

func (x *contact) TeX() string {
  p, c := ! x.phonenumber.Empty(), ! x.cellnumber.Empty()
  var s string
  if p {
    if env.E() {
      s = "phone " + x.phonenumber.TeX()
    } else {
      s = "Tel.~" + x.phonenumber.TeX()
    }
  }
  if c {
    if p { s += ", " }
    if env.E() {
      s += "cell " + x.cellnumber.TeX()
    } else {
      s += "Funk " + x.cellnumber.TeX()
    }
  }
  if ! x.email.Empty() {
    em := x.email.TeX()
    str.ReplaceAll (&em, '_', "\\_")
    if p || c {s += ", " }
    s += "E-Mail: {\\tt " + em + "}"

  }
  return s
}
