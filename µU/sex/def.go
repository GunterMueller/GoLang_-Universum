package sex

// (c) Christian Maurer   v. 221003 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  Sex interface {

  Editor
  col.Colourer
  Stringer
}

func New() Sex { return new_() }
