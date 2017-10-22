package persaddr

// (c) Christian Maurer   v. 170918 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/pers"
  "µU/addr"
  "µU/atom"
  "µU/mol"
  "µU/masks"
)
type
  personAddress struct {
                       mol.Molecule
                       }

func new_() PersonAddress {
  x := new (personAddress)
  x.Molecule = mol.New()
  a := atom.New (pers.New())
  a.SetFormat (pers.LongTB)
  a.Colours (col.Yellow(), col.Black())
  x.Ins (a, 0, 0)
  a = atom.New (addr.New())
  a.Colours (col.LightGreen(), col.Black())
  x.Ins (a, 2, 0)
  m := masks.New()
  m.Ins ("Str./Nr.:", 2,  0) // TODO: von addr übernehmen
  m.Ins ("Tel.:",     2, 39)
  m.Ins ("PLZ:",      3,  0)
  m.Ins ("Ort:",      3, 11)
  m.Ins ("Funk:",     3, 39)
  x.SetMask(m)
  return x
}

func (x *personAddress) imp (Y Any) *personAddress {
  y, ok := Y.(*personAddress)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *personAddress) Eq (Y Any) bool {
  return x.Molecule.Eq (x.imp (Y).Molecule)
}

func (x *personAddress) Copy (Y Any) {
  x.Molecule.Copy (x.imp (Y).Molecule)
}

func (x *personAddress) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

func (x *personAddress) Less (Y Any ) bool {
  return x.Molecule.Less (x.imp (Y).Molecule)
}

func (x *personAddress) Index() Func {
  return func (X Any) Any {
    x, ok := X.(*personAddress)
    if ! ok { TypeNotEqPanic (x, X) }
    return x.Component(0).(atom.Atom)
  }
}

var
  nn = pers.New()

func (x *personAddress) RotOrder() {
  nn.RotOrder()
}
