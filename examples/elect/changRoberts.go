package main

// (c) Christian Maurer   v. 241007 - license see µU.go

import (
  "µU/ego"
  "µU/env"
  "µU/scr"
  "µU/errh"
  "µU/N"
  "µU/dgra"
)

func main() {
  me := ego.Me()
  x, y := (me % 4) * (256 + 8), (me / 4) * (192 + 8 + 16)
  scr.NewWH (x, y, 256, 192); defer scr.Fin()
  scr.Name (env.Call() + " " + N.String (me))
  g := dgra.G8ringdir (ego.Me())
  g.SetRoot(4)
  g.ChangRoberts()
  errh.Error ("the leader is", g.Leader())
}
