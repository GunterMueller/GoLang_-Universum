package dgra

// (c) Christian Maurer   v. 171112 - license see µU.go

type // algorithms to traverse the network
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
