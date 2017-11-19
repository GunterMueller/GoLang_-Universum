package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import (
  . "µU/obj"
  "µU/adj"
)

func (x *distributedGraph) passmatrix() {
  x.connect (x.top)
  defer x.fin()
  if x.demo { x.top.Write(0, 0) }
  for r:= uint(1); r <= x.diameter; r++ {
    x.log ("after round", r)
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.top)
    }
    for i := uint(0); i < x.n; i++ {
      m := x.ch[i].Recv()
      x.addMatrix (m, i)
      if x.demo { x.top.Write(0, 0) }
    }
  }
}

func (x *distributedGraph) addMatrix (a Any, i uint) Any {
  x.top.Add (a.(adj.AdjacencyMatrix))
  return x.top
}
