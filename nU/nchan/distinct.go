package nchan

// (c) Christian Maurer   v. 171125 - license see nU.go

import ("strconv"; "time"; "net"; . "nU/obj"; "nU/env")

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
  ps := ":" + strconv.Itoa(int(Port0 + p))
  n0, n1 := env.Localhost() + ps, h + ps
  x.isServer = n0 < n1
  if x.isServer {
    x.Listener, x.error = net.Listen(network, n0)
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial(network, n1); x.error == nil {
        break
      }
      time.Sleep (500 * 1e6)
    }
  }
  return x
}
