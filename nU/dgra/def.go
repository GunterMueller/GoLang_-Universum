package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

import "nU/gra"

type DistributedGraph interface {

  gra.Graph

// r is the root of x.
  SetRoot (r uint)

// Pre for all following methods: The hosts of x are set.
  Me() uint
  Root() uint
  ParentChildren() string
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
func G4 (i uint) DistributedGraph { return g4(i) }
func G4flat (i uint) DistributedGraph { return g4flat(i) }
func G4full (i uint) DistributedGraph { return g4full(i) }
func G8 (i uint) DistributedGraph { return g8(i) }
func G8dir (i uint) DistributedGraph { return g8dir(i) }
func G8cyc (i uint) DistributedGraph { return g8cyc(i) }
func G8ring (i uint) DistributedGraph { return g8ring(i) }
func G8ringdir (i uint) DistributedGraph { return g8ringdir(i) }
func G8full (i uint) DistributedGraph { return g8full(i) }
func G12 (i uint) DistributedGraph { return g12(i) }
func G12ringdir (i uint) DistributedGraph { return g12ringdir(i) }
func G12full (i uint) DistributedGraph { return g12full(i) }
