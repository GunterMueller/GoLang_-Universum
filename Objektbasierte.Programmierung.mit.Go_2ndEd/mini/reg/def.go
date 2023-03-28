package reg

// (c) Christian Maurer   v. 221213 - license see µU.go

import
  . "µU/obj"
const (
  D = 9
  M = 1e9
  R = 26
)
type
  Register interface { // Register mit einem Kleinbuchstaben "a", ..., "z"
                       // als Namen und einer natürlichen Zahl < M als Wert.
                       // Daneben wird die Menge aller 26 Register verwaltet.
  Clearer
  Stringer
  Valuator

  Write (l, c uint)
  Edit (l, c uint)
}

func New() Register { return new_() }

func WriteAll (l, c uint) { writeAll(l,c) }
func EditAll (l, c uint) { editAll(l,c) }
