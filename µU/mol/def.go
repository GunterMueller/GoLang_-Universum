package mol

// (c) Christian Maurer   v. 170919 - license see µU.go

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
  Del (n uint)
  Num () uint
  Component (n uint) Any
  SetMask (m masks.MaskSequence)

// Equiv (Y Any) bool
// Sort ()
}

// Returns a new empty molecule.
// TODO more information
func New() Molecule { return new_() }
