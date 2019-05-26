package main

// (c) Christian Maurer   v. 190402 - license see nsp.go

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
  g.SetElectAlgorithm(dgra.KorachMoranZaks)
*/
  g.SetElectAlgorithm(dgra.KorachMoranZaks)
  scr.Write ("Der Leiter ist", 0, 0)
  scr.ColourF (col.Yellow())
  scr.WriteNat (g.Leader(), 0, 15)
}
