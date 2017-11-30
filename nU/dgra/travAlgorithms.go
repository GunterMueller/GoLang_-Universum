package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

import . "nU/obj"

type TravAlg byte; const (DFS = TravAlg(iota); DFS1; FmDFSA; FmDFSRing; BFS; FmBFS)

func (x *distributedGraph) SetTravAlgorithm (a TravAlg) {
  x.TravAlg = a
}

func (x *distributedGraph) TravAlgorithm() TravAlg {
  return x.TravAlg
}

func (x *distributedGraph) Trav (o Op) {
  if x.Graph.Directed() { panic("forget it: Graph is directed") }
  switch x.TravAlg {
  case DFS:
    x.dfs (o)
  case DFS1:
    x.dfs1 (o)
  case FmDFSA:
    x.fmdfsa (o)
  case FmDFSRing:
    x.fmdfsring()
  case BFS:
    x.bfs (o)
  case FmBFS:
    x.fmbfs (o)
  }
}
