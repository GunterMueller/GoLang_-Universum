package gra

// (c) Christian Maurer   v. 220609 - license see µU.go

import (
//  "µU/bn"
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
    x.writeV (v.any, true)
    wait()
  }
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && n.to != v.predecessor && p ((n.to).any) {
      if n.to.time0 == 0 {
        if x.demo [Depth] {
          x.writeE (n.edgePtr.any, true)
        }
        n.to.predecessor = v
        x.search (v0, n.to, p)
        if x.demo [Depth] {
          x.writeE (n.edgePtr.any, false)
          wait()
        }
      } else if n.to.time1 == 0 {
        x.vAnchor.acyclic = false // a cycle was found
        if x.demo [Cycle] { // also x.demo [Depth], see Set
          x.writeE (n.edgePtr.any, true)
//          errh.Error0("Kreis gefunden")
          x.writeE (n.edgePtr.any, false)
          wait()
        }
      }
    }
  }
  x.vAnchor.time0++
  v.time1 = x.vAnchor.time0
  if x.demo [Depth] {
    x.writeV (v.any, false)
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
func (x *graph) Dfs1 (p Pred) {
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

func (x *graph) search1 (v *vertex, o Op, p Pred, s Stmt) {
  (*v).bool = true
  o ((*v).any)
  if p ((*v).any) { s() }
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if ! (n.to).bool {
      x.search1 (n.to, o, p, s)
    }
  }
}

func (x *graph) Dfs (o Op, p Pred, s Stmt) {
//  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
//    if v.time0 == 0 { x.search1 (v, o, p, s) }
//    return
//  }
  x.search1 (x.local, o, p, s)
}

func (x *graph) Acyclic() bool {
  if x.Empty() { return true }
  x.vAnchor.acyclic = true
  x.Dfs1 (AllTrue)
  return x.vAnchor.acyclic
}
