package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

func (x *distributedGraph) dolevKlaweRodeh() {
  x.connect(uint(0))
  defer x.fin()
  var left uint
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  max := x.me
  x.ch[out].Send (max) // (x.me)
  for {
    i := x.ch[in].Recv().(uint)
    if i == max {
      x.ch[out].Send (max + inf)
      x.leader = max
      return
    } else { // i != max
      left = i
      x.ch[out].Send (i)
    }
    i = x.ch[in].Recv().(uint)
    if left > i && left > max {
      max = left
      x.ch[out].Send (max)
    } else {
      x.leader = max
      break
    }
  }
  for {
    i := x.ch[in].Recv().(uint)
    x.ch[out].Send (i)
    if i > inf {
      x.leader = i - inf
      break
    }
  }
}
