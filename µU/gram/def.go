package gram

// (c) Christian Maurer   v. 241016 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
  "µU/gra"
)
type
  GraphModel interface {

  gra.Graph

// The ppm-image with name n is the background.
  Background (n string)

// f, b are the colour for normal vertices/edges,
// fm, bm are the colour for marked vertices/edges.
  Colours (f, b, fm, bm col.Colour)

// x is changed interactively by the user.
  Edit()

// The posactual and actual vertex of x are defined by the user; they are different.
  VerticesSelected () bool

// The postactual vertex of x is defined by the user; it coincides with its actual vertex.
  VertexSelected () bool

// The number of Ways for the execution of Hamilton == 0.
  ResetNWays ()

// Returns the number of Ways in the execution of Hamilton.
  NWays () uint

// TODO Spec
  DFS (a bool)

// TODO Spec
  BFS (a bool)

// TODO Spec
  Hamilton (r, o Cond, w *bool)

  Demo (gra.Demo)
}

// Returns a new empty GraphModel for vertices and edges
// of type n and e resp., that is directed, iff d == true.
// (See also Spec of Graph in µU/gra.)
func New (d bool, n, e any) GraphModel { return new_(d,n,e) }
