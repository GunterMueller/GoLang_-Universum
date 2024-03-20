package fmon

// (c) Christian Maurer   v. 240108 - license see µU.go

import (
//  "reflect"
  . "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/nchan"
)
type
  farMonitor struct {
      input, output any "pattern objects"
                    uint "number of monitor functions"
                 ch []nchan.NetChannel
                    FuncSpectrum; PredSpectrum
           isServer bool // "true iff the monitor is a server"
                    }

func new_(a any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
  return new1 (a, n, fs, ps, h, p, s, Nothing)
}

func new1 (a any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h string, p uint16, s bool, stmt Stmt) FarMonitor {
  if n == 0 { Panic ("fmon.New must be called with 2nd arg > 0") }
  x := new(farMonitor)
  x.input = Clone(a)
  x.uint = n
  x.isServer = s
  x.ch = make([]nchan.NetChannel, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ch[i] = nchan.NewN (x.input, h, p + uint16(i), s)
  }
//print ("fmon.new1: type of a is ", reflect.TypeOf(a).String(), "  n == ", n, "  port == ", p)
//if s { println (" as server") } else { println (" as client") }
  return x.common (fs, ps, stmt)
}

func new2 (a, b any, n uint, fs FuncSpectrum, ps PredSpectrum,
           h string, p uint16, s bool) FarMonitor {
  if n == 0 { Panic ("fmon.New2 must be called with 2nd arg > 0") }
  x := new(farMonitor)
  x.input = Clone(a)
  x.output = Clone(b)
  x.uint = n
  x.isServer = s
  x.ch = make([]nchan.NetChannel, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ch[i] = nchan.NewN (nil, h, p + uint16(i), s)
  }
  return x.common (fs, ps, Nothing)
}

func (x *farMonitor) common (fs FuncSpectrum, ps PredSpectrum, stmt Stmt) FarMonitor {
  in, out := make([]chan any, x.uint), make([]chan any, x.uint)
  for i := uint(0); i < x.uint; i++ {
    in[i], out[i] = x.ch[i].Chan()
  }
  if ! x.isServer {
    return x // x is a client
  }
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  stmt()
  any := make([]any, x.uint)
  for i := uint(0); i < x.uint; i++ {
    go func (j uint) {
      for {
        select {
        case any[j] = <-When (x.PredSpectrum (x.input, j), in[j]):
          if x.PredSpectrum (any[j], j) {
            out[j] <- x.FuncSpectrum (any[j], j)
          } else {
            out[j] <- x.FuncSpectrum (x.output, j)
          }
        default:
        }
        time.Msleep(10)
      }
    }(i)
  }
  return x
}

func (x *farMonitor) F (a any, i uint) any {
  if x.ch[i] == nil { panic("no comm-channel") }
  x.ch[i].Send (a)
  return x.ch[i].Recv()
}

func (x *farMonitor) Fin() {
  for i := uint(0); i < x.uint; i++ {
    x.ch[i].Fin()
  }
}
