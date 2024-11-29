package dgra

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
  "µU/fmon"
)

func (x *distributedGraph) b1 (a any, i uint) any {
  x.awaitAllMonitors()
  s := a.(Stream)
  x.distance = Decode(uint(0), s[:C0]).(uint)
  x.tree = x.decodedGraph(s[C0:])
  x.tree.Write()
  n := nrLocal(x.tree)
  j := x.channel(n)
  switch i {
  case 0:
    if x.distance == 0 {
      if x.parent < inf {
        return nil
      }
      x.parent = n // == x.nr[j]
      if ! x.tree.Ex (x.actVertex) {
        x.tree.Ins (x.actVertex)
      }
      x.tree.Edge (x.edge(x.nb[j], x.actVertex)) // XXX colocal == local
      x.tree.Ex (x.actVertex)
      x.tree.Write()
      x.Op (x.me)
      return append(Encode(x.distance), x.tree.Encode()...)
    }
    c := uint(0) // x.distance > 0
    for k := uint(0); k < x.n; k++ {
      if k != j && ! x.visited[k] {
        x.tree.Ex (x.actVertex)
        s := append(Encode(x.distance - 1), x.tree.Encode()...)
        s = x.mon[k].F(s, 0).(Stream)
        if len(s) == 0 {
          x.visited[k] = true
        } else {
          x.tree = x.decodedGraph(s[C0:])
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
    s = append(Encode(uint(0)), x.tree.Encode()...)
  case 1:
    x.tree.Ex (x.actVertex)
    s := append(Encode(uint(0)), x.tree.Encode()...)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(s, 1)
      }
    }
    done <- 0
  }
  return s
}

func (x *distributedGraph) Bfsfm1() {
  scr.Cls()
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
  x.tree.Mark (x.actVertex, true)
  x.tree.Write()
  if x.me == x.root {
    x.parent = x.root
    for {
      c := uint(0)
      for k := uint(0); k < x.n; k++ {
        if ! x.visited[k] {
          x.tree.Ex (x.actVertex)
          s := append(Encode(x.distance), x.tree.Encode()...)
          s = x.mon[k].F(s, 0).(Stream)
          if len(s) == 0 {
            x.visited[k] = true
          } else {
            x.child[k] = true
            c++
            x.tree = x.decodedGraph(s[C0:])
            x.tree.Write()
          }
        }
      }
      if c == 0 {
        break
      }
      x.distance++
    }
    s := append(Encode(uint(0)), x.tree.Encode()...)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(s, 1)
      }
    }
  } else {
    <-done // wait until root finished
  }
}
