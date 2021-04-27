package audio

// (c) Christian Maurer   v. 210415 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/enum"
  "µU/text"
  "µU/masks"
  "µU/atom"
  "µU/mol"
)
const (
  len0 = 30
  len1 = 60
)
type
  order byte; const (
  subject = order(iota)
  composer_group
  medium
  nOrders
)
type
  audio struct {
               mol.Molecule
               }
var (
  actOrd order
  cF, cB = col.LightWhite(), col.Black()
)

func (x *audio) imp (Y Any) mol.Molecule {
  y, ok:= Y.(*audio)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.Molecule
}

func new_() Audio {
  x := new (audio)
  x.Molecule = mol.New()
  a := atom.New (enum.New (enum.Audio)) // Gebiet
  a.Colours (col.Yellow(), col.Black())
  x.Ins (a, 0, 12)

  a = atom.New (text.New (len0)) // Komponist/Gruppe
  a.Colours (col.LightRed(), col.Black())
  x.Ins (a, 0, 46)

  a = atom.New (enum.New (enum.AudioMedium)) // Medium
  a.Colours (col.LightBlue(), col.Black())
  x.Ins (a, 1, 12)

  a = atom.New (text.New (len0)) // Dirigent/Solist
  a.Colours (col.LightBlue(), col.Black())
  x.Ins (a, 1, 46)

  a = atom.New (text.New (len1)) // Werk
  a.Colours (col.Cyan(), col.Black())
  x.Ins (a, 2, 12)
/*/       1         2         3         4         5         6         7
01234567890123456789012345678901234567890123456789012345678901234567890123456789

    Gebiet: ___________     Komponist/Gruppe: ______________________________
    Medium: ___              Dirigent/Solist: ______________________________
      Werk: ________________________________________________________________
/*/
  m := masks.New()
  m.Ins ("Gebiet:",           0,  4)
  m.Ins ("Komponist/Gruppe:", 0, 28)
  m.Ins ("Medium:",           1,  4)
  m.Ins ("Dirigent/Solist:",  1, 29)
  m.Ins ("Werk:",             2,  6)
  x.SetMasks (m)
  return x
}

func (x *audio) Eq (Y Any) bool {
  return x.Molecule.Eq (x.imp (Y))
}

func (x *audio) Copy (Y Any) {
  x.Molecule.Copy (x.imp(Y))
}

func (x *audio) Clone() Any {
  y := new_()
  y.Copy (x)
  return y
}

// func (x *audio) Less (Y Any) bool {
//   return x.Molecule.Less (x.imp (Y))
// }

func (x *audio) Less (Y Any) bool {
  y := x.imp(Y)
  xs := x.Component(0).(atom.Atom)
  xc := x.Component(1).(atom.Atom)
//  xd := x.Component(2).(atom.Atom)
  xm := x.Component(3).(atom.Atom)
//  xw := x.Component(4).(atom.Atom)
  ys := y.Component(0).(atom.Atom)
  yc := y.Component(1).(atom.Atom)
//  yd := y.Component(2).(atom.Atom)
  ym := y.Component(3).(atom.Atom)
//  yw := y.Component(4).(atom.Atom)
/*/
    Gebiet: ___________     Komponist/Gruppe: ______________________________
    Medium: ___
/*/
  switch actOrd {
  case subject:
    if xs.Eq (ys) {
      if xc.Eq (yc) {
        return xm.Less (ym)
      }
      return xc.Less (yc)
    }
    return xs.Less (ys)
  case composer_group:
    if xc.Eq (yc) {
      if xs.Eq (ys) {
        return x.Less (ym)
      }
      return x.Less (ys)
    }
    return xc.Less (yc)
  case medium:
    if xm.Eq (ym) {
      if xs.Eq (ys) {
        return x.Less (yc)
      }
      return x.Less (ys)
    }
    return xm.Less (ym)
  }
  return false
}

func (x *audio) RotOrder() {
  actOrd = (actOrd + 1) % nOrders
}

func (x *audio) Index() Func {
  return func (a Any) Any {
    return a
  }
}
