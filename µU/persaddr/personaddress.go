package persaddr

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/font"
  "µU/pers"
  "µU/addr"
)

type
  personAddress struct {
                       pers.Person
                       addr.Address
                       }

func new_() PersonAddress {
  x := new (personAddress)
  x.Person = pers.New()
  x.Person.SetFormat (pers.LongTB)
  x.Person.Colours (col.Yellow(), col.Black())
  x.Address = addr.New()
  x.Address.Colours (col.LightGreen(), col.Black())
  return x
}

func (x *personAddress) imp (Y any) *personAddress {
  y, ok := Y.(*personAddress)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *personAddress) Empty() bool {
  return x.Person.Empty()
}

func (x *personAddress) Clr() {
  x.Person.Clr()
  x.Address.Clr()
}

func (x *personAddress) Copy (Y any) {
  y := x.imp (Y)
  x.Person.Copy (y.Person)
  x.Address.Copy (y.Address)
}

func (x *personAddress) Clone() any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *personAddress) Eq (Y any) bool {
  y := x.imp (Y)
  return x.Person.Eq (y.Person) &&
         x.Address.Eq (y.Address)
}

func (x *personAddress) Less (Y any) bool {
  y := x.imp (Y)
  return x.Person.Less (y.Person)
}

func (x *personAddress) Sub (Y any) bool {
  y := x.imp (Y)
  return x.Person.Sub (y.Person)
}

func (x *personAddress) TeX() string {
  s := x.Person.TeX() + x.Address.TeX() + "\n\\smallskip"
  return s
}

func (x *personAddress) Write (l, c uint) {
  l++
  x.Person.Write (l, c)
  x.Address.Write (l + 3, c)
}

func (x *personAddress) Edit (l, c uint) {
  x.Write (l, c)
  l++
  x.Person.Edit (l, c)
  x.Address.Edit (l + 3, c)
}

func (x *personAddress) SetFont (f font.Font) {
  x.Person.SetFont (f)
  x.Address.SetFont (f)
}

func (x *personAddress) Print (l, c uint) {
  x.Person.Print (l, c)
  x.Address.Print (l + 3, c)
}

func (x *personAddress) Codelen() uint {
  return x.Person.Codelen() + x.Address.Codelen()
}

func (x *personAddress) Encode() Stream {
  s := make (Stream, x.Codelen ())
  i, a := uint(0), x.Person.Codelen()
  copy (s[i:i+a], x.Person.Encode())
  i += a
  a = x.Address.Codelen()
  copy (s[i:i+a], x.Address.Encode())
  return s
}

func (x *personAddress) Decode (s Stream) {
  i, a := uint(0), x.Person.Codelen()
  x.Person.Decode (s[i:i+a])
  i += a
  a = x.Address.Codelen()
  x.Address.Decode (s[i:i+a])
}

func (x *personAddress) Index() Func {
  return func (a any) any {
    x, ok := a.(*personAddress)
    if ! ok { TypeNotEqPanic (x, a) }
    return x.Person
  }
}

func (x *personAddress) Rotate() {
  x.Person.Rotate()
}
