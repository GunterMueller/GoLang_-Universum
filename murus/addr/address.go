package addr

// (c) murus.org  v. 170419 - license see murus.go

import (
  . "murus/obj"
  "murus/kbd"
  "murus/col"
  "murus/box"
  "murus/font"
  "murus/pbox"
  "murus/masks"
  "murus/text"
  "murus/bnat"
  "murus/phone"
)
const (
  lenStreet = uint(28)
  lenCity   = uint(22)
)
type
  address struct {
          street text.Text
                 bnat.Natural "postcode"
            city text.Text
     phonenumber,
      cellnumber phone.PhoneNumber
                 }
var (
  bx = box.New()
  pbx = pbox.New()
  cF, cB = col.LightCyan, col.Black
  mask masks.MaskSequence = masks.New()
//  mask = [2]masks.MaskSequence { masks.New(), masks.New() }
  cst, cpc, cci, cph uint = 10, 5, 16, 45 // TODO parametrize
)

func new_() Address {
  x:= new(address)
  x.street, x.city = text.New (lenStreet), text.New (lenCity)
  x.Natural = bnat.New (5)
  x.phonenumber, x.cellnumber = phone.New(), phone.New()
  x.Colours (cF, cB)
  return x
}

func (x *address) imp(Y Any) *address {
  y, ok:= Y.(*address)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *address) Empty() bool {
  return x.street.Empty() &&
         x.Natural.Empty() &&
         x.city.Empty() &&
         x.phonenumber.Empty() &&
         x.cellnumber.Empty()
}

func (x *address) Clr() {
  x.street.Clr()
  x.Natural.Clr()
  x.city.Clr()
  x.phonenumber.Clr()
  x.cellnumber.Clr()
}

func (x *address) Clone() Any {
  y:= new_()
  y.Copy (x)
  return y
}

func (x *address) Copy (Y Any) {
  y:= x.imp (Y)
  x.street.Copy (y.street)
  x.Natural.Copy (y.Natural)
  x.city.Copy (y.city)
  x.phonenumber.Copy (y.phonenumber)
  x.cellnumber.Copy (y.cellnumber)
}

func (x *address) Eq (Y Any) bool {
  y:= x.imp (Y)
  return x.street.Eq (y.street) &&
         x.Natural.Eq (y.Natural) &&
         x.city.Eq (y.city) &&
         x.phonenumber.Eq (y.phonenumber) &&
         x.cellnumber.Eq (y.cellnumber)
}

func (x *address) Equiv (Y Any) bool {
  if x.Natural.Eq (x.imp (Y).Natural) {
    return true
  }
  return false
}

func (x *address) Less (Y Any) bool {
  y:= x.imp (Y)
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
}

func (x *address) Write (l, c uint) {
  mask.Write (l, c)
  x.street.Write (l, c + cst)
  x.Natural.Write (l + 1, c + cpc)
  x.city.Write (l + 1, c + cci)
  x.phonenumber.Write (l, c + cph)
  x.cellnumber.Write (l + 1, c + cph)
}

func (x *address) Edit (l, c uint) {
  x.Write (l, c)
  i:= 0
  if C, _:= kbd.LastCommand(); C == kbd.Up {
    i = 4
  }
  loop: for {
    switch i { case 0:
      x.street.Edit (l, c + cst)
    case 1:
      x.Natural.Edit (l + 1, c + cpc)
    case 2:
      x.city.Edit (l + 1, c + cci)
    case 3:
      x.phonenumber.Edit (l, c + cph)
    case 4:
      x.cellnumber.Edit (l + 1, c + cph)
    }
    switch C, d:= kbd.LastCommand(); C {
    case kbd.Esc:
      break loop
    case kbd.Enter:
      if d == 0 {
        if i < 4 { i++ } else { break loop }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < 4 { i ++ } else { break loop }
    case kbd.Up, kbd.Left:
      if i > 0 { i -- } else { break loop }
    }
  }
}


func (x *address) SetFont (f font.Font) {
  x.street.SetFont (f)
  x.Natural.SetFont (f)
  x.city.SetFont (f)
  x.phonenumber.SetFont (f)
  x.cellnumber.SetFont (f)
}

func (x *address) Print (l, c uint) {
  mask.Print (l, c)
  x.street.Print (l, c + cst)
  pbx.Print ("Tel.:", l, c + cph - 6)
  x.phonenumber.Print (l, c + cph)
  x.Natural.Print (l + 1, c + cpc)
  x.city.Print (l + 1, c + cci)
  pbx.Print ("Funk:", l + 1, c + cph - 5)
  x.cellnumber.Print (l + 1, c + cph)
}

func (x *address) Codelen() uint {
/*
  return lenStreet +                 // 28
         x.Natural.Codelen() +       //  8
         lenCity +                   // 22
         2 * x.phonenumber.Codelen() // 12
*/
  return                                70
}

func (x *address) Encode() []byte {
  b:= make ([]byte, x.Codelen())
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
  return b
}

func (x *address) Decode (b []byte) {
  i, a:= uint(0), lenStreet
  x.street = Decode (x.street, b[i:i+a]).(text.Text)
  i += a
  a = x.Natural.Codelen()
  x.Natural = Decode (x.Natural, b[i:i+a]).(bnat.Natural)
  i += a
  a = lenCity
  x.city = Decode (x.city, b[i:i+a]).(text.Text)
  i += a
  a = x.phonenumber.Codelen()
  x.phonenumber = Decode (x.phonenumber, b[i:i+a]).(phone.PhoneNumber)
  i += a
  x.cellnumber = Decode (x.cellnumber, b[i:i+a]).(phone.PhoneNumber)
}

func init() {
//           1         2         3         4         5         6         7
// 01234567890123456789012345678901234567890123456789012345678901234567890123456789 // Compact
// Str./Nr.: ____________________________ Tel.: ________________
// PLZ: _____ Ort: ______________________ Funk: ________________
  mask.Ins ("Str./Nr.:", 0,  0)
  mask.Ins ("Tel.:",     0, 39)
  mask.Ins ("PLZ:",      1,  0)
  mask.Ins ("Ort:",      1, 11)
  mask.Ins ("Funk:",     1, 39)
//           1         2         3         4         5         6         7         8         9        10        11        12 
// 0123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890 // Wide
// Wide:
// Str./Nr.: ____________________________  PLZ: _____ Ort: ______________________  Tel.: ________________, ________________
//  mask.Ins ("Str./Nr.:", 0,   0)
//  mask.Ins ("PLZ:",      0,  40)
//  mask.Ins ("Ort:",      0,  51)
//  mask.Ins ("Tel:",      0,  80)
//  mask.Ins (",",         0, 102)
}
