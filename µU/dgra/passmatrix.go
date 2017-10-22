package dgra

// (c) Christian Maurer   v. 170504 - license see µU.go

import (
  "µU/scr"
  "µU/adj"
)

func (x *distributedGraph) passmatrix() {
  x.top = adj.New (x.rank, uint(1))
  for i := uint(0); i < x.n; i++ {
    x.top.Set (x.me, x.nr[i], uint(1))
    x.top.Set (x.nr[i], x.me, uint(1))
  }
  x.connect (x.top)
  defer x.fin()
  if x.demo { x.top.Write(0, 0) }
  for r:= uint(0); r < x.diameter; r++ {
    x.enter (r + 1)
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send (x.top)
    }
    for i := uint(0); i < x.n; i++ {
      m := x.ch[i].Recv()
      x.addMatrix (m, i)
      if x.demo { x.top.Write(0, scr.NColumns() / 2) }
    }
  }
}
