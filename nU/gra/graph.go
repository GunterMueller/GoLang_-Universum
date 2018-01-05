package gra

// (c) Christian Maurer   v. 171227 - license see nU.go

import ("sort"; . "nU/obj")

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
    if p (v.Any) {
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

func (x *graph) insertedVertex (a Any, marked bool) *vertex {
  v := newVertex (a)
  v.bool = marked
  v.nbPtr = newNeighbour (nil, v, nil, false)
  v.nextV, v.prevV = x.vAnchor, x.vAnchor.prevV
  v.prevV.nextV = v
  x.vAnchor.prevV = v
  return v
}

func (x *graph) insMarked (a Any, marked bool) {
  CheckTypeEq (a, x.vAnchor.Any)
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

func (x *graph) Ins (a Any) {
  x.insMarked (a, false)
}

func newEdge (a Any) *edge {
  e := new(edge)
  e.Any = Clone(a)
  e.nextE, e.prevE = e, e
  return e
}

// Pre: nb.from == e.
// nb is appended in n.nbPtr
func insertNeighbour (n *neighbour, v *vertex) {
  n.nextNb, n.prevNb = v.nbPtr, v.nbPtr.prevNb
  n.prevNb.nextNb = n
  v.nbPtr.prevNb = n
}

// TODO Spec
func (x *graph) insertedEdge (a Any, marked bool) *edge {
  CheckTypeEq (a, x.eAnchor.Any)
  e := newEdge (a)
  e.bool = marked
  e.nbPtr0 = newNeighbour (e, x.colocal, x.local, true)
  insertNeighbour (e.nbPtr0, x.colocal)
  e.nbPtr1 = newNeighbour (e, x.local, x.colocal, ! x.bool)
  insertNeighbour (e.nbPtr1, x.local)
  e.nextE, e.prevE = x.eAnchor, x.eAnchor.prevE
  e.prevE.nextE = e
  x.eAnchor.prevE = e
  return e
}

func (x *graph) Edge (a Any) {
  x.edgeMarked (a, false)
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

func (x *graph) edgeMarked (a Any, marked bool) {
  if x.Empty() { return }
  if x.colocal == x.local { panic ("gra.Edge: colocal == local") }
  if a == nil { a = uint(1) }
  CheckTypeEq (a, x.eAnchor.Any)
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
  }
// and its content is replaced:
  if ! x.bool {
    n.to.nbPtr.outgoing = true
  }
  n.edgePtr.Any = Clone(a)
  n.edgePtr.bool = marked
  x.nEdges++
}

func (x *graph) Edge2 (v, v1, e Any) {
  if x.Empty() || Eq (v, v1) ||
    ! TypeEq (v, v1) || ! TypeEq (v, x.vAnchor.Any) {
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

type nSeq []*neighbour

func (ns nSeq) Less (i, j int) bool {
  return ns[i].to.Any.(Valuator).Val() < ns[j].to.Any.(Valuator).Val()
}

func (x *graph) SortNeighbours() {
  switch x.vAnchor.Any.(type) {
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
    if Eq (v.Any, a) {
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
    if p (v.Any) {
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
    if p (e.Any) {
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
  return Clone (x.local.Any)
}

func (x *graph) Get2() (Any, Any) {
  return Clone (x.colocal.Any), Clone (x.local.Any)
}

func (x *graph) Get1() Any {
  if x.local == x.vAnchor || x.local == x.colocal {
    return Clone (x.eAnchor.Any) // XXX
  }
  e := x.connection (x.colocal, x.local)
  if e == nil {
    e = x.connection (x.local, x.colocal)
  }
  if e == nil || e.Any == nil {
    return Clone (x.eAnchor.Any) // XXX
  }
  return Clone (e.Any)
}

func (x *graph) Put (v Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  x.local.Any = Clone (v)
}

func (x *graph) Put1 (e Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.nbPtr.edgePtr.Any = Clone (e)
}

func (x *graph) Put2 (v, v1 Any) {
  if x.vAnchor == x.vAnchor.nextV { return }
  if x.colocal == x.local { return }
  x.colocal.Any = Clone (v)
  x.local.Any = Clone (v1)
}

func (x *graph) ClrMarked() {
  for v := x.vAnchor.nextV; v != x.vAnchor; v = v.nextV {
    v.bool = false
  }
  for e := x.eAnchor.nextE; e != x.eAnchor; e = e.nextE {
    e.bool = false
  }
}

func (x *graph) Mark (v Any) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { panic("bluse"); return }
  x.local.bool = true
}

func (x *graph) Mark1 (v Any) {
  if x.local == x.vAnchor { return }
  if ! x.Ex (v) { return }
  x.local.bool = true
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    n.edgePtr.bool = true
  }
}

func (x *graph) Mark2 (v, v1 Any) {
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

/*
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
*/

func (x *graph) locate (colocal2local bool) {
  x.ClrMarked() // XXX
  if colocal2local {
    x.colocal = x.local
  } else {
    x.local = x.colocal
  }
}

func (x *graph) Locate (colocal2local bool) {
  x.locate (colocal2local)
  x.colocal.bool = true
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

func (x *graph) NeighbourOut (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if n.outgoing {
      if i == 0 {
        return Clone (n.to.Any)
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

func (x *graph) NeighbourIn (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if ! n.outgoing {
      if i == 0 {
        return Clone (n.to.Any)
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

func (x *graph) Neighbour (i uint) Any {
  if x.vAnchor.nextV == x.vAnchor {
    return nil
  }
  for n := x.local.nbPtr.nextNb; n != x.local.nbPtr; n = n.nextNb {
    if i == 0 {
      return Clone (n.to.Any)
    } else {
      i--
    }
  }
  return nil
}

func (x *graph) Star() Graph {
  if x.vAnchor == x.vAnchor.nextV { return nil }
  y := new_(x.bool, x.vAnchor.Any, x.eAnchor.Any).(*graph)
  y.Ins (x.local.Any)
  y.local.bool = true
  local := y.local // Rettung der lokalen Ecke von y
  if ! x.bool {
    for n, i := x.local.nbPtr.nextNb, uint(0);
        n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.Any)) // a jetzt colokal
      y.edgeMarked (Clone(n.edgePtr.Any), true) // Kante
                  // von a zur lokalen eingesetzten Ecke
                  // mit gleichem Inhalt wie in x
      y.local = local // a jetzt wieder lokale Ecke in y
    }
  } else { // x.bool
    for n, i := x.local.nbPtr.nextNb, uint(0);
        n != x.local.nbPtr; n, i = n.nextNb, i + 1 {
      y.Ins (Clone(n.to.Any)) // a ist jetzt wieder colokale
   // Ecke und die eingesetzte Ecke ist die lokale Ecke in y
      if n.outgoing { // wir brauchen eine Kante von a
                      // zu der eingesetzten Ecke
      } else { // ! n.outgoing: brauchen wir eine Kante
               // von der eingesetzten Ecke nach a
        y.local, y.colocal = y.colocal, y.local
      }
      y.edgeMarked (Clone(n.edgePtr.Any), true) // Kante in y
                        // von der colokalen zur lokalen Ecke
      y.local = local
    }
  }
  return y
}

func (x *graph) Add (Ys ...Graph) {
  for _, Y := range Ys {
    y := x.imp(Y)
    for v := y.vAnchor.nextV; v != y.vAnchor; v = v.nextV {
      if x.Ex (v.Any) && ! x.local.bool {
        x.local.bool = v.bool
      } else {
        x.insMarked (v.Any, v.bool)
      }
    }
    for e := y.eAnchor.nextE; e != y.eAnchor; e = e.nextE {
      if x.Ex2 (e.nbPtr0.from.Any, e.nbPtr1.from.Any) {
        e1 := x.connection (x.colocal, x.local)
        e2 := x.connection (x.local, x.colocal)
        if e1 == nil && e2 == nil {
          x.edgeMarked (e.Any, e.bool)
        } else if e1 != nil && ! e1.bool { // x.colocal already connected with x.local
          e1.bool = e.bool
        } else if e2 != nil && ! e2.bool {
          e2.bool = e.bool
        }
      }
    }
  }
}

func (x *graph) SetWrite (w CondOp, w2 CondOp2) {
  x.write, x.write2 = w, w2
}

func (x *graph) Write() {
  x.trav2Cond (x.write2)
  x.travCond (x.write)
}
