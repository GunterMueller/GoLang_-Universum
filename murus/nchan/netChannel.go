package nchan

// (c) murus.org  v. 170507 - license see murus.go

import (
  "net"
  . "murus/ker"
  . "murus/obj"
  "murus/errh"
//  "murus/nat"
  "murus/host"
  "murus/naddr"
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
//       cport, sport uint
                    }

func (x *netChannel) panicIfErr() {
  if x.error != nil {
    Panic (x.error.Error())
  }
}

func (x *netChannel) serve (c net.Conn) {
  for {
    if x.Any == nil {
      x.int, x.error = c.Read (x.buf)
//      if x.error != nil { println("Error:", x.error.Error()) } // TODO better error handling
      if x.int == 0 { break }
      x.uint = uint(Decode (uint(0), x.buf[:C0]).(uint))
// println(nat.String(x.cport), nat.String(x.sport), "<<", x.uint)
      if uint(x.int) != C0 + x.uint {
        errh.Error2("serve: x.int =", uint(x.int), "!=", C0 + x.uint)
      }
      x.in <- x.buf[C0:C0+x.uint]
// the calling process is blocked until until the server in the far monitor,
// that had called newcs, has sent his reply
      a := <-x.out
      x.uint = Codelen(a)
      x.int, x.error = c.Write(append(Encode(x.uint), Encode(a)...))
      if uint(x.int) != C0 + x.uint { Shit() }
    } else {
      x.int, x.error = c.Read (x.buf)
//      if x.error != nil { println("Error:", x.error.Error()) } // TODO better error handling
      if x.int == 0 { break }
      x.checkRecv()
      x.in <- Decode (Clone (x.Any), x.buf[:x.int])
      a := <-x.out
      x.int, x.error = c.Write (Encode(a))
      x.checkSend()
    }
  }
  x.nClients--
//  errh.Hint ("number of clients: " + nat.String(x.nClients))
  c.Close()
}

func newd (a Any, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic("nchan.newd: h.Empty()") }
  x := new(netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  x.oneOne = true
  n0, n1 := naddr.NewLocal(p), naddr.New2(h, p)
  x.isServer = n0.Less(n1)
  if x.isServer {
    x.Listener, x.error = net.Listen(network, naddr.New(p).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial(network, n1.String()); x.error == nil {
        break
      }
      Msleep(500)
    }
  }
  return x
}

func new_(a Any, me, i uint, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic("nchan.new_: h.Empty()") }
  if me == i { Panic("nchan.new_: me == i") }
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
  if x.isServer {
    x.Listener, x.error = net.Listen(network, naddr.New2(host.New(), p).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    dialaddr := naddr.New2(h, p).String()
    for {
      if x.Conn, x.error = net.Dial(network, dialaddr); x.error == nil {
        break
      }
      errh.Hint(x.error.Error())
      Msleep (500)
    }
  }
  return x
}

func newcs (a Any, h host.Host, p uint16, s bool /* , in, out chan Any */) NetChannel {
  x := new(netChannel)
  x.Any = Clone(a)
  x.uint = Codelen(a)
  if a == nil {
    x.uint = maxWidth
  }
  x.buf = make([]byte, x.uint)
  x.in, x.out = make(chan Any), make(chan Any) // in, out
  x.isServer = s
  if x.isServer {
//    x.cport = uint(p) - 50000
    x.Listener, x.error = net.Listen(network, naddr.New(p).String())
    x.panicIfErr()
    go func() {
      for {
        if c, e:= x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
//          nn, _ := nat.Natural(x.Listener.Addr().String()); x.cport = nn
          x.nClients++
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
  } else { // client
    for {
      if x.Conn, x.error = net.Dial(network, naddr.New2(h,p).String()); x.error == nil {
        break
      }
      Msleep(500)
    }
//    nn, _ := nat.Natural(x.Conn.RemoteAddr().String()[14:]); x.sport = nn - 50000
//    nn, _ = nat.Natural(x.Conn.LocalAddr().String()[14:]); x.cport = nn
  }
  return x
}

// XXX Do not yet use this function !
// >>> Purely experimental attempt to integrate aspects of fmon into nchan
func news (a Any, n uint, fs FuncSpectrum, ps PredSpectrum, p uint16) NetChannel {
  if n == 0 { Panic ("nchan.NewS must be called with 2nd arg > 0") }
  x := new (netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  n0 := naddr.New2(host.Localhost(), p)
  x.uint = n
  x.FuncSpectrum, x.PredSpectrum = fs, ps
  x.Listener, x.error = net.Listen (network, n0.String())
  x.panicIfErr()
  for i := uint(0); i < x.uint; i++ {
    go func (j uint) {
      for {
        select {
        case any, ok := <-When (x.PredSpectrum (x.Any, j), x.in):
          if ok {
            any = x.FuncSpectrum (any, j)
            x.out <- any
          } else {
//            println ("client off")
          }
        default:
          Msleep(100)
        }
      Msleep(100)
      }
    }(i)
  }
  return nil
}

/*
func newM (as []Any, me, i uint, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic("nchan.new_: h.Empty()") }
  if me == i { Panic("nchan.new_: me == i") }
  x := new(netChannel)
  x.as, x.ns = make([]Any, len(as)), make([]uint, len(as))
  for i, a := range as {
    if a == nil {
      x.as[i], x.ns[i] = nil, maxWidth
    } else {
      x.as[i], x.ns[i] = Clone(a), Codelen(a)
    }
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.buf = make([]byte, x.uint)
  x.oneOne = true
  x.isServer = me < i
  if x.isServer {
    x.Listener, x.error = net.Listen(network, naddr.New2(host.New(), p).String())
    x.panicIfErr()
    list := x.Listener
    x.Conn, x.error = list.Accept()
  } else { // client
    dialaddr := naddr.New2(h, p).String()
    for {
      if x.Conn, x.error = net.Dial(network, dialaddr); x.error == nil {
        break
      }
      errh.Hint(x.error.Error())
      Msleep (500)
    }
  }
  return x
}
*/

func (x *netChannel) Chan() (chan Any, chan Any) {
  return x.in, x.out
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
    if x.uint > maxWidth { Panic("object to send is too large") }
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

/*
func (x *netChannel) SendM (a ...Any) {
  if x.Conn == nil { Shit() }
  if x.Any == nil {
    x.uint = Codelen(a)
    if x.uint > maxWidth { Panic("object to send is too large") }
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
*/

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

/*
func (x *netChannel) RecvM() []Any {
  if x.Conn == nil { Panic("nchan.Recv: x.Conn == nil") }
  rs := make([]Any, len(x.as))
  for i, a := range x.as {
    if a == nil {
      x.int, x.error = x.Conn.Read(x.buf[:C0])
      if x.int == 0 {
        errh.Error0("nothing read")
        rs[i] = Clone(a)
      } else {
        x.uint = Decode (uint(0), x.buf[:C0]).(uint)
        x.int, x.error = x.Conn.Read(x.buf[C0:C0+x.uint])
        u, v := uint(x.int), x.uint // + C0
        if u != v {
          println("shit: u ==", u, "!= v ==", v)
        }
        rs[i] = x.buf[C0:C0+x.uint]
      }
    } else {
      x.int, x.error = x.Conn.Read(x.buf)
      x.checkRecv()
      if x.error == nil {
        rs[i] = Decode(Clone(a), x.buf)
      } else {
        rs[i] = Clone(x.Any)
      }
    }
  }
  return rs
}
*/

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

func port (n, i, j, a uint) uint16 {
  if a > 0 { Panic("a > 0") }
  const p0 = uint16(50000)
//  k := uint16(n * (n + 1)/ 2)
  if i > j { i, j = j, i } // i <= j
  return p0 + uint16(n * i - i * (i + 1) / 2 + j) // + uint16(a) * k
}

func nPorts (n, a uint) uint {
  return 1 * n * (n + 1) / 2
}
