package main

// (c) Christian Maurer   v. 241011 - license see µU.go

// Open 8 small windows and start the program in the 1st one with the parameter "0",
// in the 2nd with the parameter "1", in the 3rd with the parameter "2" and so on.

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
  g := dgra.G8 (me)
  g.SetRoot(4)
  g.Ring1()
  errh.Error2 ("number", g.Time(), "in the ring is", g.Me())
}
