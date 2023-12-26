package nchan

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "net"
  "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/errh"
  "µU/host"
  "µU/naddr"
)

func new1 (a any, me, i uint, n string, p uint16, s bool) NetChannel {
  if me == i { ker.Panic ("me == i") }
  x := new(netChannel)
  if a == nil {
    x.any, x.uint = nil, maxWidth
  } else {
    x.any, x.uint = Clone (a), Codelen (a)
  }
  x.port = p
  x.in, x.out = make(chan any), make(chan any)
  x.Stream = make(Stream, x.uint)
  x.oneOne = true // ? XXX
  x.isServer = s
  ht, port := host.NewS(n), Port0 + p
  if x.isServer {
    x.Listener, x.error = net.Listen (network, naddr.New2 (ht, port).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    dialaddr := naddr.New2 (ht, port).String()
    for {
      if x.Conn, x.error = net.Dial (network, dialaddr); x.error == nil {
        break
      }
      errh.Hint (x.error.Error())
      time.Msleep (500)
    }
  }
  return x
}
