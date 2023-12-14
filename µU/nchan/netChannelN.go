package nchan

// (c) Christian Maurer   v. 231213 - license see µU.go

import (
  "net"
  "µU/ker"
  "µU/time"
  . "µU/obj"
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
    if x.any == nil {
      x.uint = uint(Decode (uint(0), x.Stream[:C0]).(uint))
      x.in <- x.Stream[C0:C0+x.uint]
// the calling process is blocked until until the server in the far monitor,
// that had called newn, has sent his reply
      a := <-x.out
      _, x.error = c.Write(append(Encode(Codelen(a)), Encode(a)...))
      if x.error != nil { ker.Panic (x.error.Error()) }
    } else {
      x.in <- Decode (Clone (x.any), x.Stream[:r])
      _, x.error = c.Write (Encode(<-x.out))
      if x.error != nil { ker.Panic (x.error.Error()) } // provisorial
    }
  }
  c.Close()
}

func newn (a any, h string, p uint16, s bool) NetChannel {
  x := new(netChannel)
  if a == nil {
    x.any, x.uint = Clone(a), maxWidth
  } else {
    x.any, x.uint = Clone(a), Codelen(a)
  }
  x.port, x.Stream, x.in, x.out = p, make(Stream, x.uint), make(chan any), make(chan any)
  x.isServer = s
  if x.isServer { // println ("server", ego.Me())
    x.Listener, x.error = net.Listen (network, naddr.New (p).String())
    x.panicIfErr()
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
          go x.serve (c) // see above remark
        } else {
          break // Panic(e.Error())
        }
      }
    }()
  } else { // println ("client", ego.Me())
    ht := host.NewS (h)
    for {
      if x.Conn, x.error = net.Dial (network, naddr.New2 (ht, p).String()); x.error == nil {
        break
      }
      time.Msleep(500)
    }
  }
  return x
}

func (x *netChannel) Chan() (chan any, chan any) {
  return x.in, x.out
}
