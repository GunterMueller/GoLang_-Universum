package masks

// (c) Christian Maurer   v. 210408 - license see µU.go

import
  . "µU/obj"
type
  MaskSet interface {

  Object
  Write (l, c uint)
  Printer
  Num() uint
  Ins (m string, l, c uint)
}

// Returns a new empty mask set.
func New() MaskSet { return new_() }
