package gra

// (c) Christian Maurer   v. 171009 - license see µU.go

import (
  . "µU/obj"
  "µU/adj"
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
// the function Val (they have to be atomic or of type Valuator).
// The outgoing edges of a vertex are enumerated (starting with 0);
// the vertex, with which a vertex is connected by its n-th outgoing edge,
// is denoted as its n-th neighbourvertex.
//
// A subgraph U of a graph G consists of all vertices of G
// and of those edges of G, that connect only vertices of U.
//
// A path in a graph is a sequence of vertices and from each of those
// - excluding from the last one - an outgoing edge to the next vertex.
// A simple path is a path of pairwise disjoint vertices.
// An Euler path is a path that traverses each edge exactly once
// (it may pass any vertex more than once).
// A cycle is a path with an additional edge
// from the last vertex of the path to its first.
// Paths and cycles are subgraphs.
//
// A graph G is (strongly) connected, if for any two vertices
// v, v1 of G there is a path from v to v1 or (and) vice versa;
// so for undirected graphs this is the same.
//
// In any nonempty graph exactly one vertex is distinguished as colocal
// and exactly one as local vertex.
// Each graph has an actual subgraph, an actual path and a vertexstack.

type
  Graph interface {

  Object

// marks the local vertex, if x is not empty.
  Marker

  Persistor

// Returns true, iff x is directed.
  Directed() bool

// Returns the number of vertices of x.
  Num() uint

// Returns the number of vertices in the actual subgraph of x.
  NumSub() uint

// Pre: p is defined on vertices.
// Returns the number of vertices of x, for which p returns true.
  NumPred (p Pred) uint

// Returns the number of edges of x.
  Num1() uint

// Returns the number of edges in the actual subgraph of x.
  NumSub1() uint

// If v is not of the vertextype of x or if v is already contained
// as vertex in x, nothing has happend. Otherwise:
// v is inserted as vertex in x.
// If x was empty, then n is now the colocal and local vertex of x,
// otherwise, v is now the local vertex and the former local vertex
// is now the colocal vertex of x.
  Ins (v Any)

// If x is empty or if the colocal vertex of x coincides with the local vertex or
// if e is not of the edgetype of x, nothing has happened. Otherwise:
// e is inserted in x as edge from the colocal to the local vertex of x
// (if these two vertices were already connected by an edge,
// that edge is replaced by e).
// For e == nil e is replaced by uint(1).
  Edge (e Any)

// If x is empty or has an edgetype or
// if v or v1 is not of the vertextype of x or
// if v or v1 is not contained in x or
// if v and v1 coincide or
// if a is not of the type of the reference edge of x or
// if there is already an edge from v to v1,
// nothing has happened. Otherwise:
// v is now the colocal and v1 the local vertex of x
// and e is inserted is an edge from v to v1.
  Edge2 (v, v1, e Any)

// If x is empty of if the vertices of x do not implement Valuator,
// nothing has happened. Otherwise,
// the set of vertices of x is sorted due to their value.
  SortNeighbours()

// Returns the representation of x as adjacency matrix with
// an arbitrary order of the vertices (their content gets lost).
  Matrix() adj.AdjacencyMatrix

// Pre: uint(len(n)) == a.Num().
// x is the graph with the vertices n[i] (i = 0, ..., len(n) - 1). 
// The edges between the vertices are exactly those, which are defined by m.
  SetMatrix (n []Any, m adj.AdjacencyMatrix)

// Returns true, iff
// the colocal vertex does not coincide with the local vertex of x and
// there is an edge from the colocal to the local vertex in x.
  Edged() bool

// Returns true, iff
// the colocal vertex does not coincide with the local vertex of x
// and there is an edge from the local to the colocal vertex in x.
  CoEdged() bool

// Returns true, iff v is contained as vertex in x.
// In this case v is now the local vertex of x.
// The colocal vertex of x is the same as before.
  Ex (v Any) bool

// Returns true, if v and v1 are contained as vertices in x and do not coincide.
// In this case v now is the colocal and v1 the local vertex of x.
  Ex2 (v, v1 Any) bool

// Pre: p is defined on vertices.
// Returns true, if there is a vertex in x, for which p returns true.
// In this case some such vertex is now the local vertex of x.
// The colocal vertex of x is the same as before.
  ExPred (p Pred) bool

// Returns true, iff e is contained as edge in x.
// In this case the neighbour vertices of some such edge are now
// the colocal and the local vertex of x (if x is directed,
// the vertex, from which the edge goes out, is the colocal vertex.
  Ex1 (e Any) bool

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

// Returns a clone of the local vertex of x.
  Get() Any

// Returns a clone of the reference edge of x, if x is empty or
// if there is no edge from the colocal to the local vertex of x or
// if the colocal vertex of x coincides with the local vertex.
// Returns otherwise a clone of the edge
// from the colocal to the local vertex of x.
  Get1() Any

// Returns clones of the colocal and of the local vertex of x.
  Get2() (Any, Any)

// If x is empty or
// if v is not of the vertextype of x, nothing has happened. Otherwise:
// The local vertex of x is replaced by v.
  Put (v Any)

// If x is empty or if e has no edgetype or
// if e is not of the edgetype of x or
// if there is no edge from the colocal to the local vertex of x,
// nothing has happened. Otherwise:
// The edge from the colocal to the local vertex of x is replaced by e.
  Put1 (e Any)

// If x is empty or
// if v or v1 is not of the vertextype of x or
// if the colocal vertex of x coincides with the local vertex,
// nothing had happened. Otherwise:
// The colocal vertex of x is replaced by v
// and the local vertex of x is replaced by v1.
  Put2 (v, v1 Any)

// The subgraph of x is empty.
  ClrSub()

// The subgraph of x equals x.
  Sub()

// If x is empty, nothing had happened. Otherwise,
// the subgraph of x consists only of the local vertex.
  SubLocal()

// If x is empty, nothing had happened. Otherwise,
// the subgraph of x consists of the local and the colocal vertex
// and the edge between them.
  Sub2()

// Returns true, if x equals its subgraph.
  EqSub() bool

// If x is empty, nothing has happened. Otherwise:
// The former local vertex of x and
// all its outgoing and incoming edges are deleted.
// If x is now not empty, some other vertex is now the local vertex
// and coincides with the colocal vertex of x.
// The actual path and the actual subgraph of x are empty. 
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
// The actual subgraph of x is the actual path of x.
  Act()

// Pre: p is defined on vertices.
// If x is empty, nothing had happened. Otherwise:
// If p returns true for the local vertex and there is a path
// from the colocal to the local vertex of x, that contains
// - apart from the colocal vertex - only vertices, for which p returns true,
// the actual path of x is a shortest such path
// w.r.t. the sum of the values of its edges
// (hence, if x has no edgetype, w.r.t. their number).
// Otherwise the actual path consists only of the colocal vertex.
// The actual subgraph of x is the actual path of x.
  ActPred (p Pred)

// Returns the sum of the values of all edges of x
// (hence, if x has no edgetype, the number of the edges of x).
  Len() uint

// Returns the sum of the values of all edges in the actual subgraph of x
// (hence, if x has no edgetype, the number of the edges in the subgraph).
  LenSub() uint

// Returns 0, if x is empty.
// Returns otherwise the number of the outgoing edges of the local vertex of x.
  NumNeighboursOut() uint

// Returns 0, if x is empty.
// Returns otherwise the number of the incoming edges to the local vertex of x.
  NumNeighboursIn() uint

// Returns 0, if x is empty.
// Returns otherwise the number of all edges to the local vertex of x.
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
// The actual path of x consists only of the colocal vertex of x
// and the actual subgraph of x is the actual path of x.
  Relocate()

// If x is empty, nothing had happened. Otherwise:
// The colocal vertex coincides with the local vertex of x,
// where for f that is the vertex, that was the former local vertex of x,
// and for !f the vertex, that was the former colocal vertex of x.
// The actual path of x consists only of this vertex
// and the actual subgraph of x is the actual path.
  Locate (f bool)

// Returns true, iff x is empty or the local vertex of x
// coincides with the colocal vertex of x.
  Located() bool

// If x is empty, nothing had happened. Otherwise:
// The local and the colocal vertex of x are exchanged;
// actual path and actual subgraph of x are not changed.
  Colocate()

// If x is empty or directed, nothing has happened.
// Otherwise the actual path of x is inverted, particularly
// the local and the colocal vertex of x are exchanged.
// The actual subgraph of x is not changed.
  InvertPath()

// If x is empty or if i >= number of vertices outgoing from the local vertex
// nothing had happened. Otherwise:
// For f:  The i-th neighbour vertex of the last vertex of the actual path
//         of x is appended to it as new last vertex.
// For !f: The last vertex of the actual path of x is deleted from it,
//         if it had not only one vertex (i does not play any role in this case).
// The last vertex of the actual path of x is the local vertex of x and
// the actual subgraph of x is its actual path.
  Step (i uint, f bool)

// Returns false, if x is empty or if i >= NumNeighbourOut();
// returns otherwise true, iff the edge to its i-th neighbour is an outgoing one.
  Outgoing (i uint) bool

// Returns nil, if x is empty or if i >= NumNeighbourOut();
// returns otherwise a copy of its i-th outgoing neighbour.
  NeighbourOut (i uint) Any

// Returns false, if x is empty or if i >= NumNeighboursIn();
// returns otherwise true, iff the edge to its i-th neighbour is an incoming one.
  Incoming (i uint) bool

// Returns nil, if x is empty or if i >= NumNeighboursIn();
// returns otherwise a copy of the its i-th incoming neighbour.
  NeighbourIn (i uint) Any

// Returns nil, if x is empty or if i >= NumNeighbours();
// returns otherwise a copy of its i-th neighbour vertex.
  Neighbour (i uint) Any

// Returns 0, if x is empty or if i >= number of outgoing neighbours of the local vertex;
// returns otherwise the value of the edge to the i-th outgoing neighbour vertex.
// experimental ///////////////////////////////////////////////////  ValOut (i uint) uint

// Returns 0, if x is empty or if i >= number of incoming neighbours of the local vertex;
// returns otherwise the value of the edge to the i-th incoming neighbour vertex.
// experimental ///////////////////////////////////////////////////  ValIn (i uint) uint

// Returns 0, if x is empty or if i >= number of all neighbours of the local vertex;
// returns otherwise the value of the edge to the i-th neighbour vertex.
// experimental ///////////////////////////////////////////////////  Val (i uint) uint

// The local vertex of x is pushed on the vertexstack of x.
  Save()
  CoSave()
  LenStack() uint

// If the vertexstack of x is empty, nothing had happened. Otherwise:
// The local vertex is now the top of the vertexstack
// and this vertex is pulled from the vertexstack of x.
  Restore()
  CoRestore()

// localvertex.dist = 0.
// experimental //////////////////////// A()

// Returns true, iff localvertex.dist = infinite.
// experimental //////////////////////// B() bool

// localvertex.dist = colocalvertex.dist + 1
// localvertex.hinten = colocalvertex.
// experimental //////////////////////// C()

// Pre: p is defined on vertices.
// Returns true, if x is empty or
// if p returns true for all vertices of x.
  True (p Pred) bool

// Pre: p is defined on vertices.
// Returns true, iff x is empty or
// if p returns true for all vertices in the actual subgraph of x.
  TrueSub (p Pred) bool

// Pre: o is defined on vertices.
// o is applied to all vertices of x.
// Colocal and local vertex of x are the same as before;
// subgraph and vertexstack of x are not changed.
  Trav (o Op)

// Pre: o is defined on vertices.
// o is applied to all vertices of x, where
// o is called with 2nd parameter "true", iff
// the corresponding vertex is contained in the actual subgraph of x.
// Colocal and local vertex of x are the same as before;
// subgraph and vertexstack of x are not changed.
  TravCond (o CondOp)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of x.
// Colocal and local vertex of x are the same as before;
// subgraph and vertexstack of x are not changed.
  Trav1 (o Op)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of x with 2nd parameter "true", iff
// the correspoding edge is contained in the actual subgraph of x.
// Colocal and local vertex of x are the same as before;
// subgraph and vertexstack of x are not changed.
  Trav1Cond (o CondOp)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of the local vertex of x.
  Trav1Loc (o Op)

// Pre: o is defined on edges.
// If x has no edgetype, nothing had happened. Otherwise:
// o is applied to all edges of the colocal vertex of x.
  Trav1Coloc (o Op)

// Returns nil, if x does not contain a vertex for which p is true.
// Returns otherwise the Graph consisting with the first such vertex found
// als only vertex and of all edges outgoing from and incoming to this vertex.
// This vertex is the local vertex in Star.
  Star (/* p Pred */) Graph

// Returns true, iff there are no cycles in x.
  Acyclic() bool

// If x is empty, nothing has happened. Otherwise:
// The following equivalence relation is defined on x:
// Two vertices v and v1 of x are equivalent, iff there is
// a path in x from v to v1 and vice versa (hence the set of
// equivalence classes is a directed graph without cycles).
  Isolate() // TODO name

// The actual subgraph of x consists of exactly those vertices, that are
// equivalent to the local vertex and of exactly all edges between them.
// The actual path of x is now empty.
  IsolateSub() // TODO name

// Returns true, iff x is not empty and
// if the local and the colocal vertex of x are equivalent,
// i.e. for both of them there is a path in x to the other one.
  Equiv() bool

// Returns false, if x is not totally connected.
// Returns otherwise true, iff there is an Euler path or cycle in x.
  Euler() bool

// If x is directed, nothing has happened. Otherwise:
// The actual subgraph of x is a minimal spanning tree in
// the connected component, that contains the colocal vertex
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

// Pre: x is directed iff all graphs y are.
// x consists of all vertices and edges of x and all graphs y.
  Add (y ...Graph)

// The demofunction for d is switched on, iff s[d] == true.
  SetDemo (d Demo)

// Pre: v is defined on vertices and e on edges.
// wv and we are the actual write functions for the vertices and edges of x.
  SetWrite (wv, we CondOp)

// Returns the write functions for the vertices and edges of x.
  Writes() (CondOp, CondOp)

// x is written on the screen by means of the actual write functions.
  Write()
}

// Pre: v is atomic or imlements Object.
//      e == nil or e is of type uint or implements Valuator;
// x is Empty (with undefined local and colocal vertex,
// empty actual subgraph, empty actual path and empty vertexstack).
// x is directed, iff d (i.e. otherwise undirected).
// x has v as reference vertex defining the vertextype of x.
// For e == nil, e is replaced by uint(1) and all edges of x have value 1.
// x has e as reference edge defining the edgetype of x.
func New (d bool, v, e Any) Graph { return new_(d,v,e) }
