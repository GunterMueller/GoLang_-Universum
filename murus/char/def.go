package char

// (c) murus.org  v. 161216 - license see murus.go

import
  . "murus/obj"
type
  Character interface {

  Editor
  Stringer
  Printer
  Valuator

  Equiv (y Any) bool
  SetByte (b byte)
  ByteVal() byte
}

// Returns a new empty character (= space).
func New() Character { return new_() }
