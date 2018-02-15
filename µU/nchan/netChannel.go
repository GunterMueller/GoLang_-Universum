package nchan

// (c) Christian Maurer   v. 180212 - license see µU.go

import (
//  "strconv"
  "net"
  . "µU/ker"
  "µU/time"
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
                    net.Listener
                    Stream "buffer"
                    error
                    }

func (x *netChannel) panicIfErr() {
  if x.error != nil {
    Panic (x.error.Error())
  }
}

func new_(a Any, me, i uint, n string, p uint16) NetChannel {
  if me == i { Panic("me == i") }
  x := new(netChannel)
  if a == nil {
    x.Any, x.uint = nil, maxWidth
  } else {
    x.Any, x.uint = Clone(a), Codelen(a)
  }
  x.in, x.out = make(chan Any), make(chan Any)
  x.Stream = make(Stream, x.uint)
  x.oneOne = true
  x.isServer = me < i
  h, port := host.NewS(n), Port0 + p
  if x.isServer {
    x.Listener, x.error = net.Listen (network, naddr.New2 (h, port).String())
    x.panicIfErr()
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network, naddr.New2 (h, port).String()); x.error == nil {
        break
      }
      errh.Hint (x.error.Error())
      time.Msleep (500)
    }
  }
  return x
}

func (x *netChannel) Send (a Any) {
  if x.Conn == nil { panic("no Conn") }
  if x.Any == nil {
    _, x.error = x.Conn.Write (append (Encode (Codelen(a)), Encode(a)...))
    if x.error != nil { println (x.error.Error()) }
  } else {
    CheckTypeEq (x.Any, a)
    _, x.error = x.Conn.Write (Encode(a))
  }
/*
  n := uint(len(bs)) / maxWidth
  for i := uint(0); i < n; i++ {
    _, x.err = x.Conn.Write (bs[i*maxWidth:(i+1)*maxWidth])
    if x.err != nil { println ("2. " + x.err.Error()) }
  }
  if uint(len(bs)) % maxWidth > 0 {
    s, x.err = x.Conn.Write (bs[n*maxWidth:])
    if x.err != nil { println ("3. (" + strconv.Itoa(s) + ") " + x.err.Error()) }
  }
*/
}

func (x *netChannel) Recv() Any {
//  var r int
  if x.Conn == nil { panic("no Conn") }
  if x.Any == nil {
    _, x.error = x.Conn.Read (x.Stream[:C0])
    if x.error != nil {
//      println (x.error.Error())
      return nil
    }
    x.uint = Decode (uint(0), x.Stream[:C0]).(uint)
    _, x.error = x.Conn.Read (x.Stream[C0:C0+x.uint])
    if x.error != nil { return nil }
    return x.Stream[C0:C0+x.uint]
  }
//  bs := make(Stream, x.uint)
//  _, x.errR = x.Conn.Read(x.Stream)
//  if x.errR != nil {
//    println ("6. " + x.errR.Error())
//    return Clone(x.Any)
//  }
//  copy (bs, x.Stream[:])
  _, x.error = x.Conn.Read (x.Stream)
  return Decode (Clone(x.Any), x.Stream)
//  return Decode(Clone(x.Any), bs)
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
