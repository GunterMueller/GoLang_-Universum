package pat

// (c) Christian Maurer   v. 221021 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/font"
  "µU/pers"
  "µU/addr"
  "µU/telmail"
)

type
  personAddressTelMail struct {
                              pers.Person
                              addr.Address
                              telmail.TelMail
                              }

func new_() PersonAddressTelMail {
  x := new (personAddressTelMail)
  x.Person = pers.New()
  x.Person.SetFormat (pers.LongTB)
  x.Person.Colours (col.Yellow(), col.Black())
  x.Address = addr.New()
  x.Address.Colours (col.LightGreen(), col.Black())
  x.TelMail = telmail.New()
  x.TelMail.Colours (col.LightGreen(), col.Black())
  return x
}

func (x *personAddressTelMail) imp (Y any) *personAddressTelMail {
  y, ok := Y.(*personAddressTelMail)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *personAddressTelMail) Empty() bool {
  return x.Person.Empty()
}

func (x *personAddressTelMail) Clr() {
  x.Person.Clr()
  x.Address.Clr()
  x.TelMail.Clr()
}

func (x *personAddressTelMail) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Person.Eq (y.Person) &&
         x.Address.Eq (y.Address) &&
         x.TelMail.Eq (y.TelMail)
}

func (x *personAddressTelMail) Copy (Y any) {
  y := x.imp (Y)
  x.Person.Copy (y.Person)
  x.Address.Copy (y.Address)
  x.TelMail.Copy (y.TelMail)
}

func (x *personAddressTelMail) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *personAddressTelMail) Less (Y any) bool {
  y := x.imp (Y)
  return x.Person.Less (y.Person)
}

func (x *personAddressTelMail) Leq (Y any) bool {
  return x.Less (Y) || x.Eq (Y)
}

func (x *personAddressTelMail) Sub (Y any) bool {
  y := x.imp (Y)
  return x.Person.Sub (y.Person)
}

func (x *personAddressTelMail) TeX() string {
//  s := x.Person.TeX() + x.Address.TeX() + x.TelMail.TeX() + "\n\\smallskip\n"
  s := "{\\vbox{\\hbox{" + x.Person.TeX() + "}"
  s += "\\hbox{" + x.Address.TeX() + "}"
  s += "\\hbox{" + x.TelMail.TeX() + "}}}"
  s += "\n\\bigskip\n"
  return s
}

func (x *personAddressTelMail) Write (l, c uint) {
  l++ // XXX
  x.Person.Write (l, c)
  x.Address.Write (l + 2, c)
  x.TelMail.Write (l + 4, c)
}

func (x *personAddressTelMail) Edit (l, c uint) {
  x.Write (l, c)
  l++ // XXX
  x.Person.Edit (l, c)
  x.Address.Edit (l + 2, c)
  x.TelMail.Edit (l + 4, c)
}

func (x *personAddressTelMail) SetFont (f font.Font) {
  x.Person.SetFont (f)
  x.Address.SetFont (f)
  x.TelMail.SetFont (f)
}

func (x *personAddressTelMail) Codelen() uint {
  return x.Person.Codelen() +
         x.Address.Codelen() +
         x.TelMail.Codelen()
}

func (x *personAddressTelMail) Encode() Stream {
  s := make (Stream, x.Codelen ())
  i, a := uint(0), x.Person.Codelen()
  copy (s[i:i+a], x.Person.Encode())
  i += a
  a = x.Address.Codelen()
  copy (s[i:i+a], x.Address.Encode())
  i += a
  a = x.TelMail.Codelen()
  copy (s[i:i+a], x.TelMail.Encode())
  return s
}

func (x *personAddressTelMail) Decode (s Stream) {
  i, a := uint(0), x.Person.Codelen()
  x.Person.Decode (s[i:i+a])
  i += a
  a = x.Address.Codelen()
  x.Address.Decode (s[i:i+a])
  i += a
  a = x.TelMail.Codelen()
  x.TelMail.Decode (s[i:i+a])
}

func (x *personAddressTelMail) Index() Func {
  return func (a any) any {
    x, ok := a.(*personAddressTelMail)
    if ! ok { TypeNotEqPanic (x, a) }
    return x.Person
  }
}

func (x *personAddressTelMail) Rotate() {
  x.Person.Rotate()
}
