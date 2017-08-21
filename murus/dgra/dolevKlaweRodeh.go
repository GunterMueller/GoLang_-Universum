package dgra

// (c) murus.org  v. 170423 - license see murus.go

// >>> D. Dolev, M. Klawe, M. Rodeh: An O(n log n) Unidirectional Algorithm
//     for Extrema Finding in a Circle. J. Algorithms 3 (1982), 245-260

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
//      x.ch[out].Send (max + x.uint)
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
//    if i > x.uint {
    if i > inf {
//      x.leader = i - x.uint
      x.leader = i - inf
      break
    }
  }
}
