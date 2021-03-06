package dgra

// (c) Christian Maurer   v. 171120 - license see µU.go

// see G. Andrews: Concurrent Programming (1991) p. 375

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) fmgraph1() {
  go func() {
    fmon.NewM (nil, 1, x.n, x.nr, x.addG1, AllTrueSp, x.actHost, x.sport, true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.monM[i] = fmon.NewM (nil, 1, i, x.nr, x.addG1, AllTrueSp, x.host[i], x.cport, false)
  }
  defer x.finMonM()
  x.awaitAllMonitorsM()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Sub (x.actVertex)
  x.tmpGraph.Write()
  lock <- 0
  known := false
  active := make([]bool, x.n)
  for i := uint(0); i < x.n; i++ {
    active[i] = true
  }
  r := uint(1)
  for ! known {
    x.log ("after round", r)
    for i := uint(0); i < x.n; i++ {
      <-lock
      bs := x.tmpGraph.Encode()
      lock <- 0
      bs = x.monM[i].Fm (append(Encode(false), bs...), 0, i).(Stream)
      g := x.decodedGraph(bs[1:])
      <-lock
      if Decode(false, bs[:1]).(bool) { active[i] = false }
      x.tmpGraph.Add (g)
      x.tmpGraph.Sub2 (x.actVertex, x.nb[i])
      x.tmpGraph.Write()
      lock <- 0
    }
    if ! known && x.tmpGraph.EqSub() {
      known = true
      x.log("topology in round", r)
    }
    r++
  }
  for i := uint(0); i < x.n; i++ {
    if active[i] {
      x.monM[i].Fm (append(Encode(false), x.tmpGraph.Encode()...), 0, i)
    }
  }
}

func (x *distributedGraph) addG1 (a Any, i uint) Any {
  x.awaitAllMonitorsM()
  <-lock
  x.tmpGraph.Add (x.decodedGraph(a.(Stream)[1:]))
  bs := append(Encode(true), x.tmpGraph.Encode()...)
  lock <- 0
  return bs
}
