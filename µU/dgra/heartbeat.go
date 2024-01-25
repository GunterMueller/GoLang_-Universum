package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) HeartbeatGraph() {
  x.connect (nil)
  defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.Write()
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.send (i, x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
      s := x.ch[i].Recv().(Stream)
      g := x.decodedGraph (s)
      x.tmpGraph.Add (g)
      x.tmpGraph.Write()
    }
  }
}
