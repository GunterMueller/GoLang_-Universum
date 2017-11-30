package nchan

// (c) Christian Maurer   v. 171125 - license see nU.go

import ("strconv"; "time"; "net"; . "nU/obj")

const (
  maxWidth = uint(1<<16)
  network = "tcp"
)
type netChannel struct {
  Any "type of objects in the channel"
  uint "byte capacity of the channel"
  in, out chan Any
  FuncSpectrum
  PredSpectrum
  isServer, oneOne bool
  net.Conn
  nClients uint
  net.Listener
  buf []byte
  int "number of sent/received bytes"
  error
}

func new_(a Any, me, i uint, h string, p uint16) NetChannel {
  if me == i { panic("nchan.New: me == i") }
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
  ps := ":" + strconv.Itoa(int(Port0 + p))
  if x.isServer {
    x.Listener, x.error = net.Listen (network, h + ps)
    x.Conn, x.error = x.Listener.Accept()
  } else { // client
    for {
      if x.Conn, x.error = net.Dial (network, h + ps); x.error == nil {
        break
      }
      time.Sleep (500 * 1e6)
    }
  }
  return x
}

func (x *netChannel) Send (a Any) {
  if x.Any == nil {
    x.uint = Codelen(a)
    x.int, x.error = x.Conn.Write (append (Encode (x.uint), Encode(a)...))
  } else {
    CheckTypeEq (x.Any, a)
    x.int, x.error = x.Conn.Write(Encode(a))
  }
}

func (x *netChannel) Recv() Any {
  if x.Conn == nil { panic("nchan.Recv: x.Conn == nil") }
  if x.Any == nil {
    x.int, x.error = x.Conn.Read(x.buf[:C0])
    x.uint = Decode (uint(0), x.buf[:C0]).(uint)
    x.int, x.error = x.Conn.Read(x.buf[C0:C0+x.uint])
    return x.buf[C0:C0+x.uint]
  }
  x.int, x.error = x.Conn.Read(x.buf)
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
