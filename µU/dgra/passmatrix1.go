package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

// see G. Andrews: Concurrent Programming (1991) p. 375

import (
  . "µU/obj"
  "µU/adj"
)

func (x *distributedGraph) passmatrix1() {
  x.connect (nil)
  defer x.fin()
  wait := make(chan int, x.n)
  lock := make(chan int, 1)
  lock <- 0
  done := false
  r := uint(1)
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  if x.demo { x.top.Write (0, 0) }
  for ! done {
    bs := x.top.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(append(Encode(false), bs...))
    }
    qdone := false
    for i := uint(0); i < x.n; i++ {
      go func (k uint) {
        rs := x.ch[k].Recv().([]byte)
        m := adj.New (8, uint(0), uint(0))
        m.Decode (rs[1:])
        <-lock
        qdone = Decode(false, rs[:1]).(bool)
        if qdone {
          active[k] = false
        }
        x.addMatrix (m, i)
        x.top.Write (0, 0)
        lock <-0
        wait <-0
      }(i)
    }
    for i := uint(0); i < x.n; i++ {
      <-wait
    }
    if ! done && x.top.Full() {
      done = true
      x.log0 ("only one more round")
    } else {
      x.log ("after round", r)
    }
    r++
  }
  for i := uint(0); i < x.n; i++ {
    if active[i] {
      x.ch[i].Send(append(Encode(true), x.top.Encode()...))
    }
  }
  for i := uint(0); i < x.n; i++ {
    if active[i] {
      x.ch[i].Recv()
    }
  }
}
