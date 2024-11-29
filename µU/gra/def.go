package gra

// (c) Christian Maurer   v. 241030 - license see µU.go

import (
  . "µU/obj"
  "µU/adj"
  "µU/pseq"
  "µU/vtx"
)
type
  Demo byte // for demonstration purposes
const (
  Depth = Demo(iota)
  Cycle
  Euler
  TopSort
  ConnComp
  Breadth
  SpanTree
  nDemos
)
type
  Demoset [nDemos]bool

// Sets of vertices with an irreflexive relation:
// Two vertices are related, iff they are connected by an edge, where there
// are no loops (i.e. no vertex is connected with itself by an edge).
// If the relation is symmetric, the graph is called "undirected",
// if it is strict, "directed" (i.e. all edges have a direction).
//
// The edges have a number of type uint as value ("weight");
// either all edges have the value 1 or their value is given by
// the function Val (they have to be of an uint-type or of type Valuator).
// The outgoing edges of a vertex are enumerated (starting with 0);
// the vertex, with which a vertex is connected by its n-th outgoing edge,
// is denoted as its n-th neighbourvertex.
//
// In any graph some vertices and edges might be marked.
//
// A path in a graph is a sequence of vertices and from each of those
// - excluding from the last one - an outgoing edge to the next vertex.
// A simple path is a path of pairwise disjoint vertices.
// An Euler path is a path that traverses each edge exactly once
// (it may pass any vertex more than once).
// A cycle is a path with an additional edge
// from the last vertex of the path to its first.
//
// A graph G is (strongly) connected, if for any two vertices
// v, v1 of G there is a path from v to v1 or (and) vice versa;
// so for undirected graphs this is the same.
//
// In any nonempty graph exactly one vertex is distinguished as colocal
// and exactly one as local vertex.
// Each graph has an actual path.

type
  Graph interface {

  Object

  Persistor

// Returns true, iff x is directed.
  Directed() bool

  SetDir (b bool)

// Returns a copy of the graph that is not directed.
  Indir() Graph

// Returns the number of vertices of x.
  Num() uint

// Returns the number of edges of x.
  Num1() uint

// Returns the number of marked vertices of x.
  NumMarked() uint

// Returns the number of marked edges of x.
  NumMarked1() uint

// Pre: p is defined on vertices.
// Returns the number of vertices of x, for which p returns true.
  NumPred (p Pred) uint

// If v is not of the vertextype of x or if v is already contained
// as vertex in x, nothing has happend. Otherwise:
// v is inserted as vertex in x.
// If x was empty, then v is now the colocal and local vertex of x,
// otherwise, v is now the local vertex and the former local vertex
// is now the colocal vertex of x.
  Ins (v any)

// If x was empty or if the colocal vertex of x coincides
// with the local vertex of x or if e is not of the edgetype of x,
// nothing has happened. Otherwise:
// e is inserted into x as edge from the colocal to the local vertex of x
// (if these two vertices were already connected by an edge,
// that edge is replaced by e).
// For e == nil e is replaced by uint(1).
  Edge (e any)

// If x is empty or has an edgetype or
// if v or v1 is not of the vertextype of x or
// if v or v1 is not contained in x or
// if v and v1 coincide or
// if a is not of the type of the pattern edge of x or
// if there is already an edge from v to v1,
// nothing has happened. Otherwise:
// v is now the colocal and v1 the local vertex of x
// and e is inserted is an edge from v to v1.
  Edge2 (v, v1, e any)

// Returns the representation of x as adjacency matrix.
  Matrix() adj.AdjacencyMatrix

// Pre: m is symmetric iff x is directed.
// x is the graph with the vertices a.Vertex(i) and edges from
// a.Vertex(i) to a.Vertex(k), iff a.Val(i,k) > 0 (i, k < a.Num()).
  SetMatrix (a adj.AdjacencyMatrix)

// Returns true, iff the colocal vertex of x does not
// coincide with the local vertex of x and there is
// an edge in x from the colocal to the local vertex.
  Edged() bool

// Returns true, iff
// the colocal vertex does not coincide with the local vertex of x
// and there is an edge from the local to the colocal vertex in x.
  CoEdged() bool

// Returns true, iff v is contained as vertex in x.
// In this case, v is now the local vertex of x.
// The colocal vertex of x is the same as before.
  Ex (v any) bool

// Returns true, if v and v1 are contained as vertices in x
// and do not coincide. In this case now
// v is the colocal and v1 the local vertex of x.
  Ex2 (v, v1 any) bool

// Pre: p is defined on vertices.
// Returns true, iff there is a vertex in x, for which p returns true.
// In this case some now such vertex is the local vertex of x.
// The colocal vertex of x is the same as before.
  ExPred (p Pred) bool

/// Returns true, iff e is contained as edge in x.
// In this case the neighbour vertices of some such edge are now
// the colocal and the local vertex of x (if x is directed,
// the vertex, from which the edge goes out, is the colocal vertex.
  Ex1 (e any) bool

// Pre: p is defined on edges.
// Returns true, iff there is an edge in x, for which p returns true.
// In this case the neighbour vertices of some such edge are now
// the colocal and the local vertex of x (if x is directed,
// the vertex, from which the edge goes out, is the colocal vertex.
  ExPred1 (p Pred) bool

// Pre: p and p1 are defined on vertices.
// Returns true,
// iff there are two different vertices v and v1 with p(v) and p(v1).
// In this case now some vertex v with p(v) is the colocal vertex
// and some vertex v1 with p1(v1) is the local vertex of x.
  ExPred2 (p, p1 Pred) bool

// Returns the value of the local vertex of x,
// if it has the type Valuator; return otherwise 1.
  Val() uint

// Returns true, iff x contains a vertex with the value n.
// In this case, such a vertex is the local vertex of x.
// The colocal vertex of x is the same as before.
  ExVal (n uint) bool

// Pre: x contains a vertex with value n.
// Returns that vertex.
  Vertex (n uint) vtx.Vertex

// Returns true, iff x contains a vertex v with the value n
// and a vertex v1 with the value n1. In this case,
// v is the colocal vertex of x and v1 is the local vertex of x.
  ExVal2 (n, n1 uint) bool

// Returns the pattern vertex of x, if x is empty;
// returns otherwise a clone of the local vertex of x.
  Get() any

// Returns a clone of the pattern edge of x, if x is empty
// or if there is no edge from the colocal vertex to the
// local vertex of x or if these two vertices coincide.
// Returns otherwise a clone of the edge from the
// colocal vertex of x to the local vertex of x.
  Get1() any

// Returns (nil, nil), if x is empty.
// Returns otherwise a pair, consisting of clones
// of the colocal and of the local vertex of x.
  Get2() (any, any)

// If x is empty or if v is not of the vertex type of x, nothing has happened. Otherwise:
// The local vertex of x is replaced by v.
  Put (v any)

// If x is empty or if e has no edge type or
// if e is not of the edgetype of x or
// if there is no edge from the colocal to the local vertex of x,
// nothing has happened. Otherwise:
// The edge from the colocal to the local vertex of x is replaced by e.
  Put1 (e any)

// If x is empty or if v or v1 is not of the vertextype of x or
// if the colocal vertex of x coincides with the local vertex,
// nothing had happened. Otherwise:
// The colocal vertex of x is replaced by v
// and the local vertex of x is replaced by v1.
  Put2 (v, v1 any)

// No vertex and no edge in x is marked.
  ClrMarked()

// If x is empty or if v is not of the vertex type of x
// or if v is not contained in x, nothing has happened.
// Otherwise, v is now the local vertex of x and is marked iff m.
// The colocal vertex of x is the same as before.
  Mark (v any, m bool)

//  Mark1 (e any, m bool)

// If x is empty or if v or v1 is not of the vertex type of x
// or if v or v1 is not contained in x
// or if v and v1 conincide, nothing had happened.
// Otherwise, v is now the colocal and v1 the local vertex
// of x and these two vertices and the edge between them are now marked.
  Mark2 (v, v1 any)

// Returns true, if all vertices and all edges of x are marked.
  AllMarked() bool

// If x is empty, nothing has happened. Otherwise:
// The former local vertex of x and
// all its outgoing and incoming edges are deleted.
// If x is now not empty, some other vertex is now the local vertex
// and coincides with the colocal vertex of x.
// The actual path is empty. 
  Del()

// If there was an edge between the colocal and the local vertex of x,
// it is now deleted from x.
  Del1()

// Returns true, iff x is empty or
// if the colocal vertex coincides with the local vertex of x or
// if there is a path from the colocal to the local vertex in x.
  Conn() bool

// Pre: p is defined on vertices.
// Returns true, iff x is empty or
// the colocal vertex coincides with the local vertex of x or
// if p returns true for the local vertex and there is a path
// from the colocal vertex of x to the local vertex, that contains
// - apart from the colocal vertex - only vertices, for which p returns true.
  ConnCond (p Pred) bool

// If x is empty, nothing had happened. Otherwise:
// If there is a path from the colocal to the local vertex of x,
// the actual path of x is a shortest such path
// (shortest w.r.t. the sum of the values of its edges,
// hence, if x has no edgetype, w.r.t. their number).
// If there is no path from the colocal to the local vertex of x,
// the actual path consists only of the colocal vertex.
// The marked vertices and edges of x are
// the vertices and edges in the actual path of x.
  FindShortestPath()

// Pre: p is defined on vertices.
// If x is empty, nothing had happened. Otherwise:
// If p returns true for the local vertex and there is a path
// from the colocal to the local vertex of x, that contains
// - apart from the colocal vertex - only vertices, for which p returns true,
// the actual path of x is a shortest such path
// w.r.t. the sum of the values of its edges
// (hence, if x has no edgetype, w.r.t. their number).
// Otherwise the actual path consists only of the colocal vertex.
// The marked vertices and edges of x are
// the vertices and edges in the actual path of x.
  FindShortestPathPred (p Pred)

// Pre: Act or ActPred was called before.
// Returns the slice of the vertices of the actual path.
  ShortestPath() []any

// Returns the sum of the values of all edges of x
// (hence, if x has no edgetype, the number of the edges of x).
  Len() uint

// Returns the sum of the values of all marked edges in x
// (hence, if x has no edgetype, the number of the marked edges).
  LenMarked() uint

// Returns 0, if x is empty.
// Returns otherwise the number of the outgoing edges of the local vertex of x.
  NumNeighboursOut() uint

// Pre: x is directed.
// Returns 0, if x is empty.
// Returns otherwise the number of the incoming edges to the local vertex of x.
  NumNeighboursIn() uint

// Returns 0, if x is empty.
// Returns otherwise the number of all edges of the local vertex of x.
  NumNeighbours() uint

// If x is not directed, nothing had happened. Otherwise:
// The directions of all edges of x are reversed.
  Inv()

// If x is not directed, nothing had happened. Otherwise:
// The directions of all outgoing and incoming edges
// of the local vertex of x are reversed.
  InvLoc()

// If x is empty, nothing had happened. Otherwise:
// The local and the colocal vertex of x are exchanged.
// The actual path of x consists only of the colocal vertex of x.
// The only marked is the colocal vertex; no edges are marked.
  Relocate()

// If x is empty, nothing had happened. Otherwise:
// The colocal vertex of x coincides with the local vertex of x,
// where for f == true that is the vertex, that was the former local vertex of x,
// and for !f the vertex, that was the former colocal vertex of x.
// The actual path of x consists only of this vertex.
// The only marked vertex is this vertex; no edges are marked.
  Locate (f bool)

// Returns true, iff x is empty or the local vertex of x
// coincides with the colocal vertex of x.
  Located() bool

// If x is empty, nothing had happened. Otherwise:
// The local and the colocal vertex of x are exchanged;
// the actual path is not changed and
// the marked vertices and edges are unaffected.
  Colocate()

// If x is empty or directed, nothing has happened.
// Otherwise the actual path of x is inverted, particularly
// the local and the colocal vertex of x are exchanged.
// The marked vertices and edges are unaffected.
  InvertPath()

// If x is empty or if i >= number of vertices outgoing from the local vertex
// nothing had happened. Otherwise:
// For f:  The i-th neighbour vertex of the last vertex of the actual path
//         of x is appended to it as new last vertex.
// For !f: The last vertex of the actual path of x is deleted from it,
//         if it had not only one vertex (i does not play any role in this case).
// The last vertex of the actual path of x is the local vertex of x and
// Vertices and edges in x are marked, if the belong to its actual path.
  Step (i uint, f bool)

// Returns false, if x is empty or if i >= NumNeighbours();
// returns otherwise true, iff the edge to the i-th neighbour
// of the local vertex is an outgoing edge.
  Outgoing (i uint) bool

// Returns nil, if x is empty or if i >= NumNeighboursOut();
// returns otherwise a clone of the i-th outgoing neighbour of the local vertex.
  NeighbourOut (i uint) any

// Returns false, if x is empty or if i >= NumNeighbours();
// returns otherwise true, iff the edge to the i-th neighbour
// of the local vertex is an incoming one.
  Incoming (i uint) bool

// Returns nil, if x is empty or if i >= NumNeighboursIn();
// returns otherwise a copy of the its i-th incoming neighbour of the local vertex.
  NeighbourIn (i uint) any

// Returns nil, if x is empty or if i >= NumNeighbours();
// returns otherwise a clone of its i-th neighbour vertex
// of the local vertex of x.
  Neighbour (i uint) any

// Returns the diameter of x.
  Diameter() uint

// d is the diameter of x.
  SetDiameter (d uint)

// Pre: p is defined on vertices.
// Returns true, if x is empty or
// if p returns true for all vertices of x.
  True (p Pred) bool

// Pre: p is defined on vertices.
// Returns true, iff x is empty or
// if p returns true for all marked vertices in x.
  TrueMarked (p Pred) bool

// Pre: o is defined on vertices.
// o is applied to all vertices of x.
// The colocal and the local vertex of x are the same as before;
// the marked vertices and edges are unaffected.
  Trav (o Op)

// Pre: o is defined on vertices.
// o is applied to all vertices of x, where
// o is called with 2nd parameter "true", iff
// the corresponding vertex is marked.
// Colocal and local vertex of x are the same as before;
// The marked edgges are unaffected.
  TravCond (o CondOp)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of x.
// Colocal and local vertex of x are the same as before;
// the marked vertices and edges are unaffected.
  Trav1 (o Op)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of x with 2nd parameter "true",
// iff the correspoding edge is marked.
// Colocal and local vertex of x are the same as before;
// the marked vertices and edges are unaffected.
  Trav1Cond (o CondOp)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of the local vertex of x.
  Trav1Loc (o Op)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of the colocal vertex of x.
  Trav1Coloc (o Op)

// Returns nil, if x is empty.
// Returns otherwise the graph consisting of the local
// vertex of x, all its neighbour vertices and of all edges
// outgoing from it and incoming to it.
// The local vertex of x is the local vertex of the star.
// It is the only marked vertex in the star;
// all edges in the star are marked.
  Star() Graph

// Returns true, iff there are no cycles in x.
  Acyclic() bool

// If x is empty, nothing has happened. Otherwise:
// The following equivalence relation is defined on x:
// Two vertices v and v1 of x are equivalent, iff there is
// a path in x from v to v1 and vice versa (hence the set of
// equivalence classes is a directed graph without cycles).
  Isolate() // TODO name

// Exactly those vertices in x are marked, that are equivalent
// to the local vertex and of exactly all edges between them.
// No edges in x are marked.
  IsolateMarked() // TODO name

// Returns true, iff x is not empty and
// if the local and the colocal vertex of x are equivalent,
// i.e. for both of them there is a path in x to the other one.
  Equiv() bool

// Returns false, if x is not totally connected.
// Returns otherwise true, iff there is an Euler path or cycle in x.
  Euler() bool

// If x is directed, nothing has happened. Otherwise:
// Exactly those vertices and edges in x are marked,
// that build a minimal spanning tree in the connected component
// containing the colocal vertex
// (minimal w.r.t. the values of the sum of its edges;
// hence, if x has no edgetype, w.r.t. the number of its vertices)
// The actual path is not changed.
  MST()

// If x is empty or undirected or
// if x is directed and has cycles, nothing has happened. Otherwise:
// The vertices of x are ordered s.t. at each subsequent traversal of x
// each vertex with outgoing edges is always handled before the vertices,
// at which those edges come in.
  Sort()

// Pre: x is directed, iff all graphs y are directed.
// x consists of all vertices and edges of x before
// and of all graphs y. Thereby all marks of y are overtaken.
  Add (y ...Graph)

// The demofunction for d is switched on, iff s[d] == true.
  SetDemo (d Demo)

// Pre: wv is defined on vertices and we on edges.
// wv and we are the actual write functions for the vertices and edges of x.
  SetWrite (wv, we Op)

// Returns the write functions for the vertices and edges of x.
  Writes() (Op, Op)

// x is written on the screen by means of the actual write functions.
  Write()

// Pre: x.Name was called.
// Returns the corresponding file.
  File() pseq.PersistentSequence

// Pre: x.Name was called.
// x is loaded from the corresponding file.
  Load()

// Pre: x.Name was called.
// x is stored in the corresponding file.
  Store()

// Pre: x is connected.
// Returns true, iff x is a ring.
// If x is not connected, true is returned, iff every connected component is a ring.
  IsRing() bool
}

// Pre: v is atomic or imlements Object.
//      e == nil or e is of type uint or implements Valuator;
// Returns an empty graph, 
// x is directed, iff d (i.e. otherwise undirected).
// v is the pattern vertex of x defining the vertex type of x.
// For e == nil, e is replaced by uint(1) and all edges of x have the value 1.
// Otherwise e is the pattern edge of x defining the edgetype of x.
func New (d bool, v, e any) Graph { return new_(d,v,e) }

// Examples of distributed Graphs
// G_ returns the Graph defined by g_.
func G3() Graph { return g3() }
func G3dir() Graph { return g3dir() }
func G4() Graph { return g4() }
func G4flat() Graph { return g4flat() }
func G4ring() Graph { return g4ring() }
func G4ringdir() Graph { return g4ringdir() }
func G4full() Graph { return g4full() }
func G4star() Graph { return g4star() }
func G4ds() Graph { return g4ds() }
func G5() Graph { return g5() }
func G5ring() Graph { return g5ring() }
func G5ringdir() Graph { return g5ringdir() }
func G5full() Graph { return g5full() }
func G6() Graph { return g6() }
func G6full() Graph { return g6full() }
func G8() Graph { return g8() }
func G8a() Graph { return g8a() }
func G8dir() Graph { return g8dir() }
func G8ring() Graph { return g8ring() }
func G8ringdir() Graph { return g8ringdir() }
func G8ringdirord() Graph { return g8ringdirord() }
func G8full() Graph { return g8full() }
func G8ds() Graph { return g8ds() }
func G9a() Graph { return g9a() }
func G9b() Graph { return g9b() }
func G9dir() Graph { return g9dir() }
func G10() Graph { return g10() }
func G12() Graph { return g12() }
func G12ringdir() Graph { return g12ringdir() }
func G12full() Graph { return g12full() }
func G16() Graph { return g16() }
func G16dir() Graph { return g16dir() }
func G16ring() Graph { return g16ring() }
func G16ringdir() Graph { return g16ringdir() }
func G16full() Graph { return g16full() }
