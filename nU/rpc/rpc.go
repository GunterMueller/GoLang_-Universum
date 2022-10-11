package rpc

// (c) Christian Maurer   v. 220702 - license see nU.go

import (. "nU/obj"; "nU/fmon")

type
  rpc struct {
             fmon.FarMonitor
             }

func new_(a, b any, n uint, h string, port uint16, s bool, g FuncSpectrum) RPC {
  x := new(rpc)
  f := func (a any, i uint) any { return g (a, n) }
  x.FarMonitor = fmon.New2 (a, b, n, f, AllTrueSp, h, port, s)
  return x
}

func (x *rpc) F (a any, i uint) any {
  return x.FarMonitor.F (a, i)
}

func (x *rpc) Fin() {
  x.FarMonitor.Fin()
}
