package char

// (c) Christian Maurer   v. 171007 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  Character interface {

  Object
  col.Colourer
  EditorGr
  Stringer
  Printer
  Valuator

  Equiv (y Any) bool
  SetByte (b byte)
  ByteVal() byte
}

// Returns a new empty character (= space).
func New() Character { return new_() }
