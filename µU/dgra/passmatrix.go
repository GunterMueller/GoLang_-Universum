package dgra

// (c) Christian Maurer   v. 171112 - license see µU.go

func (x *distributedGraph) passmatrix() {
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
      if x.demo { x.top.Write(0, 0) }
    }
  }
}
