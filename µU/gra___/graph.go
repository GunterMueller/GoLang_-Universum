package gra

// (c) Christian Maurer   v. 231110 - license see µU.go
//
// >>>  References:
// >>>  CLR  = Cormen, Leiserson, Rivest        (1990)
// >>>  CLRS = Cormen, Leiserson, Rivest, Stein (2001)

import (
  "µU/ker"
  . "µU/obj"
  "µU/kbd"
  "µU/pseq"
  "µU/vtx"
  "µU/edg"
//  "µU/time"
)

/*    vertex           neighbour                        neighbour            vertex
    _________                                                              _________  
   /         \             /--------------------------------------------> /         \
  |     v     |          /                                               |     v     |
  |___________|<------------------------------------------------\        |___________|
  |           |       /                                           \      |           |
  |   nbPtr---|-----------\                                  /-----------|---nbPtr   |
  |___________|    /       |                                |        \   |___________|
  |           |   |        |                                |         |  |           |
  |   bool    |   |        v              edge              v         |  |   bool    |
  |___________|   |    _________        ________        _________     |  |___________|
  |     |     |   |   /         \      /        \      /         \    |  |     |     |
  |dist | time|   |  | edgePtr---|--->|    e     |<---|--edgePtr  |   |  |dist | time|
  |_____|_____|   |  |___________|    |__________|    |___________|   |  |_____|_____|
  |           |   |  |           |    |          |    |           |   |  |           |
  |predecessor|<-----|---from    |<---|--nbPtr0  |    |   from----|----->|predecessor|
  |___________|   |  |___________|    |__________|    |___________|   |  |___________|
  |           |   |  |           |    |          |    |           |   |  |           |
  |   repr    |    \_|____to     |    |  nbPtr1 -|--->|    to_____|__/   |   repr    |
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
  |   nextV---|->    | outgoing  |    |   bool   |    | outgoing  |      |   nextV---|->
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
<-|---prevV   |      |  nextNb---|->  |  nextE---|->  |  nextNb---|->  <-|---prevV   |
   \_________/       |___________|    |__________|    |___________|       \_________/
                     |           |    |          |    |           |
                   <-|---prevNb  |  <-|---prevE  |  <-|---prevNb  |
                      \_________/      \________/      \_________/

  The vertices of a graph are represented by structs,
  whose field "any" represents the "real" vertex.
  All vertices are connected in a doubly linked list with anchor cell,
  that can be traversed to execute some operation on all vertices of the graph.

  The edges are also represented by structs,
  whose field "any" is a variable of a type that implements Valuator.
  Also all edges are connected in a doubly linked list with anchor cell.

  For a vertex v one finds all outgoing and incoming edges
  with the help of a further doubly linked ringlist of neighbour(hoodrelation)s
    nb1 = v.nbPtr, nb2 = v.nbPtr.nextNb, nb3 = v.nbPtr.nextNb.nextNb etc.
  by the links outgoing from the nbi (i = 1, 2, 3, ...)
    nb1.edgePtr, nb2.edgePtr, nb3.edgePtr etc.
  In directed graphs the edges outgoing from a vertex are exactly those ones
  in the neighbourlist, for which outgoing == true.

  For an edge e one finds its two vertices by the links
    e.nbPtr0.from = e.nbPtr1.to und e.nbPtr0.to = e.nbPtr1.from.

  Semantics of some variables, that are "hidden" in fields of vAnchor:
    vAnchor.time0: in that the "time" is incremented for each search step
    vAnchor.acyclic: (after call of search1) == true <=> graph has no cycles. */

const (
  suffix = "gra"
  inf = uint32(1<<32 - 1)
)
type
  vertex struct {
              v vtx.Vertex
          nbPtr *neighbour
                bool "marked"
        acyclic bool   // for the development of design patterns by clients
           dist,       // for breadth first search/Dijkstra and use in En/Decode
   time0, time1 uint32 // for applications of depth first search
    predecessor,       // for back pointers in depth first search and in ways
           repr,       // for the computation of connected components
          nextV,
          prevV *vertex
                }
type
  vCell struct {
          vPtr *vertex
          next *vCell
               }
type
  edge struct {
            e edg.Edge "attribute of the edge"
       nbPtr0,
       nbPtr1 *neighbour
              bool "marked"
        nextE,
        prevE *edge
              }
type
  neighbour struct {
           edgePtr *edge
              from,
					      to *vertex
          outgoing bool
            nextNb,
		        prevNb *neighbour
                   }
type
  graph struct {
          name,
      filename string
          file pseq.PersistentSequence
          bool "directed"
     nVertices,
        nEdges uint32
       vAnchor,
       colocal,
         local *vertex
       eAnchor *edge
          path []*vertex
     eulerPath []*neighbour
          demo Demoset
        writeV,
        writeE CondOp
          esel [][]uint
               }
type
  nSeq []*neighbour

func newVertex (v vtx.Vertex) *vertex {
  w := new(vertex)
  w.v = v.Clone().(vtx.Vertex)
  w.time1 = inf // for applications of depth first search
  w.dist = inf
  w.repr = w
  w.nextV, w.prevV = w, w
  return w
}

func newEdge (e edg.Edge) *edge {
  f := new(edge)
  f.e = e.Clone().(edg.Edge)
  f.nextE, f.prevE = f, f
  return f
}

func new_(d bool, v vtx.Vertex, e edg.Edge) Graph {
  CheckAtomicOrObject(v)
  x := new(graph)
  x.bool = d
  x.vAnchor = newVertex (v)
//  if e == nil { e = uint(1) } // XXX
  CheckUintOrValuator (e)
  x.eAnchor = newEdge (e)
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.writeV, x.writeE = CondIgnore, CondIgnore
  x.esel = make([][]uint, 0)
  return x
}

func (x *graph) imp (Y any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *graph) Directed() bool {
  return x.bool
}

func (x *graph) Num() uint {
  return uint(x.nVertices)
}

func (x *graph) NumMarked() uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.bool {
      n++
    }
  }
  return n
}

func (x *graph) NumMarked1() uint {
  n := uint(0)
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.bool {
      n++
    }
  }
  return n
}

func (x *graph) NumPred (p Pred) uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.v) {
      n++
    }
  }
  return n
}

func (x *graph) Num1() uint {
  return uint(x.nEdges)
}

func newNeighbour (e *edge, v, v1 *vertex, f bool) *neighbour {
  nb := new(neighbour)
  nb.edgePtr = e
  nb.from, nb.to = v, v1
  nb.outgoing = f
  nb.nextNb, nb.prevNb = nb, nb
  return nb
}

func (x *graph) insertedVertex (v vtx.Vertex, marked bool) *vertex {
  w := newVertex (v)
  w.bool = marked
  w.nbPtr = newNeighbour (nil, w, nil, false)
  w.nextV, w.prevV = x.vAnchor, x.vAnchor.prevV
  w.prevV.nextV = w
  x.vAnchor.prevV = w
  return w
}

func (x *graph) insMarked (v vtx.Vertex, marked bool) {
  if x.vAnchor.v == nil { ker.Oops() }
  CheckTypeEq (v, x.vAnchor.v)
  if x.Ex (v) { // local is set
    return
  }
  w := x.insertedVertex (v, marked)
  x.nVertices++
  if x.nVertices == 1 {
    x.colocal = w
  } else {
    x.colocal = x.local
  }
  x.local = w
}

func (x *graph) Ins (v vtx.Vertex) {
  x.insMarked (v, false)
}

// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (n *neighbour, v *vertex) {
  n.nextNb, n.prevNb = v.nbPtr, v.nbPtr.prevNb
  n.prevNb.nextNb = n
  v.nbPtr.prevNb = n
}

// TODO Spec
func (x *graph) insertedEdge (a edg.Edge, marked bool) *edge {
//  if ! TypeEq (a, x.eAnchor.any) { ker.Panic ("gra.insertedEdge: ! TypeEq") }
  f := newEdge (a)
  f.bool = marked
  f.nbPtr0 = newNeighbour (f, x.colocal, x.local, true)
  insertNeighbour (f.nbPtr0, x.colocal)
  f.nbPtr1 = newNeighbour (f, x.local, x.colocal, ! x.bool)
  insertNeighbour (f.nbPtr1, x.local)
  if f.nbPtr1 == nil { ker.Panic ("gra.insertedEdge: e.nbPtr1 == nil") }
  f.nextE, f.prevE = x.eAnchor, x.eAnchor.prevE
  f.prevE.nextE = f
  x.eAnchor.prevE = f
  return f
}

// Returns nil, iff there is no edge from n to n1;
// returns otherwise the corresponding pointer.
func (x *graph) connection (v, v1 *vertex) *edge {
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && n.to == v1 {
      return n.edgePtr
    }
  }
  return nil
}

func (x *graph) edgeMarked (e edg.Edge, marked bool) {
  if x.Empty() { return }
  if x.colocal == x.local { ker.Panic ("gra.Edge: colocal == local") }
//  if e == nil { e = uint(1) } // XXX
  CheckTypeEq (e, x.eAnchor.e)
// simple case: local and colocal are not yet adjacent:
  f := x.connection (x.colocal, x.local)
  f1 := x.connection (x.local, x.colocal)
  if f == nil && f1 == nil {
    f = x.insertedEdge (e, marked)
    x.nEdges++
    return
  }
// otherwise: an existing edge must not be cleared:
// if there is an edge from colocal to local, it is looked for:
  n := x.colocal.nbPtr
  n.outgoing = true
  n = n.nextNb
  for n.to != x.local {
    n = n.nextNb
//    if n == x.colocal.nbPtr { ker.Oops() }
  }
  if n.to != x.local { ker.Oops() }
// and its v is replaced:
  if ! x.bool {
    n.to.nbPtr.outgoing = true
  }
  n.edgePtr.e = e.Clone().(edg.Edge)
  n.edgePtr.bool = marked
  x.nEdges++
}

func (x *graph) Edge (e edg.Edge) {
  if x.Empty() { return }
  if x.colocal == x.local { ker.Panic ("gra.Edge: colocal == local") }
  x.edgeMarked (e, false)
}

func (x *graph) Edge2 (v, v1 vtx.Vertex, e edg.Edge) {
  if x.Empty() || Eq (v, v1) ||
    ! TypeEq (v, v1) || ! TypeEq (v, x.vAnchor.v) {
    return
  }
  if n, ok := x.found (v); ! ok {
    return
  } else {
    x.colocal = n
  }
  if n1, ok := x.found (v1); ! ok {
    return
  } else {
    x.local = n1
  }
  if x.colocal == x.local ||
    x.connection (x.colocal, x.local) != nil {
    return
  }
  x.Edge (e)
}

func (ns nSeq) Less (i, j int) bool {
  return ns[i].to.v.(Valuator).Val() < ns[j].to.v.(Valuator).Val()
}

/*/
func (x *graph) SortNeighbours() {
  switch x.vAnchor.any.(type) {
  case Valuator:
    // see below
  default:
    return
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    c := x.numEdges (v)
    if c > 1 {
      ns := make(nSeq, c)
      for n, i := v.nbPtr.nextNb, 0; n != v.nbPtr; n, i = n.nextNb, i + 1 {
        ns[i] = n
      }
      sort.Slice (ns, ns.Less)
      v.nbPtr.nextNb = ns[0]
      for i := uint(0); i < c - 1; i++ {
        ns[i].nextNb = ns[i + 1]
      }
      ns[c-1].nextNb = v.nbPtr
      v.nbPtr.prevNb = ns[c - 1]
      ns[0].prevNb = v.nbPtr
      for i := uint(1); i < c; i++ {
        ns[i].prevNb = ns[i - 1]
      }
    }
  }
}
/*/

func (x *graph) Edged() bool {
  return x.connection (x.colocal, x.local) != nil
}

func (x *graph) CoEdged() bool {
  return x.connection (x.local, x.colocal) != nil
}

// Returns (nil, false), iff a there is no vertex in x with v;
// returns otherwise (n, true), where n is the pointer to that vertex
// (which is unique because of effect of Ins).
func (x *graph) found (a any) (*vertex, bool) {           // a has type *vtx.vertex
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV { // v has type *gra.vertex
    if Eq (v.v, a) {                                    // v.v has type *vtx.Vertex
      return v, true
    }
  }
  return nil, false
}

func (x *graph) Ex (v vtx.Vertex) bool {
  if w, ok := x.found (v); ok {
    x.local = w
    return true
  }
  return false
}

func (x *graph) Ex2 (v, v1 vtx.Vertex) bool {
  if Eq (v, v1) {
    return false
  }
  if w, ok := x.found (v); ok {
    if w1, ok1 := x.found (v1); ok1 {
      x.colocal, x.local = w, w1
      return true
    }
  }
  return false
}

// Returns true, iff there is no vertex in x, for which p returns true,
// Returns otherwise a pointer to such a vertex.
func (x *graph) foundPred (p Pred) (*vertex, bool) {
//  print ("foundPred ")
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.v) {
//      println (v.any.(Stringer).String()); time.Msleep (100)
      return v, true
    }
  }
//  println()
  return nil, false
}

func (x *graph) ExPred (p Pred) bool {
//  x.vAnchor.predecessor = nil
  if v, ok := x.foundPred (p); ok {
    x.local = v
    return true
  }
  return false
}

func (x *graph) Ex1 (e edg.Edge) bool {
  return x.ExPred1 (func (a any) bool { return Eq (a, e) })
}

func (x *graph) ExPred1 (p Pred) bool {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if p (e.e) {
      if e.nbPtr0.outgoing && e.nbPtr1.outgoing {
        x.colocal = e.nbPtr0.from
        x.local = e.nbPtr1.from
      } else {
        x.colocal = e.nbPtr1.from
        x.local = e.nbPtr0.from
      }
      return true
    }
  }
  return false
}

func (x *graph) ExPred2 (p, p1 Pred) bool {
  if v, ok := x.foundPred (p); ok {
    if v1, ok1 := x.foundPred (p1); ok1 {
      if v == v1 {
        tmp := x.vAnchor.nextV
        x.vAnchor.nextV = v
        v1, ok1 = x.foundPred (p1) // n1 != n
        x.vAnchor.nextV = tmp
        if ! ok1 {
          return false
        }
      }
      x.colocal, x.local = v, v1
      return true
    }
  }
  return false
}

func (x *graph) Get() vtx.Vertex {
  return x.local.v.Clone().(vtx.Vertex)
}

func (x *graph) Get2() (vtx.Vertex, vtx.Vertex) {
  return x.colocal.v.Clone().(vtx.Vertex), x.local.v.Clone().(vtx.Vertex)
}

func (x *graph) Get1() edg.Edge {
  if x.local == x.vAnchor || x.local == x.colocal {
    return x.eAnchor.e.Clone().(edg.Edge)
  }
  e := x.connection (x.colocal, x.local)
  if e == nil {
    e = x.connection (x.local, x.colocal)
  }
  if e == nil || e.e == nil {
    return x.eAnchor.e.Clone().(edg.Edge)
  }
  return e.e.Clone().(edg.Edge)
}

func (x *graph) Put (v vtx.Vertex) {
  if x.vAnchor == x.vAnchor.nextV { return }
  x.local.v= v.Clone().(vtx.Vertex)
}

func (x *graph) Put1 (e edg.Edge) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.nbPtr.edgePtr.e = e.Clone().(edg.Edge)
}

func (x *graph) Put2 (v, v1 vtx.Vertex) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.v = v.Clone().(vtx.Vertex)
  x.local.v = v1.Clone().(vtx.Vertex)
}

func (x *graph) ClrMarked() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.bool = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.bool = false
  }
}

func (x *graph) Mark (v vtx.Vertex) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { return }
  x.local.bool = true
}

func (x *graph) Mark1 (v vtx.Vertex) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { return }
  x.local.bool = true
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    n.edgePtr.bool = true
  }
}

func (x *graph) Mark2 (v, v1 vtx.Vertex) {
  if x.local == x.vAnchor { return }
  if ! x.Ex2 (v, v1) { return }
  x.local.bool = true
  x.colocal.bool = true
  e := x.connection (x.colocal, x.local)
  if e != nil {
    e.bool = true
  }
}

func (x *graph) AllMarked() bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! v.bool { return false }
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if ! e.bool { return false }
  }
  return true
}

func (x *graph) Del() {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.local == x.vAnchor { return }
//  delete all edges and their neighbour lists
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    delEdge (n.edgePtr)
    x.nEdges--
  }
  x.path = nil
//  x.ClrMarked() // XXX
  v := x.local
  delVertex (x.local)
  x.nVertices--
  x.local = v.prevV
  if x.local == x.vAnchor {
    x.local = x.vAnchor.nextV
  }
  x.colocal = x.local
}

func (x *graph) Del1() {
  if x.colocal == x.vAnchor || x.colocal == x.local {
    return
  }
  n := x.colocal.nbPtr.nextNb
  for n.to != x.local {
    if n == x.colocal.nbPtr {
      return // local no neighbour of colocal
    } else {
      n = n.nextNb
    }
  }
  delEdge (n.edgePtr)
  x.nEdges--
}

func wait() {
  return
  kbd.Wait (true)
}

func (x *graph) Len() uint {
  l := uint(0)
  if x.vAnchor == x.vAnchor.nextV {
    return l
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    l += Val (e.e)
  }
  return l
}

func (x *graph) LenMarked() uint {
  l := uint(0)
  if x.vAnchor == x.vAnchor.nextV {
    return l
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.bool {
      l += Val (e.e)
    }
  }
  return l
}

func (x *graph) NumNeighboursOut() uint {
  c := uint(0)
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      c++
    }
  }
  return c
}

func (x *graph) NumNeighboursIn() uint {
  c := uint(0)
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if ! n.outgoing {
      c++
    }
  }
  return c
}

func (x *graph) NumNeighbours() uint {
  c := uint(0)
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    c++
  }
  return c
}

/* func (x *graph) Leaf() {
  if ! x.bool || x.Empty() || ! x.Acyclic() {
    return
  }
  v := x.vAnchor.nextV
  for {
    x.local = v
    if x.NumNeighboursOut() == 0 {
      break
    }
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing {
        v = n.to
        break
      }
    }
  }
} */

func (x *graph) Inv() {
  if x.bool {
    for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
      e.nbPtr0.outgoing = ! e.nbPtr0.outgoing
      e.nbPtr1.outgoing = ! e.nbPtr1.outgoing
    }
  }
}

func (x *graph) InvLoc() {
  if x.local != x.vAnchor {
    if x.bool {
      for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
        n.edgePtr.nbPtr0.outgoing = ! n.edgePtr.nbPtr0.outgoing
        n.edgePtr.nbPtr1.outgoing = ! n.edgePtr.nbPtr1.outgoing
      }
    }
  }
}

func (x *graph) Relocate() {
  x.ClrMarked() // XXX
  x.local, x.colocal = x.colocal, x.local
  x.colocal.bool = true
  x.path = nil
  x.path = append (x.path, x.colocal)
}

func (x *graph) locate (colocal2local bool) {
  x.ClrMarked() // XXX
  if colocal2local {
    x.colocal = x.local
  } else {
    x.local = x.colocal
  }
  x.path = nil
}

func (x *graph) Locate (colocal2local bool) {
  x.locate (colocal2local)
  x.colocal.bool = true
  x.path = append (x.path, x.colocal)
}

func (x *graph) Located() bool {
  if x.vAnchor == x.vAnchor.nextV {
    return true
  }
  return x.local == x.colocal
}

func (x *graph) Colocate() {
  if x.Empty() { return }
  x.local, x.colocal = x.colocal, x.local
}

func (x *graph) InvertPath() {
  if x.bool { return }
  n := uint(len(x.path))
  if n == 0 { return }
  for i:= uint(0); i < n / 2; i++ {
    x.path[i], x.path[n - 1 - i] = x.path[n - 1 - i], x.path[i]
  }
  x.local, x.colocal = x.colocal, x.local
}

func contains (vs []*vertex, v *vertex) bool {
  l := uint(len (vs))
  c := l
  for i, a := range vs {
    if a == v {
      c = uint(i)
      break
    }
  }
  return c < l
}

func remove (vs []*vertex, i uint) []*vertex {
  l := uint(len (vs))
  if l == 0 { return nil }
  if i >= l { return vs }
  vs1 := make ([]*vertex, l - 1)
  copy (vs1[:i], vs[:i])
  copy (vs1[i:], vs[i+1:])
  return vs1
}

func (x *graph) Step (i uint, outgoing bool) {
  if x.vAnchor == x.vAnchor.nextV {
    return
  }
  if outgoing {
    if i >= x.NumNeighboursOut() { return }
    if x.path == nil {
      x.colocal.bool = true
      x.path = append (x.path, x.colocal)
      x.local = x.colocal
    } else {
      if x.path[0] != x.colocal { ker.Oops() }
    }
    c := uint(len (x.path))
    n := x.path[c - 1]
    if x.local != n { ker.Shit() }
    nb := x.local.nbPtr.nextNb
    for {
      if nb.outgoing {
        if i == 0 {
          break
        } else {
          i--
        }
      }
      nb = nb.nextNb
    }
    nb.edgePtr.bool = true
    nb.to.bool = true
    x.local = nb.to
    x.path = append (x.path, x.local) // insert (x.path, x.local, c)
  } else { // backward
    c := uint(len (x.path))
    if c <= 1 { return }
    x.local = x.path[c - 2]
    v := x.path[c - 1]
    c--
    x.path = remove (x.path, c)
    if ! contains (x.path, v) {
      v.bool = false
    }
    e := x.connection (x.local, v)
    if e == nil { ker.Oops() }
    e.bool = false
    i = uint(0)
    for {
      if i + 1 == c { break }
      v = x.path[i]
      v1 := x.path[i+1]
      if e == x.connection (v, v1) {
        e.bool = true
        break
      } else {
        i++
      }
    }
  }
}

func (x *graph) NeighbourOut (i uint) any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      if i == 0 {
        return Clone (n.to.v)
      } else {
        i--
      }
    }
  }
  return nil
}

func (x *graph) Outgoing (i uint) bool {
  if x.vAnchor.nextV == x.vAnchor {
    return false
  }
  if ! x.bool {
    return true
  }
  for n, j := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, j = n.nextNb, j + 1 {
    if j == i {
      return n.outgoing
    }
  }
  return false
}

func (x *graph) NeighbourIn (i uint) any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if ! n.outgoing {
      if i == 0 {
        return Clone (n.to.v)
      } else {
        i--
      }
    }
  }
  return nil
}

func (x *graph) Incoming (i uint) bool {
  if x.vAnchor.nextV == x.vAnchor {
    return false
  }
  if ! x.bool {
    return true
  }
  for n, j := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, j = n.nextNb, j + 1 {
    if j == i {
      return ! n.outgoing
    }
  }
  return false
}

func (x *graph) Neighbour (i uint) any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if i == 0 {
      return Clone (n.to.v)
    } else {
      i--
    }
  }
  return nil
}

func (x *graph) Conn() bool {
  return x.ConnCond (AllTrue)
}

func (x *graph) ConnCond (p Pred) bool {
  if x.vAnchor == x.vAnchor.nextV { return true }
  if x.colocal == x.local { return true }
  x.preDfs()
  x.search (x.colocal, x.colocal, p)
  return x.local.repr == x.colocal
  return x.local.time0 > 0 // Alternative
}

func (x *graph) Star() Graph {
  if x.vAnchor == x.vAnchor.nextV { return nil }
  y := new_(x.bool, x.vAnchor.v, x.eAnchor.e).(*graph)
  y.Ins (x.local.v)
  y.local.bool = true
  local := y.local
  if y.local != y.colocal { ker.Oops() }
  if ! x.bool {
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (n.to.v.Clone().(vtx.Vertex)) // a now colocal
      y.edgeMarked (n.edgePtr.e.Clone().(edg.Edge), true) // edge from a to local inserted vertex
                                              // with same v as in x
      y.local = local // a now again local in y
    }
  } else { // x.bool
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (n.to.v.Clone().(vtx.Vertex)) // a colocal, new inserted local
      if n.outgoing { // want edge from a to new inserted
      } else { // ! n.outgoing: want edge from new inserted to a
        y.local, y.colocal = y.colocal, y.local
      }
      y.edgeMarked (n.edgePtr.e.Clone().(edg.Edge), true) // edge in y from colocal to local
      y.local = local
    }
  }
  y.SetWrite (x.Writes())
  return y
}

func (x *graph) Add (Ys ...Graph) {
  for _, Y := range Ys {
    y := x.imp(Y)
    if y.bool != x.bool { ker.Oops() }
    for v := y.vAnchor.nextV; v != y.vAnchor; v = v.nextV {
      if x.Ex (v.v) && ! x.local.bool {
        x.local.bool = v.bool
      } else {
        x.insMarked (v.v, v.bool)
      }
    }
    for e := y.eAnchor.nextE; e != y.eAnchor; e = e.nextE {
      if x.Ex2 (e.nbPtr0.from.v, e.nbPtr1.from.v) {
        e1 := x.connection (x.colocal, x.local)
        e2 := x.connection (x.local, x.colocal)
        if e1 == nil && e2 == nil {
          x.edgeMarked (e.e, e.bool)
        } else if e1 != nil && ! e1.bool { // x.colacal already connected with x.local
          e1.bool = e.bool
        } else if e2 != nil && ! e2.bool {
          e2.bool = e.bool
        }
      }
    }
  }
}

func (x *graph) SetDemo (d Demo) {
  x.demo[d] = true
  if d == Cycle { x.demo[Depth] = true } // Cycle without Depth is pointless
}

func (x *graph) SetWrite (wv, we CondOp) {
  x.writeV, x.writeE = wv, we
}

func (x *graph) Writes() (CondOp, CondOp) {
  return x.writeV, x.writeE
}

func (x *graph) Write() {
  x.Trav1Cond (x.writeE)
  x.TravCond (x.writeV)
}
