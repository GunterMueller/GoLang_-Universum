package dgra

// (c) Christian Maurer   v. 170508 - license see µu.go

// see G. Andrews: Concurrent Programming (1991) p. 375

import
  . "µu/obj"

func (x *distributedGraph) graph1() {
  x.connect (nil); defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Ex (x.actVertex)
  x.tmpGraph.SubLocal()
  x.tmpGraph.Write()
  done := make(chan int, x.n)
  lock := make(chan int, 1)
  lock <- 0
  known := false
  r := uint(1)
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  for ! known {
//    x.enter(r)
    x.tmpGraph.Ex (x.actVertex)
    bs := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(append(Encode(false), bs...))
    }
// x.log("sent round", r)
    qdone := false
    for i := uint(0); i < x.n; i++ {
      go func (k uint) {
        rs := x.ch[k].Recv().([]byte)
        g := x.decodedGraph(rs[1:])
        <-lock
        qdone = Decode(false, rs[:1]).(bool)
        if qdone {
          active[k] = false
        }
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
    r++
  }
  for i := uint(0); i < x.n; i++ {
    if active[i] {
      x.ch[i].Send(append(Encode(true), x.tmpGraph.Encode()...))
    }
  }
  for i := uint(0); i < x.n; i++ {
    if active[i] {
      x.ch[i].Recv()
    }
  }
}
