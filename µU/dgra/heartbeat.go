package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) HeartbeatGraph() {
  x.connect (nil)
  defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.Write()
  x.log0 ("initial situation")
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
//      x.ch[i].Send (x.tmpGraph)
      x.send (i, x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
      bs := x.ch[i].Recv().(Stream)
      g := x.decodedGraph (bs)
      x.tmpGraph.Add (g)
      x.tmpGraph.Write()
    }
    x.log ("situation after heartbeat", r)
  }
}
