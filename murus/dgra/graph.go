package dgra

// (c) Christian Maurer   v. 170506 - license see murus.go

func (x *distributedGraph) graph() {
  x.connect (nil); defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.SubLocal()
  x.tmpGraph.Write()
  done := make(chan int, x.n)
  lock := make(chan int, 1)
  lock <- 0
  known := false
  for r := uint(1); r <= x.diameter; r++ {
//    x.enter(r)
    x.tmpGraph.Ex (x.actVertex)
    bs := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(bs)
    }
    x.log("sent round", r)
    for i := uint(0); i < x.n; i++ {
      go func (k uint) {
        rs := x.ch[k].Recv().([]byte)
        g := x.decodedGraph(rs)
        <-lock
        x.tmpGraph.Add (g)
        x.tmpGraph.Ex2(x.actVertex, x.nb[k])
        x.tmpGraph.Sub2()
        x.tmpGraph.Write()
        lock <-0
        done <-0
      }(i)
    }
    for i := uint(0); i < x.n; i++ {
      <-done
    }
    if ! known && x.tmpGraph.EqSub() {
      known = true
      x.log("topology in round", r)
    }
  }
}
