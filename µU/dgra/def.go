package dgra

// (c) Christian Maurer   v. 171010 - license see µU.go

import (
//  . "µU/obj"
  "µU/gra"
  "µU/host"
)
type
  TopAlg byte; const ( // computation of the net topology
  PassMatrix = TopAlg(iota)
  FmMatrix
  PassGraph
  Graph0
  Graph1
  FmGraph
  FmGraph1
)
type
  ElectAlg byte; const ( // election of a leader in a ring
  ChangRoberts = ElectAlg(iota)
  Peterson
  DolevKlaweRodeh
  HirschbergSinclair
  Maurer
  FmMaurer
  DFSE
  FmDFSE // same with far monitors
)
type
  TravAlg byte; const ( // traversal of the net
  DFS = TravAlg(iota) // depth first seach showing discover and finish times
  DFS1 // depth first seach showing the DFS-tree
  FmDFS1 // depth first search with far monitors without visit phase, showing the DFS-tree
  FmDFSA // simplified DFS-algorithm of Awerbuch with far monitors
  FmDFSA1 // simplified DFS-algorithm of Awerbuch with far monitors, showing the DFS-tree
  FmDFSRing // construction of a ring using DFS, showing the vertices of the ring
  FmDFSRing1 // construction of a ring using DFS, showing the ring
  BFS // BFS-algorithm of Zhu/Cheung
  FmBFS // breadth first search with far monitors
  FmBFS1 // breadth first search with far monitor, showing the BFS-tree
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

  SetElectAlgorithm (a ElectAlg)
  ElectAlgorithm() ElectAlg
  Leader() uint

  SetTopAlgorithm (a TopAlg)
  TopAlgorithm() TopAlg
  Top()

  SetTravAlgorithm (a TravAlg)
  TravAlgorithm() TravAlg
}

// Returns a new distributed Graph with underlying Graph g.
func New (g gra.Graph) DistributedGraph { return new_(g) }

// func Value (a Any) uint { return value(a) }

func Construct() DistributedGraph { return construct() }
