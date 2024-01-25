package gra

// (c) Christian Maurer   v. 231226 - license see µU.go

import
  . "nU/obj"

/*/   vertex           neighbour                        neighbour            vertex
    _________                                                              _________  
   /         \          /-----------------------------------------------> /         \
  |    any    |        /                                                 |    any    |
  |___________|<--------------------------------------------------\      |___________|
  |           |      /                                             \     |           |
  |   nbPtr---|-----------\                                  /-----------|---nbPtr   |
  |___________|    /       |                                |        \   |___________|
  |           |   |        |                                |         |  |           |
  |  marked   |   |        v              edge              v         |  |  marked   |
  |___________|   |    _________        ________        _________     |  |___________|
  |           |   |   /         \      /        \      /         \    |  |           |
  |   dist    |   |  | edgePtr---|--->|   any    |<---|--edgePtr  |   |  |   dist    |
  |___________|   |  |___________|    |__________|    |___________|   |  |___________|
  |           |   |  |           |    |          |    |           |   |  |           |
  |predecessor|<-----|---from    |<---|--nbPtr0  |    |   from----|----->|predecessor|
  |___________|   |  |___________|    |__________|    |___________|   |  |___________|
  |           |   |  |           |    |          |    |           |   |  |           |
  |   repr    |    \_|____to     |    |  nbPtr1 -|--->|    to_____|__/   |   repr    |
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
  |   nextV---|->    | outgoing  |    |  marked  |    | outgoing  |      |   nextV---|->
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
    vAnchor.acyclic: (after call of search1) == true <=> graph has no cycles.
/*/

const (
  suffix = "gra"
  inf = uint32(1<<32 - 1)
)
type (
  vertex struct {
                any "content of the vertex"
          nbPtr *neighbour
         marked,
        acyclic bool   // for the development of design patterns by clients
           dist,       // for breadth first search/Dijkstra and use in En/Decode
   time0, time1 uint32 // for applications of depth first search
    predecessor,       // for back pointers in depth first search and in ways
           repr,       // for the computation of connected components
   nextV, prevV *vertex
                }

  vCell struct {
          vPtr *vertex
          next *vCell
               }

  edge struct {
              any "attribute of the edge"
       nbPtr0,
       nbPtr1 *neighbour
       marked bool
 nextE, prevE *edge
              }

  neighbour struct {
           edgePtr *edge
          from, to *vertex
          outgoing bool
    nextNb, prevNb *neighbour
                   }

  graph struct {
//          name,
//      filename string
//          file pseq.PersistentSequence
      directed bool
     nVertices,
        nEdges uint32
       vAnchor,
       colocal,
         local *vertex
       eAnchor *edge
          path []*vertex
     eulerPath []*neighbour
//          demo Demoset
         write CondOp
        write2 CondOp2
           val,
         sinks []uint
    succ, pred []uint
               }
)
type
  nSeq []*neighbour

func newVertex (a any) *vertex {
  v := new(vertex)
  v.any = Clone(a)
  v.time1 = inf // for applications of depth first search
  v.dist = inf
  v.repr = v
  v.nextV, v.prevV = v, v
  return v
}

func newEdge (a any) *edge {
  e := new(edge)
  e.any = Clone(a)
  e.nextE, e.prevE = e, e
  return e
}

func new_(d bool, v, e any) Graph {
  CheckAtomicOrObject(v)
  x := new(graph)
  x.directed = d
  x.vAnchor = newVertex (v)
  if e == nil {
    e = uint(1)
  }
  CheckUintOrValuator (e)
  x.eAnchor = newEdge (e)
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.write, x.write2 = CondIgnore, CondIgnore2
  x.val = make([]uint, 0)
  x.sinks = make([]uint, 0)
  x.succ = make([]uint, 0)
  x.pred = make([]uint, 0)
  return x
}

func (x *graph) imp (Y any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

func (x *graph) Directed() bool {
  return x.directed
}

func (x *graph) Num() uint {
  return uint(x.nVertices)
}

func (x *graph) NumMarked() uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.marked {
      n++
    }
  }
  return n
}

func (x *graph) NumMarked1() uint {
  n := uint(0)
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.marked {
      n++
    }
  }
  return n
}

func (x *graph) NumPred (p Pred) uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.any) {
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

func (x *graph) insertedVertex (a any, marked bool) *vertex {
  v := newVertex (a)
  v.marked = marked
  v.nbPtr = newNeighbour (nil, v, nil, false)
  v.nextV, v.prevV = x.vAnchor, x.vAnchor.prevV
  v.prevV.nextV = v
  x.vAnchor.prevV = v
  return v
}

func (x *graph) insMarked (a any, marked bool) {
//  if x.vAnchor.any == nil { ker.Oops() }
  CheckTypeEq (a, x.vAnchor.any)
  if x.Ex (a) { // local is set
    return
  }
  v := x.insertedVertex (a, marked)
  x.nVertices++
  if x.nVertices == 1 {
    x.colocal = v
  } else {
    x.colocal = x.local
  }
  x.local = v
}

func (x *graph) Ins (a any) {
  x.insMarked (a, false)
}

// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (n *neighbour, v *vertex) {
  n.nextNb, n.prevNb = v.nbPtr, v.nbPtr.prevNb
  n.prevNb.nextNb = n
  v.nbPtr.prevNb = n
}

// TODO Spec
func (x *graph) insertedEdge (a any, marked bool) *edge {
  CheckTypeEq (a, x.eAnchor.any)
//  if ! TypeEq (a, x.eAnchor.any) { ker.Panic ("gra.insertedEdge: ! TypeEq") }
  e := newEdge (a)
  e.marked = marked
  e.nbPtr0 = newNeighbour (e, x.colocal, x.local, true)
  insertNeighbour (e.nbPtr0, x.colocal)
  e.nbPtr1 = newNeighbour (e, x.local, x.colocal, ! x.directed)
  insertNeighbour (e.nbPtr1, x.local)
//  if e.nbPtr1 == nil { ker.Panic ("gra.insertedEdge: e.nbPtr1 == nil") }
  e.nextE, e.prevE = x.eAnchor, x.eAnchor.prevE
  e.prevE.nextE = e
  x.eAnchor.prevE = e
  return e
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

func (x *graph) edgeMarked (a any, marked bool) {
  if x.Empty() { return }
  if x.colocal == x.local { panic ("gra.edgeMarked: colocal == local") }
  if a == nil { a = uint(1) }
  CheckTypeEq (a, x.eAnchor.any)
// simple case: local and colocal are not yet adjacent:
  e := x.connection (x.colocal, x.local)
  e1 := x.connection (x.local, x.colocal)
  if e == nil && e1 == nil {
    e = x.insertedEdge (a, marked)
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
  if n.to != x.local { panic ("oops") }
// and its content is replaced:
  if ! x.directed {
    n.to.nbPtr.outgoing = true
  }
  n.edgePtr.any = Clone(a)
  n.edgePtr.marked = marked
  x.nEdges++
}

func (x *graph) Edge (a any) {
  if x.Empty() { return }
  if x.colocal == x.local { panic ("gra.Edge: colocal == local") }
  x.edgeMarked (a, false)
}

func (x *graph) Edge2 (v, v1, e any) {
  if x.Empty() || Eq (v, v1) ||
    ! TypeEq (v, v1) || ! TypeEq (v, x.vAnchor.any) {
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

func (x *graph) Edged() bool {
  return x.connection (x.colocal, x.local) != nil
}

func (x *graph) CoEdged() bool {
  return x.connection (x.local, x.colocal) != nil
}

// Returns (nil, false), iff a there is no vertex in x with content a;
// returns otherwise (n, true), where n is the pointer to that vertex
// (which is unique because of effect of Ins).
func (x *graph) found (a any) (*vertex, bool) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if Eq (v.any, a) {
      return v, true
    } else {
    }
  }
  return nil, false
}

func (x *graph) Ex (a any) bool {
  if v, ok := x.found (a); ok {
    x.local = v
    return true
  }
  return false
}

func (x *graph) Ex2 (a, a1 any) bool {
  if Eq (a, a1) {
    return false
  }
  if v, ok := x.found (a); ok {
    if v1, ok1 := x.found (a1); ok1 {
      x.colocal, x.local = v, v1
      return true
    }
  }
  return false
}

// Returns true, iff there is no vertex in x, for which p returns true,
// Returns otherwise a pointer to such a vertex.
func (x *graph) foundPred (p Pred) (*vertex, bool) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.any) {
      return v, true
    }
  }
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

func (x *graph) Ex1 (e any) bool {
  return x.ExPred1 (func (a any) bool { return Eq (a, e) })
}

func (x *graph) ExPred1 (p Pred) bool {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if p (e.any) {
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

//func (g *graph) Val() uint {
//  return g.local.any.(vtx.Vertex).Val()
//}

//func (g *graph) ExVal (u uint) bool {
//  var v vtx.Vertex
//  if g.ExPred (func (a any) bool { v = a.(vtx.Vertex); return v.(Valuator).Val() == u }) {
//    if ! g.Ex (v) { ker.Oops() }
//    return true
//  }
//  return false
//}

func (x *graph) Get() any {
  return Clone (x.local.any)
}

func (x *graph) Get2() (any, any) {
  return Clone (x.colocal.any), Clone (x.local.any)
}

func (x *graph) Get1() any {
  if x.local == x.vAnchor || x.local == x.colocal {
    return Clone (x.eAnchor.any)
  }
  e := x.connection (x.colocal, x.local)
  if e == nil {
    e = x.connection (x.local, x.colocal)
  }
  if e == nil || e.any == nil {
    return Clone (x.eAnchor.any)
  }
  return Clone (e.any)
}

func (x *graph) Put (v any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  x.local.any = Clone (v)
}

func (x *graph) Put1 (e any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.nbPtr.edgePtr.any = Clone (e)
}

func (x *graph) Put2 (v, v1 any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.any = Clone (v)
  x.local.any = Clone (v1)
}

func (x *graph) ClrMarked() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.marked = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.marked = false
  }
}

func (x *graph) Mark (v any) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { return }
  x.local.marked = true
}

func (x *graph) Mark1 (v any) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { return }
  x.local.marked = true
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    n.edgePtr.marked = true
  }
}

func (x *graph) Mark2 (v, v1 any) {
  if x.local == x.vAnchor { return }
  if ! x.Ex2 (v, v1) { return }
  x.local.marked = true
  x.colocal.marked = true
  e := x.connection (x.colocal, x.local)
  if e != nil {
    e.marked = true
  }
}

func (x *graph) AllMarked() bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! v.marked { return false }
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if ! e.marked { return false }
  }
  return true
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

func (x *graph) Inv() {
  if x.directed {
    for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
      e.nbPtr0.outgoing = ! e.nbPtr0.outgoing
      e.nbPtr1.outgoing = ! e.nbPtr1.outgoing
    }
  }
}

func (x *graph) InvLoc() {
  if x.local != x.vAnchor {
    if x.directed {
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
  x.colocal.marked = true
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
  x.colocal.marked = true
  x.path = append (x.path, x.colocal)
}

func (x *graph) NeighbourOut (i uint) any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      if i == 0 {
        return Clone (n.to.any)
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
  if ! x.directed {
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
        return Clone (n.to.any)
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
  if ! x.directed {
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
      return Clone (n.to.any)
    } else {
      i--
    }
  }
  return nil
}

func (x *graph) Star() Graph {
  if x.vAnchor == x.vAnchor.nextV { return nil }
  y := new_(x.directed, x.vAnchor.any, x.eAnchor.any).(*graph)
  y.Ins (x.local.any)
  y.local.marked = true
  local := y.local
  if y.local != y.colocal { panic ("Oops") }
  if ! x.directed {
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.any)) // a now colocal
      y.edgeMarked (Clone(n.edgePtr.any), true) // edge from a to local inserted vertex
                                                // with same content as in x
      y.local = local // a now again local in y
    }
  } else { // x.bool
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.any)) // a colocal, new inserted local
      if n.outgoing { // want edge from a to new inserted
      } else { // ! n.outgoing: want edge from new inserted to a
        y.local, y.colocal = y.colocal, y.local
      }
      y.edgeMarked (Clone(n.edgePtr.any), true) // edge in y from colocal to local
      y.local = local
    }
  }
//  y.SetWrite (x.Writes())
  return y
}

func (x *graph) Add (Ys ...Graph) {
  for _, Y := range Ys {
    y := x.imp(Y)
    if y.directed != x.directed { panic("Oops") }
    for v := y.vAnchor.nextV; v != y.vAnchor; v = v.nextV {
      if x.Ex (v.any) && ! x.local.marked {
        x.local.marked = v.marked
      } else {
        x.insMarked (v.any, v.marked)
      }
    }
    for e := y.eAnchor.nextE; e != y.eAnchor; e = e.nextE {
      if x.Ex2 (e.nbPtr0.from.any, e.nbPtr1.from.any) {
        e1 := x.connection (x.colocal, x.local)
        e2 := x.connection (x.local, x.colocal)
        if e1 == nil && e2 == nil {
          x.edgeMarked (e.any, e.marked)
        } else if e1 != nil && ! e1.marked { // x.colacal already connected with x.local
          e1.marked = e.marked
        } else if e2 != nil && ! e2.marked {
          e2.marked = e.marked
        }
      }
    }
  }
}

func (x *graph) SetWrite (w CondOp, w2 CondOp2) {
  x.write, x.write2 = w, w2
}

//func (x *graph) Writes() (CondOp, CondOp) {
//  return x.writeV, x.writeE
//}

func (x *graph) Write() {
  x.Trav2Cond (x.write2)
  x.TravCond (x.write)
}
