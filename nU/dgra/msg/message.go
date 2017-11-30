package msg

// (c) Christian Maurer   v. 171127 - license see nU.go

import (. "nU/obj"; "nU/adj")

type message struct {
  uint "sender"
  byte "kind"
  adj.AdjacencyMatrix
}

func new_(n uint) Message {
  x := new(message)
  x.AdjacencyMatrix = adj.New (n, uint(0), uint(0))
  return x
}

func new2 (n, s uint, k byte) Message {
  x := new_(n).(*message)
  x.uint = s
  x.byte = k
  return x
}

func new3 (n, s uint, k byte, a adj.AdjacencyMatrix) Message {
  x := new_(n).(*message)
  x.uint = s
  x.byte = k
  x.AdjacencyMatrix = adj.New (n, uint(0), uint(0))
  return x
}

func (x *message) Codelen() uint {
  return C0 + 1 + x.AdjacencyMatrix.Codelen()
}

func (x *message) Encode() Stream {
  bs := make([]byte, x.Codelen())
  copy (bs[:C0], Encode(x.uint))
  bs[C0] = x.byte
  copy (bs[C0+1:], x.AdjacencyMatrix.Encode())
  return bs
}

func (x *message) Decode (bs Stream) {
  x.uint = Decode(uint(0), bs[:C0]).(uint)
  x.byte = bs[C0]
  x.AdjacencyMatrix.Decode (bs[C0+1:])
}

func (x *message) Sender() uint {
  return x.uint
}

func (x *message) Kind() byte {
  return x.byte
}

func (x *message) Matrix() adj.AdjacencyMatrix {
  return x.AdjacencyMatrix.Clone().(adj.AdjacencyMatrix)
}
