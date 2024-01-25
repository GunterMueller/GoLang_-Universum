package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

// see G. Andrews: Concurrent Programming (1991) p. 375
// >>> The condition to leave the for-loop by the break statement is
//     not correct; so unfortunately Andrews algorithm is not correct !

import (
  . "nU/obj"
  "nU/adj"
)

func (x *distributedGraph) HeartbeatMatrix1() {
  x.connect (nil)
  defer x.fin()
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  for r := uint(1); true; r++ {
    bs := x.matrix.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(append(Encode(true), bs...))
    }
    for i := uint(0); i < x.n; i++ {
      bs = x.ch[i].Recv().(Stream)
      if ! Decode (false, bs[:1]).(bool) {
        active[i] = false
      }
      a := adj.New (x.size, uint(0), uint(0))
      a.Decode (bs[1:])
      x.matrix.Add (a)
    }
    if x.matrix.Full() { // there might be undetected edges !
      break
    }
  }
}
