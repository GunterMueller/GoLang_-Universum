package main

import ("nU/scr"; "nU/ego"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  if me >= 8 { println("Gebrauch: \"me n\" mit 0 <= n <= 7"); return }
  g := dgra.G8 (me)
  g.Write()
}
