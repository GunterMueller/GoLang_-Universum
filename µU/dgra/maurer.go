package dgra

// (c) Christian Maurer   v. 170423 - license see µU.go
//
// >>> root starts sending the graph consisting only of his own vertex.
//     1st round: Each vertex receives a part of the ring from its predecessor,
//     appends its own vertex and sends that to its successor.
//     2nd round: Each vertex receives the completed ring from the predecessor
//     and propagates it to its successor.
//     So finally each vertex has the whole ring and can easily compute the leader.

func (x *distributedGraph) maurer() {
  x.connect(nil)
  defer x.fin()
  out, in := uint(0), uint(1); if x.Graph.Outgoing(1) { in, out = out, in }
  x.ring.Clr() // not above !
  if x.me == x.root { // starts sending its singleton graph
    x.ring.Ins (x.actVertex)
    x.ring.Write()
    x.ch[out].Send(x.ring.Encode())
  }
  bs := x.ch[in].Recv().([]byte)
  x.ring = x.decodedGraph(bs)
  x.ring.Write()
  if x.me == x.root {
    x.ring.Locate (true) // colocal vertex is now the local vertex
    x.ring.Ex (x.actVertex) // x.meVertex is local
  } else {
    x.ring.Ins (x.actVertex) // x.meVertex is local, former local vertex is colocal
  }
  x.ring.Edge (x.directedEdge(x.nb[in], x.actVertex))
  x.ring.Write()
  x.ch[out].Send (x.ring.Encode())
// 2nd round:
  bs = x.ch[in].Recv().([]byte)
  x.ring = x.decodedGraph(bs)
  x.ring.Write()
  if x.me != x.root {
    x.ch[out].Send (x.ring.Encode())
  }
  x.leader = valueMax (x.ring)
}
