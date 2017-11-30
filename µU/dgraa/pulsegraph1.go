package dgra

// (c) Christian Maurer   v. 171121 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) pulsegraph1() {
  x.connect (nil)
  defer x.fin()
  ready := make([]bool, x.n)
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Sub (x.actVertex)
  x.tmpGraph.Sub1()
  if x.demo { x.tmpGraph.Write() }
  x.log0 ("initial situation")
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
      x.tmpGraph.Sub2 (x.actVertex, x.nb[i])
      if x.demo { x.tmpGraph.Write() }
    }
    if x.tmpGraph.EqSub() {
      break
    } else {
      x.log ("situation after pulse", r)
    }
  }
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
      x.ch[i].Send(append(Encode(true), x.tmpGraph.Encode()...)) // ready
    }
  }
}
