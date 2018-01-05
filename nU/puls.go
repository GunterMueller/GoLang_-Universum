package main

// (c) Christian Maurer   v. 171227 - license see nsp.go

import ("nU/ego"; "nU/col"; "nU/scr"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8 (me)
  g.SetRoot (4)
/*
  g.SetPulseAlgorithm (dgra.PulseMatrix)
  g.SetPulseAlgorithm (dgra.PulseGraph)
  g.SetPulseAlgorithm (dgra.PulseGraph1)
*/
  g.SetPulseAlgorithm (dgra.PulseMatrix)
  g.Pulse()
  switch g.PulseAlgorithm() {
  case dgra.PulseMatrix:
    scr.ColourF (col.Red())
    scr.Write ("vollständige Adjazenzmatrix", 8, 0)
    scr.ColourF (col.White())
  }
}
