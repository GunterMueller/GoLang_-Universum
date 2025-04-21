package pac

// (c) Christian Maurer   v. 250407 - license see µU.go

import (
  . "µU/obj"
  "µU/kbd"
  "µU/col"
  "µU/font"
  "µU/str"
  "µU/pers"
  "µU/addr"
  "µU/cont"
)
const (
  sep = ','
  seps = ","
)
type
  personAddressContact struct {
                              pers.Person
                              addr.Address
                              cont.Contact
                              }

func new_() PersonAddressContact {
  x := new (personAddressContact)
  x.Person = pers.New()
  x.Person.SetFormat (pers.NameBT)
  x.Person.Colours (col.FlashWhite(), col.DarkRed())
  x.Address = addr.New()
  x.Address.Colours (col.FlashWhite(), col.DarkGreen())
  x.Contact = cont.New()
  x.Contact.Colours (col.FlashWhite(), col.DarkBlue())
  return x
}

func (x *personAddressContact) imp (Y any) *personAddressContact {
  y, ok := Y.(*personAddressContact)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *personAddressContact) Empty() bool {
  return x.Person.Empty()
}

func (x *personAddressContact) Clr() {
  x.Person.Clr()
  x.Address.Clr()
  x.Contact.Clr()
}

func (x *personAddressContact) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Person.Eq (y.Person) &&
         x.Address.Eq (y.Address) &&
         x.Contact.Eq (y.Contact)
}

func (x *personAddressContact) Copy (Y any) {
  y := x.imp (Y)
  x.Person.Copy (y.Person)
  x.Address.Copy (y.Address)
  x.Contact.Copy (y.Contact)
}

func (x *personAddressContact) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *personAddressContact) Less (Y any) bool {
  y := x.imp (Y)
  return x.Person.Less (y.Person)
}

func (x *personAddressContact) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *personAddressContact) String() string {
  s := x.Person.String()
  str.OffSpc1 (&s)
  t := x.Address.String()
  str.OffSpc1 (&t)
  s += t
  t = x.Contact.String()
  str.OffSpc1 (&t)
  s += t
  return s
}

func (x *personAddressContact) Defined (s string) bool {
  ss, n := str.SplitByte (s, sep)
  const (
    p = pers.N
    a = addr.N
    t = cont.N
  )
  if n != p + a + t { return false }
  s = ""; for i := uint(0); i < p; i++ { s += ss[i] + seps }
  if ! x.Person.Defined (s) { return false }
  s = ""; for i := uint(p); i < p + a; i++ { s += ss[i] + seps }
  if ! x.Address.Defined (s) { return false }
  s = ""; for i := uint(p + a); i < p + a + t; i++ { s += ss[i] + seps }
  if ! x.Contact.Defined (s) { return false }
  return true
}

func (x *personAddressContact) Sub (Y any) bool {
  y := x.imp (Y)
  return x.Person.Sub (y.Person)
}

func (x *personAddressContact) TeX() string {
  s := "{\\vbox{\\hbox{" + x.Person.TeX() + "}"
  s += "\\hbox{" + x.Address.TeX() + "}"
  s += "\\hbox{" + x.Contact.TeX() + "}}}"
  s += "\n\\bigskip\n"
  return s
}

func (x *personAddressContact) Write (l, c uint) {
  x.Person.Write (l + 1, c)
  x.Address.Write (l + 4, c)
  x.Contact.Write (l + 6, c)
}

func (x *personAddressContact) Edit (l, c uint) {
  x.Write (l, c)
  i := uint(0)
  loop:
  for {
    switch i {
    case 0:
      x.Person.Edit (l + 1, c)
    case 1:
      x.Address.Edit (l + 4, c)
    case 2:
      x.Contact.Edit (l + 6, c)
    }
    switch C, d := kbd.LastCommand(); C {
    case kbd.Esc: 
      break loop
    case kbd.Enter: 
      if d == 0 {
        if i < 2 {
          i++
        } else {
          break loop
        }
      } else {
        break loop
      }
    case kbd.Down, kbd.Right:
      if i < 2 {
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

func (x *personAddressContact) SetFont (f font.Font) {
  x.Person.SetFont (f)
  x.Address.SetFont (f)
  x.Contact.SetFont (f)
}

func (x *personAddressContact) Codelen() uint {
  return x.Person.Codelen() +
         x.Address.Codelen() +
         x.Contact.Codelen()
}

func (x *personAddressContact) Encode() Stream {
  s := make (Stream, x.Codelen ())
  i, a := uint(0), x.Person.Codelen()
  copy (s[i:i+a], x.Person.Encode())
  i += a
  a = x.Address.Codelen()
  copy (s[i:i+a], x.Address.Encode())
  i += a
  a = x.Contact.Codelen()
  copy (s[i:i+a], x.Contact.Encode())
  return s
}

func (x *personAddressContact) Decode (s Stream) {
  i, a := uint(0), x.Person.Codelen()
  x.Person.Decode (s[i:i+a])
  i += a
  a = x.Address.Codelen()
  x.Address.Decode (s[i:i+a])
  i += a
  a = x.Contact.Codelen()
  x.Contact.Decode (s[i:i+a])
}

func (x *personAddressContact) Index() Func {
  return func (a any) any {
    x, ok := a.(*personAddressContact)
    if ! ok { TypeNotEqPanic (x, a) }
    return x.Person
  }
}

func (x *personAddressContact) Rotate() {
  x.Person.Rotate()
}
