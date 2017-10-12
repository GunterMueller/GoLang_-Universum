package text

// (c) Christian Maurer   v. 170919 - license see µu.go

import (
  . "µu/obj"
  "µu/col"
)
type
  Text interface {

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
  Pos (b byte) (uint, bool)

// experimental:
// starting with position p in x, n bytes are removed; tail filled with spaces up to the original length
  Rem (p, n uint)
// x besteht aus Y ab Position p der Länge n, Rest Leerzeichen
  Cut (Y Text, p, n uint)

  IsCap0() bool
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
