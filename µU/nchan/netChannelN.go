package nchan

// (c) Christian Maurer   v. 171107 - license see µU.go

import (
  "net"
  . "µU/ker"
  . "µU/obj"
  "µU/errh"
//  "µU/nat"
  "µU/host"
  "µU/naddr"
)

func (x *netChannel) serve (c net.Conn) {
  for {
    x.int, x.error = c.Read (x.buf)
//    if x.error != nil { println("Error:", x.error.Error()) } // TODO better error handling
    if x.int == 0 { break }
    if x.Any == nil {
      x.uint = uint(Decode (uint(0), x.buf[:C0]).(uint))
// println(nat.String(x.cport), nat.String(x.sport), "<<", x.uint)
      if uint(x.int) != C0 + x.uint {
        errh.Error2("serve: x.int =", uint(x.int), "!=", C0 + x.uint)
      }
      x.in <- x.buf[C0:C0+x.uint]
// the calling process is blocked until until the server in the far monitor,
// that had called newn, has sent his reply
      a := <-x.out
      x.uint = Codelen(a)
      x.int, x.error = c.Write(append(Encode(x.uint), Encode(a)...))
      if uint(x.int) != C0 + x.uint { Shit() }
    } else {
      x.checkRecv()
      x.in <- Decode (Clone (x.Any), x.buf[:x.int])
      x.int, x.error = c.Write (Encode(<-x.out))
      x.checkSend()
    }
  }
  x.nClients--
//  errh.Hint ("number of clients: " + nat.String(x.nClients))
  c.Close()
}

func newn (a Any, h string, p uint16, s bool) NetChannel {
  x := new(netChannel)
  x.Any = Clone(a)
  x.uint = Codelen(a)
  if a == nil {
    x.uint = maxWidth
  }
  x.buf = make([]byte, x.uint)
  x.in, x.out = make(chan Any), make(chan Any)
  x.isServer = s
  if x.isServer {
//    x.cport = uint(p) - 50000
//    x.Listener, x.error = net.Listen (network, naddr.New (port).String())
    x.Listener, x.error = net.Listen (network, naddr.New (p).String())
    x.panicIfErr()
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
//          nn, _ := nat.Natural(x.Listener.Addr().String()); x.cport = nn
          x.nClients++
// println(x.nClients)
//          errh.Hint("number of clients: " + nat.String(x.nClients))
//                   port von c.LocalAddr == x.cport
//          nn, _ := nat.Natural(c.RemoteAddr().String()[14:]); x.sport = nn
// println("x.sport", x.sport)
          go x.serve (c) // see above remark
        } else {
          Panic(e.Error())
        }
      }
    }()
  } else { // clientA
    ht := host.NewS (h)
    for {
      if x.Conn, x.error = net.Dial (network, naddr.New2 (ht, p).String()); x.error == nil {
        break
      }
      Msleep(500)
    }
//    nn, _ := nat.Natural(x.Conn.RemoteAddr().String()[14:]); x.sport = nn - 50000
//    nn, _ = nat.Natural(x.Conn.LocalAddr().String()[14:]); x.cport = nn
  }
  return x
}

func (x *netChannel) Chan() (chan Any, chan Any) {
  return x.in, x.out
}
