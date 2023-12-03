package dgra

// (c) Christian Maurer   v. 231110 - license see µU.go

type
  HeartbeatAlg byte; const (
  HeartbeatMatrix = HeartbeatAlg(iota)
  HeartbeatMatrix1
  HeartbeatGraph
  HeartbeatGraph1
)

func (x *distributedGraph) SetHeartbeatAlgorithm (a HeartbeatAlg) {
  x.HeartbeatAlg = a
}

func (x *distributedGraph) HeartbeatAlgorithm() HeartbeatAlg {
  return x.HeartbeatAlg
}

func (x *distributedGraph) Heartbeat() {
  if x.Graph.Directed() { panic ("forget it: Graph is directed") }
  switch x.HeartbeatAlg {
  case HeartbeatMatrix:
    x.heartbeatmatrix()
//  case HeartbeatMatrix1:
//    x.heartbeatmatrix1()
  case HeartbeatGraph:
    x.heartbeatgraph()
  case HeartbeatGraph1:
    x.heartbeatgraph1()
  }
}
