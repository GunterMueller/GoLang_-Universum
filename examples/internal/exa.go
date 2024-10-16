package exa

// (c) Christian Maurer   v. 201102 - license see µU.go

import (
  "math"
  "µU/col"
  "µU/gl"
)

func f (x, y float64) float64 {
  return math.Sin(x) * math.Sin(y) - 4
}

func draw() {
  gl.Clear()
//  gl.InitLight (0,             // # of light
//                3, 1, 3,       // position
//                0.9, 0.9, 0.9, // ambience
//                127, 127, 0)   // r, g, b
  gl.Colour (col.Yellow())
  gl.Sphere (2, 1, 2, 0.2)
  gl.Colour (col.Orange())
  gl.Sphere (0, 0, 0, 1)
  gl.Colour (col.Blue())
  gl.Torus (0, 0, 0, 2, 0.2)
  gl.Colour (col.Green())
  gl.VerTorus (0, 0, 0, 3, 0.2, 1)
  gl.Colour (col.Red())
  gl.Surface (f, 4, 4)
}
