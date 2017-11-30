package msg

// (c) Christian Maurer   v. 171119 - license see nU.go

import (. "nU/obj"; "nU/adj")

type Message interface {

  Coder

  Sender() uint
  Kind() byte
  Matrix() adj.AdjacencyMatrix
}

func New (n uint) Message { return new_(n) }
func New2 (n, s uint, k byte) Message { return new2(n,s,k) }
func New3 (n, s uint, k byte, a adj.AdjacencyMatrix) Message { return new3(n,s,k,a) }
