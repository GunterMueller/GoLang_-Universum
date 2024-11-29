package main

// (c) Christian Maurer  v. 230326 - license see µU.go

import (
  "µU/col"
  "µU/gl"
  "µU/scr"
  . "µU/fig3"
)

func main() {
  s := scr.NewWH (0, 0, 800, 600); defer s.Fin()
  gl.ClearColour (col.FlashWhite()); 
  s.Go (draw, 0, -12, 0, 0, 0, 0, 0, 0, 1)
}

func draw() {
  m, o := col.Magenta(), col.Orange()
  Sphere (col.Red(), -1.25, -0.5, 0, 2)
  Torus (col.Green(), 0, 0, 0, 5, 1)
  VerTorus (col.Blue(), 5, -2, 0, 3, 0.5, 65)
  CylinderC ([]col.Colour{m,o}, 2.0, 1.5, -4, 1, 8)
}
