package main

import ("nU/ego"; "nU/col"; "nU/scr"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8dirring(me)
  g.SetRoot(4)
/*
  g.SetElectAlgorithm(dgra.ChangRoberts)
  g.SetElectAlgorithm(dgra.Peterson)
  g.SetElectAlgorithm(dgra.HirschbergSinclair)
*/
  g.SetElectAlgorithm(dgra.ChangRoberts)
  scr.Write ("The leader is", 0, 0)
  scr.ColourF (col.Yellow())
  scr.WriteNat (g.Leader(), 0, 14)
}
