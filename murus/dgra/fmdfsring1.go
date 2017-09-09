package dgra

// (c) Christian Maurer   v. 170507 - license see murus.go
//
// >>> Construction of a directed ring using the idea of Awerbuch's DFS-algorithm

import (
  . "murus/obj"
  "murus/vtx"
  "murus/fmon"
)

func (x *distributedGraph) fmdfsring1() {
  go func() { fmon.New (nil, 3, x.dr1, AllTrueSp, x.actHost, p0 + uint16(3 * x.me), true) }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 3, x.dr1, AllTrueSp, x.host[i], p0 + uint16(3 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy(x.Graph)
  x.ring.Clr()
  if x.me == x.root {
    x.ring.Ins(x.actVertex)
    x.ring.SubLocal()
    for k := uint(0); k < x.n; k++ {
      bs := Encode(x.me)
      x.mon[k].F(bs, visit)
    }
    for k := uint(0); k < x.n; k++ {
      if ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        x.ring.Ex(x.actVertex)
        bs := append(Encode(x.me), x.ring.Encode()...)
        bs = x.mon[k].F(bs, discover).([]byte)
        x.ring = x.decodedGraph(bs[C0:])
        x.ring.Write()
      }
    }
    v := x.ring.Get().(vtx.Vertex)
    x.ring.Ex2 (v, x.actVertex)
    x.ring.Edge (x.directedEdge (v, x.actVertex))
    x.ring.Write()
    bs := x.ring.Encode()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, 2)
      }
    }
  } else {
    <-done
  }
}

func (x *distributedGraph) dr1 (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.([]byte)
  s := Decode(uint(0), bs[:C0]).(uint)
  j := x.channel(s)
  switch i {
  case visit:
    x.visited[j] = true
  case discover:
    for k := uint(0); k < x.n; k++ {
      if k != j {
        x.mon[k].F(Encode(x.me), visit)
      }
    }
    x.ring = x.decodedGraph(bs[C0:])
    x.ring.Ins(x.actVertex)
    a0, a1 := x.ring.Get2()
    e := x.directedEdge (a0.(vtx.Vertex), a1.(vtx.Vertex))
    x.ring.Edge (e)
    x.ring.Write()
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.visited[k] = true
        x.child[k] = true
        bs := append(Encode(x.me), x.ring.Encode()...)
        bs = x.mon[k].F(bs, discover).([]byte)
        x.ring = x.decodedGraph(bs[C0:])
        x.ring.Write()
      }
    }
    return append(Encode(x.me), x.ring.Encode()...)
  case 2:
    x.ring = x.decodedGraph(bs)
    x.ring.Write()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, 2)
      }
    }
    done <- 0
  }
  return nil
}
