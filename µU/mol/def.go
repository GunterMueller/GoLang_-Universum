package mol

// (c) Christian Maurer   v. 210414 - license see µU.go

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

// a is inserted into x.
  Ins (a atom.Atom, l, c uint)

// Returns the n-th atom of x.
  Component (n uint) Any

// m is the set of masks of x.
  SetMasks (m masks.MaskSet)
}

// Returns a new empty molecule.
func New() Molecule { return new_() }
