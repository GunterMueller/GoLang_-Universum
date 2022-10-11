package mask

// (c) Christian Maurer   v. 220804 - license see µU.go

import
  . "µU/obj"
const
  M = uint(16) // maximal mask text length
type
  Mask interface { // constant text

  Object

  Place (l, c uint)
  Pos() (uint, uint)
  Wd() uint
  Del()
  Write()
  Edit()
  Print()
}

func New() Mask { return new_() }
