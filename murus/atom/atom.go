package atom

// (c) murus.org  v. 161216 - license see murus.go

import (
  "reflect"
  "murus/str"
  . "murus/obj"
  "murus/ker"
  "murus/font"
  "murus/col"
  "murus/enum"
  "murus/tval"
  "murus/char"
  "murus/text"
  "murus/bnat"
  "murus/breal"
  "murus/clk"
  "murus/day"
  "murus/euro"
  "murus/cntry"
  "murus/pers"
  "murus/phone"
  "murus/addr"
)
type
  atom struct {
              Object
              uint
              }

func new_(o Object) Atom {
  x := new(atom)
  s := reflect.TypeOf (o).String()
  if p, ok := str.Pos (s, '.'); ok { s = s[1:p] }
  switch s {
  case "enum":
    x.Object, x.uint = enum.New (o.(enum.Enumerator).Typ()), Enumerator
  case "tval":
    x.Object, x.uint = tval.New(), TruthValue
  case "char":
    x.Object, x.uint = char.New(), Character
  case "text":
    x.Object, x.uint = text.New (o.(text.Text).Len()), Text
  case "bnat":
    x.Object, x.uint = bnat.New (o.(bnat.Natural).Width()), Natural
  case "breal":
    x.Object, x.uint = breal.New (4), Real
  case "clk":
    x.Object, x.uint = clk.New(), Clocktime
  case "day":
    x.Object, x.uint = day.New(), Calendarday
  case "euro":
    x.Object, x.uint = euro.New(), Euro
  case "cntry":
    x.Object, x.uint = cntry.New(), Country
  case "pers":
    x.Object, x.uint = pers.New(), Person
////    persaddr
  case "phone":
    x.Object, x.uint = phone.New(), PhoneNumber
  case "addr":
    x.Object, x.uint = addr.New(), Address
////    host, conn
  default:
    ker.Panic ("atom.New: parameter does not characterize an atomtype")
  }
  return x
}

func (x *atom) imp(Y Any) *atom {
  y, ok := Y.(*atom)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *atom) Equiv (Y Any) bool {
  return x.uint == x.imp (Y).uint
}

func (x *atom) Empty() bool {
  return x.Object.Empty()
}

func (x *atom) Clr() {
  x.Object.Clr()
}

func (x *atom) Eq (Y Any) bool {
  return x.Object.Eq (x.imp (Y).Object)
}

func (x *atom) Copy (Y Any) {
  x.Object.Copy (x.imp (Y).Object)
}

func (x *atom) Clone() Any {
  y := new_(x.Object)
  y.Copy(x)
  return y
}

func (x *atom) Less (Y Any) bool {
  return x.Object.Less (x.imp (Y).Object)
}

func (x *atom) GetFormat() Format {
  return x.Object.(Formatter).GetFormat()
}

func (x *atom) SetFormat (f Format) {
  x.Object.(Formatter).SetFormat (f)
}

func (x *atom) Colours (f, b col.Colour) {
  x.Object.(Editor).Colours (f, b)
}

func (x *atom) Write (l, c uint) {
  x.Object.(Editor).Write (l, c)
}

func (x *atom) Edit (l, c uint) {
  x.Object.(Editor).Edit (l, c)
}

func (x *atom) SetFont (f font.Font) {
  x.Object.(Printer).SetFont (f)
}

func (x *atom) Print (l, c uint) {
  x.Object.(Printer).Print (l, c)
}

func (x *atom) Codelen() uint {
  return x.Object.Codelen()
}

func (x *atom) Encode() []byte {
  return x.Object.Encode()
}

func (x *atom) Decode (bs []byte) {
  x.Object.Decode (bs)
}
