package main

// (c) Christian Maurer  v. 230326 - license see µU.go

import (
  "µU/col"
  "µU/gl"
  "µU/scr"
)

func draw() {
  gl.ClearColour (col.Black())
  gl.Clear()
  gl.Begin (gl.Quads)
  gl.Colour (col.Red())
  gl.Vertex (-1.,-1., 1.) // front
  gl.Vertex ( 1.,-1., 1.)
  gl.Vertex ( 1.,-1.,-1.)
  gl.Vertex (-1.,-1.,-1.)
  gl.Colour (col.Yellow())
  gl.Vertex ( 1.,-1., 1.) // right
  gl.Vertex ( 1., 1., 1.)
  gl.Vertex ( 1., 1.,-1.)
  gl.Vertex ( 1.,-1.,-1.)
  gl.Colour (col.Cyan())
  gl.Vertex (-1., 1., 1.) // back
  gl.Vertex ( 1., 1., 1.)
  gl.Vertex ( 1., 1.,-1.)
  gl.Vertex (-1., 1.,-1.)
  gl.Colour (col.Magenta())
  gl.Vertex (-1.,-1., 1.) // left
  gl.Vertex (-1., 1., 1.)
  gl.Vertex (-1., 1.,-1.)
  gl.Vertex (-1.,-1.,-1.)
  gl.Colour (col.Blue())
  gl.Vertex (-1., 1., 1.) // top
  gl.Vertex ( 1., 1., 1.)
  gl.Vertex ( 1.,-1., 1.)
  gl.Vertex (-1.,-1., 1.)
  gl.Colour (col.Green())
  gl.Vertex (-1., 1.,-1.) // bottom
  gl.Vertex ( 1., 1.,-1.)
  gl.Vertex ( 1.,-1.,-1.)
  gl.Vertex (-1.,-1.,-1.)
  gl.End()
}

func main() {
  s := scr.NewWH (0, 0, 1594, 1150); defer s.Fin()
  s.Cls()
  s.Go (draw, 3, -3, 2, 0, 0, 0, 0, 0, 1)
}
