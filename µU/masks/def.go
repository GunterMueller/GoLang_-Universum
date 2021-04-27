package masks

// (c) Christian Maurer   v. 210415 - license see µU.go

import
  . "µU/obj"
type
  MaskSet interface {

  Object
  Printer

// x is written to the screen, starting at position (l, c).
  Write (l, c uint)

// Returns the number of masks in x.
  Num() uint

// m with start position (l, c) on the screen is inserted as mask into x.
  Ins (m string, l, c uint)
}

// Returns a new empty mask set.
func New() MaskSet { return new_() }
