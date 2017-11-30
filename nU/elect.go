package main

// (c) Christian Maurer   v. 171127 - license see nsp.go

import ("nU/ego"; "nU/dgra")

func main() {
  me := ego.Me()
  g := dgra.G8ringdir(me)
  g.SetRoot(4)
/*
  g.SetElectAlgorithm(dgra.ChangRoberts)
  g.SetElectAlgorithm(dgra.Peterson)
  g.SetElectAlgorithm(dgra.DolevKlaweRodeh)
  g.SetElectAlgorithm(dgra.HirschbergSinclair)
  g.SetElectAlgorithm(dgra.DFSE)
  g.SetElectAlgorithm(dgra.FmDFSE)
*/
  g.SetElectAlgorithm(dgra.ChangRoberts)
  println ("der Leiter ist", g.Leader())
}
