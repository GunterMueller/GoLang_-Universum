package telmail

// (c) Christian Maurer   v. 240407 - license see µU.go

import (
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
  telmail struct {
     phonenumber,
      cellnumber phone.PhoneNumber
           email text.Text
                 }
var (
  bx = box.New()
  pbx = pbox.New()
)

func new_() TelMail {
  x := new(telmail)
  x.phonenumber = phone.New()
  x.cellnumber = phone.New()
  x.email = text.New (lenEmail)
  x.Colours (col.LightCyan(), col.Black())
  return x
}

func (x *telmail) imp(Y any) *telmail {
  y, ok := Y.(*telmail)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *telmail) Empty() bool {
  return x.phonenumber.Empty() &&
         x.cellnumber.Empty() &&
         x.email.Empty()
}

func (x *telmail) Clr() {
  x.phonenumber.Clr()
  x.cellnumber.Clr()
  x.email.Clr()
}

func (x *telmail) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *telmail) Copy (Y any) {
  y := x.imp (Y)
  x.phonenumber.Copy (y.phonenumber)
  x.cellnumber.Copy (y.cellnumber)
  x.email.Copy (y.email)
}

func (x *telmail) Eq (Y any) bool {
  y := x.imp (Y)
  return x.phonenumber.Eq (y.phonenumber) &&
         x.cellnumber.Eq (y.cellnumber) &&
         x.email.Eq (y.email)
}

func (x *telmail) Less (Y any) bool {
  return false
}

func (x *telmail) Leq (Y any) bool {
  return false
}

func (x *telmail) String() string {
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

func (x *telmail) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != N { return false }
  if ! x.phonenumber.Defined (ss[0]) { return false }
  if ! x.cellnumber.Defined (ss[1]) { return false }
  if ! x.email.Defined (ss[2]) { return false } // TODO check @
  return true
}

func (x *telmail) Colours (f, b col.Colour) {
  x.phonenumber.Colours (f, b)
  x.cellnumber.Colours (f, b)
  x.email.Colours (f, b)
}

func (x *telmail) Cols() (col.Colour, col.Colour) {
  return x.phonenumber.Cols()
}

const (
  ct = 10; cf = 34; cx = 57
)

func writeMask (l, c uint) {
//           1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789
//     Tel.: ________________  Funk: ________________
//   E-Mail: ________________________________________
  bx.Wd (5)
  bx.Write ("Tel.:",   l, c + ct - 6)
  bx.Write ("Funk:",   l, c + cf - 6)
  bx.Wd (7)
  bx.Write ("E-Mail:", l + 1, c + ct - 8)
}

func (x *telmail) Write (l, c uint) {
  writeMask (l, c)
  x.phonenumber.Write (l, c + ct)
  x.cellnumber.Write (l, c + cf)
  x.email.Write (l + 1, c + ct)
}

func (x *telmail) Edit (l, c uint) {
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

func (x *telmail) SetFont (f font.Font) {
  x.phonenumber.SetFont (f)
  x.cellnumber.SetFont (f)
  x.email.SetFont (f)
}

func (x *telmail) Print (l, c uint) {
// printMask()
  pbx.Print ("Tel.:", l, c + ct - 6)
  pbx.Print ("Funk:", l + 1, c + ct - 6)
  x.phonenumber.Print (l, c + ct)
  x.cellnumber.Print (l + 1, c + ct)
  pbx.Print ("E-Mail:", l + 2, c + ct - 8)
  x.email.Print (l + 2, c + ct)
}

func (x *telmail) Codelen() uint {
  return 2 * x.phonenumber.Codelen() +
         lenEmail
}

func (x *telmail) Encode() Stream {
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

func (x *telmail) Decode (s Stream) {
  i, a := uint(0), x.phonenumber.Codelen()
  x.phonenumber = Decode (x.phonenumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  x.cellnumber = Decode (x.cellnumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  a = lenEmail
  x.email = Decode (x.email, s[i:i+a]).(text.Text)
}

func (x *telmail) TeX() string {
  p, c := ! x.phonenumber.Empty(), ! x.cellnumber.Empty()
  var s string
  if p {
    s = "Tel.~" + x.phonenumber.TeX()
  }
  if c {
    if p { s += ", " }
    s += "Funk " + x.cellnumber.TeX()
  }
  if ! x.email.Empty() {
    em := x.email.TeX()
    str.ReplaceAll (&em, '_', "\\_")
    if p || c {s += ", " }
    s += "E-Mail: {\\tt " + em + "}"

  }
  return s
}
