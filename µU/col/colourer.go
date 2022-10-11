package col

// (c) Christian Maurer   v. 220804 - license see µU.go

type
  Colourer interface {

// x has the fore-/background colours f, b.
  Colours (f, b Colour)

// Returns the fore- and background colour of x.
  Cols() (Colour, Colour)
}

func IsColourer (a any) bool {
  if a == nil { return false }
  _, ok := a.(Colourer)
  return ok
}
