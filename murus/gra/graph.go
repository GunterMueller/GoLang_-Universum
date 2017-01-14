package gra

// (c) murus.org  v. 170107 - license see murus.go

// >>> yet some things TODO

import (
//  "fmt"; "strconv";
  "sort"
  . "murus/obj"; "murus/ker"; "murus/str"; "murus/rand"
  "murus/kbd"; "murus/errh"
  "murus/pseq"; "murus/adj"
)

// References:
// CLR  = Cormen, Leiserson, Rivest        1990
// CLRS = Cormen, Leiserson, Rivest, Stein 2001

/*    vertex                                                                   vertex
                       neighbour                        neighbour
  [-----------]                                                          [-----------]
  [  content  ]       /------------------------------------------------->[  content  ]
  [-----------]<-----/---------------------------------------------\     [-----------]
  [   nbPtr --]-----/-----\                                  /------\----[-- nbPtr   ]
  [-----------]    /       |              edge              |        \   [-----------]
  [inSubgraph ]   /        V                                V         |  [inSubgraph ]
  [-----------]  |   [-----------]    [----------]    [-----------]   |  [-----------]
  [  marked   ]  |   [  edgePtr -]--->[  attrib  ]<---[- edgePtr  ]   |  [  marked   ]
  [-----|-----]  |   [-----------]    [----------]    [-----------]   |  [-----|-----]
  [dist |time ]<-----[-- from    ]<---[- nbPtr0  ]    [   from ---]----->[dist |time ]
  [-----|-----]  |   [-----------]    [----------]    [-----------]   |  [-----|-----]
  [predecessor]   \--[--- to     ]    [  nbPtr1 -]--->[    to ----]--/   [predecessor]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
  [    repr   ]      [  forward  ]    [inSubgraph]    [  forward  ]      [    repr   ]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
  [   nextV --]->    [  nextNb --]->  [  nextE --]->  [  nextNb --]->    [   nextV --]->
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]
<-[-- prevV   ]    <-[-- prevNb  ]  <-[-- prevE  ]  <-[-- prevNb  ]    <-[-- prevV   ]
  [-----------]      [-----------]    [----------]    [-----------]      [-----------]

The vertices of a graph are represented by structs,
whose field "content" represents the "real" vertex.
All vertices are connected in a doubly linked list with anchor cell,
that can be traversed to execute some operation on all vertices of the graph.

The edges are also represented by structs,
whose field "attrib" is either nil (with value 1)
or carries a variable of type Valuator.
Also all edges are connected in a doubly linked list with anchor cell.

For a vertex n one finds all outgoing and incoming edges
with the help of a further doubly linked ringlist of neighbour(hoodrelation)s
  nb1 = n.nbPtr, nb2 = n.nbPtr.nextNb, nb3 = n.nbPtr.nextNb.nextNb etc.
by the links outgoing from the nbi (i = 1, 2, 3, ...)
  nb1.edgePtr, nb2.edgePtr, nb3.edgePtr etc.
In directed graphs the edges outgoing from a vertex are exactly those ones
in the neighbourlist, for which forward == true.

For an edge e one finds its two vertices by the links
  e.nbPtr0.from = e.nbPtr1.to und e.nbPtr0.to = e.nbPtr1.from.

Semantics of some variables, that are "hidden" in fields of vAnchor:
  vAnchor.time0: in that the "time" is incremented for each search step
  vAnchor.marked: (after call of search) == true <=> graph has no cycles
*/

const (
  pack = "gra"
  suffix = pack
  inf = uint32(1<<32 - 1)
)
type (
  vertex struct {
      content Any
        nbPtr *neighbour
   inSubgraph,       // characteristic function of the vertices in the actual subgraph
       marked bool   // for the development of design patterns by clients
         dist,       // for breadth first search/Dijkstra and use in En/Decode
 time0, time1 uint32 // for applications of depth first search
  predecessor,       // for back pointers in depth first search and in ways
         repr,       // for the computation of connected components
 nextV, prevV *vertex
              }

  vSet []*vertex // to be able to apply sort.Sort to []*vertex

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

  eSet []*edge // to be able to apply sort.Sort to []*vertex

  neighbour struct {
           edgePtr *edge
          from, to *vertex
           forward bool
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
         write CondOp
        write3 CondOp3
               }
)

func newNode (a Any) *vertex {
  n := new (vertex)
  n.content = Clone(a)
  n.time1 = inf // for applications of depth first search
  n.dist = inf
  n.repr = n
  n.nextV, n.prevV = n, n
  return n
}

func (n *vertex) Clone() *vertex {
  return n
}

func insert (s []*vertex, n *vertex, i uint) []*vertex {
  l := uint(len (s))
  if i > l { i = l }
  s1 := make ([]*vertex, l + 1)
  copy (s1[:i], s[:i])
  s1[i] = n
  copy (s1[i+1:], s[i:])
  return s1
}

func contains (s []*vertex, n *vertex) bool {
  l := uint(len (s))
  c := l
  for i, a := range s {
    if a == n {
      c = uint(i)
      break
    }
  }
  return c < l
}

func exists (s []*vertex, p Pred) (*vertex, bool) {
  for _, a := range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}

func remove (s []*vertex, i uint) []*vertex {
  l := uint(len (s))
  if l == 0 { return nil }
  if i >= l { return s }
  s1 := make ([]*vertex, l - 1)
  copy (s1[:i], s[:i])
  copy (s1[i:], s[i+1:])
  return s1
}

func (s vSet) Len() int {
  return len (s)
}

func (s vSet) Swap (i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s vSet) Less (i, j int) bool {
  if s[i].dist == s[j].dist {
    if s[i] == s[j] {
      return false
    }
    return i < j
  }
  return s[i].dist < s[j].dist
}

func newEdge() *edge {
  e := new (edge)
  e.nextE, e.prevE = e, e
  return e
}

func (s eSet) Len() int {
  return len (s)
}

func (s eSet) Swap (i, j int) {
  s[i], s[j] = s[j], s[i]
}

func (s eSet) Less (i, j int) bool {
  return Val (s[i]) < Val (s[j])
}

func newNeighbour (e *edge, n, n1 *vertex, f bool) *neighbour {
  nb := new (neighbour)
  nb.edgePtr = e
  nb.from, nb.to = n, n1
  nb.forward = f
  nb.nextNb, nb.prevNb = nb, nb
  return nb
}

func existsnb (s []*neighbour, p Pred) (*neighbour, bool) {
  for _, a := range s {
    if p (a) {
      return a, true
    }
  }
  return nil, false
}

func newGra (d bool, n, e Any) Graph {
  CheckAtomicOrObject(n)
  x := new (graph)
  x.directed = d
  x.vAnchor, x.eAnchor = newNode(n), newEdge()
  if e == nil {
    x.eAnchor.attrib = nil
  } else {
    CheckUintOrValuator (e)
    x.eAnchor.attrib = Clone(e)
  }
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.write = CondIgnore
  x.write3 = CondIgnore3
  return x
}

func (x *graph) imp (Y Any) *graph {
  y, ok := Y.(*graph)
  if ! ok { TypeNotEqPanic (x, Y) }
  return y
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
  if e.nbPtr0 == nil { ker.Panic ("gra.delEdge: e.nbPtr0 == nil") }
  e.prevE.nextE, e.nextE.prevE = e.nextE, e.prevE
  e.nbPtr0.prevNb.nextNb, e.nbPtr0.nextNb.prevNb = e.nbPtr0.nextNb, e.nbPtr0.prevNb // bug
  e.nbPtr1.prevNb.nextNb, e.nbPtr1.nextNb.prevNb = e.nbPtr1.nextNb, e.nbPtr1.prevNb
}

func delNode (n *vertex) {
  N := n.nbPtr.nextNb
  for N != n.nbPtr {
    N = n.nbPtr
    N.to.predecessor = nil
    n.nbPtr = n.nbPtr.nextNb
  }
  n.prevV.nextV, n.nextV.prevV = n.nextV, n.prevV
  n = n.nextV
}

func (x *graph) Clr() {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    delEdge (e)
  }
  x.nEdges = 0
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    delNode (v)
  }
  x.nVertices = 0
  x.colocal, x.local = x.vAnchor, x.vAnchor
  x.path, x.eulerPath = nil, nil
}

func (x *graph) nE (n *vertex) uint {
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
  for xn := x.vAnchor.nextV; xn != x.vAnchor; xn = xn.nextV {
    if ! y.Ex (xn.content) {
      eq = false
      break
    }
    yn := y.local // y.local was changed
    if x.nE (xn) != y.nE (yn) {
      eq = false
      break
    }
    for xnb := xn.nbPtr; xnb.nextNb != xn.nbPtr; xnb = xnb.nextNb {
      for ynb := yn.nbPtr; ynb.nextNb != yn.nbPtr; ynb = ynb.nextNb {
        if ynb.to == xnb.to {
          aa := true
          if x.eAnchor.attrib != nil {
            if xnb.edgePtr == nil { break }
            if ynb.edgePtr == nil { break }
            aa = Eq (xnb.edgePtr.attrib, ynb.edgePtr.attrib)
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

// XXX The actual subgraph, the actual path and the actual verticestack are not copied.
func (x *graph) Copy (Y Any) {
  y := x.imp(Y)
  x.Clr()
  x.Decode (y.Encode())
}

func (x *graph) Clone() Any {
  y := newGra(x.directed, x.vAnchor.content, x.eAnchor.attrib)
  y.Copy (x)
  return y
}

func (x *graph) Num() uint {
  return uint(x.nVertices)
}

func (x *graph) NumAct() uint {
  c := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if v.inSubgraph {
      c++
    }
  }
  return c
}

func (x *graph) Num1Act() uint {
  c := uint(0)
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if e.inSubgraph {
      c++
    }
  }
  return c
}

func (x *graph) NumPred (p Pred) uint {
  c := uint(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.content) {
      c++
    }
  }
  return c
}

func (x *graph) Num1() uint {
  return uint(x.nEdges)
}

func (x *graph) insertedNode (a Any) *vertex {
  n := newNode (a)
  n.nbPtr = newNeighbour (nil, n, nil, false)
  n.nextV, n.prevV = x.vAnchor, x.vAnchor.prevV
  n.prevV.nextV = n
  x.vAnchor.prevV = n
  return n
}

func (x *graph) Ins (a Any) {
  if x.vAnchor.content == nil { ker.Oops() }
  CheckTypeEq (a, x.vAnchor.content)
  if x.Ex (a) { // local is set
    return
  }
  n := x.insertedNode (a)
  x.nVertices++
  if x.nVertices == 1 {
    x.colocal = n
  } else {
    x.colocal = x.local
  }
  x.local = n
}

func (x *graph) Edge() {
  if x.eAnchor.attrib != nil {
    ker.Panic ("cannot 'Edge' vertices with edges with an edgetype - use 'Edge1' ! ")
  }
  x.Edge1 (nil)
}

// Pre: n and n1 are vertices in the same graph.
// Returns nil, iff there is no edge from n to n1;
// returns otherwise the corresponding pointer.
func connection (n, n1 *vertex) *edge {
  for nb := n.nbPtr.nextNb; nb != n.nbPtr; nb = nb.nextNb {
    if nb.forward && nb.to == n1 {
      return nb.edgePtr
    }
  }
  return nil
}

// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (nb *neighbour, n *vertex) {
  nb.nextNb, nb.prevNb = n.nbPtr, n.nbPtr.prevNb
  nb.prevNb.nextNb = nb
  n.nbPtr.prevNb = nb
}

// TODO Spec
func (x *graph) insertEdge (a Any) {
  if ! TypeEq (a, x.eAnchor.attrib) {
    ker.Panic ("gra.insertEdge: ! TypeEq")
  }
  e := newEdge()
  if a == nil {
    e.attrib = nil
  } else {
    e.attrib = Clone (a)
  }
  e.nbPtr0 = newNeighbour (e, x.colocal, x.local, true)
  insertNeighbour (e.nbPtr0, x.colocal)
  e.nbPtr1 = newNeighbour (e, x.local, x.colocal, ! x.directed)
  insertNeighbour (e.nbPtr1, x.local)
  if e.nbPtr1 == nil { ker.Panic ("gra.insertEdge: e.nbPtr1 == nil") }
  e.nextE, e.prevE = x.eAnchor, x.eAnchor.prevE
  e.prevE.nextE = e
  x.eAnchor.prevE = e
}

func (x *graph) Edge1 (a Any) {
  if x.Empty() { return }
  if x.colocal == x.local {
//    ker.Panic ("gra.Edge1: colocal == local")
    return
  }
  if a != nil {
    CheckTypeEq (a, x.eAnchor.attrib)
  }
  // simple case: local and colocal are not yet adjacent:
  if connection (x.colocal, x.local) == nil &&
    connection (x.local, x.colocal) == nil {
    x.insertEdge (a)
    x.nEdges++
    return
  }
  // otherwise: an existing edge must not be cleared:
  if a == nil {
    return
  }
// if there is an edge from colocal to local, it is looked for:
  nb := x.colocal.nbPtr.nextNb
  for nb.to != x.local {
    nb = nb.nextNb
    if nb == x.colocal.nbPtr { ker.Stop (pack, 1) } // not found, contradiction
  }
// and its attrib is replaced:
  nb.edgePtr.attrib = Clone (a)
  nb.forward = true
// in the directed case the edge goes from colocal to local,
// but not the other way:
  if x.directed {
    nb = x.local.nbPtr.nextNb
    for nb.to != x.colocal {
      nb = nb.nextNb
      if nb == x.local.nbPtr { ker.Stop (pack, 2) }
    }
    nb.forward = false
  }
  x.nEdges++
}

func (x *graph) Edge2 (a, a1 Any) {
  x.edge3 (a, a1, nil)
}

func (x *graph) edge3 (a, a1, b Any) {
  if x.Empty() || // x.eAnchor.attrib == nil ||
    Eq (a, a1) ||
    ! TypeEq (a, a1) ||
    ! TypeEq (a, x.vAnchor.content) ||
    ! TypeEq (b, x.eAnchor.attrib) {
    return
  }
  if n, ok := x.found (a); ! ok {
    return
  } else {
    x.colocal = n
  }
  if n1, ok := x.found (a1); ! ok {
    return
  } else {
    x.local = n1
  }
  if x.colocal == x.local ||
    connection (x.colocal, x.local) != nil {
    return
  }
  x.Edge1 (b)
}

func (x *graph) Matrix() adj.AdjacencyMatrix {
  m := x.Num()
  vertices := make ([]Any, m)
  for n, i := x.vAnchor.nextV, uint32(0); n != x.vAnchor; n, i = n.nextV, i + 1 {
    vertices[i] = Clone (n.content)
    n.time0 = i
  }
  matrix := adj.New (m, x.eAnchor.attrib)
  for n, i := x.vAnchor.nextV, 0; n != x.vAnchor; n, i = n.nextV, i + 1 {
    for nb := n.nbPtr.nextNb; nb != n.nbPtr; nb = nb.nextNb {
      if nb.forward {
        matrix.Set (uint(i), uint(nb.to.time0), nb.edgePtr.attrib)
      }
    }
  }
  return matrix
}

func (x *graph) Set (n []Any, mat adj.AdjacencyMatrix) {
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
          x.Edge1 (mat.Val (i, k)) // XXX --> not uint, but Any (Valuator)
        }
      }
    }
  }
}

func (x *graph) Define (as []Any, ns[][]uint) {
  n := uint(len(as))
  if n != uint(len(ns)) {
    return
  }
  for _, a := range as {
    CheckTypeEq (a, x.vAnchor.content)
    x.Ins (a)
  }
  for i := uint(0); i < n; i++ {
    for _, k := range ns[i] {
      if x.directed || ! x.directed && i < k {
        if x.Ex2 (as[i], as[k]) {
          x.Edge() // as[i] is content of colocal, as[k] of local vertex
        }
      }
    }
  }
}

func (x *graph) Edged() bool {
  return connection (x.colocal, x.local) != nil
}

func (x *graph) CoEdged() bool {
  return connection (x.local, x.colocal) != nil
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
      x.colocal = v
      x.local = v1
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
      if e.nbPtr0.forward && e.nbPtr1.forward {
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
  if x.local == x.vAnchor { return nil }
  return Clone (x.local.content)
}

func (x *graph) Get2() (Any, Any) {
  if x.local == x.vAnchor { return nil, nil }
  return Clone (x.colocal.content), Clone (x.local.content)
}

func (x *graph) Get1() Any {
  if x.local == x.vAnchor { ker.Panic ("gra.Get1: local == vAnchor"); return nil }
  if x.local == x.colocal { ker.Panic ("gra.Get1: colocal == vAnchor"); return nil }
  nb := x.colocal.nbPtr.nextNb
  for {
    if nb == x.colocal.nbPtr { break }
    if nb.forward && nb.to == x.local {
      break
    }
    nb = nb.nextNb
  }
  return Clone (nb.edgePtr.attrib)
}

func (x *graph) Put (a Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  x.local.content = Clone (a)
}

func (x *graph) Put1 (a Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.nbPtr.edgePtr.attrib = Clone (a)
}

func (x *graph) Put2 (a, a1 Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.content = Clone (a)
  x.local.content = Clone (a1)
}

func (x *graph) Del() {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.local == x.vAnchor { return }
//  delete all edges and their neighbour lists
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    delEdge (nb.edgePtr)
    x.nEdges--
  }
  x.path = nil
  x.clearSubgraph()
  v := x.local
  delNode (x.local)
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
  nb := x.colocal.nbPtr.nextNb
  for nb.to != x.local {
    if nb == x.colocal.nbPtr {
      return // local no neighbour of colocal
    } else {
      nb = nb.nextNb
    }
  }
  delEdge (nb.edgePtr)
  x.nEdges--
}

func wait() {
  return
  kbd.Wait (true)
}

func (x *graph) preDepth() {
  x.vAnchor.time0 = 0
/*
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.time0, v.time1 = 0, 0
    v.predecessor = nil
    v.repr = v
  }
*/
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.time0, v.time1 = 0, 0
    v.predecessor, v.repr = nil, v
  }
}

func (x *graph) w3 (n *neighbour, a bool) {
  x.write3 (n.from.content, n.edgePtr.attrib, n.to.content, a)
}

// For all vertices n, that are accessible from n0 by a path, n.repr == n0.
// vAnchor.marked == true, if x has no cycles.
func (x *graph) search (v0, v *vertex, p Pred) {
  x.vAnchor.time0++
  v.time0 = x.vAnchor.time0
  v.repr = v0
  if x.demo [Depth] {
    x.write (v.content, true)
    wait()
  }
  for nb := v.nbPtr.nextNb; nb != v.nbPtr; nb = nb.nextNb {
    if nb.forward && nb.to != v.predecessor && p (nb.to.content) {
      if nb.to.time0 == 0 {
        if x.demo [Depth] {
          x.w3 (nb, true)
        }
        nb.to.predecessor = v
        x.search (v0, nb.to, p)
        if x.demo [Depth] {
          x.w3 (nb, false)
          wait()
        }
      } else if nb.to.time1 == 0 {
        x.vAnchor.marked = false // found cycle
        if x.demo [Cycle] { // also x.demo [Depth], see Set
          x.w3 (nb, true)
//          errh.Error0("Kreis gefunden")
          x.w3 (nb, false)
          wait()
        }
      }
    }
  }
  x.vAnchor.time0++
  v.time1 = x.vAnchor.time0
  if x.demo [Depth] {
    x.write (v.content, false)
  }
}

func (x *graph) Conn() bool {
  return x.ConnCond (AllTrue)
}

func (x *graph) ConnCond (p Pred) bool {
  if x.vAnchor == x.vAnchor.nextV { return true }
  if x.colocal == x.local { return true }
  x.preDepth()
  x.search (x.colocal, x.colocal, p)
  return x.local.repr == x.colocal
  return x.local.time0 > 0 // Alternative
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
    for nb := v.nbPtr.nextNb; nb != v.nbPtr; nb = nb.nextNb {
      if nb.forward && nb.to.dist == inf && p (nb.to.content) {
        if x.demo [Breadth] {
          var nb1 *neighbour
          if nb.to.predecessor == nil {
            // TODO then what ?
          } else {
            nb1 = nb.to.predecessor.nbPtr.nextNb
            for nb1.from != nb.to.predecessor {
              nb1 = nb1.nextNb
              if nb1.nextNb == nb1 { ker.Stop (pack, 3) }
            }
            x.w3 (nb1, false)
            x.write (nb1.from.content, true)
            wait()
          }
        }
        nb.to.dist = v.dist + 1
        nb.to.predecessor = v
        qu = append (qu, nb.to)
      }
    }
  }
}

// Algorithm of Dijkstra, Lit.: CLR 25.1-2, CLRS 24.2-3
// Pre: dist == inf, predecessor == nil for all vertices.
// TODO spec
func (x *graph) searchShortestPath (p Pred) {
  v := x.vAnchor.nextV
  set := make (vSet, x.nVertices)
  for i, v := 0, x.vAnchor.nextV; v != x.vAnchor; i, v = i + 1, v.nextV {
    set[i] = v
  }
  sort.Sort (set)
  for len (set) > 0 {
    v = set[0]
    if len (set) == 1 {
      set = nil
    } else {
      set = set[1:]
    }
    var d uint32
    for nb := v.nbPtr.nextNb; nb != v.nbPtr; nb = nb.nextNb {
      if nb.forward && nb.to != v.predecessor && p (nb.to.content) {
        if v.dist == inf {
          d = inf
        } else {
          d = v.dist + uint32(Val(nb.edgePtr.attrib))
        }
        if d < nb.to.dist {
          if x.demo [Breadth] {
            if nb.to.predecessor != nil {
              nb1 := nb.to.predecessor.nbPtr.nextNb
              for nb1.from != nb.to.predecessor {
                nb1 = nb1.nextNb
                if nb1.nextNb == nb1 { ker.Stop (pack, 4) }
              }
              x.w3 (nb1, false)
              x.write (nb.to.content, false)
            }
            x.w3 (nb, true)
            x.write (nb.to.content, true)
            wait()
          }
          nb.to.dist, nb.to.predecessor = d, v
// put the changed nb.to into the right position in set:
          sort.Sort (set)
        }
      }
    }
  }
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
    nb := v.nbPtr.nextNb
    for nb.to != v.predecessor {
      nb = nb.nextNb
      if nb == v.nbPtr { ker.Stop (pack, 5) }
    }
    nb.edgePtr.inSubgraph = true
    v = v.predecessor
  }
}

func (x *graph) Act() {
  x.ActPred (AllTrue)
}

func (x *graph) clearSubgraph() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.inSubgraph = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.inSubgraph = false
  }
}

func (x *graph) ActPred (p Pred) {
  v := x.vAnchor.nextV
  if v == x.vAnchor { errh.Error ("gra ###", 10123); return }
  if ! p (x.local.content) {
// errh.Error ("#", 11)
    return
  }
  x.clearSubgraph()
  if ! x.ConnCond (p) {
// errh.Error ("#", 12)
    return
  }
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

func (x *graph) LenAct() uint {
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

func (x *graph) NumLoc() uint {
  c := uint(0)
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    if nb.forward {
      c++
    }
  }
  return c
}

func (x *graph) NumLocInv() uint {
  c := uint(0)
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    if ! nb.forward {
      c++
    }
  }
  return c
}

func (x *graph) Inv() {
  if x.directed {
    for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
      e.nbPtr0.forward = ! e.nbPtr0.forward
      e.nbPtr1.forward = ! e.nbPtr1.forward
    }
  }
}

func (x *graph) InvLoc() {
  if x.local != x.vAnchor {
    if x.directed {
      for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
        nb.edgePtr.nbPtr0.forward = ! nb.edgePtr.nbPtr0.forward
        nb.edgePtr.nbPtr1.forward = ! nb.edgePtr.nbPtr1.forward
      }
    }
  }
}

func (x *graph) Relocate() {
  x.clearSubgraph()
  x.local, x.colocal = x.colocal, x.local
  x.colocal.inSubgraph = true
  x.path = nil
  x.path = append (x.path, x.colocal)
}

func (x *graph) locate (colocal2local bool) {
  x.clearSubgraph()
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

func (x *graph) Step (i uint, forward bool) {
  if x.vAnchor == x.vAnchor.nextV {
    return
  }
  if forward {
    if i >= x.NumLoc() { return }
    if x.path == nil {
      x.colocal.inSubgraph = true
      x.path = append (x.path, x.colocal)
      x.local = x.colocal
    } else {
      if x.path[0] != x.colocal { ker.Stop (pack, 6) }
    }
    c := uint(len (x.path))
    n := x.path[c - 1]
    if x.local != n {
//      fmt.Println (x.local.content); fmt.Println (n.content)
      ker.Stop (pack, 7) // shit happens
    }
    nb := x.local.nbPtr.nextNb
    for {
      if nb.forward {
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
    e := connection (x.local, v)
    if e == nil { ker.Stop (pack, 8) }
    e.inSubgraph = false
    i = uint(0)
    for {
      if i + 1 == c { break }
      v = x.path[i]
      v1 := x.path[i+1]
      if e == connection (v, v1) {
        e.inSubgraph = true
        break
      } else {
        i++
      }
    }
  }
}

func (x *graph) Neighbour (i uint) Any {
  if i >= x.NumLoc() || x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    if nb.forward {
      if i == 0 {
        return Clone (nb.to.content)
      } else {
        i--
      }
    }
  }
  return nil
}

func (x *graph) CoNeighbour (i uint) Any {
  if i >= x.NumLocInv() || x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    if ! nb.forward {
      if i == 0 {
        return Clone (nb.to.content)
      } else {
        i--
      }
    }
  }
  return nil
}

func (x *graph) Val (i uint) uint {
  if i >= x.NumLoc() || x.vAnchor.nextV == x.vAnchor {
    return 0
  }
  for nb := x.local.nbPtr.nextNb; nb != x.local.nbPtr; nb = nb.nextNb {
    if nb.forward {
      if i == 0 {
        return Val (nb.edgePtr.attrib)
      } else {
        i--
      }
    }
  }
  ker.Stop (pack, 9)
  return 0
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

func (x *graph) True (p Pred) bool {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if ! p (v.content) {
      return false
    }
  }
  return true
}

func (x *graph) TrueAct (p Pred) bool {
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

func (x *graph) TravPred (p Pred, o Op) {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    if p (v.content) {
      o (v.content)
    }
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
  for v := x.local.nbPtr.nextNb; v != x.local.nbPtr; v = v.nextNb {
//    if v.forward { // TODO
      o (v.edgePtr.attrib)
//    }
  }
}

/* func (x *graph) Trav1Coloc (op Op) {
  for v := x.colocal.nbPtr.nextNb; v != x.colocal.nbPtr; v = v.nextNb {
    if v.forward {
      op (v.edgePtr.attrib)
    }
  }
} */

func (x *graph) Trav3 (o Op, o3 Op3) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o3 (e.nbPtr0.from.content, e.attrib, e.nbPtr1.from.content)
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.content)
  }
}

func (x *graph) Trav3Cond (o CondOp, o3 CondOp3) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    o3 (e.nbPtr0.from.content, e.attrib, e.nbPtr1.from.content, e.inSubgraph)
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.content, v.inSubgraph)
  }
}

func (x *graph) Trav3CondDir (o CondOp, o3 CondOp3bool) {
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if ! e.nbPtr0.forward { ker.Oops() }
    if e.nbPtr1.forward == x.directed { ker.Oops() }
    o3 (e.nbPtr0.from.content, e.attrib, x.directed, e.nbPtr1.from.content, e.inSubgraph)
  }
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    o (v.content, v.inSubgraph)
  }
}

func (x *graph) Star (a Any) Graph {
  CheckTypeEq (a, x.vAnchor.content)
  y := newGra (x.directed, x.vAnchor.content, x.eAnchor.attrib).(*graph)
  if x.Ex (a) { // vertex with content a is local in x
    xlocal := x.local
    y.Ins (a) // vertex with content a is local and colocal in Star
    if y.local != y.colocal { ker.Oops() }
    k, k1 := x.NumLoc(), x.NumLocInv()
    b := make ([]Any, k + k1)
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      if n.forward {
        b[i] = Clone(n.to.content)
        y.Ins (b[i])
      } else { // ! n.forward
        b[i] = Clone(n.to.content)
        y.Ins (b[i])
      }
    }
    if xlocal != x.local { ker.Oops() }
    for n, i := x.local.nbPtr.nextNb, uint(0); n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      if n.forward {
        if x.Ex2 (a, b[i]) { // a colocal, b[i] local in x
          if y.Ex2 (a, b[i]) { // a colocal, b[i] local in Star
            y.Edge1 (x.Get1())
          }
        }
        x.local = xlocal // a again local in x
      } else { // ! n.forward
        if x.Ex2 (b[i], a) { // b[i] colocal, a local in x
          if y.Ex2 (b[i], a) { // b[i] colocal, a local in Star
            y.Edge1 (x.Get1())
          }
        }
      }
    }
    if ! y.Ex (a) { ker.Oops() } // a local in Star
  }
  return y
}

func (x *graph) StarLoc() Graph {
  return x.Star (x.local.content)
}

// Returns true, iff every vertex of x is accessible from every other one by a path.
func (x *graph) TotallyConnected() bool {
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

// CLR 23.3, CLRS 22.3
func (x *graph) dfs() {
  x.preDepth()
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
  if x.vAnchor.nextV == x.vAnchor { return true }
  x.vAnchor.marked = true
  x.dfs()
  return x.vAnchor.marked
}

func notTraversedNeighbour (n *vertex) *neighbour {
  for nb := n.nbPtr.nextNb; nb != n.nbPtr; nb = nb.nextNb {
    if nb.forward && ! nb.edgePtr.inSubgraph {
      return nb
    }
  }
  return nil
}

func notTraversed (a Any) bool {
  return notTraversedNeighbour (a.(*neighbour).from) != nil
}

func (x *graph) Euler() bool {
  if ! x.TotallyConnected() {
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
      if nb.forward {
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
    ker.Stop (pack, 100 + e)
  }
  x.clearSubgraph()
  x.eulerPath = nil
  x.colocal.inSubgraph = true
  v := x.colocal
  v.inSubgraph = true
//  for j := 0; j <= 9; j** { for a := false TO true { write (E.content, a); ker.Msleep (100) } }
// attempt, to find an Euler path/cycle "by good luck":
  var nb *neighbour
  for {
    nb = notTraversedNeighbour (v)
    if nb == nil { ker.Stop (pack, 11) }
    // write1 (N.edgePtr.attrib, true)
    //  for j := 0; j <= 9; j++ { for a := false; a <= true; a++ { write (N.to.content, a); ker.Msleep (100) } } };
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
    //   x.write1 (nb.edgePtr.attrib, a); ker.Msleep (100) } }
    v = nb.from
    v1 := v
    for {
      nb = notTraversedNeighbour (v)
      if nb == nil { ker.Stop (pack, 12) }
    // write1 (N.edgePtr.attrib, true)
    // for j := 0 TO 9 { for a := false TO true { write (N.to.content, a); ker.Msleep (100) } }
      nb.edgePtr.inSubgraph = true
      v = nb.to
      v.inSubgraph = true
      x.eulerPath = append (x.eulerPath, nb)
      if v == v1 { break } // found one mor cycle
    // errh.Error0("weiterer Teil eines Eulerwegs gefunden")
    }
  }
  if x.demo [Euler] {
    x.write (x.colocal.content, true)
    wait()
    for i := uint(0); i < uint(len (x.eulerPath)); i++ {
      nb = x.eulerPath[i]
      x.w3 (nb, true)
      if nb.edgePtr.nbPtr0 == nb {
        x.write (nb.edgePtr.nbPtr1.from.content, true)
      } else {
        x.write (nb.edgePtr.nbPtr0.from.content, true)
      }
      if i + 1 < uint(len (x.eulerPath)) {
        wait()
      }
    }
  }
  return true
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
  set := make (eSet, x.nEdges)
  for i, e := uint(0), x.eAnchor.nextE; e != x.eAnchor; i, e = i + 1, e.nextE {
    set[i] = e
    e.inSubgraph = false
  }
  sort.Sort (set)
  for len (set) > 0 {
    e := set[0]
    set = set [1:]
    v, v1 := e.nbPtr0.from, e.nbPtr1.from
    if x.demo [SpanTree] {
      x.write3 (v.content, e.attrib, v1.content, true) // TODO
      x.write (v.content, true)
      x.write (v1.content, true)
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
        x.write3 (v.content, e.attrib, v1.content, false) // TODO
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

func (x *graph) IsolateAct() {
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

func (x *graph) Equiv() bool {
  if x.Empty() {
    return false
  }
  x.Isolate()
  return x.local.repr == x.colocal.repr
}

func (x *graph) Codelen() uint {
  c := uint(1)
  c += 4
  if x.nVertices > 0 {
    for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
      c += 4 + Codelen (v.content)
    }
    c += 3 * 4
    if x.nEdges > 0 {
      for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
        if x.eAnchor.attrib != nil {
          c += 4 + Codelen (e.attrib)
        }
        c += 2 * (4 + Codelen (true))
      }
    }
  }
  return c
}

func (x *graph) Encode() []byte {
  bs := make ([]byte, x.Codelen())
  i, a := uint32(0), uint32(1)
  bs[i] = 0; if x.directed { bs[i] = 1}
  i += a
  a = 4
  copy (bs[i:i+a], Encode (x.nVertices))
  i += a
  if x.nVertices == 0 { return bs }
  z := uint32(0)
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    k := uint32(Codelen (v.content))
    copy (bs[i:i+a], Encode (k))
    i += a
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
  i += a
  if x.nEdges == 0 { return bs }
  b := uint32(Codelen (true))
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    if x.eAnchor.attrib != nil {
      k := uint32(Codelen (e.attrib))
      copy (bs[i:i+a], Encode (k))
      i += a
      copy (bs[i:i+k], Encode (e.attrib))
      i += k
    }
    copy (bs[i:i+a], Encode (e.nbPtr0.from.dist))
    i += a
    copy (bs[i:i+b], Encode (e.nbPtr0.forward))
    i += b
    copy (bs[i:i+a], Encode (e.nbPtr1.from.dist))
    i += a
    copy (bs[i:i+b], Encode (e.nbPtr1.forward))
    i += b
  }
  return bs
}

func (x *graph) Decode (bs []byte) {
  i, a := uint32(0), uint32(1)
  x.directed = bs[0] == 1
  i += a
  a = 4
  x.nVertices = Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  if x.nVertices == 0 {
    return
  }
  for n := uint32(0); n < x.nVertices; n++ {
    k := Decode (uint32(0), bs[i:i+a]).(uint32)
    i += a
    cont := Decode (Clone (x.vAnchor.content), bs[i:i+k])
    x.insertedNode (cont)
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
    v.inSubgraph = false
  }
  x.nEdges = Decode (uint32(0), bs[i:i+a]).(uint32)
  i += a
  if x.nEdges == 0 { return }
  b := uint32(Codelen (true))
  for z := uint32(0); z < x.nEdges; z++ {
    e := newEdge()
    if x.eAnchor.attrib == nil {
      e.attrib = nil
    } else {
      k := Decode (uint32(0), bs[i:i+a]).(uint32)
      i += a
      e.attrib = Decode (Clone (x.eAnchor.attrib), bs[i:i+k])
      i += k
    }
    z1 := Decode (uint32(0), bs[i:i+a]).(uint32)
    if z1 > x.nVertices { ker.Stop (pack, 13) }
    i += a
    v0 := x.vAnchor.nextV
    for z1 > 0 {
      v0 = v0.nextV
      z1--
    }
    bo := Decode (true, bs[i:i+b]).(bool)
    i += b
    e.nbPtr0 = newNeighbour (e, v0, nil, bo) // e.nbPtr0.to see below
    insertNeighbour (e.nbPtr0, v0)
    z1 = Decode (uint32(0), bs[i:i+a]).(uint32)
/*
    if z1 > x.nVertices { ker.Panic ("nVertices == " + strconv.Itoa(int(x.nVertices)) +
                                  " but z1 ==" + strconv.Itoa(int(z1))) }
*/
    i += a
    v0 = x.vAnchor.nextV
    for z1 > 0 {
      v0 = v0.nextV
      z1--
    }
    e.nbPtr0.to = v0
    bo = Decode (true, bs[i:i+b]).(bool)
    i += b
    d := e.nbPtr0.forward != bo
    if d != x.directed {
      s0, s1 := "decoded Graph is", " directed"
      s := s0 + s1; if x.directed { s = s0 + "not " + s1 }
      ker.Panic (s)
    }
    e.nbPtr1 = newNeighbour (e, v0, e.nbPtr0.from, bo)
    insertNeighbour (e.nbPtr1, v0)
    e.inSubgraph = false
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
    y.Trav (func (n Any) { if ! x.Ex(n) { x.Ins (n) } })
    for e := y.eAnchor.nextE; e != y.eAnchor; e = e.nextE {
      if x.Ex2 (e.nbPtr0.from.content, e.nbPtr1.from.content) {
//        if ! x.Ex1 (e.attrib) {
        if connection (x.local, x.colocal) == nil {
          x.Edge1 (e.attrib)
        }
      }
    }
  }
}

func (x *graph) SetDemo (d Demo) {
  x.demo[d] = true
  if d == Cycle { x.demo[Depth] = true } // Cycle without Depth is pointless
}

func (x *graph) Install (o CondOp, o3 CondOp3) {
  x.write, x.write3 = o, o3
}
