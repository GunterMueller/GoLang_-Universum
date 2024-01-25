package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go
//
// >>> Algorithm of Chang and Roberts: An Improved Algorithm for Decentralized Extrema-
//     Finding in Circular Configurations of Processes. Comm. ACM 22 (1979), 281 - 283

// import "µU/errh"

func (x *distributedGraph) ChangRoberts() {
  x.connect (uint(0))
  defer x.fin()
  x.Graph.ExVal (x.me) // my vertex is now the local one
  in, out := uint(0), uint(1)
  if x.Outgoing(1) { in, out = out, in}
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
