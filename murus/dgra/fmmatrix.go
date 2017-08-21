package dgra

// (c) murus.org  v. 170507 - license see murus.go

import (
  . "murus/obj"
  "murus/scr"
  "murus/fmon"
  "murus/adj"
)

func (x *distributedGraph) fmmatrix() {
  go func() { fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.actHost, p0 + uint16(x.me), true) }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.host[i], p0 + uint16(x.nr[i]), false)
  }
  defer x.finMon()
  if x.demo { x.top.Write(0, 0) }
  for r := uint(0); r < 1 * x.diameter; r++ {
    x.enter (r + 1)
    for i := uint(0); i < x.n; i++ {
      x.addMatrix (x.mon[i].F(x.top, 0), i)
    }
    if x.demo { x.top.Write(0, scr.NColumns() / 2) }
  }
}

func (x *distributedGraph) addMatrix (a Any, i uint) Any {
  x.top.Add (a.(adj.AdjacencyMatrix))
  return x.top
}
