package stru

// (c) Christian Maurer   v. 220818 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Structure interface {
// Sextuples of an atom-typ, a position on the screen, a fore- and a background colour
// and a boolean value indicating the structure is an index.

  Object

  Colours (f, b col.Colour)
  Cols() (col.Colour, col.Colour)
  Define (t int, n uint)
  Typ() int
  Index (b bool)
  IsIndex () bool
  Place (l, c uint)
  Pos() (uint, uint)
  Width() (uint)
}

func New() Structure { return new_() }
