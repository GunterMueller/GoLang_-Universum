package gra

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/errh"
)

// For all vertices v, that are accessible from v0 by a path, v.repr == v0.
// vAnchor.acyclic == true, if x has no cycles.
func (x *graph) search (v0, v *vertex, p Pred) {
  x.vAnchor.time0++
  v.time0 = x.vAnchor.time0
  v.repr = v0
  if x.demo [Depth] {
//    v.Mark (true)
    x.Mark (v, true)
    x.writeV (v)
    wait()
  }
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && n.to != v.predecessor && p ((n.to).any) {
      if n.to.time0 == 0 {
        if x.demo [Depth] {
//          n.edgePtr.Mark (true)
          x.Mark1 (n.edgePtr, true)
          x.writeE (n.edgePtr)
        }
        n.to.predecessor = v
        x.search (v0, n.to, p)
        if x.demo [Depth] {
//          n.edgePtr.Mark (false)
          x.Mark1 (n.edgePtr, false)
          x.writeE (n.edgePtr)
          wait()
        }
      } else if n.to.time1 == 0 {
        x.vAnchor.acyclic = false // a cycle was found
        if x.demo [Cycle] { // also x.demo [Depth], see Set
//          n.edgePtr.Mark (true)
          x.Mark1 (n.edgePtr, true)
          x.writeE (n.edgePtr)
//          errh.Error0 ("Kreis gefunden")
//          n.edgePtr.Mark (false)
          x.Mark1 (n.edgePtr, false)
          x.writeE (n.edgePtr)
          wait()
        }
      }
    }
  }
  x.vAnchor.time0++
  v.time1 = x.vAnchor.time0
  if x.demo [Depth] {
//    n.edgePtr.Mark (false)
    x.writeV (v)
  }
}

func (x *graph) preDfs() {
  x.vAnchor.time0 = 0
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.time0, v.time1 = 0, 0
    v.predecessor, v.repr = nil, v
  }
}

// CLR 23.3, CLRS 22.3
func (x *graph) dfs (p Pred) {
  x.preDfs()
  if x.demo [Depth] {
    errh.Hint ("weiter mit Eingabetaste")
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.time0 == 0 { x.search (v, v, p) }
  }
  if x.demo [Depth] {
    errh.DelHint()
  }
}

func (x *graph) Acyclic() bool {
  if x.Empty() { return true }
  x.vAnchor.acyclic = true
  x.dfs (AllTrue)
  return x.vAnchor.acyclic
}
