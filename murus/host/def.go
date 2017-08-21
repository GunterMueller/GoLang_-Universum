package host

// (c) murus.org  v. 161222 - license see murus.go

import
  . "murus/obj"
const ( // Format
  Hostname = iota
  IPnumber
  NFormats
)
type
  Host interface { // Hostnames and their IPv6-numbers

  Editor
  Formatter
  Stringer
//  Printer TODO
  Marker

// Returns the IP-number of x as byte sequence.
  IP() []byte

// Returns true, if x has the name s.
  Equiv (s string) bool

// Returns true, iff x is the local host, that runs the calling process.
  Local() bool
}

// Retirms a new empty host.
func New() Host { return new_() }

func Localhost() Host { return localhost() }
