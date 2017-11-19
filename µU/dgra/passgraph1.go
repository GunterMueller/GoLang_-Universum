package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

//import
//  . "µU/obj"

func (x *distributedGraph) passgraph1() {
  x.connect (nil)
  defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Sub (x.actVertex)
  x.tmpGraph.SubAllEdges()
  x.tmpGraph.Write()
  x.log0 ("initial situation")
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
      g := x.ch[i].Recv()
      x.add (g, i)
      x.tmpGraph.Sub2 (x.actVertex,x.nb[i])
      x.tmpGraph.Write()
    }
    if x.tmpGraph.EqSub() {
      if r < x.diameter {
        x.log0 ("network known")
      }
    } else {
      x.log ("after round", r)
    }
  }
}
