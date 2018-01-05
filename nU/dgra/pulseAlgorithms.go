package dgra

// (c) Christian Maurer   v. 171227 - license see nU.go

type PulseAlg byte
const (PulseMatrix = PulseAlg(iota); PulseGraph; PulseGraph1)

func (x *distributedGraph) SetPulseAlgorithm (a PulseAlg) {
  x.PulseAlg = a
}

func (x *distributedGraph) PulseAlgorithm() PulseAlg {
  return x.PulseAlg
}

func (x *distributedGraph) Pulse() {
  if x.Graph.Directed() { panic ("Graph is directed") }
  switch x.PulseAlg {
  case PulseMatrix:
    x.pulsematrix()
  case PulseGraph:
    x.pulsegraph()
  case PulseGraph1:
    x.pulsegraph1()
  }
}
