package col

// (c) Christian Maurer   v. 201004 - license see µU.go

import
  . "µU/obj"
type
  Colourer interface {

// x has the fore-/background colours f, b.
  Colours (f, b Colour)
}

func IsColourer (a Any) bool {
  if a == nil { return false }
  _, ok := a.(Colourer)
  return ok
}
