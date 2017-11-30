package dgra

// (c) Christian Maurer   v. 171125 - license see nU.go

type
  PulseAlg byte; const (
  PulseMatrix = PulseAlg(iota)
//  PulseMatrix1
  PulseGraph
  PulseGraph1
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
//  case PulseMatrix1:
//    x.pulsematrix1()
  case PulseGraph:
    x.pulsegraph()
  case PulseGraph1:
    x.pulsegraph1()
  }
}
