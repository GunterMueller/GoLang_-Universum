package main

import ("nU/scr"; "nU/ego"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  if me > 8 { panic("no or wrong argument") }
  g := dgra.G8 (me)
  g.Write()
}
