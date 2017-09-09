package adj

// (c) Christian Maurer   v. 170217 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
)
type
  AdjacencyMatrix interface {
// Sqare matrices with pairs (d, v) as entries,
// where d is of type bool and v of type Valuator or a uint-type.
// The entry of a matrix x in row i and column k is called x(i,k).
// Any such matrix defines a graph (possibly with loops): Its nodes
// are just numbered from 0 to n-1, but do not have any other content.
// x(i,k) == (true, v) means, that in the corresponding graph there
// is an edge outgoing from its i-th and incoming at its k-th node
// with weight v (if x(i,k) = (false, v), v does not play any role).

  Object

// Pre: x is not weighted.
// If i or k >= x.Num(), nothing has happened.
// Otherwise: x(i,k) == (true, v) with v of value 1.
  Edge (i, k uint)

// Pre: x is weighted. v is of the same type as the second
//      parameter in the call of the constructor New of x.
// If i or k >= x.Num(), nothing has happened.
// Otherwise: x(i,k) == (true, v).
//  Set (i, k uint, v uint)
  Set (i, k uint, v Any)

// x(i,k) == (false, v) with undefined v.
  Del (i, k uint)

// Returns the number of rows/columns of x, defined by New.
  Num() uint

// see Editor
  Colours (f, g col.Colour)

// A matrix corresponding to the edges of x is in the colours of x
// written to the screen with left top corner at line, column = l, c.
// Its values in line i, column j are "*", iff there is an edge
// from node i to node j, and "." otherwise.
  Write (l, c uint)

// Returns true, if ! x(i,i) for all i < x.Num().
// i.e. the corresponding graph does not contain loops.
  Ok() bool // Loopfree() ? // Ok() = Loop() == Num()

// Returns the smallest i s.t. x(i,i) != e, if such exists;
// returns x.Num otherwise.
  Loop() uint

// Pre: i, k < x.Num().
// Returns true, iff x(i,k) = (true, v) for some v.
  Edged (i, k uint) bool

// Pre: i, k < x.Num().
// Returns v, iff x(i,k) = (true, v).
  Val (i, k uint) uint

// Returns true, iff x(i,i) == e for all i < x.Num()
// and x(i,k) == x(k,i) for all i, k < x.Num(), i.e. iff
// the corresponding graph does not contain loops
// and is undirected.
  Symmetric() bool

// Returns true, iff x(i,i) == (true, v) for some v for all i < x.Num()
// and there is no pair with x(i,k) == x(k,i), i.e. iff the
// corresponding graph does not contain loops and is directed.
  Directed() bool

// TODO Spec
  Equiv (y AdjacencyMatrix) bool

// TODO Spec
  Add (y AdjacencyMatrix)

// x is mirrored at its diagonal, i.e. x(i,k) equals f(k,i) for the
// former entries f of x, i.e. if the corresponding graph is undirected
// nothing has happened, otherwise, all its edges are inverted.
  Invert()

// Returns true, if all rows of x contain at least one entry with
// value (true, v) for some v, i.e. iff in the corresponding graph
// every node has at least one outgoing edge.
  Full () bool
}

// Pre: n > 0; e is of type Valuator or a uint-type or e == nil. 
// Returns an n*n-matrix with all entries of (false, 0).
func New (n uint, e Any) AdjacencyMatrix { return new_(n,e) }
//          ^^^^ XXX generalize to allow also node.Node !!!
