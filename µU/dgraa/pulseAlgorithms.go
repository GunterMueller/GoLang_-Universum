package dgra

// (c) Christian Maurer   v. 171120 - license see µU.go

type
  PulseAlg byte; const (
  PulseMatrix = PulseAlg(iota)
  PulseMatrix1
  PulseGraph
  PulseGraph1
  FmMatrix
  FmGraph
  FmGraph1
)

func (x *distributedGraph) SetPulseAlgorithm (a PulseAlg) {
  x.PulseAlg = a
}

func (x *distributedGraph) PulseAlgorithm() PulseAlg {
  return x.PulseAlg
}

func (x *distributedGraph) Pulse() {
  if x.Graph.Directed() { panic ("forget it: Graph is directed") }
  switch x.PulseAlg {
  case PulseMatrix:
    x.pulsematrix()
  case PulseMatrix1:
    x.pulsematrix1()
  case PulseGraph:
    x.pulsegraph()
  case PulseGraph1:
    x.pulsegraph1()
  case FmMatrix:
    x.fmmatrix()
  case FmGraph:
    x.fmgraph()
  case FmGraph1:
    x.fmgraph1()
  }
}
