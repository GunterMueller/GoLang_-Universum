package main

// (c) Christian Maurer  v. 230313 - license see µU.go

import (
  "µU/col"
  "µU/gl"
  "µU/scr"
  . "µU/fig3"
)

func main() {
  s := scr.NewWH (0, 0, 800, 600); defer s.Fin()
  gl.ClearColour (col.LightWhite()); 
  s.Go (scr.Look, draw, 0, -6, 1, 0, 0, 1, 0, 0, 1)
}

func draw() {
  r, o, y, g := col.Red(), col.Orange(), col.Yellow(), col.Green()
  b, m, s := col.Blue(), col.Magenta(), col.Black()
  OctopusC ([]col.Colour{r,g,m,o,b,y,s},
            0, 0, 3,
            3, 0, 0,
            2, 1, 2,
            0, 2, 1,
           -2, 1, 0,
           -2,-1, 0,
            1,-2, 1)
}
