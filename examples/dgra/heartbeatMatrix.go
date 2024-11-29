package main

// (c) Christian Maurer   v. 241030 - license see µU.go

import (
  "µU/ego"
  "µU/scr"
  "µU/wm"
  "µU/errh"
  "µU/dgra"
)

func main() {
  me := ego.Me()
  w, h := uint(15 * 8), uint(9 * 16)
  x, y := (me % 4) * (w + 2 * wm.B), (me / 4) * (h + wm.T)
  scr.NewWH (x, y, w, h); defer scr.Fin()
  g := dgra.G8 (me)
  g.Demo()
  g.SetRoot (4)
  scr.Cls()
  g.HeartbeatMatrix()
  errh.Error0 ("Adjazenzmatrix")
}
