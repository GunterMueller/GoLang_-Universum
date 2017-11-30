package dgra

// (c) Christian Maurer   v. 171120 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) pulsegraph() {
  x.connect (nil)
  defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.Write()
  x.log0 ("initial situation")
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
//      g := x.emptyGraph()
//      g.Decode (x.ch[i].Recv().(Stream))
      g := x.decodedGraph (x.ch[i].Recv().(Stream))
      x.tmpGraph.Add (g)
      x.tmpGraph.Write()
    }
    x.log ("situation after pulse", r)
  }
}
