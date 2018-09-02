package msgkmz

// (c) Christian Maurer  v. 180819 - license see µU.go

import (
  . "nU/obj"
  "nU/dgra/status"
)
const (
  Ask = byte(iota); Accept; Update; YourCitizen; Leader; NKinds // kinds of messages
)
type
  Message interface {

  Equaler
  Coder

// Returns the kind of x.
  Kind() byte

// Returns the status of x.
  Status() status.Status
}

// Returns a new message of kind k and status s.
func NewMsg (k byte, s status.Status) Message { return new_(k,s) }
