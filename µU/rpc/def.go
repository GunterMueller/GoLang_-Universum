package rpc

// (c) Christian Maurer   v. 180813 - license see µU.go

import
  . "µU/obj"

type
  RPC interface {
  F (a Any) Any
  Fin()
}

func New (a, b Any, h string, port uint16, s bool, f Func) RPC {
  return new_(a, b, h, port, s, f)
}
