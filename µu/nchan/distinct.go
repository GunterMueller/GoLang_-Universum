package nchan

// (c) Christian Maurer   v. 170923 - license see µu.go

import (
  "net"
  . "µu/ker"
  . "µu/obj"
  "µu/host"
  "µu/naddr"
)

func newd (a Any, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic("nchan.newd: h.Empty()") }
  x := new(netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  x.oneOne = true
  n0, n1 := naddr.NewLocal(p), naddr.New2(h, p)
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
