package msg

// (c) Christian Maurer  v. 190402 - license see µU.go

import (
  . "nU/obj"
  "nU/dgra/status"
)
const (
  Ask = byte(iota); Accept; Update; YourCitizen; Leader; NTypes // types of messages
)
type
  Message interface { // pair of kind and Status

  Equaler
  Coder

// x has type t and status s.
  Set (t byte, s status.Status)

// Returns the type of x.
  Type() byte

// Returns the string of x.
  String() string

// Returns the status of x.
  Status() status.Status
}

// Returns a new message of undefined type and status.
func NewMsg() Message { return new_() }
