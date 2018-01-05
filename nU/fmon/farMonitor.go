package fmon

// (c) Christian Maurer   v. 180101 - license see nU.go

import ("time"; . "nU/obj"; "nU/nchan")

type farMonitor struct {
  Any "Musterobjekt"
  uint "Anzahl der Monitorfunctionen"
  ch []nchan.NetChannel
  FuncSpectrum
  PredSpectrum
  bool "true gdw. Monitor der Anbieter ist"
}

func new_(a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h string, p uint16, s bool) FarMonitor {
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
    return x // x ist ein Kunde
  }
  x.FuncSpectrum, x.PredSpectrum = fs, ps // x ist der Server
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
