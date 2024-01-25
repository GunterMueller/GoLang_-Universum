package nchan

// (c) Christian Maurer   v. 231220 - license see nU.go

import (
  "strconv"
  "net"
  "time"
  . "nU/obj"
)

func (x *netChannel) Chan() (chan any, chan any) {
  return x.in, x.out
}

func (x *netChannel) serve (c net.Conn) {
  var r int
  for {
    r, x.error = c.Read (x.Stream)
    if r == 0 {
      break
    }
    if x.any == nil {
      x.uint = uint(Decode (uint(0), x.Stream[:C0]).(uint))
      x.in <- x.Stream[C0:C0+x.uint]
// the calling process is blocked until the server in the far monitor,
// that has call newn, hast sent its reply
      a := <-x.out
      _, x.error = c.Write (append(Encode(Codelen(a)), Encode(a)...))
      if x.error != nil { panic (x.error.Error()) }
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
    if x.error != nil { panic (x.error.Error()) }
    go func() {
      for {
        if c, e := x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
          go x.serve (c) // see above remark
        }
      }
    }()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network, h + ps); x.error == nil {
        break
      }
      time.Sleep(500 * 1e6)
    }
  }
  return x
}
