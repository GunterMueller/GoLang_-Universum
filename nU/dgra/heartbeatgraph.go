package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

import
  . "nU/obj"

func (x *distributedGraph) HeartbeatGraph() {
  x.connect (nil)
  defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.Write()
  pause()
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
      g := x.emptyGraph()
      g.Decode (x.ch[i].Recv().(Stream))
      x.tmpGraph.Add (g)
      x.tmpGraph.Write()
      pause()
    }
  }
}
