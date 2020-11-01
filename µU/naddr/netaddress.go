package naddr

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  . "µU/obj"
  "µU/str"
  "µU/nat"
  "µU/host"
)
type
  netAddress struct {
                    host.Host
                    uint16 "port"
                    Format // see µU/host/def.go
                    }

func new_(p uint16) NetAddress {
  return &netAddress { host.New(), p, host.Hostname }
//                 see mu/host/def.go ^^^^^^^^^^^^^
}

func new2 (h host.Host, p uint16) NetAddress {
  x := new_(p).(*netAddress)
  x.Host.Copy(h)
  return x
}

func (x *netAddress) imp (Y Any) *netAddress {
  y, ok := Y.(*netAddress)
  if ! ok { TypeNotEqPanic(x, Y) }
  return y
}

func (x *netAddress) Empty() bool {
  return x.Host.Empty()
}

func (x *netAddress) Clr() {
  x.Host.Clr()
  x.uint16 = 0
}

func (x *netAddress) Eq (Y Any) bool {
  y := x.imp(Y)
  return x.Host.Eq(y.Host) &&
         x.uint16 == y.uint16
}

func (x *netAddress) Less (Y Any) bool {
  y := x.imp(Y)
  if x.Host.Eq(y.Host) {
    return x.uint16 < y.uint16
  }
  return x.Host.Less(y.Host)
}

func (x *netAddress) Copy (Y Any) {
  y := x.imp(Y)
  x.Host.Copy(y.Host)
  x.uint16 = y.uint16
}

func (x *netAddress) Clone() Any {
  y := new_(x.uint16).(*netAddress)
  y.Host.Copy(x.Host)
  return y
}

func (x *netAddress) Codelen() uint {
  return x.Host.Codelen() + 2
}

func (x *netAddress) Encode() Stream {
  b := make (Stream, x.Codelen())
  cl := x.Host.Codelen()
  copy (b[:cl], x.Host.Encode())
  copy (b[cl:], Encode (x.uint16))
  return b
}

func (x *netAddress) Decode (b Stream) {
  cl := x.Host.Codelen()
  x.Host.Decode (b[:cl])
  x.uint16 = Decode (x.uint16, b[cl:]).(uint16)
}

func (x *netAddress) GetFormat() Format {
  return x.Format
}

func (x *netAddress) SetFormat (f Format) {
  if f < host.NFormats {
    x.Format = f
  }
}

const
  separator = byte(':')

func (x *netAddress) Defined (s string) bool {
  n := uint(len (s))
  if i, ok := str.Pos (s, separator); ok && i < n {
    if p, ok1 := nat.Natural (str.Part (s, i + 1, n - (i + 1))); ok1 && p < 1<<16 {
      return x.Host.Defined (str.Part (s, 0, i))
    }
  }
  return false
}

func (x *netAddress) String() string {
  x.Host.SetFormat (x.Format)
  return x.Host.String() + string(separator) + nat.String (uint(x.uint16))
}

func (x *netAddress) Set (h host.Host, p uint16) {
  x.Host.Copy(h)
  x.uint16 = p
}

func (x *netAddress) SetPort (p uint16) {
  x.uint16 = p
}

func (x *netAddress) HostPort() (host.Host, uint16) {
  return x.Host.Clone().(host.Host), x.uint16
}

func (x *netAddress) IPPort() (Stream, uint16) {
  f := x.Host.GetFormat()
  x.Host.SetFormat (host.IPnumber)
  defer x.Host.SetFormat(f)
  return x.Host.IP(), x.uint16
}

func (x *netAddress) Port() uint16 {
  return x.uint16
}

func (x *netAddress) Local() bool {
  return x.Host.Local()
}
