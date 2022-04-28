package rpc

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  . "µU/obj"
  "µU/fmon"
)

type
  rpc struct {
             fmon.FarMonitor
             }

func new_(a, b any, n uint, h string, port uint16, s bool, fs FuncSpectrum) RPC {
  return &rpc { fmon.New2 (a, b, n, fs, AllTrueSp, h, port, s) }
}
