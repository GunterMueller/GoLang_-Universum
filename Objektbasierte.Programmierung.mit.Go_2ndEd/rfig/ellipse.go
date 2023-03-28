package main

// (c) Christian Maurer  v. 230312 - license see µU.go

import (
  "µU/col"
  "µU/gl"
  "µU/scr"
  . "µU/fig3"
)

func main() {
  s := scr.NewWH (0, 0, 800, 600); defer s.Fin()
  gl.ClearColour (col.LightWhite()); gl.Clear()
  s.Go (scr.Look, draw, 0, -6, 2, 0, 0, 2, 0, 0, 1)
}

func draw() {
  Cone (col.Blue(), 0, 0, 0, 2, 5)
  Plane (col.Orange(), 0.8, 0.8, 2.5, 2.5, 2.5)
}
