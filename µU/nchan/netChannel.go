package nchan

// (c) Christian Maurer   v. 171125 - license see µU.go

import (
  "net"
  . "µU/ker"
  . "µU/obj"
  "µU/errh"
//  "µU/nat"
  "µU/host"
  "µU/naddr"
)
const (
  maxWidth = uint(1<<16)
  network = "tcp"
)
type
  netChannel struct {
                    Any "type of objects in the channel"
                    uint "byte capacity of the channel"
            in, out chan Any
                    FuncSpectrum
                    PredSpectrum
           isServer,
             oneOne bool
                    net.Conn
           nClients uint
                    net.Listener
                buf []byte
                    int "number of sent/received bytes"
                    error
                    }

func (x *netChannel) panicIfErr() {
  if x.error != nil {
    Panic (x.error.Error())
  }
}

func new_(a Any, me, i uint, n string, p uint16) NetChannel {
  if me == i { Panic("nchan.New: me == i") }
  x := new(netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  x.oneOne = true
  x.isServer = me < i
  h, port := host.NewS(n), Port0 + p
  if x.isServer {
    x.Listener, x.error = net.Listen (network, naddr.New2 (h, port).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    dialaddr := naddr.New2 (h, port).String()
    for {
      if x.Conn, x.error = net.Dial (network, dialaddr); x.error == nil {
        break
      }
      errh.Hint(x.error.Error())
      Msleep (500)
    }
  }
  return x
}

func (x *netChannel) checkSend() {
  if x.error != nil && x.int < int(x.uint) {
    println (x.error.Error(), "(sent only", x.int, "bytes)")
  }
}

func (x *netChannel) Send (a Any) {
  if x.Conn == nil { Shit() }
  if x.Any == nil {
    x.uint = Codelen(a)
    if x.uint > maxWidth { Panic ("object to send is too large") }
    x.int, x.error = x.Conn.Write (append (Encode (x.uint), Encode(a)...))
    if x.error != nil && uint(x.int) < x.uint + C0 {
      println(x.error.Error(), "(sent only", x.int, "bytes)")
//    TODO better error handling
    }
  } else {
    CheckTypeEq (x.Any, a)
    x.int, x.error = x.Conn.Write(Encode(a))
    x.checkSend()
  }
}

func (x *netChannel) checkRecv() {
// TODO better handling
  switch x.int {
  case -1:
    println ("partner closed")
  case 0:
    println ("connection to partner broken")
  default:
    if x.error != nil && x.int < int(x.uint) {
      println (x.error.Error() + " (received only", x.int, "bytes)")
    }
  }
}

func (x *netChannel) Recv() Any {
  if x.Conn == nil { Panic("nchan.Recv: x.Conn == nil") }
  if x.Any == nil {
    x.int, x.error = x.Conn.Read(x.buf[:C0])
    if x.int == 0 { errh.Error0("nothing read"); return Clone(x.Any) }
    x.uint = Decode (uint(0), x.buf[:C0]).(uint)
    x.int, x.error = x.Conn.Read(x.buf[C0:C0+x.uint])
    u, v := uint(x.int), x.uint // + C0
    if u != v {
      println("shit: u ==", u, "!= v ==", v)
    }
    return x.buf[C0:C0+x.uint]
  }
  x.int, x.error = x.Conn.Read(x.buf)
  x.checkRecv()
  if x.error != nil {
    return Clone(x.Any)
  }
  return Decode(Clone(x.Any), x.buf)
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
