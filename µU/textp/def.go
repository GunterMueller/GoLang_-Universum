package textp

import (
  . "µU/obj"
  "µU/col"
)
type
  TextPair interface {

  Object
  Editor
  col.Colourer
  Len() (uint, uint)
}

func New (m, n uint) TextPair { return new_(m,n) }
