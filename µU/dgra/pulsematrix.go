package dgra

// (c) Christian Maurer   v. 171120 - license see µU.go

import
  "µU/adj"

func (x *distributedGraph) pulsematrix() {
  x.connect (x.matrix)
  defer x.fin()
  if x.demo { x.matrix.Write(0, 0) }
  x.log0 ("initial situation")
  for r:= uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.matrix)
    }
    for i := uint(0); i < x.n; i++ {
      a := x.ch[i].Recv().(adj.AdjacencyMatrix)
      x.matrix.Add (a)
      if x.demo { x.matrix.Write(0, 0) }
    }
    x.log ("situation after pulse", r)
  }
}
