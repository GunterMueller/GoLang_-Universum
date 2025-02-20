package gra

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  "sort"
  . "µU/obj"
)

// Kruskal's algorithm, CLR 24.1-2, CLRS 23.1-2
func (x *graph) MST() {
  if x.nVertices < 2 || x.directed || x.eAnchor.nextE == x.eAnchor {
    return
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.predecessor = nil
    v.repr = v
    v.marked = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.marked = false
  }
  if x.nVertices == 1 {
    x.local = x.vAnchor.nextV
    x.local.marked = true
    return
  }
  es := make ([]*edge, x.nEdges)
  for i, e := uint(0), x.eAnchor.nextE; e != x.eAnchor; i, e = i + 1, e.nextE {
    es[i] = e
    e.marked = false
  }
  sort.Slice (es, func (i, j int) bool { return Val(es[i]) < Val(es[j]) })
  for len(es) > 0 {
    e := es[0]
    es = es[1:]
    v, v1 := e.nbPtr0.from, e.nbPtr1.from
    if x.demo [SpanTree] {
//      x.writeE (e.any, true)
//    XXX
      x.writeE (e.any)
//      x.writeV (v.any, true)
      x.writeV (v.any)
//      x.writeV (v1.any, true)
      x.writeV (v1.any)
      wait()
    }
    if v.repr != v1.repr {
      v.marked = true
      v1.marked = true
      e.marked = true
      for v.predecessor != nil {
        v = v.predecessor
      }
      v1 = v1.repr
      v.predecessor = v1
      v = v.repr
      for v1.predecessor != nil {
        v1.repr = v
        v1 = v1.predecessor
      }
      v1.repr = v
    } else {
      if x.demo [SpanTree] {
//        x.writeE (e.any, false)
//      XXX
        x.writeE (e.any)
//        x.writeV (v.any, false)
        x.writeV (v.any)
//        x.writeV (v1.any, false)
        x.writeV (v1.any )
        wait()
      }
    }
  }
}
