package col

// (c) Christian Maurer   v. 170919 - license see µU.go

type
  Colourer interface {

// x has the fore-/background colours f, b.
  Colours (f, b Colour)
}
