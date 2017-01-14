package nchan

// (c) murus.org  v. 170106 - license see murus.go

import (
  "errors"
  "net"
  . "murus/ker" // Panic
  . "murus/obj"
  "murus/errh"
  "murus/perm"
//  "murus/nat"
  "murus/host"; "murus/naddr"
)
const (
  maxWidth = 1<<16
  network = "tcp"
)
type
  netChannel struct {
                    Any "type of objects in the channel"
              width uint
             nFuncs uint // XXX
              chans []NetChannel // XXX
                    FuncSpectrum // XXX
                    PredSpectrum // XXX
                    perm.Permutation
           isServer,
             oneOne bool
                    net.Conn
           nClients uint
                    net.Listener
                buf []byte
        ccin, ccout chan Any
                    int "number of sent/received bytes"
                    error
              conns map[uint64]net.Conn
                    }

func (x *netChannel) panicIfErr() {
  if x.error != nil {
    Panic (x.error.Error())
  }
}

func (x *netChannel) Chan() (chan Any, chan Any) {
  if x.isServer && ! x.oneOne { // 1:n-server
    return x.ccin, x.ccout
  }
  return nil, nil
}

func (x *netChannel) serve (c net.Conn) {
  for {
    if x.Any == nil {
      x.int, x.error = c.Read (x.buf[:4])
      if x.int == 0 { break }
      x.width = uint(Decode (uint32(0), x.buf[:4]).(uint32))
// errh.Error ("serve recv", x.width)
      x.int, x.error = c.Read (x.buf[:x.width])
      x.ccin <- x.buf[:x.width]
// println ("in ok, out ?", x.width); errh.Error ("in ok, out ?", x.width)
      a := <-x.ccout
      x.width = Codelen(a)
errh.Error ("not reached", x.width)
      x.int, x.error = c.Write (Encode (uint32(x.width)))
      if x.int != 4 { Oops() }
errh.Error ("not reached", Codelen(a))
      x.int, x.error = c.Write (Encode(a))
      if x.int != int(x.width) { Oops() }
    } else {
      x.int, x.error = c.Read (x.buf)
      x.checkRecv()
      if x.int == 0 { break }
      x.ccin <- Decode (Clone (x.Any), x.buf)
      a := <-x.ccout
// calling process blocks until the function of the client is executed
      x.int, x.error = c.Write (Encode (a))
      x.checkSend()
    }
  }
  x.nClients--
//  errh.Hint ("number of clients: " + nat.String(x.nClients))
  c.Close()
}

func newd (a Any, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic(errors.New("nchan.newd: h.Empty()").Error()) }
  x := new(netChannel)
  if a == nil {
    x.Any, x.width = nil, maxWidth
  } else {
    x.Any, x.width = Clone(a), Codelen(a)
  }
  x.buf = make([]byte, x.width)
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
        errh.DelHint()
        break
      }
      Msleep (500)
    }
  }
  return x
}

func new_ (a Any, me, i uint, h host.Host, p uint16) NetChannel {
  if h.Empty() { Panic(errors.New("nchan.new_: h.Empty()").Error()) }
  if me == i { Panic(errors.New("nchan.new_: me == i").Error()) }
  x := new(netChannel)
  if a == nil {
    x.Any, x.width = nil, maxWidth
  } else {
    x.Any, x.width = Clone(a), Codelen(a)
  }
  x.buf = make([]byte, x.width)
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
      Msleep (500)
    }
  }
  return x
}

func (x *netChannel) checkSend() {
// TODO better handling
  if x.error != nil && x.int < int(x.width) {
    println (x.error.Error(), "(sent only", x.int, "bytes)")
  }
}

func newCS (a Any, h host.Host, p uint16, s bool) NetChannel {
  x := new(netChannel)
  if a == nil {
    x.Any, x.width = nil, maxWidth
  } else {
    x.Any, x.width = Clone(a), Codelen(a)
  }
  x.buf = make([]byte, x.width)
  x.isServer = s
  if x.isServer {
    x.Listener, x.error = net.Listen(network, naddr.New(p).String())
    x.panicIfErr()
    x.ccin, x.ccout = make(chan Any), make(chan Any)
    go func() {
      for {
        if c, e:= x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
          x.nClients ++
//          errh.Hint("number of clients: " + nat.String(x.nClients))
          go x.serve(c) // see above remark
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
  }
  return x
}

func (x *netChannel) Send (a Any) {
  if x.Conn == nil { Panic("nchan.Send: x.Conn == nil") }
  if x.Any == nil { // XXX
    x.width = Codelen(a)
    if x.width > maxWidth { Panic("object to send is too large") }
    x.int, x.error = x.Conn.Write (Encode (uint32(x.width)))
    if x.error != nil && x.int < 4 {
      println(x.error.Error(), "(sent only", x.int, "bytes)")
//    TODO better error handling
    }
  } else {
    CheckTypeEq (x.Any, a)
  }
  x.int, x.error = x.Conn.Write(Encode(a))
  x.checkSend()
}

func (x *netChannel) checkRecv() {
// TODO better handling
  switch x.int {
  case -1:
    println ("partner closed")
  case 0:
    println ("connection to partner broken")
  default:
    if x.error != nil && x.int < int(x.width) {
      println (x.error.Error() + " (received only", x.int, "bytes)")
    }
  }
}

func (x *netChannel) Recv() Any {
  if x.Conn == nil { Panic("nchan.Recv: x.Conn == nil") }
  if x.Any == nil { // XXX
// errh.Error ("wanna read", 4)
    x.int, x.error = x.Conn.Read(x.buf[:4])
    x.width = uint(Decode (uint32(0), x.buf[:4]).(uint32))
// errh.Error ("have read", x.width)
    x.int, x.error = x.Conn.Read(x.buf[:x.width])
    if x.int == 0 {
      return make([]byte, 0)
    }
    x.checkRecv()
    if x.error != nil {
      return Clone(x.Any)
    }
    return x.buf[:x.width]
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
      close(x.ccin)
      close(x.ccout)
    }
  }
}
