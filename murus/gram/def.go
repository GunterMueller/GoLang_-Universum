package gram

// (c) murus.org  v. 151116 - license see murus.go

import (
  . "murus/obj"
  "murus/col"
  "murus/gra"
)
type
  GraphModel interface {

  gra.Graph

// The ppm-image with name n is the background.
  Background (n string)

// f, b are the colour for normal nodes/edges.
  Colours (f, b col.Colour)

// f, b are the colour for active nodes/edges.
  ColoursA (f, b col.Colour)

// x is written to the screen.
  Write()

// x is changed interactively by the user.
  Edit()

// The posactual and actual node of x are defined by the user; they are different.
  NodesSelected () bool

// The postactual node of x is defined by the user; it coincides with its actual node.
  NodeSelected () bool

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
