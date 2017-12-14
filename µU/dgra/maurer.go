package dgra

// (c) Christian Maurer   v. 171203 - license see µU.go
//
// >>> root starts sending the graph consisting only of his own vertex.
//     1st round: Each vertex receives a part of the cycle from its predecessor,
//     appends its own vertex and sends that to its successor.
//     2nd round: Each vertex receives the completed cycle from the predecessor
//     and propagates it to its successor.
//     So finally each vertex has the whole cycle and can easily compute the leader.

// XXX funzt nicht, sondern hört nach zwei auf

import
  . "µU/obj"

func (x *distributedGraph) maurer() {
  x.connect(nil)
  defer x.fin()
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  x.cycle.Clr() // not above !
  if x.me == x.root { // starts sending its singleton graph
    x.cycle.Ins (x.actVertex)
    x.cycle.Write()
    x.ch[out].Send(x.cycle.Encode())
  }
  bs := x.ch[in].Recv().(Stream)
x.log0("recvd")
  x.cycle = x.decodedGraph (bs)
  x.cycle.Write()
  if x.me == x.root {
    x.cycle.Locate (true) // colocal vertex is now the local vertex
    x.cycle.Ex (x.actVertex) // x.actVertex is local
  } else {
    x.cycle.Ins (x.actVertex) // x.actVertex is local, former local vertex is colocal
  }
  x.cycle.Edge (x.directedEdge(x.nb[in], x.actVertex))
  x.cycle.Write()
  x.ch[out].Send (x.cycle.Encode())
// 2nd round:
  x.cycle = x.decodedGraph (x.ch[in].Recv().(Stream))
  x.cycle.Write()
  if x.me != x.root {
    x.ch[out].Send (x.cycle.Encode())
  }
  x.leader = valueMax (x.cycle)
}
