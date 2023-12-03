package nchan

// (c) Christian Maurer   v. 220420 - license see µU.go

import (
  "errors"
  "net"
  "µU/ker"
  "µU/time"
  . "µU/obj"
  "µU/errh"
  "µU/host"
  "µU/naddr"
)
const (
  maxWidth = uint(1<<16)
  network = "tcp"
)
type
  netChannel struct {
                    any "type of objects in the channel"
               port uint16
                    uint "byte capacity of the channel"
            in, out chan any
                    FuncSpectrum
                    PredSpectrum
           isServer,
             oneOne bool
                    net.Conn
                    net.Listener
                    Stream "buffer"
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
  x.port, x.Stream, x.in, x.out = p, make(Stream, x.uint), make(chan any), make(chan any)
  x.oneOne = true
  x.isServer = me < i
  ht, port := host.NewS (n), Port0 + p
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

func (x *netChannel) Send (a any) error {
  if x.Conn == nil {
    return errors.New ("no Connection:")
  }
  if x.any == nil {
    bs := Encode(a)
    bs = append (Encode(Codelen(a)), bs...)
    _, x.error = x.Conn.Write (bs)
  } else {
// print ("Send on port ", x.port, ", ")
// ta, ts := reflect.TypeOf(x.any).String(), reflect.TypeOf(a).String()
// println ("pattern has type", ta, "and object to send has type", ts)
    CheckTypeEq (x.any, a)
    _, x.error = x.Conn.Write (Encode(a))
  }
  return x.error
}

func (x *netChannel) Recv() any {
  if x.Conn == nil {
    ker.Panic ("no Conn")
  }
  if x.any == nil {
    _, x.error = x.Conn.Read (x.Stream[:C0])
    if x.error == nil {
      return Clone (x.any)
    }
    x.uint = Decode (uint(0), x.Stream[:C0]).(uint)
    _, x.error = x.Conn.Read (x.Stream[C0:C0+x.uint])
    if x.error != nil {
      return Clone (x.any)
    }
    return x.Stream[C0:C0+x.uint]
  }
  _, x.error = x.Conn.Read (x.Stream)
  return Decode (Clone(x.any), x.Stream)
}

func (x *netChannel) Fin() {
  if x.Conn != nil {
    x.Conn.Close()
  }
  if x.isServer {
    x.Listener.Close()
    if ! x.oneOne {
      close(x.in)
      close(x.out)
    }
  }
}
