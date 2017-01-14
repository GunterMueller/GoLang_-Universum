package fmon

// (c) murus.org  v. 170104 - license see murus.go

import (
  . "murus/ker"; . "murus/obj"; "murus/perm"; "murus/errh"
//  "murus/nat"
  "murus/host"
//  "murus/naddr"
  "murus/nchan"
)
const
  pack = "fmon"
type
  farMonitor struct {
                    Any
                    uint "number of monitor functions"
              chans []nchan.NetChannel
                    FuncSpectrum
                    PredSpectrum
                    perm.Permutation
           isServer bool
                    }

/*
func New1 (a Any, n uint, fs FuncSpectrum, ps PredSpectrum, na naddr.NetAddress, s bool) FarMonitor {
  h, p := na.HostPort()
  return new_(a, n, fs, ps, h, p, s)
}
*/

func new_(a Any, n uint, fs FuncSpectrum, ps PredSpectrum, h host.Host, p uint16, s bool) FarMonitor {
  return news(a, n, fs, ps, h, p, s, Null)
}

func news(a Any, n uint, fs FuncSpectrum, ps PredSpectrum,
          h host.Host, p uint16, s bool, stmt Stmt) FarMonitor {
  if n == 0 { Panic ("fmon.New must be called with 2nd arg > 0") }
  x := new(farMonitor)
  x.Any = Clone(a)
  x.uint = n
  x.chans = make([]nchan.NetChannel, x.uint)
  x.isServer = s
  for i := uint(0); i < x.uint; i++ {
    x.chans[i] = nchan.NewCS (x.Any, h, p + uint16(i), s)
  }
  if ! x.isServer {
    return x
  }
  x.Permutation = perm.New (x.uint)
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  stmt()
  for {
    x.Permutation.Permute()
    for i := uint(0); i < x.uint; i++ {
      j := x.Permutation.F (i)
      in, out := x.chans[j].Chan()
      select {
      case a, p := <-When (x.PredSpectrum (x.Any, j), in):
// XXX                                     ^^^^^
//    Hier brauchen wir neben diesem Typ im Grunde noch einen weiteren,
//    was dazu führt, dass zwei verschiedene Kanäle gebraucht werden.
        if p {
          a = x.FuncSpectrum (a, j) // interface is nil ?
          out <- a // ist am Schluss nil
        } else {
          errh.Error0 ("client off")
        }
      default:
      }
    }
    Sleep (1)
  }
  return nil
}

/*
func (x *farMonitor) F0 (a Any, i uint) { // XXX
  if x.chans[i] == nil { Stop (pack, 2) }
  x.chans[i].Send (a)
}
*/

func (x *farMonitor) F (a Any, i uint) Any {
  if x.chans[i] == nil { Stop (pack, 1) }
  x.chans[i].Send (a)
  return x.chans[i].Recv()
}

// experimental // XXX
func (x *farMonitor) S (a Any, i uint, c chan Any) {
  c <- x.F (a, i)
}

func (x *farMonitor) Fin() {
  for i := uint(0); i < x.uint; i++ {
    x.chans[i].Fin()
  }
}
