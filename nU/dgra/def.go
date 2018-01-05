package dgra

// (c) Christian Maurer   v. 171128 - license see nU.go

import "nU/gra"

type DistributedGraph interface {

  gra.Graph

// r is the root of x.
  SetRoot (r uint)

// Pre for all following methods: The hosts of x are set.
  Me() uint
  Root() uint
  Parent() uint
  Children() string
  Time() uint
  Time1() uint

  SetPulseAlgorithm (a PulseAlg) // see pulseAlgorithms.go
  PulseAlgorithm() PulseAlg
  Pulse()

  SetTravAlgorithm (a TravAlg) // see travAlgorithms.go
  TravAlgorithm() TravAlg

  SetElectAlgorithm (a ElectAlg) // see electAlgorithms.go
  ElectAlgorithm() ElectAlg
  Leader() uint
}

// Pre: The values of the edges of g + nchan.Port0 are the ports
//      for 1-1-connections between the vertices connected by them.
// Returns a new distributed Graph with underlying Graph g.
func New (g gra.Graph) DistributedGraph { return new_(g) }

// Examples of distributed Graphs

// G_ returns the star of the distributed Graph defined by g_
// with the vertex with the identity i as center.
func G8 (i uint) DistributedGraph { return g8(i) }
func G8dirring (i uint) DistributedGraph { return g8dr(i) }
func G12 (i uint) DistributedGraph { return g12(i) }
func G12dirring (i uint) DistributedGraph { return g12dr(i) }
