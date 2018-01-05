package dgra

// (c) Christian Maurer   v. 171209 - license see nU.go

import . "nU/obj"

type TravAlg byte
const (DFS = TravAlg(iota); DFS1; DFSfm1; Awerbuch;
       Awerbuch1; Ring; Ring1; BFS; BFSfm; BFSfm1)

func (x *distributedGraph) SetTravAlgorithm (a TravAlg) {
  x.TravAlg = a
}

func (x *distributedGraph) TravAlgorithm() TravAlg {
  return x.TravAlg
}

func (x *distributedGraph) Trav (o Op) {
  if x.Graph.Directed() { panic("Graph is directed") }
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
  }
}
