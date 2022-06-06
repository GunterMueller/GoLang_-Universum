package text

// (c) Christian Maurer   v. 220524 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/font"
)
type
  Text interface { // strings of fixed length

  Object
  col.Colourer
  Editor
  Stringer
  TeXer
  Printer

  Transparence (t bool)
  SetFont (f font.Font) // only to print
  SetFontsize (s font.Size)

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

// starting with position p in x, n bytes are removed;
// tail filled with spaces up to the original length
  Rem (p, n uint)

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
