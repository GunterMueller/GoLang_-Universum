package ntrav

// (c) murus.org  v. 170106 - license see murus.go

import (
  . "murus/obj"
  "murus/gra"
  "murus/host"
)
type
  Algorithm byte; const ( // computation methods
  StartEndTimes = Algorithm(iota)
  SpanningTree
  Moni
)

// TODO Spec
func New(g gra.Graph, h []host.Host, n, r uint) NetTraversal { return new_(g,h,n,r) }

type
  NetTraversal interface {

  Demo()

  Trav (a Algorithm, o Op)

  Graph() gra.Graph

  Write()
}
