package main

import (. "nU/obj"; "nU/ego"; "nU/col"; "nU/scr"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8(me)
  g.SetRoot(4)
/*
  g.SetTravAlgorithm (dgra.DFS)
  g.SetTravAlgorithm (dgra.DFS1)
  g.SetTravAlgorithm (dgra.DFSfm1)
  g.SetTravAlgorithm (dgra.Awerbuch)
  g.SetTravAlgorithm (dgra.Awerbuch1)
  g.SetTravAlgorithm (dgra.HelaryRaynal)
  g.SetTravAlgorithm (dgra.Ring)
  g.SetTravAlgorithm (dgra.Ring1)
  g.SetTravAlgorithm (dgra.BFS)
  g.SetTravAlgorithm (dgra.BFSfm)
  g.SetTravAlgorithm (dgra.BFSfm1)
*/
  g.SetTravAlgorithm (dgra.DFS)
  g.Trav (Ignore)
  switch g.TravAlgorithm() {
  case dgra.DFS, dgra.Awerbuch, dgra.HelaryRaynal,
       dgra.BFS, dgra.BFSfm:
    scr.Write ("father:     child[ren]:", 0, 0)
    scr.ColourF (col.LightBlue())
    scr.WriteNat (g.Parent(), 0, 8)
    scr.ColourF (col.Orange())
    scr.Write (g.Children(), 0, 24)
    if g.TravAlgorithm() == dgra.DFS {
      scr.ColourF (col.White())
      scr.Write ("arrival    / departure", 1, 0)
      scr.ColourF (col.Green())
      scr.WriteNat (g.Time(), 1, 8)
      scr.ColourF (col.Red())
      scr.WriteNat (g.Time1(), 1, 23)
    }
  case dgra.Ring:
    scr.ColourF (col.Yellow())
    scr.Write ("   is number    in the ring.", 0, 0)
    scr.WriteNat (g.Me(), 0, 0)
    scr.ColourF (col.Green())
    scr.WriteNat (g.Time(), 0, 13)
  }
}
