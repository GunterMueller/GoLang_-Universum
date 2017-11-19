package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
  "µU/fmon"
)

func (x *distributedGraph) fmmatrix() {
//  go func() { fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.actHost, p0 + uint16(x.me), true) }()
  go func() { fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.actHost, uint16(x.me), true) }()
  for i := uint(0); i < x.n; i++ {
//    x.mon[i] = fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.host[i], p0 + uint16(x.nr[i]), false)
    x.mon[i] = fmon.New (x.top, 1, x.addMatrix, AllTrueSp, x.host[i], uint16(x.nr[i]), false)
  }
  defer x.finMon()
  if x.demo { x.top.Write(0, 0) }
  for r := uint(1); r <= x.diameter; r++ {
    x.log("after round", r)
    for i := uint(0); i < x.n; i++ {
      x.addMatrix (x.mon[i].F(x.top, 0), i)
    }
    if x.demo { x.top.Write(0, scr.NColumns() / 2) }
  }
}
