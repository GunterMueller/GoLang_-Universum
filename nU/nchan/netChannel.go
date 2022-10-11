package nchan

// (c) Christian Maurer   v. 220702 - license see nU.go

import ("time"; "strconv"; "net"; . "nU/obj")

const (
  maxWidth = uint(1<<12)
  network = "tcp"
)

type netChannel struct {
  any "Musterobjekt"
  uint "Kapazität des Kanals"
  in, out chan any
  FuncSpectrum
  PredSpectrum
  isServer,
  oneOne bool
  net.Conn
  net.Listener
  Stream "Puffer zur Datenübertragung"
  error
}
var c0 uint

func init() {
  c0 = C0()
}

func new_(a any, me, i uint, n string, p uint16) NetChannel {
  if me == i { panic("me == i") }
  x := new(netChannel)
  if a == nil {
    x.any, x.uint = nil, maxWidth
  } else {
    x.any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan any), make(chan any)
  x.Stream = make(Stream, x.uint)
  x.oneOne = true
  x.isServer = me < i
  ps := ":" + strconv.Itoa(int(Port0 + p))
  if x.isServer {
    x.Listener, x.error = net.Listen (network, n + ps)
if x.error != nil { println ("Listen error", n, ps) }
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network, n + ps); x.error == nil {
        break
      }
      time.Sleep (500 * 1e6)
    }
  }
  return x
}

func (x *netChannel) Send (a any) {
  if x.Conn == nil { panic("no Conn") }
  if x.any == nil {
    x.Conn.Write (append (Encode (Codelen(a)), Encode(a)...))
  } else {
    x.Conn.Write (Encode(a))
  }
}

func (x *netChannel) Recv() any {
  if x.Conn == nil { panic("no Conn") }
  if x.any == nil {
    _, x.error = x.Conn.Read (x.Stream[:c0])
    if x.error != nil { return nil }
    x.uint = Decode (uint(0), x.Stream[:c0]).(uint)
    _, x.error = x.Conn.Read (x.Stream[c0:c0+x.uint])
    if x.error != nil { return nil }
    return x.Stream[c0:c0+x.uint]
  }
  x.Conn.Read (x.Stream)
  return Decode(Clone(x.any), x.Stream)
}

func (x *netChannel) Fin() {
  x.Conn.Close()
  if x.isServer {
    x.Listener.Close()
    if ! x.oneOne {
      close(x.in)
      close(x.out)
    }
  }
}
