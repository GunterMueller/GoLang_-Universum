package main

// (c) Christian Maurer   v. 241006 - license see µU.go

import (
  "µU/ego"
  "µU/scr"
  "µU/errh"
  "µU/dgra"
)

func main() {
  me := ego.Me()
  w, h := uint(120), uint(144)
  x, y := (me % 4) * (w + 8), (me / 4) * (h + 28)
  scr.NewWH (x, y, w, h); defer scr.Fin()
  g := dgra.G8 (me)
  g.Demo()
  g.SetRoot (4)
  scr.Cls()
  g.HeartbeatMatrix()
  errh.Error0 ("Adjazenzmatrix")
}
