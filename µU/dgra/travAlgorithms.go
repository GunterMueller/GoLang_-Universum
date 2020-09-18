package dgra

// (c) Christian Maurer   v. 171203 - license see µU.go

import
  . "µU/obj"
type
  TravAlg byte; const (
  DFS = TravAlg(iota) // depth first seach showing discover and finish times
  DFS1 // depth first seach showing the DFS-tree
  DFSfm1 // depth first search without visit phase showing the DFS-tree, with far monitors
  Awerbuch // simplified DFS-algorithm of Awerbuch, with far monitors
  Awerbuch1 // simplified DFS-algorithm of Awerbuch showing the DFS-tree, with far monitors
  Awerbuch2 // experimental
  HelaryRaynal // experimental
  Ring // construction of a ring using DFS showing the vertices of the ring, with far monitors
  Ring1 // construction of a ring using DFS showing the ring, with far monitors
  BFS // BFS-algorithm of Zhu/Cheung
  BFSfm // breadth first search, with far monitors
  BFSfm1 // breadth first search showing the BFS-tree, with far monitors
  MA1 // experimental
)

func (x *distributedGraph) SetTravAlgorithm (a TravAlg) {
  x.TravAlg = a
}

func (x *distributedGraph) TravAlgorithm() TravAlg {
  return x.TravAlg
}

func (x *distributedGraph) Trav (o Op) {
  if x.Directed() { panic("forget it: Graph is directed") }
  switch x.TravAlg {
  case DFS:
    x.dfs (o)
  case DFS1:
    x.dfs1 (o)
  case DFSfm1:
    x.dfsfm1 (o)
  case Awerbuch:
    x.awerbuch (o)
  case Awerbuch1:
    x.awerbuch1 (o)
  case Awerbuch2:
    x.awerbuch2 (o)
  case HelaryRaynal:
    x.helaryRaynal (o)
  case Ring:
    x.ring()
  case Ring1:
    x.ring1()
  case BFS:
    x.bfs (o)
  case BFSfm:
    x.bfsfm (o)
  case BFSfm1:
    x.bfsfm1 (o)
  case MA1:
    x.ma1 (o)
  }
}
