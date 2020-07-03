package rpc

// (c) Christian Maurer   v. 200213 - license see nU.go

import . "nU/obj"

type
  RPC interface {
  F (a Any, i uint) Any
  Fin()
}

func New (a, b Any, n uint, h string, port uint16, s bool, f FuncSpectrum) RPC {
  return new_(a, b, n, h, port, s, f)
}
