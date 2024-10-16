package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

// see G. Andrews: Concurrent Programming (1991) p. 375
// >>> The condition to leave the for-loop by the break statement is
//     not correct; so unfortunately Andrews algorithm is not correct !

import (
  . "µU/obj"
  "µU/adj"
)

func (x *distributedGraph) HeartbeatMatrix1() {
  x.connect (nil)
  defer x.fin()
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  if x.demo { x.matrix.Write (0, 0) }
  for r := uint(1); true; r++ {
    s := x.matrix.Encode()
    for i := uint(0); i < x.n; i++ {
      x.send (i, append(Encode(true), s...))
    }
    for i := uint(0); i < x.n; i++ {
      s = x.ch[i].Recv().(Stream)
      if ! Decode (false, s[:1]).(bool) {
        active[i] = false
      }
      a := adj.New (x.size, uint(0), uint(0))
      a.Decode (s[1:])
      x.matrix.Add (a)
      if x.demo { x.matrix.Write (0, 0) }
    }
    if x.matrix.Full() { // there might be undetected edges !
      break
    }
  }
}
