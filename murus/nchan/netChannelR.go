package nchan

// (c) murus.org  v. 161228 - license see murus.go
//
// XXX NOT TO BE USED - probably not correct (at least a proof of correctness is missing) !!!

import (
  "syscall"
  "strings"
  "errors"
  "time"
  "net"
  . "murus/ker" // Panic
  . "murus/obj"
  "murus/errh"
  "murus/host"; "murus/naddr"
)

func newR (a Any, h host.Host, p uint16, o bool) NetChannel {
  return new1R (a, naddr.New2 (h, p), o)
}

func new1R (a Any, n naddr.NetAddress, o bool) NetChannel {
  x:= new (netChannel)
  if a == nil {
    x.Any, x.width = nil, maxWidth
  } else {
    x.Any, x.width = Clone(a), Codelen(a)
  }
  x.buf = make ([]byte, x.width)
  n0 := naddr.New2 (host.Localhost(), n.Port())
  n0.SetFormat (host.Hostname)
  nEq := false
  x.oneOne = o
//  errh.Error2 (n0.String() + " ->", 0, n.String(), 0) // just to test
  if x.oneOne { // 1:1
    if n.Empty() {
      x.error = errors.New ("nchan.New: n.Empty()")
      Panic (x.error.Error())
    }
//    x.port = n.Port()
    nEq = n0.Eq (n)
    if nEq { ///////////////////////////////////////////////////////////////////
      errh.Hint (n0.String() + " listens")
      if x.Listener, x.error = net.Listen (network, n0.String()); x.error == nil {
        x.isServer = true
      } else {
        a := syscall.Errno(syscall.EADDRINUSE).Error()
        if strings.Contains (x.error.Error(), a) {
          x.isServer = false
        } else {
          Panic (x.error.Error())
        }
      }
      errh.DelHint() ////////////////////////////////////////////////////////////
    } else {
      x.isServer = n0.Less (n)
// if x.isServer { println ("n0 < n") } else { println ("n0 >= n") }
    }
  } else { // 1:n
    x.isServer = n.Empty()
  }
  if x.isServer {
    if x.oneOne {
      if ! nEq { // if nEq, Listen was already called
        errh.Hint (n0.String() + " listens")
        x.Listener, x.error = net.Listen (network, n0.String())
        x.panicIfErr()
      }
      x.Conn, x.error = x.Listener.Accept()
      x.panicIfErr()
    } else { // 1:n
      x.Listener, x.error = net.Listen (network, n0.String())
      x.panicIfErr()
      x.ccin, x.ccout = make (chan Any), make (chan Any)
      go func() {
        for {
          if c, e:= x.Listener.Accept(); e == nil { // NOT x.Conn, x.error !
//            println ("server accepted", c.RemoteAddr().(*net.TCPAddr).String()) // just to test
            x.nClients ++
//            errh.Hint ("number of clients: " + nat.String(x.nClients))
            go x.serve (c) // see above remark
          } else {
            Panic (e.Error())
          }
        }
      }()
    }
  } else { // client
    if n.Empty() { Oops() } // check Pre
    dialaddr:= n.String() // ! n.Host().Empty()
    for {
      errh.Hint ("waiting for " + dialaddr)
      if x.Conn, x.error = net.Dial (network, dialaddr); x.error == nil {
        errh.DelHint()
        break
      }
      time.Sleep (2 * 1e9)
      errh.Hint (x.error.Error()) // just to test
    }
  }
  return x
}
