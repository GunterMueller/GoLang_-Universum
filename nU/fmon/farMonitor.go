package fmon

// (c) Christian Maurer   v. 171125 - license see nU.go

import ("time"; . "nU/obj"; "nU/nchan")

type farMonitor struct {
  uint "Anzahl der Monitor-Functionen"
  ch []nchan.NetChannel
  FuncSpectrum; PredSpectrum
  bool "Monitor ist der Anbieter"
}

func new_(a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
  x := new(farMonitor)
  pattern := Clone(a)
  x.uint = n
  x.ch = make([]nchan.NetChannel, x.uint)
  x.bool = s
  in := make([]chan Any, x.uint)
  out := make([]chan Any, x.uint)
  any := make([]Any, x.uint)
  for i := uint(0); i < x.uint; i++ {
    x.ch[i] = nchan.NewN (pattern, h, p + uint16(i), s)
    in[i], out[i] = x.ch[i].Chan()
  }
  if ! x.bool {
    return x // x ist ein Kunde
  }
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  for i := uint(0); i < x.uint; i++ {
    go func (j uint) {
      for {
        select {
        case any[j] = <-When (x.PredSpectrum (pattern, j), in[j]):
          if x.PredSpectrum (any[j], j) {
            out[j] <- x.FuncSpectrum (any[j], j)
          } else {
            out[j] <- x.FuncSpectrum (pattern, j)
          }
        default:
        }
        time.Sleep(1e9)
      }
    }(i)
  }
  return x
}

func (x *farMonitor) F (a Any, i uint) Any {
  x.ch[i].Send (a)
  return x.ch[i].Recv()
}

func (x *farMonitor) Fin() {
  for i := uint(0); i < x.uint; i++ {
    x.ch[i].Fin()
  }
}
