package gra

// (c) Christian Maurer   v. 171122 - license see µU.go

import (
  . "µU/obj"
  "µU/errh"
)

// For all vertices n, that are accessible from n0 by a path, n.repr == n0.
// vAnchor.acyclic == true, if x has no cycles.
func (x *graph) search (v0, v *vertex, p Pred) {
  x.vAnchor.time0++
  v.time0 = x.vAnchor.time0
  v.repr = v0
  if x.demo [Depth] {
    x.writeV (v.Any, true)
    wait()
  }
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && n.to != v.predecessor && p (n.to.Any) {
      if n.to.time0 == 0 {
        if x.demo [Depth] {
          x.writeE (n.edgePtr.Any, true)
        }
        n.to.predecessor = v
        x.search (v0, n.to, p)
        if x.demo [Depth] {
          x.writeE (n.edgePtr.Any, false)
          wait()
        }
      } else if n.to.time1 == 0 {
        x.vAnchor.acyclic = false // found cycle
        if x.demo [Cycle] { // also x.demo [Depth], see Set
          x.writeE (n.edgePtr.Any, true)
//          errh.Error0("Kreis gefunden")
          x.writeE (n.edgePtr.Any, false)
          wait()
        }
      }
    }
  }
  x.vAnchor.time0++
  v.time1 = x.vAnchor.time0
  if x.demo [Depth] {
    x.writeV (v.Any, false)
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
func (x *graph) dfs() {
  x.preDfs()
  if x.demo [Depth] {
    errh.Hint ("weiter mit Eingabetaste")
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.time0 == 0 {
      x.search (v, v, AllTrue)
    }
  }
  if x.demo [Depth] {
    errh.DelHint()
  }
}

func (x *graph) Acyclic() bool {
  if x.Empty() { return true }
  x.vAnchor.acyclic = true
  x.dfs()
  return x.vAnchor.acyclic
}
