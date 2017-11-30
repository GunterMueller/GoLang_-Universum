package nchan

// (c) Christian Maurer   v. 170925 - license see µU.go

import (
  "net"
  . "µU/ker"
  . "µU/obj"
  "µU/host"
  "µU/naddr"
)

func newd (a Any, h string, p uint16) NetChannel {
  x := new(netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  x.oneOne = true
  n0, n1 := naddr.NewLocal(p), naddr.New2(host.NewS(h), p)
  x.isServer = n0.Less(n1)
  if x.isServer {
    x.Listener, x.error = net.Listen(network, naddr.New(p).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial(network, n1.String()); x.error == nil {
        break
      }
      Msleep(500)
    }
  }
  return x
}
