package atom

// (c) Christian Maurer   v. 170918 - license see µU.go

import (
  "µU/scr"
  "µU/sel"
  "µU/enum"
  "µU/tval"
  "µU/char"
  "µU/text"
  "µU/bnat"
  "µU/breal"
  "µU/clk"
  "µU/day"
  "µU/euro"
  "µU/cntry"
  "µU/pers"
  "µU/phone"
  "µU/addr"
)
const
  M = 14
var
  name = []string {
           "Enumerator",
           "TruthValue", "Character", "Text", "Natural", "Real",
           "Clocktime", "Calendarday", "Euro", "Country",
           "Person", "PhoneNumber", "Address" }


func Selected (l, c uint) Atom {
  cF, cB := scr.ScrCols()
  cB.Contrast()
  n := uint(0)
  z, s := scr.MousePos()
  x := new(atom)
  sel.Select1 (name, uint(Ntypes), M, &n, z, s, cF, cB)
  if n < Ntypes {
    x.uint = n
  } else {
    return nil
  }
  switch x.uint {
  case Enumerator:
    e := enum.Title // TODO e per select-menue aussuchen
    x.Object = enum.New (e)
  case TruthValue:
    x.Object = tval.New()
  case Character:
    x.Object = char.New()
  case Text:
    n := uint(10) // TODO n editieren
    x.Object = text.New (n)
  case Natural:
    n := uint(2) // TODO n editieren
    x.Object = bnat.New (n)
  case Real:
    n := uint(6) // TODO n editieren
    x.Object = breal.New (n)
  case Clocktime:
    x.Object = clk.New()
  case Calendarday:
    x.Object = day.New()
  case Euro:
    x.Object = euro.New()
  case Country:
    x.Object = cntry.New()
  case Person:
    x.Object = pers.New()
  case PhoneNumber:
    x.Object = phone.New()
  case Address:
    x.Object = addr.New()
  }
  return x
}
