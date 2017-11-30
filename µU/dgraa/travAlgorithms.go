package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

import
  . "µU/obj"
type
  TravAlg byte; const (
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
  case FmDFS1:
    x.fmdfs1 (o)
  case FmDFSA:
    x.fmdfsa (o)
  case FmDFSA1:
    x.fmdfsa1 (o)
  case FmDFSRing:
    x.fmdfsring()
  case FmDFSRing1:
    x.fmdfsring1()
  case BFS:
    x.bfs (o)
  case FmBFS:
    x.fmbfs (o)
  case FmBFS1:
    x.fmbfs1 (o)
  }
}
