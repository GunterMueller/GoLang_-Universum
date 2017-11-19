package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) passgraph() {
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
      g := x.ch[i].Recv()
      x.add (g, i)
      x.tmpGraph.Write()
    }
    if false { // x.tmpGraph.EqSub() {
      x.log0 ("network known")
    } else {
      x.log ("after round", r)
    }
  }
}

func (x *distributedGraph) add (a Any, i uint) Any {
  x.tmpGraph.Add (x.decodedGraph(a.([]byte)))
  return x.tmpGraph
}
