package dgra

// (c) Christian Maurer   v. 171227 - license see nU.go

import . "nU/obj"

func (x *distributedGraph) pulsegraph1() {
  x.connect (nil)
  defer x.fin()
  ready := make([]bool, x.n)
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Write()
  pause()
  for r := uint(1); true; r++ {
    bs := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
      x.ch[i].Send(append(Encode(false), bs...)) // not ready
    }
    for i := uint(0); i < x.n; i++ {
      bs = x.ch[i].Recv().(Stream)
      if Decode (false, bs[:1]).(bool) {
        ready[i] = true
      }
      g := x.emptyGraph()
      g.Decode (bs[1:])
      x.tmpGraph.Add (g)
      x.tmpGraph.Mark2 (x.actVertex, x.nb[i])
      x.tmpGraph.Write()
      pause()
    }
    if x.tmpGraph.AllMarked() {
      break
    }
  }
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
      x.ch[i].Send(append(Encode(true), x.tmpGraph.Encode()...)) // ready
    }
  }
}
