package nchan

// (c) Christian Maurer   v. 170923 - license see µu.go

import (
  "net"
  . "µu/ker"
  . "µu/obj"
  "µu/host"
  "µu/naddr"
)

// XXX Do not yet use this function !
// >>> Purely experimental attempt to integrate aspects of fmon into nchan
func news (a Any, n uint, fs FuncSpectrum, ps PredSpectrum, p uint16) NetChannel {
  if n == 0 { Panic ("nchan.NewS must be called with 2nd arg > 0") }
  x := new (netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  n0 := naddr.New2(host.Localhost(), p)
  x.uint = n
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  x.Listener, x.error = net.Listen (network, n0.String())
  x.panicIfErr()
  for i := uint(0); i < x.uint; i++ {
    go func (j uint) {
      for {
        select {
        case any, ok := <-When (x.PredSpectrum (x.Any, j), x.in):
          if ok {
            any = x.FuncSpectrum (any, j)
            x.out <- any
          } else {
//            println ("client off")
          }
        default:
          Msleep(100)
        }
      Msleep(100)
      }
    }(i)
  }
  return nil
}
