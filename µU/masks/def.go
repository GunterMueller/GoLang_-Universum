package masks

// (c) Christian Maurer   v. 161216 - license see µU.go

import
  . "µU/obj"
type
  MaskSequence interface {

  Object
  Write (l, c uint)
  Printer
  Line (n uint)
  Ins (m string, l, c uint)
}

// Returns a new empty mask sequence.
func New() MaskSequence { return new_() }
