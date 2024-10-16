package dgra

// (c) Christian Maurer   v. 240927 - license see µU.go
//
// >>> Algorithm of Chang and Roberts: An Improved Algorithm for Decentralized Extrema-
//     Finding in Circular Configurations of Processes. Comm. ACM 22 (1979), 281 - 283

const
  P = 8 // number of involved processes

func (x *distributedGraph) ChangRoberts() {
  defer x.fin()
  x.connect(uint(0))
  in, out := uint(0), uint(1)
  if x.Graph.Outgoing(1) { in, out = out, in }
  x.send (out, x.me)
  for {
    id := x.recv (in).(uint)
    if id < inf {
      if id > x.me {
        x.send (out, id)
      } else if id == x.me {
        x.leader = x.me
        x.send (out, inf + x.me)
        return
      }
    } else { // id >= inf
      x.leader = id - inf
      if x.leader != x.me {
        x.send (out, id)
      }
      return
    }
  }
}
