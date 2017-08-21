package internal

// (c) murus.org  v. 170112 - license see murus.go

import (
  . "murus/obj"
)
type Type byte; const (
  Candidate = Type(iota)
  Reply
  Won
)
func NewMsg() Message { return new_() }

type
  Message interface {

  Object

  Content() (Type, uint, uint, uint, bool)
  PassCandidate (i, r, d uint)
  Reply (t bool)
  PassWon (i uint)
  Pass ()
}
