package dgra

// (c) Christian Maurer   v. 171112 - license see µU.go

import (
  "µU/gra"
  "µU/host"
)
type
  DistributedGraph interface {

  gra.Graph // At the current stage of development absolutely lunatic.
            // But I am convinced that the structure of the idea is clear,
            // so I hope you do not put me into a loony bin :-)

// Pre: hs must have been globally set to avoid conflicts.
// The hs are the hosts of x.
  SetHosts (hs []host.Host)

// r is the root of x.
  SetRoot (r uint)

  SetDiameter (d uint) // TODO see below

// Returns the diameter of the net.
  Diameter() uint // TODO corresponding graph algorithm

// The rank of the matrices for PassMatrix is set to r.
  SetRank (r uint)

// The demo modus for graphical output is set.
  Demo()

// Pre for all following methods: The hosts of x are set.
  Me() uint
  Root() uint
  ParentChildren() string
  Time() uint
  Time1() uint

  SetElectAlgorithm (a ElectAlg) // see electAlgorithms.go
  ElectAlgorithm() ElectAlg
  Leader() uint

  SetTopAlgorithm (a TopAlg) // see topAlgorithms.go
  TopAlgorithm() TopAlg
  Top()

  SetTravAlgorithm (a TravAlg) // see travAlgorithms.go
  TravAlgorithm() TravAlg
}

// Returns a new distributed Graph with underlying Graph g.
func New (g gra.Graph) DistributedGraph { return new_(g) }

// func Value (a Any) uint { return value(a) }

func Construct() DistributedGraph { return construct() }
