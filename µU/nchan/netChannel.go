package nchan

// (c) Christian Maurer   v. 220702 - license see nU.go

import (
  "net"
  "µU/ker"
  "µU/time"
  "strconv"
  . "µU/obj"
  "µU/errh"
)
const (
  maxWidth = uint(1<<12)
  network = "tcp"
)
type
  netChannel struct {
                    any "Musterobjekt"
//               port uint16       
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

func (x *netChannel) panicIfErr() {
  if x.error != nil {
    ker.Panic (x.error.Error())
  }
}

func new_(a any, me, i uint, n string, p uint16) NetChannel {
  if me == i { ker.Panic ("me == i") }
  x := new(netChannel)
  if a == nil {
    x.any, x.uint = nil, maxWidth
  } else {
    x.any, x.uint = Clone (a), Codelen (a)
  }
  x.Stream = make(Stream, x.uint)
  x.in, x.out = make(chan any), make(chan any)
//  x.port = p
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
      errh.Hint (x.error.Error())
      time.Msleep (500)
    }
  }
  return x
}

func (x *netChannel) Send (a any) {
  if x.Conn == nil { ker.Panic ("no Conn") }
  if x.any == nil {
    x.Conn.Write (append (Encode (Codelen(a)), Encode(a)...))
  } else {
    x.Conn.Write (Encode(a))
  }
}

func (x *netChannel) Recv() any {
  if x.Conn == nil { ker.Panic ("no Conn") }
  if x.any == nil {
    _, x.error = x.Conn.Read (x.Stream[:C0])
    if x.error != nil { return nil }
    x.uint = Decode (uint(0), x.Stream[:C0]).(uint)
    _, x.error = x.Conn.Read (x.Stream[C0:C0+x.uint])
    if x.error != nil { return nil }
    return x.Stream[C0:C0+x.uint]
  }
  x.Conn.Read (x.Stream)
  return Decode (Clone(x.any), x.Stream)
}

func (x *netChannel) Fin() {
  x.Conn.Close()
  if x.isServer {
    x.Listener.Close()
    if ! x.oneOne {
      close (x.in)
      close (x.out)
    }
  }
}
