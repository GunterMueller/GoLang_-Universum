package persaddr

// (c) Christian Maurer   v. 201010 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/pers"
  "µU/addr"
  "µU/atom"
  "µU/mol"
)
type
  personAddress struct {
                       mol.Molecule
                       }

func new_() PersonAddress {
  x := new (personAddress)
  x.Molecule = mol.New (2)
  a := atom.New (pers.New())
  a.SetFormat (pers.LongTB)
  a.Colours (col.Yellow(), col.Black())
  x.Ins (a, 0, 0)
  a = atom.New (addr.New())
  a.Colours (col.LightGreen(), col.Black())
  x.Ins (a, 2, 0)
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
  return func (a Any) Any {
    x, ok := a.(*personAddress)
    if ! ok { TypeNotEqPanic (x, a) }
    return x.Component(0).(atom.Atom)
  }
}

func (x *personAddress) Rotate() {
  x.Component(0).(Rotator).Rotate()
}
