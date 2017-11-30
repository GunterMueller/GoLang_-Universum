package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

import "nU/adj"

func (x *distributedGraph) pulsematrix() {
  x.connect (x.matrix)
  defer x.fin()
  for r:= uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.matrix)
    }
    for i := uint(0); i < x.n; i++ {
      a := x.ch[i].Recv().(adj.AdjacencyMatrix)
      x.matrix.Add (a)
    }
  }
}
