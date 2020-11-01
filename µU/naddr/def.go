package naddr

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  . "µU/obj"
  "µU/host"
)
type
  NetAddress interface { // host and IPnumber

  Object
  Formatter // see "µU/host"
  Stringer

// Returns true, iff (h, p) defines an IP4-net address.
  Set (h host.Host, p uint16)

  SetPort (p uint16)

// Returns the host and the port of x.
  HostPort() (host.Host, uint16)

// Returns the IPv6-number and the port of x.
//  IPPort() (Stream, uint16) // ? XXX

// Returns the port of x.
  Port() uint16

// Returns true, if the host of x is the calling host.
  Local() bool
}

// Returns a new net address :p (without host, for servers).
func New (p uint16) NetAddress { return new_(p) }

// Returns the new net address h:p.
func New2 (h host.Host, p uint16) NetAddress { return new2(h,p) }

func NewLocal (p uint16) NetAddress { return new2(host.Localhost(), p) }
