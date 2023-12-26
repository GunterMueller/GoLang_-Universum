package dgra

// (c) Christian Maurer   v. 231215 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) Maurerfm() {
  in, out := uint(1), uint(0); if x.Outgoing(1) { in, out = out, in }
  go func() {
    fmon.New (nil, 1, x.m, AllTrueSp,
              x.actHost, p0 + uint16(x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 1, x.m, AllTrueSp,
                         x.host[i], p0 + uint16(x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.cycle.Clr()
  if x.me == x.root {
    x.cycle.Ins (x.actVertex)
    x.cycle.Write()
    bs := x.cycle.Encode()
    bs = x.mon[out].F(bs, 0).(Stream)
    x.cycle = x.decodedGraph (bs)
    x.cycle.Write()
  } else {
    <-done
  }
  x.cycle.Write()
  x.leader = valueMax (x.cycle)
  vm := vertexMax (x.cycle)
  exValue (x.cycle, x.leader)
  x.cycle.Mark (vm)
  x.cycle.Write()
}

func (x *distributedGraph) m (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.cycle = x.decodedGraph(bs)
  s := nrLocal(x.cycle)
  j := x.channel(s)
  out := uint(0); if j == 0 { out++ }
  x.cycle.Ins (x.actVertex)
  if x.me != x.root {
    x.tmpGraph.Ex2 (x.nb[j], x.actVertex)
    x.cycle.Edge (x.tmpGraph.Get1())
    if x.nr[out] == x.root {
      x.cycle.Ex2 (x.actVertex, x.nb[out])
      x.tmpGraph.Ex2 (x.actVertex, x.nb[out])
      x.cycle.Edge (x.tmpGraph.Get1())
    }
    x.cycle.Write()
    bs = x.cycle.Encode()
    bs = x.mon[out].F(bs, 0).(Stream)
    x.cycle = x.decodedGraph(bs)
    x.cycle.Write()
  }
  done <- 0
  return bs
}
