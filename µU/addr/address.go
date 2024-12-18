package addr

// (c) Christian Maurer   v. 241025 - license see µU.go

import (
  "µU/env"
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/str"
  "µU/box"
  "µU/font"
  "µU/pbox"
  "µU/text"
  "µU/bn"
  "µU/cntry"
)
const (
  lenStreet = uint(28)
  lenCity   = uint(22)
  lenEmail  = uint(40)
  sep = ','
  seps = ","
)
type
  address struct {
          street text.Text
                 bn.Natural "zip" // "PLZ"
            city text.Text
                 cntry.Country
                 }
var (
  bx = box.New()
  pbx = pbox.New()
  cF, cB = col.FlashWhite(), col.DarkBlue()
)

func new_() Address {
  x := new(address)
  f, b := col.FlashWhite(), col.DarkGreen()
  x.street = text.New (lenStreet)
  x.street.Colours (f, b)
  x.city =  text.New (lenCity)
  x.city.Colours (f, b)
  x.Natural = bn.New (5)
  x.Natural.Colours (f, b)
  x.Country = cntry.New()
  x.Country.SetFormat (cntry.Long)
  x.Colours (cF, cB)
  return x
}

func (x *address) imp(Y any) *address {
  y, ok := Y.(*address)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *address) Empty() bool {
  return x.street.Empty() &&
         x.Natural.Empty() &&
         x.city.Empty() &&
         x.Country.Empty()
}

func (x *address) Clr() {
  x.street.Clr()
  x.Natural.Clr()
  x.city.Clr()
  x.Country.Clr()
}

func (x *address) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *address) Copy (Y any) {
  y := x.imp (Y)
  x.street.Copy (y.street)
  x.Natural.Copy (y.Natural)
  x.city.Copy (y.city)
  x.Country.Copy (y.Country)
}

func (x *address) Eq (Y any) bool {
  y := x.imp (Y)
  return x.street.Eq (y.street) &&
         x.Natural.Eq (y.Natural) &&
         x.city.Eq (y.city) &&
         x.Country.Eq (y.Country)
}

func (x *address) Equiv (Y any) bool {
  if x.Natural.Eq (x.imp (Y).Natural) {
    return true
  }
  return false
}

func (x *address) Less (Y any) bool {
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

func (x *address) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *address) String() string {
  s := x.street.String()
  str.OffSpc1 (&s)
  s += seps
  t := x.Natural.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.city.String()
  str.OffSpc1 (&t)
  s += t + seps
  t = x.Country.String()
  str.OffSpc1 (&t)
  s += t + seps
  return s
}

func (x *address) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  if n != N { return false }
  if ! x.street.Defined (ss[0]) { return false }
  if ! x.Natural.Defined (ss[1]) { return false }
  if ! x.city.Defined (ss[2]) { return false }
  if ! x.Country.Defined (ss[3]) { return false }
  return true
}

func (x *address) Colours (f, b col.Colour) {
  x.street.Colours (f, b)
  x.Natural.Colours (f, b)
  x.city.Colours (f, b)
  x.Country.Colours (f, b)
}

func (x *address) Cols() (col.Colour, col.Colour) {
  return x.street.Cols()
}

const (
  cs = 10; cz = 45; cc = 57; cl = cc
)

func writeMask (l, c uint) {
/*        1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789
  street: ____________________________  zip: _____ city: ______________________
                                                country: ________________
 Str./Nr: ____________________________  PLZ: _____  Ort: ______________________
                                                   Land: ________________
*/
  if env.E() {
    bx.Wd (7)
    bx.Write ("street:",  l, c)
    bx.Wd (4)
    bx.Write ("zip:",     l, c + cz - 5)
    bx.Wd (5)
    bx.Write ("city:",    l, c + cc - 6)
    bx.Wd (8)
    bx.Write ("country:", l + 1, c + cc - 9)
  } else {
    bx.Wd (8)
    bx.Write ("Str./Nr:", l, c + cs - 9)
    bx.Wd (4)
    bx.Write ("PLZ:",     l, c + cz - 5)
    bx.Write ("Ort:",     l, c + cc - 5)
    bx.Wd (5)
    bx.Write ("Land:",    l + 1, c + cc - 6)
  }
}

func (x *address) Write (l, c uint) {
  writeMask (l, c)
  x.street.Write (l, c + cs)
  x.Natural.Write (l, c + cz)
  x.city.Write (l, c + cc)
  co := cntry.New()
  co.Colours (col.FlashWhite(), col.DarkGreen())
  co.Write (l + 1, c + cl)
  x.Country.Write (l + 1, c + cl)
}

func (x *address) Edit (l, c uint) {
  x.Write (l, c)
  i := 0
  if C, _ := kbd.LastCommand(); C == kbd.Up {
    i = N - 1
  }
  loop:
  for {
    switch i {
    case 0:
      x.street.Edit (l, c + cs)
    case 1:
      x.Natural.Edit (l, c + cz)
    case 2:
      x.city.Edit (l, c + cc)
    case 3:
      x.Country.Edit (l + 1, c + cl)
    }
    switch cmd, d := kbd.LastCommand(); cmd {
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

func (x *address) SetFont (f font.Font) {
  x.street.SetFont (f)
  x.Natural.SetFont (f)
  x.city.SetFont (f)
  x.Country.SetFont (f)
}

func (x *address) Print (l, c uint) {
// printMask()
  x.street.Print (l, c + cs)
  x.Natural.Print (l, c + cz)
  x.city.Print (l, c + cc)
  x.Country.Print (l + 1, c + cl)
}

func (x *address) Codelen() uint {
  return lenStreet +
         x.Natural.Codelen() +
         lenCity +
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
  a = x.Country.Codelen()
  x.Country = Decode (x.Country, s[i:i+a]).(cntry.Country)
}

func (x *address) TeX() string {
  s := x.street.TeX()
  if x.Natural.Empty() {
    s += ", " + x.city.TeX()
  } else {
    s += ", " + x.Natural.String() + " " + x.city.TeX()
  }
  c := x.Country
  if ! c.Empty() && c.TLD() != "de" { s += " (" + c.TeX() + ")" }
  return s
}
