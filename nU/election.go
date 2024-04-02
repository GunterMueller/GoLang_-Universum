package main

import (
  "nU/ego"
  "nU/scr"
  "nU/dgra"
)

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8ring(me)
  g.SetRoot(4)
/*/
  g.ChangRoberts()
  g.Peterson()           // XXX
  g.DolevKlaweRodeh()    // XXX
  g.HirschbergSinclair() // XXX
  g.Dfselect()
  g.Dfselectfm()
/*/
  g.Peterson()
  if g.Leader() == me {
    scr.Write ("I am leader", 0, 0)
  } else {
    scr.Write ("leader is", 0, 0); scr.WriteNat (g.Leader(), 0, 14)
  }
}
