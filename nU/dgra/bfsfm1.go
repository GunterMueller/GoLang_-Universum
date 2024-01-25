package dgra

// (c) Christian Maurer   v. 231220 - license see nU.go

import (
  . "nU/obj"
  "nU/vtx"
  "nU/fmon"
)

func (x *distributedGraph) b1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.distance = Decode(uint(0), bs[:C0]).(uint)
  x.tree = x.decodedGraph(bs[C0:])
  x.tree.Write()
  pause()
  s := x.tree.Get().(vtx.Vertex).Val()
  j := x.channel(s)
  if i == search {
    if x.distance == 0 {
      if x.parent < inf {
        return nil
      }
      x.parent = s // == x.nr[j]
      if ! x.tree.Ex (x.actVertex) {
        x.tree.Ins (x.actVertex)
      }
      x.tree.Edge (x.edge(x.nb[j], x.actVertex))
      x.tree.Ex (x.actVertex)
      x.tree.Write()
      pause()
      return append(Encode(x.distance), x.tree.Encode()...)
    }
    c := uint(0) // x.distance > 0
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.tree.Ex (x.actVertex)
        bs = append(Encode(x.distance - 1), x.tree.Encode()...)
        bs = x.mon[k].F(bs, search).(Stream)
        if len(bs) == 0 {
          x.visited[k] = true
        } else {
          x.tree = x.decodedGraph(bs[C0:])
          x.tree.Write()
          pause()
          x.child[k] = true
          c++
        }
      }
    }
    if c == 0 {
      return nil
    }
    x.tree.Ex (x.actVertex)
    bs = append(Encode(uint(0)), x.tree.Encode()...)
  } else { // i == deliver
    x.tree.Ex (x.actVertex)
    bs = append(Encode(uint(0)), x.tree.Encode()...)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(bs, deliver)
      }
    }
    done <- 0
  }
  return bs
}

func (x *distributedGraph) Bfsfm1() {
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
  x.parent = inf
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Mark (x.actVertex)
  x.tree.Write()
  pause()
  if x.me == x.root {
    x.parent = x.root
    for {
      c := uint(0)
      for k := uint(0); k < x.n; k++ {
        if ! x.visited[k] {
          x.tree.Ex (x.actVertex)
          bs := append(Encode(x.distance), x.tree.Encode()...)
          bs = x.mon[k].F(bs, search).(Stream)
          if len(bs) == 0 {
            x.visited[k] = true
          } else {
            x.child[k] = true
            c++
            x.tree = x.decodedGraph(bs[C0:])
            x.tree.Write()
            pause()
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
        x.mon[k].F(bs, deliver)
      }
    }
  } else {
    <-done // wait until root finished
  }
  x.tree.Write()
}
