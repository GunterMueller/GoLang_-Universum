package rpc

// (c) Christian Maurer   v. 220702 - license see nU.go

import . "nU/obj"

type
  RPC interface {
  F (a any, i uint) any
  Fin()
}

func New (a, b any, n uint, h string, port uint16, s bool, f FuncSpectrum) RPC {
  return new_(a, b, n, h, port, s, f)
}
