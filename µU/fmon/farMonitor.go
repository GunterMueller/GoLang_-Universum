package fmon

// (c) Christian Maurer   v. 180101 - license see µU.go

import (
  . "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/nchan"
)
type
  farMonitor struct {
                    Any "pattern object"
                    uint "number of monitor functions"
                 ch []nchan.NetChannel
                    FuncSpectrum; PredSpectrum
                    bool "true iff the monitor is a server"
                    }

func new_(a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
  return news (a, n, fs, ps, h, p, s, Nothing)
}

func news (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h string, p uint16, s bool, stmt Stmt) FarMonitor {
  if n == 0 { Panic ("fmon.New must be called with 2nd arg > 0") }
  x := new(farMonitor)
  x.Any = Clone(a)
  x.uint = n
  x.ch = make([]nchan.NetChannel, x.uint)
  x.bool = s
  in, out := make([]chan Any, x.uint), make([]chan Any, x.uint)
  any := make([]Any, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ch[i] = nchan.NewN (x.Any, h, p + uint16(i), s)
    in[i], out[i] = x.ch[i].Chan()
  }
  if ! x.bool {
    return x // x is a client
  }
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  stmt()
  for i := uint(0); i < x.uint; i++ {
    go func (j uint) {
      for {
        select {
        case any[j] = <-When (x.PredSpectrum (x.Any, j), in[j]):
          if x.PredSpectrum (any[j], j) {
            out[j] <- x.FuncSpectrum (any[j], j)
          } else {
            out[j] <- x.FuncSpectrum (x.Any, j)
          }
        default:
        }
        time.Msleep(1e3)
      }
    }(i)
  }
  return nil
}

func (x *farMonitor) F (a Any, i uint) Any {
  if x.ch[i] == nil { panic ("no comm-channel") }
  x.ch[i].Send (a)
  return x.ch[i].Recv()
}

func (x *farMonitor) Fin() {
  for i := uint(0); i < x.uint; i++ {
    x.ch[i].Fin()
  }
}
