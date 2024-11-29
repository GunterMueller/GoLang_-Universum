package nchan

// (c) Christian Maurer   v. 241017 - license see µU.go

import (
  "strconv"
  "net"
  "µU/ker"
  "µU/time"
  . "µU/obj"
//  "µU/host"
//  "µU/naddr"
)

func (x *netChannel) Chan() (chan any, chan any) {
  return x.in, x.out
}

func (x *netChannel) serve (c net.Conn) {
  var r int
  for {
    r, x.error = c.Read (x.Stream)
    if r == 0 { // x.error != nil {
      break
    }
    if x.any == nil {
      x.uint = uint(Decode (uint(0), x.Stream[:C0]).(uint))
      x.in <- x.Stream[C0:C0+x.uint]
// the calling process is blocked until the server in the far monitor,
// that had called newn, has sent its reply
      a := <-x.out
      _, x.error = c.Write (append(Encode(Codelen(a)), Encode(a)...))
      if x.error != nil { ker.Panic (x.error.Error()) }
    } else {
      x.in <- Decode (Clone (x.any), x.Stream[:r])
      _, x.error = c.Write (Encode(<-x.out))
    }
  }
  c.Close()
}

func newn (a any, h string, p uint16, s bool) NetChannel {
  x := new(netChannel)
  x.any = Clone (a)
  x.uint = Codelen (a)
  if a == nil {
    x.uint = maxWidth
  }
  x.Stream = make(Stream, x.uint)
  x.in, x.out = make(chan any), make(chan any)
  x.isServer = s
  ps := ":" + strconv.Itoa (int(p))
  if x.isServer {
    x.Listener, x.error = net.Listen (network, ps)
    x.panicIfErr()
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
          go x.serve (c) // see above remark
        }
      }
    }()
  } else {
    for {
      if x.Conn, x.error = net.Dial (network, h + ps); x.error == nil {
        break
      }
      time.Msleep(500)
    }
  }
  return x
}
