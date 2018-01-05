package main

import ("nU/scr"; "nU/ego"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8 (me)
  g.Write()
}
