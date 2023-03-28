package main

import ("µU/col"; "µU/gl"; "µU/scr"; . "µU/fig3")

func main() {
  s := scr.NewWH (0, 0, 800, 600); defer s.Fin()
  gl.ClearColour (col.LightWhite()); 
  s.Go (scr.Look, draw, 3,-1, 10, 3,-1, 0, 0, 1, 0)
}

func draw() {
  r, o, y, g := col.Red(), col.Orange(), col.Yellow(), col.Green()
  c, b, m, n := col.Cyan(), col.Blue(), col.Magenta(), col.Brown()
  MultipyramidC ([]col.Colour {m, n, r, o, y, g, c, b}, 
    0, 2, 0, 2, 3, 3, 1, 4, -1, 4, -2, 3, -2, 2, -1, 0, 2, 0)
  OctahedronC ([]col.Colour{r,o,y,g,c,b,m,n}, 6, 2, 0, 1.4)
  PrismC ([]col.Colour {c, b, m, r, o, y, g}, 1, 0, 1.5, 1, -2, 0,
    -1, -1, 0, -2, -2, 0, 0, -3, 0, -1, -4, 0, 0, -5, 0, 2, -4, 0)
  ParallelepipedC ([]col.Colour{r,o,y,g,b,m}, 5, -2, -1,
                   2, 0, 1, 1, -2, 0, -1.5, 0, 2)
}
