package dgra

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/vtx"; "nU/fmon")

const (search = iota; deliver)

func (x *distributedGraph) dfsfm1 (o Op) {
  go func() {
    fmon.New (nil, 2, x.d1, AllTrueSp,
              x.actHost, p0 + uint16(2 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 2, x.d1, AllTrueSp,
                         x.host[i], p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.Op = o
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Mark (x.actVertex)
  x.tree.Write()
  pause()
  if x.me == x.root {
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.tree.Ex (x.actVertex) // x.actVertex ist lokale Ecke in x.tree
      bs := x.mon[k].F(x.tree, search).(Stream)
      if len(bs) == 0 {
        x.visited[k] = true
      } else {
        x.child[k] = true
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
        pause()
      }
    }
    x.tree.Ex (x.actVertex) // x.actVertex ist lokale Ecke in x.tree
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, deliver)
      }
    }
    x.Op (x.me)
  } else {
    <-done // root abwarten
  }
  x.tree.Write()
}

func (x *distributedGraph) d1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.tree = x.decodedGraph(bs)
  x.tree.Write()
  pause()
  if i == search {
    if x.tree.Ex (x.actVertex) {
      return nil
    }
    s := x.tree.Get().(vtx.Vertex).Val()
    j := x.channel(s)
    x.parent = x.nr[j]
    x.tree.Ins (x.actVertex)
    x.tree.Edge (x.edge (x.nb[j], x.actVertex))
    x.tree.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if k != j {
        if ! x.tree.Ex (x.nb[k]) {
          x.tree.Ex (x.actVertex)
          bs = x.mon[k].F(x.tree, search).(Stream)
          if len(bs) == 0 {
            return nil
          } else {
            x.child[k] = true
            x.tree = x.decodedGraph (bs)
            x.tree.Write()
            pause()
          }
        }
      }
    }
  } else { // i == deliver
    x.tree.Ex (x.actVertex)
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, deliver)
      }
    }
    x.Op (x.me)
    done <- 0
  }
  return x.tree.Encode()
}
