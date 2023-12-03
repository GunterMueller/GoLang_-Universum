package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

// see G. Andrews: Concurrent Programming (1991) p. 375
// >>> The condition to leave the for-loop by the break statement is
//     not correct; so unfortunately Andrews algorithm is not correct !

import (
  . "µU/obj"
  "µU/adj"
)

func (x *distributedGraph) heartbeatmatrix1() {
  x.connect (nil)
  defer x.fin()
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  if x.demo { x.matrix.Write (0, 0) }
  x.log0 ("initial situation")
  for r := uint(1); true; r++ {
    bs := x.matrix.Encode()
    for i := uint(0); i < x.n; i++ {
//      x.ch[i].Send(append(Encode(true), bs...))
      x.send (i, append(Encode(true), bs...))
    }
    for i := uint(0); i < x.n; i++ {
      bs = x.ch[i].Recv().(Stream)
      if ! Decode (false, bs[:1]).(bool) {
        active[i] = false
      }
      a := adj.New (x.size, uint(0), uint(0))
      a.Decode (bs[1:])
      x.matrix.Add (a)
      if x.demo { x.matrix.Write (0, 0) }
    }
    if x.matrix.Full() { // there might be undetected edges !
      break
    } else {
      x.log ("situation after heartbeat", r)
    }
  }
}
