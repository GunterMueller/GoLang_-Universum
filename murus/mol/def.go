package mol

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/obj"
  "murus/atom"
  "murus/masks"
)
type
  Molecule interface {

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
func New() Molecule { return newMol() }
