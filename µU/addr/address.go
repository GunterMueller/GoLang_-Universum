package addr

// (c) Christian Maurer   v. 210410 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/str"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/masks"
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
      cellnumber phone.PhoneNumber
           email text.Text
                 cntry.Country
                 }
var (
  bx = box.New()
  pbx = pbox.New()
  cF, cB = col.LightCyan(), col.Black()
  mask masks.MaskSet = masks.New()
  cst, cpc, cci, cph, cce, cem, cco uint
)

func new_() Address {
  x := new(address)
  x.street, x.city = text.New (lenStreet), text.New (lenCity)
  x.Natural = bn.New (5)
  x.phonenumber, x.cellnumber = phone.New(), phone.New()
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
  return x.street.Empty() &&
         x.Natural.Empty() &&
         x.city.Empty() &&
         x.phonenumber.Empty() &&
         x.cellnumber.Empty() &&
         x.email.Empty() &&
         x.Country.Empty()
}

func (x *address) Clr() {
  x.street.Clr()
  x.Natural.Clr()
  x.city.Clr()
  x.phonenumber.Clr()
  x.cellnumber.Clr()
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
  x.email.Copy (y.email)
  x.Country.Copy (y.Country)
}

func (x *address) Eq (Y Any) bool {
  y := x.imp (Y)
  return x.street.Eq (y.street) &&
         x.Natural.Eq (y.Natural) &&
         x.city.Eq (y.city) &&
         x.phonenumber.Eq (y.phonenumber) &&
         x.cellnumber.Eq (y.cellnumber) &&
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
  x.email.Colours (f, b)
  x.Country.Colours (f, b)
}

func (x *address) Write (l, c uint) {
  mask.Write (l, c)
  x.street.Write (l, c + cst)
  x.Natural.Write (l, c + cpc)
  x.city.Write (l, c + cci)
  x.phonenumber.Write (l + 1, c + cph)
  x.cellnumber.Write (l + 1, c + cce)
  x.email.Write (l + 2, c + cem)
  x.Country.Write (l + 2, c + cco)
}

func (x *address) Edit (l, c uint) {
  const n = 6
  x.Write (l, c)
  i := 0
  if C, _:= kbd.LastCommand(); C == kbd.Up {
    i = n
  }
  loop: for {
    switch i {
    case 0:
      x.street.Edit (l, c + cst)
    case 1:
      x.Natural.Edit (l, c + cpc)
    case 2:
      x.city.Edit (l, c + cci)
    case 3:
      x.phonenumber.Edit (l + 1, c + cph)
    case 4:
      x.cellnumber.Edit (l + 1, c + cce)
    case 5:
      x.email.Edit (l + 2, c + cem)
    case 6:
      x.Country.Edit (l + 2, c + cco)
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
  x.email.SetFont (f)
  x.Country.SetFont (f)
}

func (x *address) Print (l, c uint) {
  mask.Print (l, c)
  x.street.Print (l, c + cst)
  pbx.Print ("Tel.:", l, c + cph - 6)
  x.Natural.Print (l, c + cpc)
  x.city.Print (l, c + cci)
  pbx.Print ("Funk:", l + 1, c + cph - 5)
  x.phonenumber.Print (l, c + cph)
  x.cellnumber.Print (l + 1, c + cph)
  pbx.Print ("E-Mail:", l + 2, c + cem - 7)
  x.email.Print (l + 2, c + cem)
}

func (x *address) Codelen() uint {
/*
  return lenStreet +                   // 28
         x.Natural.Codelen() +         //  8
         lenCity +                     // 22
         2 * x.phonenumber.Codelen() + // 12
         lenEmail +                    // 40
         x.Country                     //  2
*/
  return                                 112
}

func (x *address) Encode() Stream {
  b:= make (Stream, x.Codelen())
  i, a:= uint(0), lenStreet
  copy (b[i:i+a], x.street.Encode())
  i += a
  a = x.Natural.Codelen()
  copy (b[i:i+a], x.Natural.Encode())
  i += a
  a = lenCity
  copy (b[i:i+a], x.city.Encode())
  i += a
  a = x.phonenumber.Codelen()
  copy (b[i:i+a], x.phonenumber.Encode())
  i += a
  copy (b[i:i+a], x.cellnumber.Encode())
  i += a
  a = lenEmail
  copy (b[i:i+a], x.email.Encode())
  i += a
  a = x.Country.Codelen()
  copy (b[i:i+a], x.Country.Encode())
  return b
}

func (x *address) Decode (b Stream) {
  i, a:= uint(0), lenStreet
  x.street = Decode (x.street, b[i:i+a]).(text.Text)
  i += a
  a = x.Natural.Codelen()
  x.Natural = Decode (x.Natural, b[i:i+a]).(bn.Natural)
  i += a
  a = lenCity
  x.city = Decode (x.city, b[i:i+a]).(text.Text)
  i += a
  a = x.phonenumber.Codelen()
  x.phonenumber = Decode (x.phonenumber, b[i:i+a]).(phone.PhoneNumber)
  i += a
  x.cellnumber = Decode (x.cellnumber, b[i:i+a]).(phone.PhoneNumber)
  i += a
  a = lenEmail
  x.email = Decode (x.email, b[i:i+a]).(text.Text)
  i += a
  a = x.Country.Codelen()
  x.Country = Decode (x.Country, b[i:i+a]).(cntry.Country)
}

func (x *address) String() string {
  s := x.street.String()
  str.OffSpc1 (&s)
  s += ", "
  s += x.Natural.String()
  s += " "
  t := x.city.String()
  str.OffSpc1 (&t)
  s += t
  s += ", Tel. "
  t = x.phonenumber.String()
  str.OffSpc1 (&t)
  s += t
  s += ", "
  t = x.phonenumber.String()
  str.OffSpc1 (&t)
  s += t
  s += ", "
  t = x.email.String()
  str.OffSpc1 (&t)
  s += t
  s += ", "
  c := Clone (x.Country).(cntry.Country)
  c.SetFormat (cntry.Car)
  s += c.String()
  return s
}

func (x *address) Defined (s string) bool {
  t, n := str.SplitByte (s, ',')
  if n != 6 { return false }
//  s == "Keithstr. 16, 10787 Berlin, 21478429, 01785534563, maurer@maurer, D"
  x.street.Defined (t[0])
  if ! x.Natural.Defined (t[1][:5]) { return false }
  x.city.Defined (t[1][6:])
  if ! x.phonenumber.Defined (t[2]) { return false }
  if ! x.phonenumber.Defined (t[3]) { return false }
  x.email.Defined (t[4])
  x.Country.SetFormat (cntry.Car)
  if ! x.Country.Defined (t[5]) { return false }
  return true
}

func init() {
  cst, cpc, cci, cph, cce, cem, cco = 10, 45, 57, 10, 34, 10, 57
//           1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789
// Str./Nr.: ____________________________  PLZ: _____  Ort: ______________________
//     Tel.: ________________  Funk: ________________
//   E-Mail: ________________________________________ Land: ______________________
  mask.Ins ("Str./Nr.:", 0,  0)
  mask.Ins ("PLZ:",      0, 40)
  mask.Ins ("Ort:",      0, 52)
  mask.Ins ("Tel.:",     1,  4)
  mask.Ins ("Funk:",     1, 28)
  mask.Ins ("E-Mail:",   2,  2)
  mask.Ins ("Land:",     2, 51)
}
