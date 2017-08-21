package ntop

// (c) murus.org  v. 161230 - license see murus.go

import (
//  . "murus/obj"
  "murus/gra"
  "murus/host"
)
type
  Algorithm byte; const (
  Matrix = Algorithm(iota) // adjacency matrix
  Graph                    // graph
  MatrixFarMon             // adj. matrix with far monitor
//  GraphFarMon              // graph with far monitor // TODO
)

// TODO Spec
func New (g gra.Graph, h []host.Host, n, d uint) NetTopology { return new_(g,h,n,d) }

type
  NetTopology interface {

  Demo()

  Do (a Algorithm)

  Topology() gra.Graph

// Pre: Do was called (since Write depends on the Algorithm).
  Write()
}
