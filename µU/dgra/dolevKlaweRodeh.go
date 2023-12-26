package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

// >>> D. Dolev, M. Klawe, M. Rodeh: An O(n log n) Unidirectional Algorithm
//     for Extrema Finding in a Circle. J. Algorithms 3 (1982), 245-260

func (x *distributedGraph) DolevKlaweRodeh() {
  x.connect(uint(0))
  defer x.fin()
  var left uint
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  max := x.me
//  x.ch[out].Send (max) // (x.me)
  x.send (out, max) // (x.me)
  for {
    i := x.ch[in].Recv().(uint)
    if i == max {
//      x.ch[out].Send (max + inf)
      x.send (out, max + inf)
      x.leader = max
      return
    } else { // i != max
      left = i
//      x.ch[out].Send (i)
      x.send (out, i)
    }
    i = x.ch[in].Recv().(uint)
    if left > i && left > max {
      max = left
//      x.ch[out].Send (max)
      x.send (out, max)
    } else {
      x.leader = max
      break
    }
  }
  for {
    i := x.ch[in].Recv().(uint)
//    x.ch[out].Send (i)
    x.send (out, i)
    if i > inf {
      x.leader = i - inf
      break
    }
  }
}
