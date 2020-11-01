package casus

// (c) Christian Maurer   v. 200010 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const (
  Short = Format(iota)
  Long
  NFormats
)
type
  Casus interface {

  Formatter
  Object
  col.Colourer
  Editor
  Stringer
  Printer
}

// Returns a new empty casus.
func New() Casus { return new_() }
