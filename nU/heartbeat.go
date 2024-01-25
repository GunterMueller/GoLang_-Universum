package main

import (
  "nU/ego"
  "nU/scr"
  "nU/dgra"
)

func main() {
  scr.New(); defer scr.Fin()
  me := ego.Me()
  g := dgra.G8 (me)
  g.SetRoot (4)
/*/
/*/
  g.HeartbeatMatrix(); scr.Write ("complete adjacency matrix", 8, 0)
  g.HeartbeatMatrix1()
  g.HeartbeatGraph()
  g.HeartbeatGraph1()
/*/
/*/
}
