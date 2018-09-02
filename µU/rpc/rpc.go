package rpc

// (c) Christian Maurer   v. 180813 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

type
  rpc struct {
             fmon.FarMonitor
             }

func new_(a, b Any, h string, port uint16, s bool, g Func) RPC {
  x := new(rpc)
	f := func (a Any, i uint) Any {
         return g (a)
       }
  x.FarMonitor = fmon.New2 (a, b, 1, f, AllTrueSp, h, port, s, Nothing)
  return x
}

func (x *rpc) F (a Any) Any {
  return x.FarMonitor.F (a, 0)
}

func (x *rpc) Fin() {
  x.FarMonitor.Fin()
}
