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
  gl.ClearColour (col.FlashWhite()); gl.Clear()
  s.Go (draw, 5, -3, 2, 0, 0, 0, 0, 1, 0)
}

func draw() {
  DoubleCone (col.Red(), 0, 0, 0, 1, 3)
  VertRectangle (col.Green(), .3, -1, -3, .3, 1, 3)
}
