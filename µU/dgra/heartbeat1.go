package dgra

// (c) Christian Maurer   v. 200728 - license see µU.go

import
  . "µU/obj"

func (x *distributedGraph) heartbeatgraph1() {
  x.connect (nil)
  defer x.fin()
  ready := make([]bool, x.n)
  x.tmpGraph.Copy (x.Graph)
  if x.demo { x.tmpGraph.Write() }
  x.log0 ("initial situation")
  for r := uint(1); true; r++ {
    bs := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
//      x.ch[i].Send(append(Encode(false), bs...)) // not ready
      x.send (i, append(Encode(false), bs...)) // not ready
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
      if x.demo { x.tmpGraph.Write() }
    }
    if x.tmpGraph.AllMarked() {
      break
    } else {
      x.log ("situation after heartbeat", r)
    }
  }
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
//      x.ch[i].Send(append(Encode(true), x.tmpGraph.Encode()...)) // ready
      x.send (i, append(Encode(true), x.tmpGraph.Encode()...)) // ready
    }
  }
/*
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
      x.ch[i].Recv()
    }
  }
*/
}
