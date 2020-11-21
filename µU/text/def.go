package text

// (c) Christian Maurer   v. 201104 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Text interface { // strings of fixed length

  Object
  col.Colourer
  Editor
  Stringer
  Printer

// Specs see str/def.go.
  Equiv (Y Text) bool

  Sub (Y Text) bool
  Sub0 (Y Text) bool
  EquivSub (Y Text) (uint, bool)

  Len() uint
  ProperLen() uint
  Byte (n uint) byte
  Pos (b byte) (uint, bool)
  Replace1 (p uint, b byte)

// starting with position p in x, n bytes are removed; tail filled with spaces up to the original length
  Rem (p, n uint)

// x besteht aus Y ab Position p der Länge n, Rest Leerzeichen
//  Cut (Y Text, p, n uint)

  IsUpper0() bool
  ToUpper()
  ToLower()
  ToUpper0()
  ToLower0()

  Split() []Text

  WriteGr (x, y int)
  EditGr (x, y int)
}

// Returns a new empty text of length n.
func New (n uint) Text { return new_(n) }
