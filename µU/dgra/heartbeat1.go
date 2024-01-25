package dgra

// (c) Christian Maurer   v. 231229 - license see µU.go

import (
  . "µU/obj"
  "µU/vtx"
)

func (x *distributedGraph) HeartbeatGraph1() {
  x.connect (nil)
  defer x.fin()
  ready := make([]bool, x.n)
  x.tmpGraph.Copy (x.Graph)
  if x.demo { x.tmpGraph.Write() }
  for r := uint(1); true; r++ {
    s := x.tmpGraph.Encode()
    for i := uint(0); i < x.n; i++ {
      x.send (i, append(Encode(false), s...)) // not ready
    }
    for i := uint(0); i < x.n; i++ {
      s = x.recv (i).(Stream)
      if Decode (false, s[:1]).(bool) {
        ready[i] = true
      }
      g := x.emptyGraph()
      g.Decode (s[1:])
      x.tmpGraph.Add (g)
//      x.tmpGraph.Mark2 (x.actVertex, x.nb[i])
      nbi := x.Graph.Neighbour(i).(vtx.Vertex)
      x.tmpGraph.Mark2 (x.actVertex, nbi)
      if x.demo { x.tmpGraph.Write() }
    }
    if x.tmpGraph.AllMarked() {
      break
    }
  }
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
      x.send (i, append(Encode(true), x.tmpGraph.Encode()...)) // ready
    }
  }
/*
  for i := uint(0); i < x.n; i++ {
    if ! ready[i] {
      x.recv (i)
    }
  }
*/
}
