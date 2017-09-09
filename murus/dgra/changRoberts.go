package dgra

// (c) Christian Maurer   v. 170423 - license see murus.go
//
// >>> Algorithm of Chang and Roberts: An Improved Algorithm for Decentralized Extrema-
//     Finding in Circular Configurations of Processes. Comm. ACM 22 (1979), 281 - 283

func (x *distributedGraph) changRoberts() {
  x.connect(uint(0))
  defer x.fin()
  id, m := x.me, inf
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  if x.me == x.root {
    x.ch[out].Send (id)
  }
  for {
    n := x.ch[in].Recv().(uint)
    if n < m {
      if n > id {
        x.ch[out].Send (n)
      } else if n < x.me {
        x.ch[out].Send (x.me)
      } else { // r == id
        x.ch[out].Send (m + x.me)
        x.leader = x.me
        return
      }
    } else { // n >= m
      n -= m
      if n != x.me {
        x.ch[out].Send (m + n)
      }
      x.leader = n
      return
    }
  }
}
