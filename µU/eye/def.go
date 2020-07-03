package eye

// (c) Christian Maurer   v. 191117 - license see µU.go

import (
//  . "µU/obj"
  "µU/col"
  "µU/vect"
)
type
  Eye interface {

  SetLight (n uint)
  Get() (float64, float64, float64, float64, float64, float64)
  Set (x, y, z, x1, y1, z1 float64)
  DistanceFrom (v vect.Vector) float64
  Distance() float64
  Flatten (b bool)
  Move (i int, d float64)
  Turn (i int, a float64)
  Invert()
  Focus (d float64)
  TurnAroundFocus (i int, a float64)
  Push (c col.Colour)
  Top() col.Colour
}

func New() Eye { return new_() }
