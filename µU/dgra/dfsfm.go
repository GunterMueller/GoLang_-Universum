package dgra

// (c) Christian Maurer   v. 231215 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)
const (
  search = iota
  deliver
)

func (x *distributedGraph) Dfsfm() {
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
  x.tree.Clr()
  x.tree.Ins (x.actVertex)
  x.tree.Mark (x.actVertex)
  x.tree.Write()
  pause()
  if x.me == x.root {
    x.parent = x.me
    for k := uint(0); k < x.n; k++ {
      x.tree.Ex (x.actVertex) // actVertex local in x.tree
// x.log("search ", x.nr[k])
      bs := x.mon[k].F(x.tree, search).(Stream) // crash XXX interface is nil, not Stream
// x.log0(" ok")
      if len(bs) == 0 {
      } else {
        x.child[k] = true
        x.tree = x.decodedGraph(bs)
        x.tree.Write()
        pause()
      }
    }
    x.tree.Ex (x.actVertex) // actVertex local in x.tree
    x.tree.Write()
    pause()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
// x.log("deliver ", x.nr[k])
        x.mon[k].F(x.tree, deliver)
// x.log0(" ok")
      }
    }
  } else {
    <-done // wait for root's result
  }
  x.tree.Write()
}

func (x *distributedGraph) d1 (a any, i uint) any {
  x.awaitAllMonitors()
  bs := a.(Stream)
  x.tree = x.decodedGraph(bs)
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
// x.log("search ", x.nr[k])
          bs = x.mon[k].F(x.tree, search).(Stream)
// x.log0(" ok")
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
  } else { // deliver
    x.tree.Ex (x.actVertex)
    x.tree.Write()
    for k := uint(0); k < x.n; k++ {
      if x.child[k] {
// x.log("deliver ", x.nr[k])
        x.mon[k].F(x.tree, deliver)
// x.log0(" ok")
      }
    }
    x.Op (x.me)
    done <- 0
  }
  return x.tree.Encode()
}
