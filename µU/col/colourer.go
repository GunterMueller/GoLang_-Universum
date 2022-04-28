package col

// (c) Christian Maurer   v. 220420 - license see µU.go

type
  Colourer interface {

// x has the fore-/background colours f, b.
  Colours (f, b Colour)
}

func IsColourer (a any) bool {
  if a == nil { return false }
  _, ok := a.(Colourer)
  return ok
}
