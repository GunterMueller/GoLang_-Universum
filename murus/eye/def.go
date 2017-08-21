package eye

// (c) murus.org  v. 161216 - license see murus.go

import (
  . "murus/spc"
  "murus/col"
  "murus/vect"
)
type
  Eye interface {

  SetLight (n uint)
  Actualize ()
  DistanceFrom (v vect.Vector) float64
  Distance () float64
  Read (A []vect.Vector) bool
  Move (d Direction, dist float64)
  Turn (d Direction, a float64)
  Invert ()
  Focus (d float64)
  TurnAroundFocus (d Direction, a float64)
  Set (x, y, z, xf, yf, zf float64)
  Push (c col.Colour)
  Colour () col.Colour
}

func New() Eye { return new_() }
