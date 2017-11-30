package main

// (c) Christian Maurer   v. 171127 - license see µU.go

import (. "nU/obj"; "nU/ego"; "nU/dgra")

func main() {
  me := ego.Me()
  g := dgra.G8(me); g.SetRoot(4)
/*
  g.SetTravAlgorithm (dgra.DFS)
  g.SetTravAlgorithm (dgra.FmDFSA)
  g.SetTravAlgorithm (dgra.FmDFSRing)
  g.SetTravAlgorithm (dgra.BFS)
  g.SetTravAlgorithm (dgra.FmBFS)
*/
  g.SetTravAlgorithm (dgra.DFS1)
  g.Trav (Ignore)
  switch g.TravAlgorithm() {
  case dgra.DFS:
    println (g.ParentChildren())
    println ("Start-/Endzeit:", g.Time(), "/", g.Time1())
  case dgra.FmDFSA, dgra.FmBFS:
    println (g.ParentChildren())
  case dgra.FmDFSRing:
    println ("#", g.Time(), "is", g.Me())
  case dgra.BFS:
    println (g.ParentChildren())
  }
}
