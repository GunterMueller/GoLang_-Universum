package audio

// (c) Christian Maurer   v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/enum"
  "murus/text"
  "murus/masks"
  "murus/atom"
  "murus/mol"
)
const (
  lenWerk = 25
  lenOrch = 25
  lenName = 25
)
type
  audio struct {
               mol.Molecule
               }

func (x *audio) imp (Y Any) mol.Molecule {
  y, ok:= Y.(*audio)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y.Molecule
}

func new_() Audio {
  x := new (audio)
  x.Molecule = mol.New()
  a := atom.New (enum.New (enum.Composer))
  a.Colours (col.Yellow, col.Black)
  x.Ins (a, 0, 11)
  a = atom.New (text.New (lenWerk))
  a.Colours (col.LightRed, col.Black)
  x.Ins (a, 1, 11)
  a = atom.New (text.New (lenOrch)) // Orchester
  x.Ins (a, 2, 11)
  a = atom.New (text.New (lenName)) // Dirigent
  a.Colours (col.Cyan, col.Black)
  x.Ins (a, 3, 11)
  a = atom.New (text.New (lenName)) // Solist
  a.Colours (col.LightBlue, col.Black)
  x.Ins (a, 4, 11)
  a = atom.New (text.New (lenName)) // Solist
  a.Colours (col.LightBlue, col.Black)
  x.Ins (a, 5, 11)
  a = atom.New (enum.New (enum.RecordLabel))
  a.Colours (col.LightCyan, col.Black)
  x.Ins (a, 6, 11)
  a = atom.New (enum.New (enum.AudioMedium))
  x.Ins (a, 7, 11)
  a = atom.New (enum.New (enum.SparsCode))
  x.Ins (a, 8, 11)

  m:= masks.New()
  m.Ins ("Komponist:", 0, 0)
  m.Ins ("     Werk:", 1, 0)
  m.Ins ("Orchester:", 2, 0)
  m.Ins (" Dirigent:", 3, 0)
  m.Ins (" Solist 1:", 4, 0)
  m.Ins (" Solist 2:", 5, 0)
  m.Ins ("    Firma:", 6, 0)
  m.Ins ("   Platte:", 7, 0)
  m.Ins ("       ad:", 8, 0)
  x.SetMask (m)

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

func (x *audio) Less (Y Any) bool {
  return x.Molecule.Less (x.imp (Y))
}

func (x *audio) Index() Func {
  return func (X Any) Any {
    return x.imp (X).Component (0).(atom.Atom)
  }
}

func (x *audio) RotOrder() {
//  ___.RotOrder()
}
