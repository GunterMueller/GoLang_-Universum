package dgra

// (c) Christian Maurer   v. 171112 - license see µU.go

type // algorithms to compute the net topology
  TopAlg byte; const (
  PassMatrix = TopAlg(iota)
  FmMatrix
  PassGraph
  Graph0
  Graph1
  FmGraph
  FmGraph1
)
