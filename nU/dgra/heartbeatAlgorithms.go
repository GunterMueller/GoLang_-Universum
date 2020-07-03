package dgra

type HeartbeatAlg byte
const (HeartbeatMatrix = HeartbeatAlg(iota); HeartbeatGraph; HeartbeatGraph1)

func (x *distributedGraph) SetHeartbeatAlgorithm (a HeartbeatAlg) {
  x.HeartbeatAlg = a
}

func (x *distributedGraph) HeartbeatAlgorithm() HeartbeatAlg {
  return x.HeartbeatAlg
}

func (x *distributedGraph) Heartbeat() {
  if x.Graph.Directed() { panic ("Graph is directed") }
  switch x.HeartbeatAlg {
  case HeartbeatMatrix:
    x.heartbeatmatrix()
  case HeartbeatGraph:
    x.heartbeatgraph()
  case HeartbeatGraph1:
    x.heartbeatgraph1()
  }
}
