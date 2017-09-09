package dgra

// (c) Christian Maurer   v. 170423 - license see murus.go

import
  . "murus/obj"

func (x *distributedGraph) passgraph1() {
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Write()
  x.connect (nil)
  defer x.fin()
  for r:= uint(0); r < x.diameter; r++ {
    x.enter (r + 1)
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.tmpGraph)
    }
    for i := uint(0); i < x.n; i++ {
      g := x.ch[i].Recv()
      x.add1 (g, i)
      x.tmpGraph.Write()
    }
    x.tmpGraph.Write()
  }
}

func (x *distributedGraph) add1 (a Any, i uint) Any {
  x.tmpGraph.Add (x.decodedGraph(a.([]byte)))
  x.tmpGraph.Write()
  return x.tmpGraph
}
