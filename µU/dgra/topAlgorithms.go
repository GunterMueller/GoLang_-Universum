package dgra

// (c) Christian Maurer   v. 171118 - license see µU.go

type
  TopAlg byte; const (
  PassMatrix = TopAlg(iota)
  PassMatrix1
  PassGraph
  PassGraph1
  FmMatrix
  FmGraph
  FmGraph1
)

func (x *distributedGraph) SetTopAlgorithm (a TopAlg) {
  x.TopAlg = a
}

func (x *distributedGraph) TopAlgorithm() TopAlg {
  return x.TopAlg
}

func (x *distributedGraph) Top() {
  if x.Graph.Directed() { panic ("forget it: Graph is directed") }
  switch x.TopAlg {
  case PassMatrix:
    x.passmatrix()
  case PassMatrix1:
    x.passmatrix1()
  case PassGraph:
    x.passgraph()
  case PassGraph1:
    x.passgraph1()
  case FmMatrix:
    x.fmmatrix()
  case FmGraph:
    x.fmgraph()
  case FmGraph1:
    x.fmgraph1()
  }
}
