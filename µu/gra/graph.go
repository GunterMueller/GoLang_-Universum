package gra

// (c) Christian Maurer   v. 170506 - license see µu.go
//
// >>>  References:
// >>>  CLR  = Cormen, Leiserson, Rivest        (1990)
// >>>  CLRS = Cormen, Leiserson, Rivest, Stein (2001)

import (
//  "reflect"
  "sort"
  "µu/ker"
  . "µu/obj"
  "µu/str"
  "µu/rand"
  "µu/kbd"
  "µu/errh"
  "µu/pseq"
  "µu/adj"
)

/*    vertex           neighbour                        neighbour            vertex
   ___________                                                            ___________ 
  /           \         /----------------------------------------------->/           \
  |  content  |        /                                                 |  content  |
  |___________|<--------------------------------------------------\      |___________|
  |           |      /                                             \     |           |
  |   nbPtr---|-----------\                                  /-----------|---nbPtr   |
  |___________|    /       |                                |        \   |___________|
  |           |   |        |                                |         |  |           |
  |inSubgraph |   |        v              edge              V         |  |inSubgraph |
  |___________|   |   ___________      __________      ___________    |  |___________|
  |           |   |  /           \    /          \    /           \   |  |           |
  |  marked   |   |  | edgePtr---|--->|  attrib  |<---|--edgePtr  |   |  |  marked   |
  |___________|   |  |___________|    |__________|    |___________|   |  |___________|
  |     |     |   |  |           |    |          |    |           |   |  |     |     |
  |dist | time|<-----|---from    |<---|--nbPtr0  |    |   from----|----->|dist | time|
  |_____|_____|   |  |___________|    |__________|    |___________|   |  |_____|_____|
  |           |   \  |           |    |          |    |           |   |  |           |
  |predecessor|    --|----to     |    |  nbPtr1 -|--->|    to-----|--/   |predecessor|
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
  |    repr   |      | outgoing  |    |inSubgraph|    | outgoing  |      |    repr   |
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |           |
  |   nextV---|->    |  nextNb---|->  |  nextE---|->  |  nextNb---|->    |   nextV---|->
  |___________|      |___________|    |__________|    |___________|      |___________|
  |           |      |           |    |          |    |           |      |          _|
<-|---prevV   |    <-|---prevNb  |  <-|---prevE  |  <-|---prevNb  |    <-|---prevV   |
  \___________/      \___________/    \__________/    \___________/      \___________/

The vertices of a graph are represented by structs,
whose field "content" represents the "real" vertex.
All vertices are connected in a doubly linked list with anchor cell,
that can be traversed to execute some operation on all vertices of the graph.

The edges are also represented by structs,
whose field "attrib" is a variable of a type that implements Valuator.
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
  /bvAnchor.acyclic: (after call of search) == true <=> graph has no cycles. */

const (
  suffix = "gra"
  inf = uint32(1<<32 - 1)
)
type (
  vertex struct {
        content Any
          nbPtr *neighbour
     inSubgraph,       // characteristic function of the vertices in the actual subgraph
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
       attrib Any
       nbPtr0,
       nbPtr1 *neighbour
   inSubgraph bool // characteristic function of the vertices in the actual subgraph
 nextE, prevE *edge
              }

  neighbour struct {
           edgePtr *edge
          from, to *vertex
          outgoing bool
    nextNb, prevNb *neighbour
                   }

  graph struct {
          name,
      filename string
      directed bool
     nVertices,
        nEdges uint32
       vAnchor,
       colocal,
         local *vertex
       eAnchor *edge
          path []*vertex // XXX
     eulerPath []*neighbour
       nlocPtr,
     ncolocPtr *vCell
          demo Demoset
        writeV,
        writeE CondOp
               }
)

func newVertex (a Any) *vertex {
  v := new (vertex)
  v.content = Clone(a)
  v.time1 = inf // for applications of depth first search
  v.dist = inf
  v.repr = v
  v.nextV, v.prevV = v, v
  return v
}

func (n *vertex) Clone() *vertex {
  return n
}

func insert (s []*vertex, v *vertex, i uint) []*vertex {
  l := uint(len (s))
  if i > l { i = l }
  s1 := make ([]*vertex, l + 1)
  copy (s1[:i], s[:i])
  s1[i] = v
  copy (s1[i+1:], s[i:])
  return s1
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

func newEdge (a Any) *edge {
  e := new(edge)
  e.attrib = Clone(a)
  e.nextE, e.prevE = e, e
  return e
}

func newNeighbour (e *edge, v, v1 *vertex, f bool) *neighbour {
  nb := new(neighbour)
  nb.edgePtr = e
  nb.from, nb.to = v, v1
  nb.outgoing = f
  nb.nextNb, nb.prevNb = nb, nb
  return nb
}

func new_(d bool, v, e Any) Graph {
  CheckAtomicOrObject(v)
  x := new (graph)
  x.directed = d
  x.vAnchor = newVertex(v)
  if e == nil {
    e = uint(1)
  }
  CheckUintOrValuator (e)
  x.eAnchor = newEdge (e)
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.writeV, x.writeE = CondIgnore, CondIgnore
  return x
}

func (x *graph) imp (Y Any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
}

type nSeq []*neighbour

func (ns nSeq) Less (i, j int) bool {
  return ns[i].to.content.(Valuator).Val() < ns[j].to.content.(Valuator).Val()
}

func (x *graph) SortNeighbours() {
  switch x.vAnchor.content.(type) {
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
      ns[c - 1].nextNb = v.nbPtr
      v.nbPtr.prevNb = ns[c - 1]
      ns[0].prevNb = v.nbPtr
      for i := uint(1); i < c; i++ {
        ns[i].prevNb = ns[i - 1]
      }
    }
  }
}

func (x *graph) Name (s string) {
  x.name = s
  str.OffSpc (&x.name)
  if str.Empty (x.name) { x.name = "tmp" } // TODO + pid
  x.filename = x.name + "." + suffix
  n := pseq.Length (x.filename)
  if n > 0 {
    buf := make ([]byte, n)
    f := pseq.New (buf)
    f.Name (x.filename)
    buf = f.Get().([]byte)
    f.Fin()
    x.Decode (buf)
  }
}

func (x *graph) Rename (s string) {
  x.name = s
  str.OffSpc (&x.name)
  x.filename = x.name + "." + suffix
// rest of implementation TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO TODO
  n := pseq.Length (x.filename)
  if n > 0 {
    buf := make ([]byte, n)
    f := pseq.New (buf)
    f.Rename (x.name)
//    buf = f.Get().([]byte)
    f.Fin()
//    x.Decode (buf)
  }
}

func (x *graph) Fin() {
  if ! str.Empty (x.name) {
    buf := x.Encode()
    f := pseq.New (buf)
    f.Name (x.filename)
    f.Clr()
    f.Put (buf)
    f.Fin()
  }
//  x.Clr()
}

func (x *graph) Directed() bool {
  return x.directed
}

func (x *graph) Empty() bool {
  return x.vAnchor.nextV == x.vAnchor
}

func delEdge (e *edge) {
  if e.nbPtr0 == nil { ker.Panic("gra.delEdge: e.nbPtr0 == nil") }
  e.prevE.nextE, e.nextE.prevE = e.nextE, e.prevE
  e.nbPtr0.prevNb.nextNb, e.nbPtr0.nextNb.prevNb = e.nbPtr0.nextNb, e.nbPtr0.prevNb // bug
  e.nbPtr1.prevNb.nextNb, e.nbPtr1.nextNb.prevNb = e.nbPtr1.nextNb, e.nbPtr1.prevNb
}

func delVertex (v *vertex) {
  n := v.nbPtr.nextNb
  for n != v.nbPtr {
    n = v.nbPtr
    n.to.predecessor = nil
    v.nbPtr = v.nbPtr.nextNb
  }
  v.prevV.nextV, v.nextV.prevV = v.nextV, v.prevV
  v = v.nextV
}

func (x *graph) Clr() {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    delEdge (e)
  }
  x.nEdges = 0
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    delVertex (v)
  }
  x.nVertices = 0
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.path, x.eulerPath = nil, nil
}

func (x *graph) numEdges (n *vertex) uint {
  c := uint(0)
  for nb := n.nbPtr; nb.nextNb != n.nbPtr; nb = nb.nextNb {
    c++
  }
  return c
}

func (x *graph) Eq (Y Any) bool { // disgusting complexity
  y := x.imp (Y)
  if x.nVertices != y.nVertices || x.nEdges != y.nEdges ||
     ! TypeEq (x.vAnchor.content, y.vAnchor.content) ||
     ! TypeEq (x.eAnchor.attrib, y.eAnchor.attrib) {
    return false
  }
  ya := y.local // save
  eq := true
  loop:
  for xv := x.vAnchor.nextV; xv != x.vAnchor; xv = xv.nextV {
    if ! y.Ex (xv.content) {
      eq = false
      break
    }
    yv := y.local // y.local was changed
    if x.numEdges (xv) != y.numEdges (yv) {
      eq = false
      break
    }
    for xn := xv.nbPtr; xn.nextNb != xv.nbPtr; xn = xn.nextNb {
      for yn := yv.nbPtr; yn.nextNb != yv.nbPtr; yn = yn.nextNb {
        if yn.to == xn.to {
          aa := true
          if x.eAnchor.attrib != nil {
            if xn.edgePtr == nil { break }
            if yn.edgePtr == nil { break }
            aa = Eq (xn.edgePtr.attrib, yn.edgePtr.attrib)
          }
          if aa {
            break // next xnb
          } else {
            eq = false
            break loop
          }
        }
      }
    }
  }
  y.local = ya // restore
  return eq
}

func (x *graph) Less (Y Any) bool {
  return false
}

// XXX The actual path and the actual vertexstack are not copied.
func (x *graph) Copy (Y Any) {
  y := x.imp(Y)
  x.Decode (y.Encode())
  x.SetWrite (y.Writes())
}

func (x *graph) Clone() Any {
  y := new_(x.directed, x.vAnchor.content, x.eAnchor.attrib)
  y.Copy (x)
  return y
}

func (x *graph) Num() uint {
  return uint(x.nVertices)
}

func (x *graph) NumSub() uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.inSubgraph {
      n++
    }
  }
  return n
}

func (x *graph) NumSub1() uint {
  n := uint(0)
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.inSubgraph {
      n++
    }
  }
  return n
}

func (x *graph) NumPred (p Pred) uint {
  n := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.content) {
      n++
    }
  }
  return n
}

func (x *graph) Num1() uint {
  return uint(x.nEdges)
}

func (x *graph) insertedVertex (a Any, inSubgraph bool) *vertex {
  v := newVertex (a)
  v.inSubgraph = inSubgraph
  v.nbPtr = newNeighbour (nil, v, nil, false)
  v.nextV, v.prevV = x.vAnchor, x.vAnchor.prevV
  v.prevV.nextV = v
  x.vAnchor.prevV = v
  return v
}

func (x *graph) Ins (a Any) {
  x.insSub (a, false)
}

func (x *graph) insSub (a Any, inSubgraph bool) {
  if x.vAnchor.content == nil { ker.Oops() }
  CheckTypeEq (a, x.vAnchor.content)
  if x.Ex (a) { // local is set
    return
  }
  v := x.insertedVertex (a, inSubgraph)
  x.nVertices++
  if x.nVertices == 1 {
    x.colocal = v
  } else {
    x.colocal = x.local
  }
  x.local = v
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

// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (n *neighbour, v *vertex) {
  n.nextNb, n.prevNb = v.nbPtr, v.nbPtr.prevNb
  n.prevNb.nextNb = n
  v.nbPtr.prevNb = n
}

// TODO Spec
func (x *graph) insertedEdge (a Any, inSubgraph bool) *edge {
  CheckTypeEq (a, x.eAnchor.attrib)
//  if ! TypeEq (a, x.eAnchor.content) { ker.Panic ("gra.insertedEdge: ! TypeEq") }
  e := newEdge (a)
  e.inSubgraph = inSubgraph
//  e.inSubgraph = inSubgraph
  e.nbPtr0 = newNeighbour (e, x.colocal, x.local, true)
  insertNeighbour (e.nbPtr0, x.colocal)
  e.nbPtr1 = newNeighbour (e, x.local, x.colocal, ! x.directed)
  insertNeighbour (e.nbPtr1, x.local)
  if e.nbPtr1 == nil { ker.Panic ("gra.insertedEdge: e.nbPtr1 == nil") }
  e.nextE, e.prevE = x.eAnchor, x.eAnchor.prevE
  e.prevE.nextE = e
  x.eAnchor.prevE = e
  return e
}

func (x *graph) Edge (a Any) {
  x.edgeSub (a, false)
}

func (x *graph) edgeSub (a Any, inSubgraph bool) {
  if x.Empty() { return }
  if x.colocal == x.local { ker.Panic ("gra.Edge: colocal == local") }
  if a == nil { a = uint(1) }
  CheckTypeEq (a, x.eAnchor.attrib)
// simple case: local and colocal are not yet adjacent:
  e := x.connection (x.colocal, x.local)
  e1 := x.connection (x.local, x.colocal)
  if e == nil && e1 == nil {
    e = x.insertedEdge (a, inSubgraph)
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
// and its content is replaced:
  if ! x.directed {
    n.to.nbPtr.outgoing = true
  }
  n.edgePtr.attrib = Clone(a)
  n.edgePtr.inSubgraph = inSubgraph
  x.nEdges++
}

func (x *graph) Edge2 (v, v1, e Any) {
  if x.Empty() || Eq (v, v1) ||
    ! TypeEq (v, v1) || ! TypeEq (v, x.vAnchor.content) {
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

func (x *graph) Matrix() adj.AdjacencyMatrix {
  m := x.Num()
  vertices := make ([]Any, m)
  for n, i := x.vAnchor.nextV, uint32(0); n != x.vAnchor; n, i = n.nextV, i + 1 {
    vertices[i] = Clone (n.content)
    n.time0 = i
  }
  matrix := adj.New (m, x.eAnchor.attrib)
  for v, i := x.vAnchor.nextV, 0; v != x.vAnchor; v, i = v.nextV, i + 1 {
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing {
        matrix.Set (uint(i), uint(n.to.time0), n.edgePtr.attrib)
      }
    }
  }
  return matrix
}

func (x *graph) SetMatrix (n []Any, mat adj.AdjacencyMatrix) {
  if x.directed == mat.Symmetric() { ker.Panic ("gra.Set: x directed, matrix not") }
  m := mat.Num()
  for i := uint(0); i < m; i++ {
    x.Ins (n[i])
  }
  for i := uint(0); i < m; i++ {
    for k := uint(0); k < m; k++ {
      if mat.Edged (i, k) && (x.directed || i < k) {
        if x.ExPred2 (func (a Any) bool { return Eq (a, n[i]) },
                      func (a Any) bool { return Eq (a, n[k]) }) {
          x.Edge (mat.Val (i, k)) // XXX --> not uint, but Any (Valuator)
        }
      }
    }
  }
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
func (x *graph) found (a Any) (*vertex, bool) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if Eq (v.content, a) {
      return v, true
    }
  }
  return nil, false
}

func (x *graph) Ex (a Any) bool {
  if v, ok := x.found (a); ok {
    x.local = v
    return true
  }
  return false
}

func (x *graph) Ex2 (a, a1 Any) bool {
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
    if p (v.content) {
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

func (x *graph) Ex1 (e Any) bool {
  return x.ExPred1 (func (a Any) bool { return Eq (a, e) })
}

func (x *graph) ExPred1 (p Pred) bool {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if p (e.attrib) {
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

func (x *graph) Get() Any {
  return Clone (x.local.content)
}

func (x *graph) Get2() (Any, Any) {
  return Clone (x.colocal.content), Clone (x.local.content)
}

func (x *graph) Get1() Any {
  if x.local == x.vAnchor || x.local == x.colocal {
    return Clone (x.eAnchor.attrib) // XXX
  }
  e := x.connection (x.colocal, x.local)
  if e == nil {
    e = x.connection (x.local, x.colocal)
  }
  if e == nil || e.attrib == nil {
    return Clone (x.eAnchor.attrib) // XXX
  }
  return Clone (e.attrib)
}

func (x *graph) Put (v Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  x.local.content = Clone (v)
}

func (x *graph) Put1 (e Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.nbPtr.edgePtr.attrib = Clone (e)
}

func (x *graph) Put2 (v, v1 Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.content = Clone (v)
  x.local.content = Clone (v1)
}

func (x *graph) ClrSub() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.inSubgraph = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = false
  }
}

func (x *graph) Sub() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.inSubgraph = true
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = true
  }
}

func (x *graph) SubLocal() {
  if x.local == x.vAnchor { return }
  x.ClrSub()
  x.local.inSubgraph = true
}

func (x *graph) Sub2() {
  if x.local == x.vAnchor { return }
  x.local.inSubgraph = true
  x.colocal.inSubgraph = true
  e := x.connection (x.colocal, x.local)
  if e != nil {
    e.inSubgraph = true
  }
}

func (x *graph) EqSub() bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! v.inSubgraph { return false }
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if ! e.inSubgraph { return false }
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
//  x.ClrSub() // XXX
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

func (x *graph) defineSubgraph (v *vertex) {
  for v1 := v; v1 != x.colocal; v1 = v1.predecessor {
    if v1.predecessor == nil {
      return
    }
  }
  for {
    v.inSubgraph = true
    if v == x.colocal { return }
    n := v.nbPtr.nextNb
    for n.to != v.predecessor {
      n = n.nextNb
      if n == v.nbPtr { ker.Oops() }
    }
    n.edgePtr.inSubgraph = true
    v = v.predecessor
  }
}

func (x *graph) Act() {
  x.ActPred (AllTrue)
}

func (x *graph) ActPred (p Pred) {
  v := x.vAnchor.nextV
  if v == x.vAnchor { return }
  if ! p (x.local.content) { return }
  x.ClrSub()
  if ! x.ConnCond (p) { return }
  x.preBreadth()
  if x.eAnchor.attrib == nil {
    x.bfs (p)
  } else {
    x.searchShortestPath (p)
  }
  x.path = nil
  for v := x.local; v != nil; v = v.predecessor {
   x.path = insert (x.path, v, 0)
  }
  x.defineSubgraph (x.local)
}

func (x *graph) Len() uint {
  l := uint(0)
  if x.vAnchor == x.vAnchor.nextV {
    return l
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    l += Val (e.attrib)
  }
  return l
}

func (x *graph) LenSub() uint {
  l := uint(0)
  if x.vAnchor == x.vAnchor.nextV {
    return l
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.inSubgraph {
      l += Val (e.attrib)
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
  if ! x.directed || x.Empty() || ! x.Acyclic() {
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
  x.ClrSub() // XXX
  x.local, x.colocal = x.colocal, x.local
  x.colocal.inSubgraph = true
  x.path = nil
  x.path = append (x.path, x.colocal)
}

func (x *graph) locate (colocal2local bool) {
  x.ClrSub() // XXX
  if colocal2local {
    x.colocal = x.local
  } else {
    x.local = x.colocal
  }
  x.path = nil
}

func (x *graph) Locate (colocal2local bool) {
  x.locate (colocal2local)
  x.colocal.inSubgraph = true
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
  if x.directed { return }
  n := uint(len(x.path))
  if n == 0 { return }
  for i:= uint(0); i < n / 2; i++ {
    x.path[i], x.path[n - 1 - i] = x.path[n - 1 - i], x.path[i]
  }
  x.local, x.colocal = x.colocal, x.local
}

func (x *graph) Step (i uint, outgoing bool) {
  if x.vAnchor == x.vAnchor.nextV {
    return
  }
  if outgoing {
    if i >= x.NumNeighboursOut() { return }
    if x.path == nil {
      x.colocal.inSubgraph = true
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
    nb.edgePtr.inSubgraph = true
    nb.to.inSubgraph = true
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
      v.inSubgraph = false
    }
    e := x.connection (x.local, v)
    if e == nil { ker.Oops() }
    e.inSubgraph = false
    i = uint(0)
    for {
      if i + 1 == c { break }
      v = x.path[i]
      v1 := x.path[i+1]
      if e == x.connection (v, v1) {
        e.inSubgraph = true
        break
      } else {
        i++
      }
    }
  }
}

func (x *graph) NeighbourOut (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      if i == 0 {
        return Clone (n.to.content)
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

func (x *graph) NeighbourIn (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if ! n.outgoing {
      if i == 0 {
        return Clone (n.to.content)
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

func (x *graph) Neighbour (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if i == 0 {
      return Clone (n.to.content)
    } else {
      i--
    }
  }
  return nil
}

/*
func (x *graph) ValOut (i uint) uint {
  if x.vAnchor.nextV == x.vAnchor {
    return 0
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      if i == 0 {
        return Val (n.edgePtr.content)
      } else {
        i--
      }
    }
  }
  return 0
}

func (x *graph) ValIn (i uint) uint {
  if x.vAnchor.nextV == x.vAnchor {
    return 0
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if ! n.outgoing {
      if i == 0 {
        return Val (n.edgePtr.content)
      } else {
        i--
      }
    }
  }
  return 0
}

func (x *graph) Val (i uint) uint {
  if x.vAnchor.nextV == x.vAnchor {
    return 0
  }
  for n, j := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, j = n.nextNb, j + 1 {
    if j == i {
      return Val (n.edgePtr.content)
    }
  }
  return 0
}
*/

func (x *graph) LenStack() uint {
  c := uint(0)
  affe := x.nlocPtr // XXX
  for affe != nil {
    affe = affe.next
print ("lenstack  c++")
    c++
  }
  return c
}

func (x *graph) Save() {
  x.nlocPtr = new (vCell)
  x.nlocPtr.vPtr = x.local
  x.nlocPtr.next = x.nlocPtr
}

func (x *graph) Restore() {
  if x.nlocPtr == nil { return }
  x.local = x.nlocPtr.vPtr
  x.nlocPtr = x.nlocPtr.next
}

func (x *graph) CoSave() {
  x.ncolocPtr = new (vCell)
  x.ncolocPtr.vPtr = x.colocal
  x.ncolocPtr.next = x.ncolocPtr
}

func (x *graph) CoRestore() {
  if x.ncolocPtr == nil { return }
  x.colocal = x.ncolocPtr.vPtr
  x.ncolocPtr = x.ncolocPtr.next
}

func (x *graph) Mark (m bool) {
  if x.local == nil {
    return
  }
  x.local.marked = m
}

func (x *graph) Marked() bool {
  if x.vAnchor == x.vAnchor.nextV {
    return false
  }
  return x.local.marked
}

/*
func (x *graph) A() {
  x.local.dist = 0
}

func (x *graph) B() bool {
  return x.local.dist == 1 << 32 - 1
}

func (x *graph) C() {
  x.local.dist = x.colocal.dist + 1
  x.local.predecessor = x.colocal
}
*/

func (x *graph) True (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! p (v.content) {
      return false
    }
  }
  return true
}

func (x *graph) TrueSub (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.inSubgraph {
      if ! p (v.content) {
        return false
      }
    }
  }
  return true
}

func (x *graph) Trav (o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.content)
  }
}

func (x *graph) TravCond (o CondOp) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.content, v.inSubgraph)
  }
}

func (x *graph) Trav1 (o Op) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.attrib)
  }
}

func (x *graph) Trav1Cond (o CondOp) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o (e.attrib, e.inSubgraph)
  }
}

func (x *graph) Trav1Loc (o Op) {
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    o (n.edgePtr.attrib)
  }
}

func (x *graph) Trav1Coloc (o Op) {
  for n := x.colocal.nbPtr.nextNb; n != x.colocal.nbPtr; n = n.nextNb {
    o (n.edgePtr.attrib)
  }
}

func (x *graph) Star() Graph {
  y := new_(x.directed, x.vAnchor.content, x.eAnchor.attrib).(*graph)
  y.Ins (x.local.content)
  local := y.local
  if y.local != y.colocal { ker.Oops() }
  if ! x.directed {
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.content)) // a now colocal
      y.Edge (Clone(n.edgePtr.attrib)) // edge from a to local inserted vertex with same content as in x
      y.local = local // a now again local in y
    }
  } else { // x.directed
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.content)) // a colocal, new inserted local
      if n.outgoing { // want edge from a to new inserted
      } else { // ! n.outgoing: want edge from new inserted to a
        y.local, y.colocal = y.colocal, y.local
      }
      y.Edge (Clone(n.edgePtr.attrib)) // edge in y from colocal to local
      y.local = local
    }
  }
  return y
}

// For all vertices n, that are accessible from n0 by a path, n.repr == n0.
// vAnchor.acyclic == true, if x has no cycles.
func (x *graph) search (v0, v *vertex, p Pred) {
  x.vAnchor.time0++
  v.time0 = x.vAnchor.time0
  v.repr = v0
  if x.demo [Depth] {
    x.writeV (v.content, true)
    wait()
  }
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && n.to != v.predecessor && p (n.to.content) {
      if n.to.time0 == 0 {
        if x.demo [Depth] {
          x.writeE (n.edgePtr.attrib, true)
        }
        n.to.predecessor = v
        x.search (v0, n.to, p)
        if x.demo [Depth] {
          x.writeE (n.edgePtr.attrib, false)
          wait()
        }
      } else if n.to.time1 == 0 {
        x.vAnchor.acyclic = false // found cycle
        if x.demo [Cycle] { // also x.demo [Depth], see Set
          x.writeE (n.edgePtr.attrib, true)
//          errh.Error0("Kreis gefunden")
          x.writeE (n.edgePtr.attrib, false)
          wait()
        }
      }
    }
  }
  x.vAnchor.time0++
  v.time1 = x.vAnchor.time0
  if x.demo [Depth] {
    x.writeV (v.content, false)
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

func (x *graph) preBreadth() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.dist = inf
    v.predecessor = nil
  }
  x.colocal.dist = 0
}

// Lit.: CLR 23.2, CLRS 22.2
// TODO spec
func (x *graph) bfs (p Pred) {
  var qu []*vertex
  qu = append (qu, x.colocal)
  for len (qu) > 0 {
    v := qu[0]
    if len (qu) == 1 {
      qu = nil
    } else {
      qu = qu [1:]
    }
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing && n.to.dist == inf && p (n.to.content) {
        if x.demo [Breadth] {
          var n1 *neighbour
          if n.to.predecessor == nil {
            // TODO then what ?
          } else {
            n1 = n.to.predecessor.nbPtr.nextNb
            for n1.from != n.to.predecessor {
              n1 = n1.nextNb
              if n1.nextNb == n1 { ker.Oops() }
            }
            x.writeE (n1.edgePtr.attrib, false)
            x.writeV (n1.from.content, true)
            wait()
          }
        }
        n.to.dist = v.dist + 1
        n.to.predecessor = v
        qu = append (qu, n.to)
      }
    }
  }
}

type vSeq []*vertex

func (vs vSeq) Less (i, j int) bool {
  if vs[i].dist == vs[j].dist {
    if vs[i] == vs[j] { return false }
    return i < j
  }
  return vs[i].dist < vs[j].dist
}

// Algorithm of Dijkstra, Lit.: CLR 25.1-2, CLRS 24.2-3
// Pre: dist == inf, predecessor == nil for all vertices.
// TODO spec
func (x *graph) searchShortestPath (p Pred) {
  v := x.vAnchor.nextV
  vs := make (vSeq, x.nVertices)
  for i, v := 0, x.vAnchor.nextV; v != x.vAnchor; i, v = i + 1, v.nextV {
    vs[i] = v
  }
  sort.Slice (vs, vs.Less)
  for len (vs) > 0 {
    v = vs[0]
    if len (vs) == 1 {
      vs = nil
    } else {
      vs = vs[1:]
    }
    var d uint32
    for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
      if n.outgoing && n.to != v.predecessor && p (n.to.content) {
        if v.dist == inf {
          d = inf
        } else {
          d = v.dist + uint32(Val(n.edgePtr.attrib))
        }
        if d < n.to.dist {
          if x.demo [Breadth] {
            if n.to.predecessor != nil {
              n1 := n.to.predecessor.nbPtr.nextNb
              for n1.from != n.to.predecessor {
                n1 = n1.nextNb
                if n1.nextNb == n1 { ker.Oops() }
              }
              x.writeE (n1.edgePtr.attrib, false)
              x.writeV (n.to.content, false)
            }
            x.writeE (n.edgePtr.attrib, true)
            x.writeV (n.to.content, true)
            wait()
          }
          n.to.dist, n.to.predecessor = d, v
// put the changed nb.to into the right position in vs:
          sort.Slice (vs, vs.Less)
        }
      }
    }
  }
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

func (x *graph) Acyclic() bool {
  if x.Empty() { return true }
  x.vAnchor.acyclic = true
  x.dfs()
  return x.vAnchor.acyclic
}

// Kruskal's algorithm, CLR 24.1-2, CLRS 23.1-2
func (x *graph) MST() {
  if x.nVertices < 2 || x.directed || x.eAnchor.nextE == x.eAnchor {
    return
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.predecessor = nil
    v.repr = v
    v.inSubgraph = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = false
  }
  if x.nVertices == 1 {
    x.local = x.vAnchor.nextV
    x.local.inSubgraph = true
    return
  }
  es := make ([]*edge, x.nEdges)
  for i, e := uint(0), x.eAnchor.nextE; e != x.eAnchor; i, e = i + 1, e.nextE {
    es[i] = e
    e.inSubgraph = false
  }
  sort.Slice (es, func (i, j int) bool { return Val(es[i]) < Val(es[j]) })
  for len(es) > 0 {
    e := es[0]
    es = es[1:]
    v, v1 := e.nbPtr0.from, e.nbPtr1.from
    if x.demo [SpanTree] {
      x.writeE (e.attrib, true)
      x.writeV (v.content, true)
      x.writeV (v1.content, true)
      wait()
    }
    if v.repr != v1.repr {
      v.inSubgraph = true
      v1.inSubgraph = true
      e.inSubgraph = true
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
        x.writeE (e.attrib, false)
        x.writeV (v.content, false)
        x.writeV (v1.content, false)
        wait()
      }
    }
  }
}

// topological Sort, CLR 23.4, CLRS 22.4
// TODO spec
func (x *graph) Sort() {
  if x.nVertices < 2 || ! x.directed { return }
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
  if x.nVertices < 1 || ! x.directed {
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
    v.inSubgraph = true
  }
// and furthermore all edges, that connect two vertices in the same strongly connected component:
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = e.nbPtr0.from.repr == e.nbPtr1.from.repr
  }
}

func (x *graph) IsolateSub() {
  x.Isolate()
// only exactly those vertices in the actual subgraph, that
// are contained in the strong connection component of the local vertex:
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.inSubgraph = v.repr == x.local.repr
  }
// and furthermore exactly those edges, that connect these vertices:
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = e.nbPtr0.from.inSubgraph && e.nbPtr1.from.inSubgraph
  }
}

// Returns true, iff every vertex of x is accessible from every other one by a path.
func (x *graph) totallyConnected() bool {
  if x.nVertices <= 1 {
    return true
  }
  if x.directed {
    x.Isolate()
  } else {
    x.dfs()
  }
  v := x.vAnchor.nextV
  e0 := v.repr
  for v != x.vAnchor {
    if v.repr != e0 {
      return false
    }
    v = v.nextV
  }
  return true
}

func existsnb (s []*neighbour, p Pred) (*neighbour, bool) {
  for _, a := range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}

func notTraversedNeighbour (v *vertex) *neighbour {
  for n := v.nbPtr.nextNb; n != v.nbPtr; n = n.nextNb {
    if n.outgoing && ! n.edgePtr.inSubgraph {
      return n
    }
  }
  return nil
}

func notTraversed (a Any) bool {
  return notTraversedNeighbour (a.(*neighbour).from) != nil
}

func (x *graph) Euler() bool {
  if ! x.totallyConnected() {
    return false // TODO Fleury's algorithm
  }
  p := x.colocal
  a := x.local
  x.colocal = x.vAnchor
  x.local = x.vAnchor
  e := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
// check for existence of Euler cycles (iff graph
// not directed:
//   if each vertex has an even number of neighbours,
// directed:
//   if at any vertex the number of outgoing edges is equal to the number of incoming ones)
// or of Euler paths (iff graph
// not directed:
//   if there are exactly two vertices with odd number of neighbours,
// directed:
//   if exactly one vertex has one more outgoing than incoming edges
//   and exactly one vertex has one more incoming than outgoint edges)
    z := uint(0)
    z1 := uint(0)
    nb := v.nbPtr.nextNb
    for nb != v.nbPtr {
      if nb.outgoing {
        z++
      } else {
        z1++
      }
      nb = nb.nextNb
    }
    if x.directed {
      if z == z1 + 1 {
        if x.colocal == x.vAnchor {
          x.colocal = v
          e++
        } else {
          x.colocal = p
          x.local = a
          return false
        }
      } else if z1 == z + 1 {
        if x.local == x.vAnchor {
          x.local = v
          e++
        } else {
          x.colocal = p
          x.local = a
          return false
        }
      }
    } else { // ! x.directed
      if z % 2 == 1 {
        if x.colocal == x.vAnchor {
          x.colocal = v
        } else if x.local == x.vAnchor {
          x.local = v
        } else {
          x.colocal = p
          x.local = a
          return false
        }
        e++
      }
    }
  }
  switch e {
  case 0: // Euler cycle with random starting vertex
    x.colocal = x.vAnchor.nextV
    n := rand.Natural (uint(x.nVertices))
    for n > 0 {
      x.colocal = x.colocal.nextV
      n--
    }
    x.local = x.colocal
  case 1:
    x.colocal = p
    x.local = a
    return false
  case 2: // Euler path from colocal to local vertex
    ;
  default:
    ker.Shit()
  }
  x.ClrSub()
  x.eulerPath = nil
  x.colocal.inSubgraph = true
  v := x.colocal
  v.inSubgraph = true
//  for j := 0; j <= 9; j** { for a := false TO true { writeE (E.content, a); ker.Msleep (100) } }
// attempt, to find an Euler path/cycle "by good luck":
  var nb *neighbour
  for {
    nb = notTraversedNeighbour (v)
    if nb == nil { ker.Oops() }
    // writeE (N.edgePtr.content, true)
    //  for j := 0; j <= 9; j++ { for a := false; a <= true; a++ { writVe (N.to.content, a); ker.Msleep (100) } } };
    nb.edgePtr.inSubgraph = true
    v = nb.to
    v.inSubgraph = true
    x.eulerPath = append (x.eulerPath, nb)
    if v == x.local { break }
  }
// errh.Error0("erster Wegabschnitt gefunden");
// as long there are edges not yet traversed,
// look for vertices in the Euler path, from which such edges go out,
// and find more cycles starting there and insert them into the Euler path:
  for {
    nb, ok := existsnb (x.eulerPath, notTraversed)
    if ! ok { break }
    // for j := 0; j <= 9; j++ { for a := false; a <= true; a++ { // nonsense
    //   x.writeE (nb.edgePtr.content, a); ker.Msleep (100) } }
    v = nb.from
    v1 := v
    for {
      nb = notTraversedNeighbour (v)
      if nb == nil { ker.Oops() }
    // writeE (N.edgePtr.content, true)
    // for j := 0 TO 9 { for a := false TO true { writeV (N.to.content, a); ker.Msleep (100) } }
      nb.edgePtr.inSubgraph = true
      v = nb.to
      v.inSubgraph = true
      x.eulerPath = append (x.eulerPath, nb)
      if v == v1 { break } // found one mor cycle
    // errh.Error0("weiterer Teil eines Eulerwegs gefunden")
    }
  }
  if x.demo [Euler] {
    x.writeV (x.colocal.content, true)
    wait()
    for i := uint(0); i < uint(len (x.eulerPath)); i++ {
      nb = x.eulerPath[i]
      x.writeE (nb.edgePtr.attrib, true)
      if nb.edgePtr.nbPtr0 == nb {
        x.writeV (nb.edgePtr.nbPtr1.from.content, true)
      } else {
        x.writeV (nb.edgePtr.nbPtr0.from.content, true)
      }
      if i + 1 < uint(len (x.eulerPath)) {
        wait()
      }
    }
  }
  return true
}

func (x *graph) Equiv() bool {
  if x.Empty() {
    return false
  }
  x.Isolate()
  return x.local.repr == x.colocal.repr
}

func (x *graph) Codelen() uint {
  c := uint(1) + 4
  if x.nVertices > 0 {
    for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
      c += 4 + Codelen(v.content) + 1
    }
    c += 3 * 4
    if x.nEdges > 0 {
      for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
        c += 4 + Codelen(e.attrib) + 1 + 2 * (4 + 1)
      }
    }
  }
  return c
}

func (x *graph) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  bs[0] = 0; if x.directed { bs[0] = 1 }
  i, a := uint32(1), uint32(4)
  copy (bs[i:i+a], Encode (x.nVertices))
  if x.nVertices == 0 { return bs }
  i += a
  z := uint32(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    k := uint32(Codelen (v.content))
    copy (bs[i:i+a], Encode (k))
    i += a
    bs[i] = 0; if v.inSubgraph { bs[i] = 1 }
    i++
    copy (bs[i:i+k], Encode (v.content))
    i += k
    v.dist = z
    z++
  }
  copy (bs[i:i+a], Encode (x.colocal.dist))
  i += a
  copy (bs[i:i+a], Encode (x.local.dist))
  i += a
  copy (bs[i:i+a], Encode (x.nEdges))
  if x.nEdges == 0 { return bs }
  i += a
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if x.eAnchor.attrib == nil { ker.Oops() }
    k := uint32(Codelen (e.attrib))
    copy (bs[i:i+a], Encode (k))
    i += a
    copy (bs[i:i+k], Encode (e.attrib))
    i += k
    bs[i] = 0; if e.inSubgraph { bs[i] = 1 }
    i++
    copy (bs[i:i+a], Encode (e.nbPtr0.from.dist))
    i += a
    bs[i] = 0; if e.nbPtr0.outgoing { bs[i] = 1 }
    i++
    copy (bs[i:i+a], Encode (e.nbPtr1.from.dist))
    i += a
    bs[i] = 0; if e.nbPtr1.outgoing { bs[i] = 1 }
    i++
  }
  return bs
}

func (x *graph) check (s string, i, a uint32, bs []byte) {
  n := uint32(len(bs))
  if i >= n {
    errh.Error2(s + ": i =", uint(i), ">= len(bs) =", uint(n))
    as := bs[i:i+a]
    m := uint32(len(as))
    if m != a {
      errh.Error2("a =", uint(a), "!= len =", uint(m))
    }
  }
}

func (x *graph) Decode (bs []byte) {
  if len(bs) == 0 { panic("gra.Decode: len(bs) == 0") }
  x.Clr()
  x.directed = bs[0] == 1
  i, a := uint32(1), uint32(4)
  x.nVertices = Decode (uint32(0), bs[i:i+a]).(uint32)
  if x.nVertices == 0 {
    return
  }
  i += a
  for n := uint32(0); n < x.nVertices; n++ {
    k := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    inSubgraph := bs[i] == 1
    i++
    vc := Clone (x.vAnchor.content)
    content := Decode (x.vAnchor.content, bs[i:i+k])
    x.vAnchor.content = Clone (vc)
    x.insertedVertex (content, inSubgraph)
    i += k
  }
  p := Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  c := Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  for v, z := x.vAnchor.nextV, uint32(0); v != x.vAnchor; v, z = v.nextV, z + 1 {
    if z == p {
      x.colocal = v
    }
    if z == c {
      x.local = v
    }
  }
  x.nEdges = Decode (uint32(0), bs[i:i+a]).(uint32)
  if x.nEdges == 0 {
    return
  }
  i += a
  for z := uint32(0); z < x.nEdges; z++ {
    k := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    attrib := Decode (x.eAnchor.attrib, bs[i:i+k])
    e := newEdge (attrib)
    i += k
    e.inSubgraph = bs[i] == 1
    i++
    fromdist := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    v0 := x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0 = newNeighbour (e, v0, nil, bs[i] == 1) // e.nbPtr0.to see below
    i++
    insertNeighbour (e.nbPtr0, v0)
    fromdist = Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    v0 = x.vAnchor.nextV
    for fromdist > 0 {
      v0 = v0.nextV
      fromdist--
    }
    e.nbPtr0.to = v0
    dir := bs[i] == 1
    i++
    d := e.nbPtr0.outgoing != dir
    if d != x.directed {
      s := "decoded Graph is "; if x.directed { s += "not " }; ker.Panic(s + "directed")
    }
    e.nbPtr1 = newNeighbour (e, v0, e.nbPtr0.from, dir)
    insertNeighbour (e.nbPtr1, v0)
    e.nextE = x.eAnchor
    e.prevE = x.eAnchor.prevE
    e.prevE.nextE = e
    x.eAnchor.prevE = e
  }
  x.path, x.eulerPath = nil, nil
  x.nlocPtr, x.ncolocPtr = nil, nil
}

func (x *graph) Add (Ys ...Graph) {
  for _, Y := range Ys {
    y := x.imp(Y)
    if y.directed != x.directed { ker.Oops() }
    for v := y.vAnchor.nextV; v != y.vAnchor; v = v.nextV {
      if x.Ex (v.content) && ! x.local.inSubgraph {
        x.local.inSubgraph = v.inSubgraph
      } else {
        x.insSub (v.content, v.inSubgraph)
      }
    }
    for e := y.eAnchor.nextE; e != y.eAnchor; e = e.nextE {
      if x.Ex2 (e.nbPtr0.from.content, e.nbPtr1.from.content) {
        e1 := x.connection (x.colocal, x.local)
        e2 := x.connection (x.local, x.colocal)
        if e1 == nil && e1 == nil {
          x.edgeSub (e.attrib, e.inSubgraph)
        } else if e1 != nil && ! e1.inSubgraph { // x.colacal already connected with x.local
          e1.inSubgraph = e.inSubgraph
        } else if e2 != nil && ! e2.inSubgraph {
          e2.inSubgraph = e.inSubgraph
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
