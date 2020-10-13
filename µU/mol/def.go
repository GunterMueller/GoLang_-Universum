package mol

// (c) Christian Maurer   v. 201004 - license see µU.go

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

// Returns a new empty molecule with n atoms.
// TODO more information
func New (n uint) Molecule { return new_(n) }
