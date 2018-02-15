package nchan

// (c) Christian Maurer   v. 180212 - license see µU.go

import (
  "net"
  . "µU/ker"
  "µU/time"
  . "µU/obj"
//  "µU/nat"
  "µU/host"
  "µU/naddr"
)

func (x *netChannel) serve (c net.Conn) {
  var r int
  for {
    r, x.error = c.Read (x.Stream)
    if x.error != nil {
      break
    }
    if x.Any == nil {
      x.uint = uint(Decode (uint(0), x.Stream[:C0]).(uint))
// println(nat.String(x.cport), nat.String(x.sport), "<<", x.uint)
      x.in <- x.Stream[C0:C0+x.uint]
// the calling process is blocked until until the server in the far monitor,
// that had called newn, has sent his reply
      a := <-x.out
      _, x.error = c.Write(append(Encode(Codelen(a)), Encode(a)...))
      if x.error != nil { println(x.error.Error()) }
    } else {
      x.in <- Decode (Clone (x.Any), x.Stream[:r])
      _, x.error = c.Write (Encode(<-x.out))
      if x.error != nil { println(x.error.Error()) } // provisorial
    }
  } // x.nClients--
  c.Close()
}

func newn (a Any, h string, p uint16, s bool) NetChannel {
  x := new(netChannel)
  x.Any = Clone(a)
  x.uint = Codelen(a)
  if a == nil {
    x.uint = maxWidth
  }
  x.Stream = make(Stream, x.uint)
  x.in, x.out = make(chan Any), make(chan Any)
  x.isServer = s
  if x.isServer {
//    x.cport = uint(p) - 50000
//    x.Listener, x.err = net.Listen (network, naddr.New (port).String())
    x.Listener, x.error = net.Listen (network, naddr.New (p).String())
    x.panicIfErr()
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.err !
//          nn, _ := nat.Natural(x.Listener.Addr().String()); x.cport = nn
//          x.nClients++
//                   port von c.LocalAddr == x.cport
//          nn, _ := nat.Natural(c.RemoteAddr().String()[14:]); x.sport = nn
// println("x.sport", x.sport)
          go x.serve (c) // see above remark
        } else {
          Panic(e.Error())
        }
      }
    }()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network,
                                     naddr.New2 (host.NewS (h), p).String()); x.error == nil {
        break
      }
      time.Msleep(500)
    }
//    nn, _ := nat.Natural(x.Conn.RemoteAddr().String()[14:]); x.sport = nn - 50000
//    nn, _ = nat.Natural(x.Conn.LocalAddr().String()[14:]); x.cport = nn
  }
  return x
}

func (x *netChannel) Chan() (chan Any, chan Any) {
  return x.in, x.out
}
