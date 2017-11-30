package nchan

// (c) Christian Maurer   v. 171125 - license see nU.go

import ("strconv"; "time"; "net"; . "nU/obj")

func (x *netChannel) Chan() (chan Any, chan Any) {
  return x.in, x.out
}

func (x *netChannel) serve (c net.Conn) {
  for {
    x.int, x.error = c.Read (x.buf)
    if x.int == 0 { break }
    if x.Any == nil {
      x.uint = uint(Decode (uint(0), x.buf[:C0]).(uint))
      x.in <- x.buf[C0:C0+x.uint]
      a := <-x.out
      x.uint = Codelen(a)
      x.int, x.error = c.Write(append(Encode(x.uint), Encode(a)...))
    } else {
      x.in <- Decode (Clone (x.Any), x.buf[:x.int])
      x.int, x.error = c.Write (Encode(<-x.out))
    }
  }
  x.nClients--
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
  ps := ":" + strconv.Itoa(int(p))
  if x.isServer {
    x.Listener, x.error = net.Listen (network, ps)
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
