package host

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  "os"; "net"
  . "µU/ker"
  . "µU/obj"
  "µU/str"
  "µU/col"
  "µU/scr"
  "µU/box"
  "µU/errh"
)
type
  host struct {
           ip net.IP
         name []string // maximal length max see some header-file in /usr/include[/...]
              Format
       cF, cB col.Colour
              bool "marked"
              }
var (
  localHost Host = new_()
  localname string
  localIP net.IP
  ll [NFormats]uint = [NFormats]uint { 32, 39 } // 32: nackte Willkür
  bx = box.New()
  nullIP = net.IPv6zero
)

func init() {
  var e error
  localname, e = os.Hostname()
  if e != nil { Panic("hostname not defined") }
  localHost = localhost()
}

func new_() Host {
  x := new(host)
  x.Clr()
  x.Format = Hostname
  x.cF, x.cB = scr.StartCols()
  return x
}

func news (s string) Host {
  x := new(host)
  if ! x.Defined (s) { return nil }
  return x
}

func (x *host) imp (Y Any) *host {
  y, ok := Y.(*host)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *host) Empty() bool {
  return x.name[0] == ""
}

func (x *host) Clr() {
  x.name = []string { "" }
  x.ip = nullIP
}

func (x *host) Eq (Y Any) bool {
  y := x.imp(Y)
  if x.Empty() { return y.Empty() }
  if y.Empty() { return x.Empty() }
  e := x.ip.Equal(y.ip)
  if len(x.name) != len(y.name) { return false }
  if e {
    if x.name[0] != y.name[0] { Panic("shit: ip.Equal, but " + x.name[0] + " != " + y.name[0]) }
  } else {
    if x.name[0] == y.name[0] { Panic("shit: ! ip.Equal, but " + x.name[0] + " == " + y.name[0]) }
  }
  return e
}

func (x *host) Less (Y Any) bool {
  y := x.imp(Y)
  if len(x.ip) != len(y.ip) {
    println("Less: len(x.ip) =", len(x.ip), "but len(y.ip) =", len(y.ip)) // XXX
  }
  for i, b := range y.ip {
    if x.ip[i] < b {
      return true
    } else if x.ip[i] > b {
      break
    }
  }
  return false
}

func (x *host) Copy (Y Any) {
  y := x.imp(Y)
  if y.Empty() {
    x.Clr()
    return
  }
  n := len(y.name)
// println ("len y.name ==", n, "y.name[0] =", y.name[0])
  x.name = make([]string, n)
  for i:= 0; i < n; i++ {
    x.name[i] = y.name[i]
// println (">" + x.name[i] + "<")
  }
// println ("len x.name ==", n, "x.name[0] =", x.name[0])
//  for i, b := range y.ip.To4() { x.ip[i] = b } // changes ip of h0 - WHY THE HELL ???
  if n > 0 {
    if ! x.Defined(x.name[0]) { Panic("shit: x.name[0] = " + x.name[0] + "!") } // so geht's
  }
}

func (x *host) Clone() Any {
  y := new_()
// println("Clone: len(x.name) ==", len(x.name))
  y.Copy(x)
// println("Clone: len(y.name) ==", len(y.(*host).name))
  return y
}

func (x *host) Codelen() uint {
  if len(x.ip) != 16 { Panic("shit happens: len(x.ip) != 16") }
  return uint(len(x.ip))
}

func (x *host) Encode() Stream {
  b := make (Stream, x.Codelen())
  copy (b, x.ip)
  return b
}

func (x *host) Decode (b Stream) {
  copy (x.ip, b)
  if ! x.Defined (x.ip.String()) { Panic ("oops " + x.ip.String()) }
}

func (x *host) GetFormat() Format {
  return x.Format
}

func (x *host) SetFormat (f Format) {
  if f < NFormats {
    x.Format = f
  }
}

// Pre: len(n) != 0.
// Returns the smallest index with a minimal number of n.
func min (n []int) int {
  i := 0
  m := n[i]
  for j := 1; j < len(n); j++ {
    if n[j] < m {
      m = n[j]
      i = j
    }
  }
  return i
}

func (x *host) String() string {
  if x.Format == Hostname {
    n := len(x.name)
    switch n {
    case 0:
      return "<nil>"
    case 1:
      return x.name[0]
    default:
      k := make([]int, n)
      for i:= 0; i < n; i++ {
        k[i] = len(x.name[i])
      }
      return x.name[min(k)]
    }
  }
  return x.ip.String()
}

const
  providerIP = "81.223.238.231" // www2.sprit.org

func (x *host) Defined (s string) bool {
  str.OffSpc(&s)
  x.Clr()
  x.ip = net.ParseIP(s) // s daraufhin prüfen, ob es eine IP-Nummer ist
  if x.ip != nil {
    Panic ("host.Defined: s is IP-number with x.ip == " + x.ip.String())
    // s is a IP-number in the form of "a.b.c.d"
  } else { // s is no IP-number, but has to be tested for hostname
//    println ("host.Defined: x.ip == nil")
    addr, err := net.LookupHost(s) // schau'mer moi, ob's oana is
    if err == nil { // mir hoam's, s'is oana
      s = addr[0] // is the IP-number
      if s == providerIP { return false } // chock out fucking provider
      if s == nullIP.String() { // s == "0.0.0.0"
        x.Clr()
        return true
      }
      if len(addr) > 1 {
        const (
          a1 = "127.0.0.1"
          a2 = "127.0.0.2"
        )
        for _, a  := range addr {
          if a != a1 && a != a2 {
            s = a
            break
          }
        }
      }
      x.ip = net.ParseIP(s) // .To4()
//      println ("host.Defined: len(addr) > 1 => x.ip =", x.ip.String(), len(x.ip))
    } else { // err != nil
      x.Clr()
      return false
    }
  }
  var e error // x.name herausfinden
  if x.name, e = net.LookupAddr(s); e == nil {
//    println ("host.Defined: Ende => x.name[0], x.ip =", x.name[0], ", ", x.ip.String())
    if len(x.name) > 1 {
//      println ("x.name[1] = ", x.name[1])
      if len(x.name[0]) > len(x.name[1]) {
        x.name[0], x.name[1] = x.name[1], x.name[0]
      }
    }
    return true
  }
  x.Clr()
  return false
}

func (x *host) Equiv (s string) bool {
  b := net.ParseIP (s)
  if b != nil {
    return x.ip.Equal (b)
  }
  for i := 0; i < len (x.name); i++ {
    if x.name[i] == s {
      return true
    }
  }
  return false
}

func (x *host) Colours (f, b col.Colour) {
  x.cF, x.cB = f, b
}

func (x *host) Write (l, c uint) {
  bx.Wd (ll [x.Format])
  bx.Colours (x.cF, x.cB)
  bx.Write (x.String(), l, c)
}

func (x *host) Edit (l, c uint) {
  x.Write (l, c)
  s := x.String()
  for {
    bx.Edit (&s, l, c)
    if x.Defined (s) {
       break
    } else {
      errh.Error0 ("falsche Eingabe")
    }
  }
}

func (x *host) Mark (m bool) {
  x.bool = m
}

func (x *host) Marked() bool {
  return x.bool
}

func (x *host) IP() Stream {
  return x.ip
}

func (x *host) Local() bool {
  return x.Eq(localHost)
}

/*
func local (s string) bool {
  for i := 0; i < len (localHost.name); i++ {
    if localHost.name[i] == s {
      return true
    }
  }
  return false
}
*/

func localhost() Host {
  if ! localHost.Defined(localname) { Panic ("Hostname not defined") }
  localIP = localHost.(*host).ip
//  n := len(localHost.(*host).name); if n == 0 { ker.Panic("localhost: jaul") }
  return localHost.Clone().(*host)
}
