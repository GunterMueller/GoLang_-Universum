package dgra

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/scr"
  "µU/fmon"
)
const (
  search = iota
  deliver
)

func (x *distributedGraph) d (a any, i uint) any {
  x.awaitAllMonitors()
  s := a.(Stream)
  x.tree = x.decodedGraph(s)
  if i == search {
    j := x.channel(nrLocal(x.tree))
    if x.tree.Ex (x.actVertex) {
      return nil
    }
    x.parent = x.nr[j]
    x.tree.Ins (x.actVertex)
    x.tree.Edge (x.edge(x.nb[j], x.actVertex))
    x.tree.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if k != j {
        if ! x.tree.Ex (x.nb[k]) {
          x.tree.Ex (x.actVertex)
          s = x.mon[k].F(x.tree, search).(Stream)
          if len(s) == 0 {
            return nil
          } else {
            x.child[k] = true
            x.tree = x.decodedGraph (s)
            x.tree.Write()
            pause()
          }
        }
      }
    }
  } else { // i == deliver
    x.tree.Ex (x.actVertex)
    x.tree.Write()
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

func (x *distributedGraph) Dfsfm() {
  scr.Cls()
  go func() {
    fmon.New (nil, 2, x.d, AllTrueSp,
              x.actHost, p0 + uint16(2 * x.me), true)
  }()
  for i := uint(0); i < x.n; i++ {
    x.mon[i] = fmon.New (nil, 2, x.d, AllTrueSp,
                         x.host[i], p0 + uint16(2 * x.nr[i]), false)
  }
  defer x.finMon()
  x.awaitAllMonitors()
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Mark (x.actVertex, true)
  x.tree.Write()
  pause()
  if x.me == x.root {
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.tree.Ex (x.actVertex) // actVertex local in x.tree
      s := x.mon[k].F(x.tree, search).(Stream)
      if len(s) == 0 {
      } else {
        x.child[k] = true
        x.tree = x.decodedGraph(s)
        x.tree.Write()
        pause()
      }
    }
    x.tree.Ex (x.actVertex) // actVertex local in x.tree
    x.tree.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
        x.mon[k].F(x.tree, deliver)
      }
    }
  } else {
    <-done // wait for root's result
  }
  x.tree.Write()
}
