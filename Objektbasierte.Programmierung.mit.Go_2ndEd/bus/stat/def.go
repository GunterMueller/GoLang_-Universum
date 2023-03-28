package stat

// (c) Christian Maurer   v. 230309 - license see µU.go

import (
  . "µU/obj"
  . "bus/line"
)
const (
  L = 'l'
  R = 'r'
  O = 'o'
  U = 'u'
)
type
  Station interface {

  Object

  Set (l Line, nr uint, n string, b byte, y, x float64)
  Line() Line
  Number() uint
  Pos() (float64, float64)
  Umstieg()
  Renumber (l Line, nr uint)
  Equiv (Y any) bool
  EditScale()
  UnderMouse() bool
  Write (b bool)
}
