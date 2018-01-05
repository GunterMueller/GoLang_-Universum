package dgra

// (c) Christian Maurer   v. 171127 - license see nsp.go

import (. "nU/obj"; "nU/vtx"; "nU/fmon")

func (x *distributedGraph) fmMaurer() {
  in, out := uint(1), uint(0); if x.Graph.Outgoing(1) { in, out = out, in }
  go func() {
    fmon.New (nil, 1, x.addMeVertex, AllTrueSp,
              x.actHost, p0 + uint16(x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 1, x.addMeVertex, AllTrueSp,
                         x.host[i], p0 + uint16(x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.cycle.Clr()
  if x.me == x.root {
    x.cycle.Ins (x.me)
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
  exValue (x.cycle, x.leader)
  x.cycle.Mark (x.leader)
  x.cycle.Write()
}

func (x *distributedGraph) addMeVertex (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.cycle = x.decodedGraph(bs)
  s := x.cycle.Get().(vtx.Vertex).Val()
  j := x.channel(s)
  out := uint(0); if j == 0 { out++ }
  x.cycle.Ins (x.me)
  if x.me != x.root {
    x.tmpGraph.Ex2 (x.nr[j], x.me)
    x.cycle.Edge (x.tmpGraph.Get1())
    if x.nr[out] == x.root {
      x.cycle.Ex2 (x.me, x.nr[out])
      x.tmpGraph.Ex2 (x.me, x.nr[out])
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
