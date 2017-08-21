package nelect

// (c) murus.org  v. 170107 - license see murus.org
//
// >>> TODO Description of this algorithm and proof of its correctness

import (
  "murus/ker"
  . "murus/obj"
  "murus/bnat"
  "murus/nat"
  "murus/node"
  "murus/gra"
)

func (x *netElection) maurer (g0 gra.Graph, root uint) uint {
  b0 := bnat.New(nat.Wd(x.uint)); b0.SetVal(x.me)
  node0 := node.New (b0, 2, 1) // XXX
  gg := gra.New (true, node0, nil)
  if x.me == x.root {
    gg.Ins (x.meNode)
    x.chR.Send (gg.Encode())
  }
  bs := x.chL.Recv().([]byte)
  gg = Decode (gra.New (true, node0, nil), bs).(gra.Graph)
  x.logm (gg)
// examples conforming to lans.Network8ringdir with root = 0:
  if x.me == x.root { // me == 0
    // g = 0 -> 7 -> 4 -> 6 -> 3 -> 1 - > 5 -> 2 // 5 colocal, 2 local
    gg.Locate (true) // 2 local and colocal
    if ! gg.Ex (x.meNode) { ker.Oops() } // me == 0 local, 2 colocal
    gg.Edge() // g = ring 0 -> 7 -> 4 -> 6 -> 3 -> 1 -> 5 -> 2 -> 0 // 2 colocal, 0 local
    x.chR.Send (gg.Encode())
  } else {
    // e.g. for me == 6: g = 0 -> 7 -> 4 // 7 colocal, 4 local
    gg.Ins (x.meNode) // g = 0 -> 7 -> 4   6 // 4 colocal, 6 local
    gg.Edge() // g = 0 -> 7 -> 4 -> 6 // 4 colocal, 6 local
    x.chR.Send (gg.Encode())
  }
  bs = x.chL.Recv().([]byte)
  gg = Decode (gra.New (true, node0, nil), bs).(gra.Graph)
  x.logm (gg)
// g = ring 0 -> 7 -> 4 -> 6 -> 3 -> 1 -> 5 -> 2 -> 0 // 2 colocal, 0 local
  x.chR.Send (gg.Encode())
  max := uint(0)
  first := true
  gg.Trav (func (a Any) {
             n := a.(node.Node)
             if first {
               first = false
               node0.Copy (n)
             }
             m := n.Content().(bnat.Natural).Val()
             if m > max {
               node0.Copy (n)
               max = m
             }
           })
  if ! gg.Ex (node0) { ker.Oops() }
  gg.Locate (true)
  // XXX gg.Step(0) // What is local ?? XXX
  // gg.Locate() // XXX ?
  x.graph.Copy (gg)
  return max
}
