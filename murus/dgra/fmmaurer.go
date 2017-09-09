package dgra

// (c) Christian Maurer   v. 170508 - license see murus.go

import (
  . "murus/obj"
  "murus/fmon"
)

func (x *distributedGraph) fmMaurer() {
  in, out := uint(1), uint(0); if x.Graph.Outgoing(1) { in, out = out, in }
  go func() { fmon.New (nil, 1, x.addMeVertex, AllTrueSp, x.actHost, p0 + uint16(x.me), true) }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 1, x.addMeVertex, AllTrueSp, x.host[i], p0 + uint16(x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.ring.Clr()
  if x.me == x.root {
    x.ring.Ins (x.actVertex)
    x.ring.Write()
    bs := x.ring.Encode()
    bs = x.mon[out].F(bs, 0).([]byte)
    x.ring = x.decodedGraph(bs)
    x.ring.Write()
  } else {
    <-done
  }
  x.ring.Write()
  x.leader = valueMax (x.ring)
  exValue (x.ring, x.leader)
  x.ring.SubLocal()
  x.ring.Write()
}

func (x *distributedGraph) addMeVertex (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.([]byte)
  x.ring = x.decodedGraph(bs)
  s := nrLocal(x.ring)
  j := x.channel(s)
  out := uint(0); if j == 0 { out++ }
  x.ring.Ins (x.actVertex)
  if x.me != x.root {
    x.tmpGraph.Ex2(x.nb[j], x.actVertex)
    x.ring.Edge(x.tmpGraph.Get1())
    if x.nr[out] == x.root {
      x.ring.Ex2(x.actVertex, x.nb[out])
      x.tmpGraph.Ex2(x.actVertex, x.nb[out])
      x.ring.Edge(x.tmpGraph.Get1())
    }
    x.ring.Write()
    bs = x.ring.Encode()
    bs = x.mon[out].F(bs, 0).([]byte)
    x.ring = x.decodedGraph(bs)
    x.ring.Write()
  }
  done <- 0
  return bs
}
