package dgra

// (c) Christian Maurer   v. 231228 - license see nU.go

func (x *distributedGraph) ChangRoberts() {
  x.connect(uint(0))
  defer x.fin()
  out, in := 0, 1
  if x.Graph.Outgoing(1) { in, out = out, in }
  x.ch[out].Send (x.me)
  for {
    id := x.ch[in].Recv().(uint)
    if id < inf {
      if id > x.me {
        x.ch[out].Send (id)
      } else if id == x.me {
        x.leader = x.me
        x.ch[out].Send (inf + x.me)
        return
      }
    } else { // id >= inf
      x.leader = id - inf
      if x.leader != x.me {
        x.ch[out].Send (id)
      }
      return
    }
  }
}
