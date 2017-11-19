package gra

// (c) Christian Maurer   v. 171112 - license see µU.go

// topological Sort, CLR 23.4, CLRS 22.4
// TODO spec
func (x *graph) Sort() {
  if x.nVertices < 2 || ! x.bool { return }
  x.dfs()
// sort list of vertices due to decrementing times, for which we supply a slice von *vertex:
  f := make ([]*vertex, 2 * x.nVertices)
  for i := uint32(0); i < 2 * x.nVertices; i++ {
    f[i] = nil
  }
// partial function f: [0 .. 2 * nVertices - 1] -> *vertex with
// f[i] := the vertex with time1 = i, if there is such, otherwise vAnchor
  v := x.vAnchor.nextV
  for i := uint32(0); i < x.nVertices; i++ {
    f[v.time1 - 1] = v
    v = v.nextV
  }
// sort list of vertices by
// von vorne nach hinten jeweils die Ecke mit Zeit i an den Anfang der Liste holen:
  for i := uint32(0); i < 2 * x.nVertices; i++ {
    v = f[i]
    if v != nil { // put n to the head of the list:
      v.nextV.prevV, v.prevV.nextV = v.prevV, v.nextV
      v.nextV, v.prevV = x.vAnchor.nextV, x.vAnchor
      v.nextV.prevV = v
      x.vAnchor.nextV = v
    }
  }
}

// strongly connected components, CLR 23.5, CLRS 22.5
func (x *graph) Isolate() {
  if x.nVertices < 1 || ! x.bool {
    return
  }
// depth first search with sorting of the list of vertices by decrementing times:
  x.Sort()
// essence of the algorithm: invert directions of all edges:
  x.Inv()
// and now once more depth first search,
// starting with the highest time of the first depth first search:
  x.dfs()
// the depth first search trees are now the strongly connected components with common repr
// finally again invert the directions of all edges
  x.Inv()
// all vertices in the actual subgraph:
// the depth first search trees are now the strongly connected components with common repr
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.bool = true
  }
// and furthermore all edges, that connect two vertices in the same strongly connected component:
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.bool = e.nbPtr0.from.repr == e.nbPtr1.from.repr
  }
}

func (x *graph) IsolateSub() {
  x.Isolate()
// only exactly those vertices in the actual subgraph, that
// are contained in the strong connection component of the local vertex:
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.bool = v.repr == x.local.repr
  }
// and furthermore exactly those edges, that connect these vertices:
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.bool = e.nbPtr0.from.bool && e.nbPtr1.from.bool
  }
}
