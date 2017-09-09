package dgra

// (c) Christian Maurer   v. 170510 - license see murus.go

import (
  . "murus/obj"
  "murus/fmon"
)

func (x *distributedGraph) fmbfs1 (o Op) {
  go func() {
    fmon.New (nil, 2, x.b1, AllTrueSp, x.actHost, p0 + uint16(2 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 2, x.b1, AllTrueSp, x.host[i], p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tmpGraph.Copy (x.Graph)
  x.Op = o
  x.parent = inf
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.SubLocal()
  x.tree.Write()
  if x.me == x.root {
    x.parent = x.me
    for {
      c := uint(0)
      for k := uint(0); k < x.n; k++ {
        if ! x.visited[k] {
          x.tree.Ex(x.actVertex)
          bs := append(Encode(x.distance), x.tree.Encode()...)
          bs = x.mon[k].F(bs, 0).([]byte)
          if len(bs) == 0 {
            x.visited[k] = true
          } else {
            x.child[k] = true
            c++
            x.tree = x.decodedGraph(bs[C0:])
            x.tree.Write()
          }
        }
      }
      if c == 0 {
        break
      }
      x.distance++
    }
    bs := append(Encode(uint(0)), x.tree.Encode()...)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, 1)
      }
    }
    x.Op (x.actVertex)
  } else {
    <-done // wait for root
  }
}

func (x *distributedGraph) b1 (a Any, i uint) Any {
  x.awaitAllMonitors()
  bs := a.([]byte)
  x.distance = Decode(uint(0), bs[:C0]).(uint)
  x.tree = x.decodedGraph(bs[C0:])
  x.tree.Write()
  s := nrLocal(x.tree)
  j := x.channel(s)
  switch i {
  case 0:
    if x.distance == 0 {
      if x.parent < inf {
        return nil
      }
      x.parent = s
      if ! x.tree.Ex(x.actVertex) {
        x.tree.Ins(x.actVertex)
      }
      x.tree.Edge (x.directedEdge(x.nb[j], x.actVertex)) // XXX colocal == local
      x.tree.Ex(x.actVertex)
      x.tree.Write()
      x.Op (x.actVertex)
      return append(Encode(x.distance), x.tree.Encode()...)
    }
    c := uint(0) // x.distance > 0
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.tree.Ex (x.actVertex)
        bs := append(Encode(x.distance - 1), x.tree.Encode()...)
        bs = x.mon[k].F(bs, 0).([]byte)
        if len(bs) == 0 {
          x.visited[k] = true
        } else {
          x.tree = x.decodedGraph(bs[C0:])
          x.child[k] = true
          x.tree.Write()
          c++
        }
      }
    }
    if c == 0 {
      return nil
    }
    x.tree.Ex(x.actVertex)
    bs = append(Encode(uint(0)), x.tree.Encode()...)
  case 1:
    x.tree.Ex(x.actVertex)
    bs := append(Encode(uint(0)), x.tree.Encode()...)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, 1)
      }
    }
    done <- 0
  }
  return bs
}
