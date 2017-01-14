package nelect

// (c) murus.org  v. 170111 - license see murus.go

import (
  "murus/gra"
  "murus/host"
)
type
  Algorithm byte; const (
  Peterson = Algorithm(iota)
  PetersonImproved
  ChangRoberts
  HirschbergSinclair
  Maurer
)

func New (g gra.Graph, h []host.Host, n, r uint) NetElection { return new_(g,h,n,r) }

type
  NetElection interface {

  Demo()

  Do (a Algorithm)

  Graph() gra.Graph

  Write()
}
