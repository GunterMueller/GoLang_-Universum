package main

// (c) Christian Maurer   v. 171127 - license see nsp.go

import ("nU/ego"; "nU/dgra")

func main() {
  me := ego.Me()
  g := dgra.G8 (me)
  g.SetRoot (4)
/*
  g.SetPulseAlgorithm (dgra.PulseMatrix)
  g.SetPulseAlgorithm (dgra.PulseGraph)
*/
  g.SetPulseAlgorithm (dgra.PulseGraph)
  g.Pulse()
  g.Write()
}
