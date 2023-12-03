package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go
//
// >>> Algorithm of Chang and Roberts: An Improved Algorithm for Decentralized Extrema-
//     Finding in Circular Configurations of Processes. Comm. ACM 22 (1979), 281 - 283

func (x *distributedGraph) changRoberts() {
  x.connect(uint(0))
  defer x.fin()
  out, in := uint(0), uint(1)
  if x.Graph.Outgoing(1) { in, out = out, in }
//  x.ch[out].Send (x.me)
  x.send (out, x.me)
  for {
    id := x.ch[in].Recv().(uint)
    if id < inf {
      if id > x.me {
//        x.ch[out].Send (id)
        x.send (out, id)
      } else if id == x.me {
        x.leader = x.me
//        x.ch[out].Send (inf + x.me)
        x.send (out, inf + x.me)
        return
      }
    } else { // n > inf
      x.leader = id - inf
      if x.leader != x.me {
//        x.ch[out].Send (id)
        x.send (out, id)
      }
      return
    }
  }
}
