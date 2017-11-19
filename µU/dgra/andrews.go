package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

// see G. Andrews: Concurrent Programming (1991) p. 375

import
  . "µU/obj"

func (x *distributedGraph) andrews() {
  x.connect (nil); defer x.fin()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Sub (x.actVertex)
  x.tmpGraph.SubAllEdges()
  x.tmpGraph.Write()
  x.log0 ("initial situation")
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
    x.tmpGraph.Ex (x.actVertex)
    bs := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(append(Encode(false), bs...))
    }
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
        x.tmpGraph.Sub2 (x.actVertex, x.nb[k])
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
      x.log0 ("network known")
    } else {
      x.log ("after round", r)
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
