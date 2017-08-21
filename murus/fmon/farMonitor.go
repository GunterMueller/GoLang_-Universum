package fmon

// (c) murus.org  v. 170524 - license see murus.

import (
  . "murus/ker"
  . "murus/obj"
  "murus/host"
  "murus/nchan"
)
type
  farMonitor struct {
                    Any "type of objects the monitor functions operate on"
                    uint "number of monitor functions"
                 ch []nchan.NetChannel
                    FuncSpectrum; PredSpectrum
                    bool "true iff the monitor is a server"
                    }

func new_(a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h host.Host, p uint16, s bool) FarMonitor {
  return newS (a, n, fs, ps, h, p, s, Null)
}

func newS (a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h host.Host, p uint16, s bool, stmt Stmt) FarMonitor {
  if n == 0 { Panic ("fmon.New must be called with 2nd arg > 0") }
  x := new(farMonitor)
  x.Any = Clone(a)
  x.uint = n
  x.ch = make([]nchan.NetChannel, x.uint)
  x.bool = s
  in, out := make([]chan Any, x.uint), make([]chan Any, x.uint)
  any, ok := make([]Any, x.uint), make([]bool, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ch[i] = nchan.NewCS (x.Any, h, p + uint16(i), s)
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
        case any[j], ok[j] = <-When (x.PredSpectrum (x.Any, j), in[j]):
          if ok[j] {
            if x.PredSpectrum (any[j], j) {
              any[j] = x.FuncSpectrum (any[j], j)
              out[j] <- any[j]
            } else {
              out[j] <- x.Any
            }
          } else {
//            println("client off")
          }
        default:
//          println("client off")
        }
        Msleep(300)
      }
    }(i)
  }
  return nil
}

func (x *farMonitor) F (a Any, i uint) Any {
  if x.ch[i] == nil { Panic("fmon.F: x.ch == nil") }
  x.ch[i].Send (a)
  return x.ch[i].Recv()
}

func (x *farMonitor) Fin() {
  for i := uint(0); i < x.uint; i++ {
    x.ch[i].Fin()
  }
}
