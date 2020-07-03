package dgra

import "nU/adj"

func (x *distributedGraph) heartbeatmatrix() {
  x.connect (x.matrix)
  defer x.fin()
  x.matrix.Write()
  pause()
  for r:= uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.matrix)
    }
    for i := uint(0); i < x.n; i++ {
      a := x.ch[i].Recv().(adj.AdjacencyMatrix)
      x.matrix.Add (a)
      x.matrix.Write()
      pause()
    }
  }
}
