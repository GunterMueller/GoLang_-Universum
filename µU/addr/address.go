package addr

// (c) Christian Maurer   v. 211126 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/str"
  "µU/text"
  "µU/bn"
  "µU/phone"
  "µU/cntry"
)
const (
  lenStreet = uint(28)
  lenCity   = uint(22)
  lenEmail  = uint(40)
)
type
  address struct {
          street text.Text
                 bn.Natural "postcode"
            city text.Text
     phonenumber,
      cellnumber,
       faxnumber phone.PhoneNumber
           email text.Text
                 cntry.Country
                 }
var (
  bx = box.New()
  pbx = pbox.New()
  cF, cB = col.LightCyan(), col.Black()
)

func new_() Address {
  x := new(address)
  x.street, x.city = text.New (lenStreet), text.New (lenCity)
  x.Natural = bn.New (5)
  x.phonenumber, x.cellnumber, x.faxnumber = phone.New(), phone.New(), phone.New()
  x.email = text.New (lenEmail)
  x.Country = cntry.New()
  x.Country.SetFormat (cntry.Long)
  x.Colours (cF, cB)
  return x
}

func (x *address) imp(Y Any) *address {
  y, ok := Y.(*address)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *address) Empty() bool {
  return x.street.Empty() && x.Natural.Empty() && x.city.Empty() &&
         x.phonenumber.Empty() && x.cellnumber.Empty() && x.faxnumber.Empty() &&
         x.email.Empty() &&
         x.Country.Empty()
}

func (x *address) Clr() {
  x.street.Clr(); x.Natural.Clr(); x.city.Clr()
  x.phonenumber.Clr(); x.cellnumber.Clr(); x.faxnumber.Clr()
  x.email.Clr()
  x.Country.Clr()
}

func (x *address) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *address) Copy (Y Any) {
  y := x.imp (Y)
  x.street.Copy (y.street)
  x.Natural.Copy (y.Natural)
  x.city.Copy (y.city)
  x.phonenumber.Copy (y.phonenumber)
  x.cellnumber.Copy (y.cellnumber)
  x.faxnumber.Copy (y.faxnumber)
  x.email.Copy (y.email)
  x.Country.Copy (y.Country)
}

func (x *address) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.street.Eq (y.street) && x.Natural.Eq (y.Natural) && x.city.Eq (y.city) &&
         x.phonenumber.Eq (y.phonenumber) && x.cellnumber.Eq (y.cellnumber) &&
         x.faxnumber.Eq (y.faxnumber) &&
         x.email.Eq (y.email) &&
         x.Country.Eq (y.Country)
}

func (x *address) Equiv (Y Any) bool {
  if x.Natural.Eq (x.imp (Y).Natural) {
    return true
  }
  return false
}

func (x *address) Less (Y Any) bool {
  y := x.imp (Y)
  if x.Natural.Eq (y.Natural) {
    if x.city.Eq (y.city) {
      return x.street.Less (y.street)
    } else {
      return x.city.Less (y.city)
    }
  }
  return x.Natural.Less (y.Natural)
}

func (x *address) Colours (f, b col.Colour) {
  x.street.Colours (f, b)
  x.Natural.Colours (f, b)
  x.city.Colours (f, b)
  x.phonenumber.Colours (f, b)
  x.cellnumber.Colours (f, b)
  x.faxnumber.Colours (f, b)
  x.email.Colours (f, b)
  x.Country.Colours (f, b)
}

const (
  cs = 10; cp = 45; cc = 57; ct = cs; cf = 34; cx = cc; cm = cs; cl = cc
)

func writeMask (l, c uint) {
//           1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789
//  Str./Nr: ____________________________  PLZ: _____  Ort: ______________________
//     Tel.: ________________  Funk: ________________  Fax: ________________
//   E-Mail: ________________________________________ Land: ______________________
  bx.Wd (8)
  bx.Write ("Str./Nr:", l, c + 1)
  bx.Wd (4)
  bx.Write ("PLZ:",     l, c + cp - 5)
  bx.Write ("Ort:",     l, c + cc - 5)
  bx.Wd (5)
  bx.Write ("Tel.:",    l + 1, c + ct - 6)
  bx.Write ("Funk:",    l + 1, c + cf - 6)
  bx.Wd (4)
  bx.Write ("Fax:",     l + 1, c + cx - 5)
  bx.Wd (7)
  bx.Write ("E-Mail:",  l + 2, c + cm - 8)
  bx.Wd (5)
  bx.Write ("Tel.:",    l + 1, c + ct - 6)
  bx.Write ("Land:",    l + 2, c + cc - 6)
}

func (x *address) Write (l, c uint) {
  writeMask (l, c)
  x.street.Write (l, c + cs)
  x.Natural.Write (l, c + cp)
  x.city.Write (l, c + cc)
  x.phonenumber.Write (l + 1, c + ct)
  x.cellnumber.Write (l + 1, c + cf)
  x.faxnumber.Write (l + 1, c + cx)
  x.email.Write (l + 2, c + cm)
  x.Country.Write (l + 2, c + cl)
}

func (x *address) Edit (l, c uint) {
  const n = 7
  x.Write (l, c)
  i := 0
  if C, _ := kbd.LastCommand(); C == kbd.Up {
    i = n
  }
  loop:
  for {
    switch i {
    case 0:
      x.street.Edit (l, c + cs)
    case 1:
      x.Natural.Edit (l, c + cp)
    case 2:
      x.city.Edit (l, c + cc)
    case 3:
      x.phonenumber.Edit (l + 1, c + ct)
    case 4:
      x.cellnumber.Edit (l + 1, c + cf)
    case 5:
      x.faxnumber.Edit (l + 1, c + cx)
    case 6:
      x.email.Edit (l + 2, c + cm)
    case 7:
      x.Country.Edit (l + 2, c + cl)
    }
    switch C, d:= kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < n { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < n { i++ } else { break loop }
    case kbd.Up, kbd.Left:
      if i > 0 { i-- } else { break loop }
    }
  }
}

func (x *address) SetFont (f font.Font) {
  x.street.SetFont (f)
  x.Natural.SetFont (f)
  x.city.SetFont (f)
  x.phonenumber.SetFont (f)
  x.cellnumber.SetFont (f)
  x.faxnumber.SetFont (f)
  x.email.SetFont (f)
  x.Country.SetFont (f)
}

func (x *address) Print (l, c uint) {
// printMask()
  x.street.Print (l, c + cs)
  pbx.Print ("Tel.:", l, c + ct - 6)
  x.Natural.Print (l, c + cp)
  x.city.Print (l, c + cf)
  pbx.Print ("Funk:", l + 1, c + ct - 6)
  x.phonenumber.Print (l, c + ct)
  x.cellnumber.Print (l + 1, c + ct)
  pbx.Print ("Fax:", l + 1, c + cx - 6)
  x.faxnumber.Print (l + 1, c + cx)
  pbx.Print ("E-Mail:", l + 2, c + cm - 8)
  x.email.Print (l + 2, c + cm)
}

func (x *address) Codelen() uint {
  return lenStreet +
         x.Natural.Codelen() +
         lenCity +
         3 * x.phonenumber.Codelen() +
         lenEmail +
         x.Country.Codelen()
}

func (x *address) Encode() Stream {
  s := make (Stream, x.Codelen())
  i, a := uint(0), lenStreet
  copy (s[i:i+a], x.street.Encode())
  i += a
  a = x.Natural.Codelen()
  copy (s[i:i+a], x.Natural.Encode())
  i += a
  a = lenCity
  copy (s[i:i+a], x.city.Encode())
  i += a
  a = x.phonenumber.Codelen()
  copy (s[i:i+a], x.phonenumber.Encode())
  i += a
  copy (s[i:i+a], x.cellnumber.Encode())
  i += a
  copy (s[i:i+a], x.faxnumber.Encode())
  i += a
  a = lenEmail
  copy (s[i:i+a], x.email.Encode())
  i += a
  a = x.Country.Codelen()
  copy (s[i:i+a], x.Country.Encode())
  return s
}

func (x *address) Decode (s Stream) {
  i, a := uint(0), lenStreet
  x.street = Decode (x.street, s[i:i+a]).(text.Text)
  i += a
  a = x.Natural.Codelen()
  x.Natural = Decode (x.Natural, s[i:i+a]).(bn.Natural)
  i += a
  a = lenCity
  x.city = Decode (x.city, s[i:i+a]).(text.Text)
  i += a
  a = x.phonenumber.Codelen()
  x.phonenumber = Decode (x.phonenumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  x.cellnumber = Decode (x.cellnumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  x.faxnumber = Decode (x.cellnumber, s[i:i+a]).(phone.PhoneNumber)
  i += a
  a = lenEmail
  x.email = Decode (x.email, s[i:i+a]).(text.Text)
  i += a
  a = x.Country.Codelen()
  x.Country = Decode (x.Country, s[i:i+a]).(cntry.Country)
}

func (x *address) TeX() string {
  s := x.street.TeX() + ", " + x.Natural.String() + " " + x.city.TeX()
  c := x.Country.Clone().(cntry.Country)
  if ! c.Empty() { s += " (" + c.TeX() + ")" }
  if x.phonenumber.Empty() {
    if ! x.cellnumber.Empty() {
      s += "\\newline\nTel. " + x.cellnumber.TeX() + ")"
    }
  } else {
    s += "\\newline\nTel. " + x.phonenumber.TeX()
    if ! x.cellnumber.Empty() {
      s += ", " + x.cellnumber.TeX()
    }
  }
  if ! x.faxnumber.Empty() {
    s += ", Fax " + x.faxnumber.TeX()
  }
  if ! x.email.Empty() {
    em := x.email.TeX()
    str.ReplaceAll (&em, '_', "\\_")
    s += "\\newline\nE-Mail: {\\tte " + em + "}"
  }
  return s
}
