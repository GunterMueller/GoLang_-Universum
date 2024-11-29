package main

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  "µU/col"
  "µU/scr"
  "µU/vtx"
  "µU/edg"
  "µU/gra"
  "µU/bn"
  "µU/errh"
)
const
  m = uint(10)
var (
  g gra.Graph
  v = make([]vtx.Vertex, m)
  d = false
  n = bn.New(0)
  e edg.Edge
  w, b = col.FlashWhite(), col.Black()
)

func vertex (i uint, x, y int) {
  n.SetVal (i)
  v[i] = vtx.New (n, 1, 1)
  v[i].Set (x, y)
  v[i].Colours (w, b, w, b)
  g.Ins (v[i])
}

func ins (i, j, k uint) {
  g.Ex2 (v[i], v[j])
  e = edg.New (d, uint(k))
  e.SetPos0 (v[i].Pos())
  e.SetPos1 (v[j].Pos())
  e.Colours (w, b, w, b)
  e.Direct (d); g.Edge (e)
}

func main() {
  scr.NewWH (0, 0, 500, 400); defer scr.Fin()
  v0, e0 := vtx.New (bn.New(0), 1, 1), edg.New (false, uint(1))
  g = gra.New (d, v0, e0)
  vertex (0, 150, 100)
  vertex (1, 250, 100)
  vertex (2, 350, 100)
  vertex (3, 100, 200)
  vertex (4, 200, 200)
  vertex (5, 300, 200)
  vertex (6, 400, 200)
  vertex (7, 150, 300)
  vertex (8, 250, 300)
  vertex (9, 350, 300)
  ins (0, 5, 15)
  ins (1, 2, 12)
  ins (1, 5,  2)
  ins (2, 6,  8)
  ins (3, 0,  5)
  ins (3, 4, 14)
  ins (7, 4,  1)
  ins (7, 8, 10)
  ins (8, 9,  4)
  ins (9, 6, 13)
  g.SetWrite (vtx.W, edg.W)
  g.Write()
  if g.IsRing() { errh.Error0 ("is a ring") } else { errh.Error0 ("is not a ring") }
  if g.Euler() { errh.Error0 ("Euler") } else { errh.Error0 ("not Euler") }
}
