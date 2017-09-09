package car

// (c) Christian Maurer   v. 120715 - license see murus.go

import
  "murus/col"
const (
  W = 32; H = 16 // pixelsize of car
)

func Draw (toTheRight bool, c col.Colour, x, y int) { draw(toTheRight, c, x, y) }
