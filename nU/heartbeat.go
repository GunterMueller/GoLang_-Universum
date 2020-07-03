package main

import ("nU/ego"; "nU/col"; "nU/scr"; "nU/dgra")

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8 (me)
  g.SetRoot (4)
/*
  g.SetHeartbeatAlgorithm (dgra.HeartbeatMatrix)
  g.SetHeartbeatAlgorithm (dgra.HeartbeatGraph)
  g.SetHeartbeatAlgorithm (dgra.HeartbeatGraph1)
*/
  g.SetHeartbeatAlgorithm (dgra.HeartbeatMatrix)
  g.Heartbeat()
  switch g.HeartbeatAlgorithm() {
  case dgra.HeartbeatMatrix:
    scr.ColourF (col.Red())
    scr.Write ("complete adjacency matrix", 8, 0)
    scr.ColourF (col.White())
  }
}
