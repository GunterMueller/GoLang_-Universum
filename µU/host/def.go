package host

// (c) Christian Maurer   v. 201014 - license see µU.go

import (
  . "µU/obj"
  "µU/col"
)
const ( // Format
  Hostname = iota
  IPnumber
  NFormats
)
type
  Host interface { // Hostnames and their IPv4-numbers

  Object
  col.Colourer
  Editor
  Formatter
  Stringer
//  Printer TODO
  Marker

// Returns the IPv4-number of x as byte sequence.
  IP() Stream

// Returns true, if x has the name s.
  Equiv (s string) bool

// Returns true, iff x is the local host, that runs the calling process.
  Local() bool
}

// Returns a new empty host.
func New() Host { return new_() }

// Pre: s denotes a host contained in /etc/hosts
//      or reachable by DNS.
func NewS (s string) Host { return news(s) }

func Localhost() Host { return localhost() }

func LocalName() string { return localname }
