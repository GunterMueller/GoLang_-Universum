package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

import
  "µU/adj"

func (x *distributedGraph) HeartbeatMatrix() {
  x.connect (x.matrix)
  defer x.fin()
  if x.demo { x.matrix.Write(0, 0) }
  for r := uint(1); r <= x.diameter; r++ {
    for i := uint(0); i < x.n; i++ {
      x.send (i, x.matrix)
    }
    for i := uint(0); i < x.n; i++ {
      a := x.recv (i).(adj.AdjacencyMatrix)
      x.matrix.Add (a)
      if x.demo { x.matrix.Write(0, 0) }
    }
  }
}
