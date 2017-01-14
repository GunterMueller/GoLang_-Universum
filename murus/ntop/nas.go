package ntop

// (c) murus.org  v. 170101 - license see murus.go
//
// >>> experimental - proof of concept

import (
  "murus/scr"
  "murus/nchan"
  "murus/ntop/aaa"
)

func (x *netTopology) nas() {
  cm := scr.NColumns() / 2
  nas := aaa.New (x.uint)
  nas.Put (x.me, 101 * uint32(x.me))
  if x.bool {
    nas.Write (0, 0)
  }
  ch := make ([]nchan.NetChannel, x.n)
  for i := uint(0); i < x.n; i++ {
    ch[i] = nchan.New (nas, x.me, x.nn[i].Val(), x.h[i], x.port[i])
  }
  for r:= uint(0); r < x.diameter; r++ {
    x.enter (r + 1)
    for i := uint(0); i < x.n; i++ {
      ch[i].Send (nas)
    }
    for i := uint(0); i < x.n; i++ {
      nas.Add (ch[i].Recv().(aaa.AAA))
      if x.bool { nas.Write (0, cm) }
    }
  }
  for i := uint(0); i < x.n; i++ { ch[i].Fin() }
}
