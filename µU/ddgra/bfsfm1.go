package dgra

// (c) Christian Maurer   v. 231109 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

func (x *distributedGraph) bfsfm1 (o Op) {
  go func() {
    fmon.New (nil, 2, x.b1, AllTrueSp,
              x.actHost, p0 + uint16(2 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 2, x.b1, AllTrueSp,
                         x.host[i], p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  x.parent = inf
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Mark (x.actVertex)
  x.tree.Write()
  if x.me == x.root {
    x.parent = x.root
    for {
      c := uint(0)
      for k := uint(0); k < x.n; k++ {
        if ! x.visited[k] {
          x.tree.Ex (x.actVertex)
          bs := append(Encode(x.distance), x.tree.Encode()...)
          bs = x.mon[k].F(bs, 0).(Stream)
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
    <-done // wait until root finished
  }
}

func (x *distributedGraph) b1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
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
      x.parent = s // == x.nr[j]
      if ! x.tree.Ex (x.actVertex) {
        x.tree.Ins (x.actVertex)
      }
      x.tree.Edge (x.edge(x.nb[j], x.actVertex)) // XXX colocal == local
      x.tree.Ex (x.actVertex)
      x.tree.Write()
      x.Op (x.actVertex)
      return append(Encode(x.distance), x.tree.Encode()...)
    }
    c := uint(0) // x.distance > 0
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.tree.Ex (x.actVertex)
        bs := append(Encode(x.distance - 1), x.tree.Encode()...)
        bs = x.mon[k].F(bs, 0).(Stream)
        if len(bs) == 0 {
          x.visited[k] = true
        } else {
          x.tree = x.decodedGraph(bs[C0:])
          x.child[k] = true
          c++
          x.tree.Write()
        }
      }
    }
    if c == 0 {
      return nil
    }
    x.tree.Ex (x.actVertex)
    bs = append(Encode(uint(0)), x.tree.Encode()...)
  case 1:
    x.tree.Ex (x.actVertex)
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
