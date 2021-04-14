package mol

// (c) Christian Maurer   v. 210413 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/atom"
  "µU/masks"
)
type
  Molecule interface {

  Object
  col.Colourer
  Editor
  Printer

  Ins (a atom.Atom, l, c uint)
  Selected (l, c uint) bool
  Del (n uint)
  Num() uint
  Component (n uint) Any
  SetMasks (m masks.MaskSet)
}

// Returns a new empty molecule.
func New() Molecule { return new_() }
