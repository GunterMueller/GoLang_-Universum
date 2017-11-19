package adj

// (c) Christian Maurer   v. 171111 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
type
  AdjacencyMatrix interface {
// Sqare matrices with pairs (v, e) as entries,
// where v is atomic or implements Object and
// and e has a uint-type or implements Valuator.
// The entry of a matrix x in row i and column k is called x(i,k).
// Any such matrix defines a graph in the following way:
// x(i,k) = (e, v) means
// for i == k:
//     v is a vertex in the graph (in this case e is the pattern edge of x.
// for i != k:
//     There is an edge outgoing from the i-th and incoming at the k-th vertex
//     of the graph with value e, iff v is not equal to the pattern vertex of x.
//     In this case v is the pattern vertex of x.
// The patterns are those objects, that are given to a New as parameters;
// they must not be used as vertices or edges resp.

  Object
  col.Colourer

// Returns the number of rows/columns of x, defined by New.
  Num() uint

// Returns true, iff x and y have the same number of rows/columns
// and equal vertex patterns and equal edge patterns.
  Equiv (y AdjacencyMatrix) bool

// Pre: e is of the type of the pattern edge of x.
// If i or k >= x.Num(), nothing has happened.
// Otherwise: x(i,k) is the pair (v, e) with v = pattern vertex of x,
// i.e. in the corresponding graph there is an edge with the value of e
// from its i-th vertex to its k-th vertex, iff e is not equal to the pattern edge of x.
  Edge (i, k uint, e Any)

// Returns the first element in the pair x(i,i), i.e. a vertex.
  Vertex (i uint) Any

// Pre: i, k < x.Num().
// Returns 0, iff for x(i,k) = (v, e) e is the pattern edge of x,
// returns otherwise the value of e.
  Val (i, k uint) uint

// Pre: v has the type as the pattern vertex of x and
//      e has the type of the pattern edge of x.
// If i or k >= x.Num(), nothing has happened.
// Otherwise: x(i,k) == (v, e).
  Set (i, k uint, v, e Any)

// Returns true, iff x(i,k) == x(k,i) for all i, k < x.Num(),
// i.e. the corresponding graph is undirected.
  Symmetric() bool

// Pre: x and y are equivalent.
// x contains all entries of x and additionally all entries of y
// with a value > 0 (the values of entries of x are overwritten
// by the values of corresponding entries in x. // XXX ? ? ?
  Add (y AdjacencyMatrix)

// Returns true, iff each row of x contains at least one entry
// with (v, e) with Val(e) > 0, i.e. iff in the corresponding
// graph every node has at least one outgoing edge.
  Full () bool

// A matrix corresponding to the edges of x is in the colours of x
// written to the screen with left top corner at line, column = l, c.
// Its values in line i, column j are 0, if for x(i,k) = (v, e)
// e is equal to the pattern edge of x, otherwise the value of e.
  Write (l, c uint)
}

// Pre: n > 0; e is of a uint-type or implements Valuator;
//      v is atomic or implements Object.
// v is the pattern vertex of x
// Returns an n*n-matrix with the pattern vertex v
// and the pattern edge e. All it's entries have the value
// (v, e).
func New (n uint, v, e Any) AdjacencyMatrix { return new_(n,v,e) }
