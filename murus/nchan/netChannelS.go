package nchan

// (c) murus.org  v. 161228 - license see murus.go
//
// XXX DO NOT USE ! Purely experimental attempt to integrate fmon into nchan

import (
  "net"
  . "murus/ker" // Panic
  . "murus/obj"
  "murus/errh"
  "murus/perm"
  "murus/host"
  "murus/naddr"
)

func NewServer (a Any, n uint, fs FuncSpectrum, ps PredSpectrum, p uint16) NetChannel {
  if n == 0 { Panic ("nchan.NewS must be called with 2nd arg > 0") }
  x := new (netChannel)
  if a == nil {
    x.Any, x.width = nil, maxWidth
  } else {
    x.Any, x.width = Clone(a), Codelen(a)
  }
  x.buf = make([]byte, x.width)
  n0 := naddr.New2(host.Localhost(), p)
  x.nFuncs = n
  x.isServer = true
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  x.Listener, x.error = net.Listen(network, n0.String())
  x.panicIfErr()
  x.chans = make([]NetChannel, x.nFuncs)
/*
  for i := uint(0); i < x.nFuncs; i++ {
    x.chans[i] = NewCS (x.Any, h, p + uint16(i), true)
//                             ^
  }
*/
  x.ccin, x.ccout = make(chan Any), make(chan Any)
  for {
    for i := uint(0); i < x.nFuncs; i++ {
      j := x.Permutation.F (i)
      x.Permutation = perm.New (x.nFuncs)
      in, out := x.chans[j].Chan()
      select {
      case a, p := <-When (x.PredSpectrum (x.Any, j), in):
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
