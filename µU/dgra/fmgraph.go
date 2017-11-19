package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) fmgraph() {
  go func() {
    fmon.NewM (nil, 1, x.n, x.nr, x.addG, AllTrueSp, x.actHost, x.sport, true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.monM[i] = fmon.NewM (nil, 1, i, x.nr, x.addG, AllTrueSp, x.host[i], x.cport, false)
  }
  defer x.finMonM()
  x.awaitAllMonitorsM()
  x.tmpGraph.Copy (x.Graph)
  x.tmpGraph.Sub (x.actVertex)
  x.tmpGraph.Write()
  lock <- 0
  for r := uint(1); r <= x.diameter; r++ {
    x.log ("after round", r)
    for i := uint(0); i < x.n; i++ {
      <-lock
      bs := x.tmpGraph.Encode()
      lock <- 0
      bs = x.monM[i].Fm (bs, 0, i).([]byte)
      g := x.decodedGraph(bs)
      <-lock
      x.tmpGraph.Add (g)
      x.tmpGraph.Sub2 (x.actVertex, x.nb[i])
      x.tmpGraph.Write()
      lock <- 0
    }
  }
}

func (x *distributedGraph) addG (a Any, i uint) Any {
  x.awaitAllMonitorsM()
  <-lock
  x.tmpGraph.Add (x.decodedGraph(a.([]byte)))
  bs := x.tmpGraph.Encode()
  lock <- 0
  return bs
}
